// shbrowseforfolder.go
// Copyright (c) 2021 Tobotobo
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

// Copyright 2010 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// https://github.com/lxn/walk/blob/master/LICENSE

// +build windows

//lint:file-ignore SA1019 syscall.StringToUTF16 と syscall.StringToUTF16Ptr を使用します

package shbrowseforfolder

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"

	"github.com/Tobotobo/msgbox"
	"github.com/lxn/win"
)

func newError(message string) error {
	return errors.New(message)
}

func pathFromPIDL(pidl uintptr) (string, error) {
	var path [win.MAX_PATH]uint16
	if !win.SHGetPathFromIDList(pidl, &path[0]) {
		return "", newError("SHGetPathFromIDList failed")
	}

	return syscall.UTF16ToString(path[:]), nil
}

var initialized func(win.HWND, uint32, uintptr, uintptr) = nil

// We use this callback to disable the OK button in case of "invalid" selections.
func browseFolderCallback(hwnd win.HWND, msg uint32, lp, wp uintptr) uintptr {
	switch msg {
	case BFFM_INITIALIZED:
		if initialized != nil {
			initialized(hwnd, msg, lp, wp)
		}
	case BFFM_SELCHANGED:
		_, err := pathFromPIDL(lp)
		var enabled uintptr
		if err == nil {
			enabled = 1
		}
		win.SendMessage(hwnd, BFFM_ENABLEOK, 0, enabled)
	case BFFM_VALIDATEFAILED:
		msgbox.Owner(hwnd).Err().Show("フォルダー名が無効です。", "エラー")
		return 1
	}
	return 0
}

var browseFolderCallbackPtr uintptr = 0

func init() {
	browseFolderCallbackPtr = syscall.NewCallback(browseFolderCallback)
}

func Show(owner win.HWND, title string, rootDirPath string, initSelectedPath string, flags uint32) (selectedPath string, accepted bool, err error) {
	defer func() {
		initialized = nil
	}()

	// Calling OleInitialize (or similar) is required for BIF_NEWDIALOGSTYLE.
	if hr := win.OleInitialize(); hr != win.S_OK && hr != win.S_FALSE {
		return "", false, newError(fmt.Sprint("OleInitialize Error: ", hr))
	}
	defer win.OleUninitialize()

	bi := win.BROWSEINFO{
		HwndOwner: owner,
		LpszTitle: syscall.StringToUTF16Ptr(title),
		UlFlags:   flags | BIF_VALIDATE,
		Lpfn:      browseFolderCallbackPtr,
	}

	initialized = func(hwnd win.HWND, msg uint32, lp, wp uintptr) {
		if len(initSelectedPath) > 0 {
			ptrSelectedPath := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(initSelectedPath)))
			win.SendMessage(hwnd, BFFM_SETSELECTION, 1, ptrSelectedPath)
			win.SendMessage(hwnd, BFFM_SETEXPANDED, 1, ptrSelectedPath)
		}
	}

	// We need to put the initial path into a buffer of at least MAX_LENGTH
	// length, or we may get random crashes.
	var buf [win.MAX_PATH]uint16
	copy(buf[:], syscall.StringToUTF16(rootDirPath))
	win.SHParseDisplayName(&buf[0], 0, &bi.PidlRoot, 0, nil)
	defer win.CoTaskMemFree(bi.PidlRoot)

	pidl := win.SHBrowseForFolder(&bi)
	if pidl == 0 {
		return "", false, nil
	}
	defer win.CoTaskMemFree(pidl)

	selectedPath, err = pathFromPIDL(pidl)
	accepted = selectedPath != ""
	return
}
