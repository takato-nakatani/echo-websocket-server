package main

import (
	"./route"
)

func main() {
	e := route.Init()
	e.Logger.Fatal(e.Start(":80"))
}