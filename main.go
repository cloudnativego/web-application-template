package main

import "os"

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8101"
	}

	server := NewServer()
	server.Run(":" + port)
}
