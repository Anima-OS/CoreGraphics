// Copyright (c) 2018, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.
//

package main

import (
	".."
	"math"
)

var hoist_width = 300.00
var fly_length = 570.00
var stripe_width = 0.0769 * hoist_width

var union_width = 0.5385 * hoist_width
var union_length = 0.76 * hoist_width

var outerRadius = 0.0616 * hoist_width / 2
var innerRadius = outerRadius * math.Sin(math.Pi/10) / math.Sin(7*math.Pi/10)

var xStarSeparation = hoist_width * 0.063
var yStarSeparation = hoist_width * 0.054

func main() {
	context := cg.CreatePDFContext("USA.pdf", fly_length, hoist_width)
	//	context := cg.CreateImageContext(C.CAIRO_FORMAT_ARGB32, fly_length, hoist_width)

	drawFlag(context)
	drawStripes(context)

	context.SetSourceRGB(250, 250, 250)
	context.MoveTo(10.00, 40.00)
	context.LineTo(100.00, 40.00)

	for row := 1.00; row <= 9.00; row++ {
		for col := 1.00; col <= 11.00; col++ {
			if math.Mod(row+col, 2) == 0 {
				drawStar(context, xStarSeparation*col, yStarSeparation*row, 5.00, outerRadius, innerRadius)
			}
		}
	}

	context.WriteToPNG("USA.png")
	context.DrawToJPG("USA1.jpeg", "jpeg")
	context.Close()
}

func drawFlag(context *cg.Context) {
	context.SetSourceRGB(200, 0, 0)
	context.Rect(0, 0, fly_length, hoist_width)
	context.Fill()
}

func drawStripes(context *cg.Context) {
	context.SetSourceRGB(250, 250, 250)

	for stripes := 1.00; stripes < 13.00; stripes++ {
		context.Rect(0, stripe_width*stripes, fly_length, stripe_width)
		context.Fill()
		stripes = stripes + 1.00
	}

	context.SetSourceRGB(0, 0, 150)
	context.Rect(0, 0, union_length, union_width)
	context.Fill()
}

func drawStar(context *cg.Context, xCenter, yCenter, nPoints, outerRadius, innerRadius float64) {
	//CGContextBeginPath
	context.BeginPath()
	for ixVertex := 0.00; ixVertex <= 2.00*nPoints; ixVertex++ {
		var angle = ixVertex*math.Pi/nPoints - math.Pi/2
		var radius float64

		if math.Mod(ixVertex, 2.00) == 0.00 {
			radius = outerRadius
		} else {
			radius = innerRadius
		}
		context.LineTo(xCenter+radius*math.Cos(angle), yCenter+radius*math.Sin(angle))
	}
	context.ClosePath()
	context.Fill()
}
