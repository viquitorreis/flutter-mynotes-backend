package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      Storage
}

var tasks []Task

func NewApíServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

// func AllTasks() {
// 	locationTime, err := GetBrazilCurrentTimeHelper()
// 	if err != nil {
// 		fmt.Printf("Error getting time in Brazil %s", err)
// 	}

// 	task := Task{
// 		ID:         1,
// 		TaskName:   "Ler",
// 		TaskDetail: "Ler 10 páginas do livro hoje",
// 		Date:       locationTime,
// 	}

// 	task2 := Task{
// 		ID:         2,
// 		TaskName:   "Desenvolver",
// 		TaskDetail: "Desenvolver o app",
// 		Date:       locationTime,
// 	}

// 	tasks = append(tasks, task, task2)
// 	fmt.Println(tasks)
// }

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/gettasks", getTasks).Methods("GET")
	router.HandleFunc("/gettask/{id}", getTask).Methods("GET")
	router.HandleFunc("/create", MakeHttpHandlerHelper(s.handleCreateTask))
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

func (s *APIServer) handleCreateTask(w http.ResponseWriter, r *http.Request) error {

	if r.Method != "POST" {
		return fmt.Errorf("Método de request não permitido")
	}

	createTaskReq := CreateTaskReq{}
	if err := json.NewDecoder(r.Body).Decode(&createTaskReq); err != nil {
		return err
	}

	validate := validator.New()
	err := validate.Struct(createTaskReq)
	if err != nil {
		errMsg := fmt.Errorf("Not all fields were given %s: ", err)
		fmt.Println(errMsg)
		return WriteJsonHelper(w, http.StatusBadRequest, ApiError{Error: errMsg.Error()})

	}

	newTask, err := NewTask(createTaskReq.TaskName, createTaskReq.TaskDetail)
	if err != nil {
		return err
	}

	if err := s.store.CreateTask(newTask); err != nil {
		return err
	}

	return WriteJsonHelper(w, http.StatusOK, newTask)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint da home page")
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint da home page")
}
