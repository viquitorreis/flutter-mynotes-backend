package main

func main() {
	AllTasks()
	server := NewApíServer("0.0.0.0:3030")
	server.Run()
}
