// Copyright 2013 Jongmin Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"flag"
)

var (
	src = flag.String("src", "", "source file for replacement")
)

func main() {
	flag.Parse()
}
