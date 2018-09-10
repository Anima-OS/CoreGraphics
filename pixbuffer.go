// Copyright (c) 2018, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.
//

package cg

// #cgo pkg-config: cairo cairo-gobject gdk-3.0
// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

type Pixbuf struct {
	GObject *C.GdkPixbuf
}
