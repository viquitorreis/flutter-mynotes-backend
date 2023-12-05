package main

import "log"

func main() {
	// AllTasks()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewApíServer("0.0.0.0:3030", store)
	server.Run()
}
