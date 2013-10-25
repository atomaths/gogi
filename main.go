// Copyright 2013 Jongmin Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"flag"
)

var (
	src        = flag.String("src", "", "source file for replacement")
	recursive  = flag.Bool("r", false, "recursive flag for target directory")
	targetFile = flag.Strng("f", "", "target file")
	pagePath   = flag.Strng("pagepath", "", "Static page path to generate or serve")
	host       = flag.String("host", "127.0.0.1", "Host IP to listen")
	port       = flag.String("port", "3999", "Listen port")
)

func main() {
	flag.Parse()
}
