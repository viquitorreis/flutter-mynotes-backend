package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
}

var tasks []Tasks

func NewApíServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func AllTasks() {
	locationTime, err := GetBrazilCurrentTimeHelper()
	if err != nil {
		fmt.Printf("Error getting time in Brazil %s", err)
	}

	task := Tasks{
		ID:         1,
		TaskName:   "Ler",
		TaskDetail: "Ler 10 páginas do livro hoje",
		Date:       locationTime,
	}

	task2 := Tasks{
		ID:         2,
		TaskName:   "Desenvolver",
		TaskDetail: "Desenvolver o app",
		Date:       locationTime,
	}

	tasks = append(tasks, task, task2)
	fmt.Println(tasks)
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/gettasks", getTasks).Methods("GET")
	router.HandleFunc("/gettask/{id}", getTask).Methods("GET")
	router.HandleFunc("/create", createTask).Methods("POST")
	router.HandleFunc("/delete/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/update/{id}", updateTask).Methods("PUT")

	log.Println("Escutando API JSON na porta:", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint da home page")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint de getTasks")

	err := WriteJsonHelper(w, http.StatusOK, tasks)
	if err != nil {
		fmt.Println("erro passando o JSON => ", err)
		return
	}

	return
}

func getTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint da getTask (apenas uma task)")

	taskId := mux.Vars(r)
	flag := false
	for i, task := range tasks {
		if taskId["id"] == fmt.Sprint(tasks[i].ID) {
			err := WriteJsonHelper(w, http.StatusOK, task)
			if err != nil {
				fmt.Println("erro passando o JSON => ", err)
				return
			}
			flag = true
			break
		}
	}

	if flag == false {
		err := WriteJsonHelper(w, http.StatusBadRequest, map[string]string{"status": "Erro, task não encontrada"})
		if err != nil {
			fmt.Println("erro passando o JSON => ", err)
			return
		}
	}
	return
}

func createTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint da home page")
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint da home page")
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint da home page")
}
