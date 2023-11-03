package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Models
type Task struct {
	Todo      string
	Username  string
	Timestamp time.Time
	Done      bool
}

type TaskNode struct {
	NodeId   int       `json:"id"`
	Task     Task      `json:"task"`
	NextTask *TaskNode `json:"next"`
}

type Body struct {
	Content map[string]string
}

// Functions
func CreateNewTask(todo string, username string) Task {
	var newtask Task = Task{
		Todo:      todo,
		Username:  username,
		Timestamp: time.Now(),
		Done:      false,
	}

	return newtask
}

func AddTaskToList(current *TaskNode, task Task) {
	if current.NextTask != nil {
		AddTaskToList(current.NextTask, task)
	} else {
		current.NextTask = &TaskNode{
			NodeId:   current.NodeId + 1,
			Task:     task,
			NextTask: nil,
		}
	}
}

func RemoveTaskFromList(current *TaskNode, taskid int) {
	// n1 -> n2 -> nil
	// n1-> nil
	if current.NextTask.NodeId == taskid {
		current.NextTask = current.NextTask.NextTask
	}else{
		RemoveTaskFromList(current.NextTask,taskid)
	}
}

func MapBody(body io.ReadCloser)Body{
	var bodymap Body = Body{Content: make(map[string]string)}

	bodyraw,err := io.ReadAll(body)

	if err != nil{
		panic(err)
	}

	var bodycleaned string = strings.ReplaceAll(string(bodyraw),"%20"," ")
	var bodyparsed []string = strings.Split(bodycleaned,"&")

	for i:=0;i<len(bodyparsed);i++ {
		var kv []string = strings.Split(bodyparsed[i],"=")
		bodymap.Content[kv[0]] = kv[1]
	}

	return bodymap
}

// Main
func main() {
	var port string = ":8080"
	var mux *http.ServeMux = http.NewServeMux()
	var head *TaskNode = &TaskNode{NodeId: 0}

	// ADD TASK
	mux.HandleFunc("/task/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var body Body = MapBody(r.Body)
			var newtask Task = CreateNewTask(body.Content["todo"],body.Content["username"])
			AddTaskToList(head, newtask)
			w.WriteHeader(http.StatusOK)
		w.Write([]byte("New task added."))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 page not found."))
		}
	})

	// DELETE TASK
	mux.HandleFunc("/task/remove", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			var body Body = MapBody(r.Body)
			taskid,err := strconv.Atoi(body.Content["id"])

			if err != nil {
				panic(err)
			}

			RemoveTaskFromList(head, taskid)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Task removed."))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 page not found."))
		}
	})

	// GET ALL TASKS
	mux.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			r, err := json.Marshal(head.NextTask)

			if err != nil {
				panic(err)
			}

			w.Header().Add("Content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(r)
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 page not found."))
		}
	})

	fmt.Println("Server up at ",port)
	http.ListenAndServe(port, mux)
}
