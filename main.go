package main

import (
	s "./src"
)

func main() {

	go s.AiCheckerInit()
	s.ServerInit()
}
