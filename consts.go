// consts.go
// Copyright (c) 2021 Tobotobo
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

package shbrowseforfolder

import "github.com/lxn/win"

// -------------------------------------------------------
// BROWSEINFO
// http://chokuto.ifdef.jp/urawaza/struct/BROWSEINFO.html
//

// ファイルシステムディレクトリのみを返します。
// それ以外のアイテムが選択されているときには、[OK]ボタンは灰色表示になります。
const BIF_RETURNONLYFSDIRS = 0x00000001

// ダイアログボックスのツリービューコントロールにドメインレベルのネットワークフォルダを含めないようにします。
const BIF_DONTGOBELOWDOMAIN = 0x00000002

// ダイアログボックスにステータス領域を表示します。
// 表示テキストを設定するには、コールバック関数からダイアログボックスにメッセージを送信します。
const BIF_STATUSTEXT = 0x00000004

// シェルネームスペース階層構造の中でルートフォルダの下にあるファイルシステムサブフォルダのみを返します。
// それ以外のアイテムが選択されているときには、[OK]ボタンは灰色表示になります。
const BIF_RETURNFSANCESTORS = 0x00000008

// Version 4.71 以降： ユーザーがアイテム名を書き込むことができるエディットコントロールを表示します。
const BIF_EDITBOX = 0x00000010

// Version 4.71 以降： ユーザーがエディットコントロールに無効な名前を入力した場合に、
// BFFM_VALIDATEFAILED メッセージとともにコールバック関数が呼び出されます。
// BIF_EDITBOXフラグが指定されていない場合は、このフラグは無視されます。
const BIF_VALIDATE = 0x00000020

// Version 5.0 以降： 新しいユーザーインターフェースを使用します。
// 従来のダイアログボックスよりも大きい、リサイズ可能なダイアログボックスが表示され、
// ダイアログボックスへのドラッグアンドドロップ、フォルダの再整理、ショートカットメニュー、
// 新しいフォルダ作成、削除、その他のショートカットメニューコマンドが追加されます。
// このフラグを使用するには、あらかじめOleInitialize関数またはCoInitialize関数を呼び出して
// COMを初期化しておく必要があります。
const BIF_NEWDIALOGSTYLE = 0x00000040

// Version 5.0 以降： エディットコントロールを持つ、新しいユーザーインターフェースを使用します。
// このフラグはBIF_EDITBOX|BIF_NEWDIALOGSTYLEと同等です。
// このフラグを使用するには、あらかじめOleInitialize関数またはCoInitialize関数を呼び出して
// COMを初期化しておく必要があります。
const BIF_USENEWUI uint32 = 0x00000050

// Version 6.0 以降： エディットコントロールの代わりに、ダイアログボックスに用法ヒントを追加します。
// BIF_NEWDIALOGSTYLEフラグとともに指定しなければなりません。
const BIF_UAHINT = 0x00000100

// Version 6.0 以降： ダイアログボックスに「新しいフォルダ」ボタンを表示しないようにします。
// BIF_NEWDIALOGSTYLEフラグとともに指定しなければなりません。
const BIF_NONEWFOLDERBUTTON = 0x00000200

// Version 6.0 以降： 選択されたアイテムがショートカットであるとき、そのリンク先ではなく、
// ショートカットファイル自体のPIDLを返します。
const BIF_NOTRANSLATETARGETS = 0x00000400

// コンピュータのみを返します。それ以外のアイテムが選択されているときには、
// [OK]ボタンは灰色表示になります。
const BIF_BROWSEFORCOMPUTER = 0x00001000

// プリンタのみを返します。それ以外のアイテムが選択されているときには、
// [OK]ボタンは灰色表示になります。
const BIF_BROWSEFORPRINTER = 0x00002000

// Version 4.71 以降： フォルダとファイルを表示します。
const BIF_BROWSEINCLUDEFILES = 0x00004000

// Version 5.0 以降： リモートシステム上にある共有リソースを表示できるようにします。
// BIF_USENEWUIフラグとともに指定しなければなりません。
const BIF_SHAREABLE = 0x00008000

// --------------------------------------------------
// 第２７６章 フォルダ選択ダイアログを出す
// http://www.kumei.ne.jp/c_lang/sdk3/sdk_276.htm
//

// ダイアログボックスの初期化が終わりました。
// lParamは0です。
const BFFM_INITIALIZED = 1

// フォルダの選択が変化しました。
// lParamには新しく選択された フォルダのアイテムIDが入ります。
const BFFM_SELCHANGED = 2

// ユーザーがエディットボックスに無効な名前をタイプしました。
// lParamは無効な名前を含んでいるバッファのアドレスを示します。
const BFFM_VALIDATEFAILEDA = 3

// ユーザーがエディットボックスに無効な名前をタイプしました。
// lParamは無効な名前を含んでいるバッファのアドレスを示します。
const BFFM_VALIDATEFAILEDW = 4

// ユーザーがエディットボックスに無効な名前をタイプしました。
// lParamは無効な名前を含んでいるバッファのアドレスを示します。
const BFFM_VALIDATEFAILED = BFFM_VALIDATEFAILEDW

// ステータステキストをセットします。
const BFFM_SETSTATUSTEXTA = win.WM_USER + 100

// ステータステキストをセットします。
const BFFM_SETSTATUSTEXTW = win.WM_USER + 104

// ステータステキストをセットします。
const BFFM_SETSTATUSTEXT = BFFM_SETSTATUSTEXTW

// ダイアログボックスのOKボタンを使用可能、不能にします。
// lParamをTRUEにするかFALSEにする事でボタン使用の可否を設定します。
const BFFM_ENABLEOK = win.WM_USER + 101

// 特定のフォルダを選択します。フォルダはPIDLで指定し、lParam で送ります。この時wParamはFALSEにします。
// もしくはフォルダを文字列として表しlParamで送ります。この時wParamはTRUEにします。
const BFFM_SETSELECTIONA = win.WM_USER + 102

// 特定のフォルダを選択します。フォルダはPIDLで指定し、lParam で送ります。この時wParamはFALSEにします。
// もしくはフォルダを文字列として表しlParamで送ります。この時wParamはTRUEにします。
const BFFM_SETSELECTIONW = win.WM_USER + 103

// 特定のフォルダを選択します。フォルダはPIDLで指定し、lParam で送ります。この時wParamはFALSEにします。
// もしくはフォルダを文字列として表しlParamで送ります。この時wParamはTRUEにします。
const BFFM_SETSELECTION = BFFM_SETSELECTIONW

// -----------------------------------------------------
// 【Win32】SHBrowseForFolder 関数
// https://ameblo.jp/sgl00044/entry-12547637468.html
//

// ダイアログボックスの[OK]ボタンに表示されるテキストを設定します。
// wParam : 未使用です。
// lParam : 終端NULL文字の表示文字列へのポインタを指定します。
const BFFM_SETOKTEXT = win.WM_USER + 105

// ダイアログボックスで展開するフォルダーのパスを指定します。
// パスはUnicode文字列またはPIDLとして指定できます。
// wParam : 文字列を使用する場合はTRUE、PIDLを使用する場合はFALSEを指定します。
// lParam : パスを指定する文字列またはPIDLへのポインタを指定します。
const BFFM_SETEXPANDED = win.WM_USER + 106
