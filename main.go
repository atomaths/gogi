// Copyright 2013 Jongmin Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	src        = flag.String("src", "", "source file for replacement")
	recursive  = flag.Bool("r", false, "recursive flag for target directory")
	targetFile = flag.String("f", "", "target file")
	pagePath   = flag.String("path", "", "Static page path to serve or replace")
	host       = flag.String("host", "127.0.0.1", "Host IP to listen")
	port       = flag.String("port", "3999", "Listen port")
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello, World")
}

func main() {
	flag.Parse()

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(*host+":"+*port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
