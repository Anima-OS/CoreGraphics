// Copyright (c) 2018, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.
//

package main

import "C"
import "runtime"
import ".."

//export CGMainDisplay
func CGMainDisplay() *C.char {
	dc := cg.MainDisplay()
	return C.CString(dc.GPU)
}

func main() {}
