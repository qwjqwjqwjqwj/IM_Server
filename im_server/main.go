package main

import (
	"im_server/server"
)

func main() {
	s1 := server.NewServer("127.0.0.1", 8888)
	s1.Start()
}
