package main

func main() {
	startReporters()

	port := getPort()
	server := NewServer(port)
	server.Start()
}
