package main

func main() {
	println("hello wasm")
}

// GOOS=js GOARCH=wasm go build -o main.wasm
// GOOS=js GOARCH=wasm go build -o main.js
