package coaster

import (
	"fmt"
	"syscall/js"
)

type AppProps struct {
	RootComponent *Component
	DistDir       string
}

func RenderApp(rootComponent *Component) {
	fmt.Println("WASM Go Initialized")
	start(rootComponent)
	<-make(chan bool)
}

func start(rootComponent *Component) {
	jsDoc := js.Global().Get("document")
	rootRender, err := rootComponent.Render()
	if err != nil {
		fmt.Println(err)
	}
	jsDoc.Call("getElementById", "root").Set("innerHTML", *rootRender)
}
