//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"
	"math"
)

type Point struct {
	X, Y float64
}

type Triangle struct {
	P1, P2, P3 Point
}

func getRandomPointInTriangle(x1, y1, x2, y2, x3, y3 float64) (float64, float64) {
	r1 := js.Global().Get("Math").Call("random").Float()
	r2 := js.Global().Get("Math").Call("random").Float()

	// Ensure uniform distribution within the triangle
	if r1 + r2 > 1 {
		r1, r2 = 1 - r1, 1 - r2
	}

	x := x1 + (x2 - x1)*r1 + (x3 - x1)*r2
	y := y1 + (y2 - y1)*r1 + (y3 - y1)*r2
	return x, y
}

func drawRecursiveTriangle(this js.Value, args []js.Value) interface{} {
	// Get canvas context from JavaScript
	ctx := args[0]
	x1, y1 := args[1].Float(), args[2].Float()
	x2, y2 := args[3].Float(), args[4].Float()
	x3, y3 := args[5].Float(), args[6].Float()
	depth := args[7].Int()

	// Draw the current triangle
	ctx.Call("beginPath")
	ctx.Call("moveTo", x1, y1)
	ctx.Call("lineTo", x2, y2)
	ctx.Call("lineTo", x3, y3)
	ctx.Call("closePath")
	ctx.Call("stroke")

	if depth == 0 {
		// Draw random lines
		for i := 0; i < 15; i++ {
			// Get two random points in the triangle
			xa, ya := getRandomPointInTriangle(x1, y1, x2, y2, x3, y3)
			xb, yb := getRandomPointInTriangle(x1, y1, x2, y2, x3, y3)

			ctx.Call("beginPath")
			ctx.Call("moveTo", xa, ya)
			ctx.Call("lineTo", xb, yb)
			ctx.Call("stroke")
		}

		// Draw random circles
		for i := 0; i < 10; i++ {
			cx, cy := getRandomPointInTriangle(x1, y1, x2, y2, x3, y3)
			radius := js.Global().Get("Math").Call("random").Float()*5 + 1 // Radius between 1-6

			ctx.Call("beginPath")
			ctx.Call("arc", cx, cy, radius, 0, 2*math.Pi)
			ctx.Call("stroke")
		}
	}

	if depth > 0 {
		// Calculate midpoints
		mx1, my1 := (x1+x2)/2, (y1+y2)/2
		mx2, my2 := (x2+x3)/2, (y2+y3)/2
		mx3, my3 := (x3+x1)/2, (y3+y1)/2

		// Recursively draw smaller triangles
		drawRecursiveTriangle(this, []js.Value{
			ctx,
			js.ValueOf(x1), js.ValueOf(y1),
			js.ValueOf(mx1), js.ValueOf(my1),
			js.ValueOf(mx3), js.ValueOf(my3),
			js.ValueOf(depth - 1),
		})
		drawRecursiveTriangle(this, []js.Value{
			ctx,
			js.ValueOf(mx1), js.ValueOf(my1),
			js.ValueOf(x2), js.ValueOf(y2),
			js.ValueOf(mx2), js.ValueOf(my2),
			js.ValueOf(depth - 1),
		})
		drawRecursiveTriangle(this, []js.Value{
			ctx,
			js.ValueOf(mx3), js.ValueOf(my3),
			js.ValueOf(mx2), js.ValueOf(my2),
			js.ValueOf(x3), js.ValueOf(y3),
			js.ValueOf(depth - 1),
		})
	}

	return nil
}

func calculateTrianglePoints(this js.Value, args []js.Value) interface{} {
	centerX := args[0].Float()
	centerY := args[1].Float()
	radius := args[2].Float()
	rotation := args[3].Float()

	angle1 := rotation + math.Pi/2
	angle2 := rotation + (math.Pi*7)/6
	angle3 := rotation + (math.Pi*11)/6

	points := make([]interface{}, 6)
	points[0] = centerX + radius*math.Cos(angle1)
	points[1] = centerY + radius*math.Sin(angle1)
	points[2] = centerX + radius*math.Cos(angle2)
	points[3] = centerY + radius*math.Sin(angle2)
	points[4] = centerX + radius*math.Cos(angle3)
	points[5] = centerY + radius*math.Sin(angle3)

	return js.ValueOf(points)
}

func main() {
	c := make(chan struct{}, 0)
	
	js.Global().Set("drawRecursiveTriangle", js.FuncOf(drawRecursiveTriangle))
	js.Global().Set("calculateTrianglePoints", js.FuncOf(calculateTrianglePoints))
	
	<-c
}