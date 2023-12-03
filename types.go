package main

type Tasks struct {
	ID         int    `json:"id"`
	TaskName   string `json:"taskName"`
	TaskDetail string `json:"taskDetail"`
	Date       string `json:"date"`
}
