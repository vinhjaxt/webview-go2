package main

// go build -ldflags "-H windowsgui"

// import "C"

import (
	"log"
	"os"
	"reflect"
	"runtime"
	"strings"
	"syscall"
	"unsafe"
)

// /*
func getString(s uintptr) string {
	len := uintptr(0)
	for {
		ch := *(*uint8)(unsafe.Pointer(s + len))
		if ch == 0 {
			break
		}
		len++
	}
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: s,
		Len:  int(len),
	}))
}

func myFunc(seq *int, req uintptr, _ uintptr) uintptr {
	log.Println(*seq, getString(req))
	os.Exit(1)
	return 0
}

func myFunc2(seq *int, req uintptr, _ uintptr) uintptr {
	log.Println(*seq, getString(req))
	return 0
}

/*/

/*
func myFunc(seq *int, req *C.char, _ uintptr) uintptr {
	log.Println(*seq, C.GoString(req))
	return 0
}
//*/

func checkErr(label string, err error) {
	if err == nil {
		return
	}
	if strings.Contains(err.Error(), "completed successfully") {
		return
	}
	log.Panicln(label, err)
}

var (
	wv                uintptr
	webview           = platformDll()
	webview_create    = webview.NewProc("webview_create")
	webview_set_title = webview.NewProc("webview_set_title")
	webview_set_size  = webview.NewProc("webview_set_size")
	webview_bind      = webview.NewProc("webview_bind")
	webview_navigate  = webview.NewProc("webview_navigate")
	webview_run       = webview.NewProc("webview_run")
	webview_destroy   = webview.NewProc("webview_destroy")
)

func platformDll() *syscall.LazyDLL {
	var path string
	if runtime.GOARCH == "amd64" {
		path = os.TempDir() + "\\x64_webview.dll"
		os.WriteFile(path, x64_webview, os.ModePerm)
	} else {
		path = os.TempDir() + "\\x86_webview.dll"
		os.WriteFile(path, x86_webview, os.ModePerm)
	}
	return syscall.NewLazyDLL(path)
}

func main() {
	log.Println("Starting..")
	var err error
	wv, _, err = webview_create.Call(0, 0)
	checkErr("webview_create", err)

	title := []byte("Webview example: Vịnh")
	_, _, err = webview_set_title.Call(wv, uintptr(unsafe.Pointer(&title[0])))
	checkErr("webview_set_title", err)

	_, _, err = webview_set_size.Call(wv, 480, 320, 0)
	checkErr("webview_set_size", err)

	fnName := []byte("myFunc")
	_, _, err = webview_bind.Call(wv, uintptr(unsafe.Pointer(&fnName[0])), syscall.NewCallback(myFunc), 0)
	checkErr("webview_bind", err)

	fnName = []byte("myFunc2")
	_, _, err = webview_bind.Call(wv, uintptr(unsafe.Pointer(&fnName[0])), syscall.NewCallback(myFunc2), 0)
	checkErr("webview_bind", err)

	html := []byte(`data:text/html, <meta charset="utf-8"><button onclick='myFunc("Foo bar")'>Click Vịnh</button><script>myFunc2("Foo bar2")</script>`)
	_, _, err = webview_navigate.Call(wv, uintptr(unsafe.Pointer(&html[0])))
	checkErr("webview_navigate", err)

	/*
		go func() {
			for {
				time.Sleep(time.Second)
				log.Println("OK")
			}
		}()
		// */

	_, _, err = webview_run.Call(wv)
	checkErr("webview_run", err)

	_, _, err = webview_destroy.Call(wv)
	checkErr("webview_destroy", err)
}
