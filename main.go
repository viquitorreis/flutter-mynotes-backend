package main

func main() {
	AllTasks()
	server := NewApíServer(":3030")
	server.Run()
}
