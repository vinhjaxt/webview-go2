package main

import _ "embed"

//go:embed x64_webview.dll
var x64_webview []byte

//go:embed x86_webview.dll
var x86_webview []byte
