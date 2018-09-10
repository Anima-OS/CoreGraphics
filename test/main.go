// Copyright (c) 2018, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.
//

package main

import ".."

func main() {
	//	dc := cg.MainDisplay()
	//	fmt.Println("\n" + dc.GPU)
	//	fmt.Println(dc.OGLversion)
	//	fmt.Println(dc.ID)

	context := cg.CreatePDFContext("cairo.pdf", 320, 200)
	context.MoveTo(0, 10)
	context.ShowText("Hello")
	context.Close()
}
