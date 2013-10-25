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
	targetFile = flag.Strng("pagepath", "", "Static page path to generate or serve")
)

func main() {
	flag.Parse()
}
