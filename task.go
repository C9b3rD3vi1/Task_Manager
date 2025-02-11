
package main


import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "time"
)

type Task struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Completed bool      `json:"completed"`
    CreatedAt time.Time `json:"created_at"`
}



type TaskManager struct {
    Tasks []Task `json:"tasks"`
}


func (tm *TaskManager) AddTask(title string) {
    newTask := Task{
        ID:        len(tm.Tasks) + 1,
        Title:     title,
        Completed: false,
        CreatedAt: time.Now(),
    }
    tm.Tasks = append(tm.Tasks, newTask)
    fmt.Printf("Task added: %s\n", title)
}

func (tm *TaskManager) ListTasks() {
    if len(tm.Tasks) == 0 {
        fmt.Println("No tasks available.")
        return
    }
    for _, task := range tm.Tasks {
        status := "Pending"
        if task.Completed {
            status = "Completed"
        }
        fmt.Printf("%d: %s [%s]\n", task.ID, task.Title, status)
    }
}

func (tm *TaskManager) MarkTaskDone(id int) {
    for i, task := range tm.Tasks {
        if task.ID == id {
            tm.Tasks[i].Completed = true
            fmt.Printf("Task %d marked as done.\n", id)
            return
        }
    }
    fmt.Printf("Task %d not found.\n", id)
}

func (tm *TaskManager) RemoveTask(id int) {
    for i, task := range tm.Tasks {
        if task.ID == id {
            tm.Tasks = append(tm.Tasks[:i], tm.Tasks[i+1:]...)
            fmt.Printf("Task %d removed.\n", id)
            return
        }
    }
    fmt.Printf("Task %d not found.\n", id)
}

func (tm *TaskManager) SaveToFile(filename string) error {
    data, err := json.MarshalIndent(tm, "", "  ")
    if err != nil {
        return err
    }
    return ioutil.WriteFile(filename, data, 0644)
}

func (tm *TaskManager) LoadFromFile(filename string) error {
    file, err := ioutil.ReadFile(filename)
    if err != nil {
        if os.IsNotExist(err) {
            return nil // No saved tasks, start fresh
        }
        return err
    }
    return json.Unmarshal(file, tm)
}
