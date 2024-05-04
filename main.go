package main

import (
	"ascenda/handler"
	"ascenda/logic"
)

func main() {
	logic.PrepareDataSource()
	handler.StartHTTPServer()
}
