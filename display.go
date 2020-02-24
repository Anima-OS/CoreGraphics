// Copyright (c) 2018, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.
//

package cg

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Display struct {
	GPU        string `json:"gpu"`
	OGLversion string `json:"openGL_version"`
	ID         string `json:"displayID"`
	XID        string `json:"X11displayID"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	Gamma      Gamma
}

type Gamma struct {
	Red   []uint16
	Green []uint16
	Blue  []uint16
}

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func MainDisplay() Display {
	var dc Display

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Visible, glfw.False)

	window, err := glfw.CreateWindow(640, 480, "Hello World!", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	GPU := gl.GoStr(gl.GetString(gl.RENDERER))
	version := gl.GoStr(gl.GetString(gl.VERSION))

	monitor := glfw.GetPrimaryMonitor()
	ID := monitor.GetName()
	//XID := glfw.GetX11Display()

	video := monitor.GetVideoMode()
	gamma := monitor.GetGammaRamp()

	dc.GPU = GPU
	dc.OGLversion = version
	dc.ID = ID
	dc.Width = video.Width
	dc.Height = video.Height

	dc.Gamma.Red = gamma.Red
	dc.Gamma.Green = gamma.Green
	dc.Gamma.Blue = gamma.Blue

	return dc
}
