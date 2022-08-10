package main

import (
	"fmt"
	"syscall/js"
	"time"
)

func main() {

	done := make(chan struct{}, 0)

	window := js.Global()
	document := window.Get("document")
	body := document.Get("body")

	// create a title
	title := document.Call("createElement", "h1")
	title.Set("innerText", "This is a go wasm example!")
	body.Call("appendChild", title)

	// create a button with click event
	div := document.Call("createElement", "div")
	body.Call("appendChild", div)
	button := document.Call("createElement", "button")
	div.Call("appendChild", button)
	span := document.Call("createElement", "span")
	div.Call("appendChild", span)

	button.Set("innerText", "click me")
	button.Set("onclick", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		greet := fmt.Sprintf("Here is a button example! %s", time.Now())
		span.Set("innerText", greet)
		return nil
	}))

	// create a canvas
	canvas := document.Call("createElement", "canvas")
	canvas.Set("width", 400)
	canvas.Set("height", 100)
	body.Call("appendChild", canvas)

	ctx := canvas.Call("getContext", "2d")

	ctx.Set("fillStyle", "#cba")
	ctx.Call("fillRect", 0, 0, 400, 100)
	ctx.Set("fillStyle", "#123")
	ctx.Set("font", "30px Arial")
	ctx.Call("fillText", "Here is a canvas example", 10, 50)

	// Here must waiting for on click event call
	<-done
}
