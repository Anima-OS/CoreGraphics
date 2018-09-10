// Copyright (c) 2018, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.
//

package cg

// #cgo pkg-config: cairo cairo-gobject gdk-3.0
// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <gdk/gdk.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "pixbuffer.h"
import "C"

import (
	"strconv"
	"unsafe"
)

type Context struct {
	surface *C.cairo_surface_t
	context *C.cairo_t
	width   float64
	height  float64
}

func (c *Context) ShowText(utf8 string) {
	text := C.CString(utf8)
	defer C.free(unsafe.Pointer(text))
	C.cairo_show_text(c.context, text)
}

// BeginPath : CGContextBeginPath
func (c *Context) BeginPath() {
	C.cairo_new_path(c.context)
}

func (c *Context) MoveTo(x, y float64) {
	C.cairo_move_to(c.context, C.double(x), C.double(y))
}

// CreatePDFContext : CGCreatePDFContext
func CreatePDFContext(file string, width, height float64) *Context {
	fileName := C.CString(file)
	defer C.free(unsafe.Pointer(fileName))
	s := C.cairo_pdf_surface_create(fileName, C.double(width), C.double(height))

	return &Context{surface: s, context: C.cairo_create(s), width: width, height: height}
}

// Close : CGContextClose
func (c *Context) Close() {
	C.cairo_surface_finish(c.surface)
}

func (c *Context) SetSourceRGB(r, g, b float64) {
	C.cairo_set_source_rgb(c.context, C.double(r), C.double(g), C.double(b))
}

func (c *Context) LineTo(x, y float64) {
	C.cairo_line_to(c.context, C.double(x), C.double(y))
}

func (c *Context) Fill() {
	C.cairo_fill(c.context)
}

func (c *Context) ClosePath() {
	C.cairo_close_path(c.context)
}

func (c *Context) Rect(x, y, width, height float64) {
	C.cairo_rectangle(c.context,
		C.double(x), C.double(y),
		C.double(width), C.double(height))
}

// TODO(Happy-Ferret): Deprecate and merge into generic "ToFile" method.
func (c *Context) WriteToPNG(fileName string) {

	file := C.CString(fileName)
	defer C.free(unsafe.Pointer(file))

	C.cairo_surface_write_to_png(c.surface, file)
}

func CreatePSContext(file string, width, height float64) *Context {
	fileName := C.CString(file)
	defer C.free(unsafe.Pointer(fileName))
	s := C.cairo_ps_surface_create(fileName, C.double(width), C.double(height))

	return &Context{surface: s, context: C.cairo_create(s)}
}

func CreateImageContext(f C.cairo_format_t, width, height float64) *Context {
	s := C.cairo_image_surface_create(f, C.int(width), C.int(height))

	return &Context{surface: s, context: C.cairo_create(s)}
}

// CGContextToFile
//
// TODO(Happy-Ferret): Rename to a more generic method.
func (c *Context) DrawToJPG(file string, format string) {
	fileName := C.CString(file)
	f := C.CString(format)
	t := C.CString("quality")
	quality := C.CString(strconv.Itoa(100))

	defer C.free(unsafe.Pointer(fileName))
	defer C.free(unsafe.Pointer(f))
	defer C.free(unsafe.Pointer(t))
	defer C.free(unsafe.Pointer(quality))

	var err *C.GError
	pixbuf := c.ToPixbuf()
	C._gdk_pixbuf_save(pixbuf.GObject, fileName, f, &err, quality)
}

// CGContextGetWidth
func (c *Context) GetWidth() float64 {
	return c.width
}

// CGContextGetHeight
func (c *Context) GetHeight() float64 {
	return c.height
}

func (c *Context) ToPixbuf() *Pixbuf {
	width := c.GetWidth()
	height := c.GetHeight()
	pixbuf := C.gdk_pixbuf_get_from_surface(c.surface, 0, 0, C.int(width), C.int(height))
	return &Pixbuf{GObject: pixbuf}
}

// CGContextClippedToPixbuf
//func (c *Context) ClippedToPixbuf(x, y, width, height int) *Pixbuf {
//	pixbuf := C.gdk_pixbuf_get_from_surface (c.surface, x, y, width, height);
//        return &Pixbuf{GObject: pixbuf}
//}

//type Pixbuf struct {
//	GObject *C.GdkPixbuf
//}
