package main

import (
	"github.com/zenazn/goji"	
	"net/http"
	
)


func main() {

    staticFilesLocation := "/tmp/127.0.0.1/"
    goji.Handle("/*", http.FileServer(http.Dir(staticFilesLocation)))

    goji.Serve()
}

