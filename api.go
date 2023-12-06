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

func NewApíServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/tasks", MakeHttpHandlerHelper(s.handleGetTasks))
	router.HandleFunc("/task/{id}", MakeHttpHandlerHelper(s.handleGetTask))
	router.HandleFunc("/create", MakeHttpHandlerHelper(s.handleCreateTask))
	router.HandleFunc("/delete/{id}", MakeHttpHandlerHelper(s.handleDeleteTask))
	router.HandleFunc("/update/{id}", MakeHttpHandlerHelper(s.updateTask))

	log.Println("Escutando API JSON na porta:", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint da home page")
}

func (s *APIServer) handleGetTasks(w http.ResponseWriter, r *http.Request) error {

	if r.Method != "GET" {
		return fmt.Errorf("Método de request não permitido")
	}

	tasks, err := s.store.GetTasks()
	if err != nil {
		return fmt.Errorf("Erro ao obter as tarefas %s", err)
	}

	err = WriteJsonHelper(w, http.StatusOK, tasks)
	if err != nil {
		return fmt.Errorf("erro passando o JSON %s", err)
	}

	return nil
}

func (s *APIServer) handleGetTask(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Endpoint da getTask (apenas uma task)")

	if r.Method != "GET" {
		return fmt.Errorf("Método de request não permitido")
	}

	taskId := mux.Vars(r)
	task, err := s.store.GetTask(taskId["id"])
	if err != nil {
		return err
	}

	err = WriteJsonHelper(w, http.StatusOK, task)
	if err != nil {
		return fmt.Errorf("erro passando o JSON %s", err)
	}

	return nil
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
		errMsg := fmt.Errorf("Não foi fornecido todos os campos %s: ", err)
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

func (s *APIServer) updateTask(w http.ResponseWriter, r *http.Request) error {

	if r.Method != "PUT" {
		return fmt.Errorf("Método de request não permitido")
	}

	taskId := mux.Vars(r)["id"]
	var updatedTask *Task
	err := json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		return err
	}

	afterUpdateTask, err := s.store.UpdateTask(taskId, updatedTask)
	if err != nil {
		return err
	}

	return WriteJsonHelper(w, http.StatusOK, afterUpdateTask)
}

func (s *APIServer) handleDeleteTask(w http.ResponseWriter, r *http.Request) error {

	if r.Method != "DELETE" {
		return fmt.Errorf("Método de request não permitido")
	}

	taskId := mux.Vars(r)["id"]
	err := s.store.DeleteTask(taskId)
	if err != nil {
		return err
	}

	return WriteJsonHelper(w, http.StatusOK, map[string]string{"Task deletada:": taskId})
}
