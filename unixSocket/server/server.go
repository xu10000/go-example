package main

func main() {
	filename := "/tmp/us.socket"
	u := NewUnixSocket(filename)
	u.startServer()
}
