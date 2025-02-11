package main


import (
    "flag"
    "fmt"
    "os"
)

const taskFile = "tasks.json"


func main() {
    // Command-line flags
    add := flag.String("add", "", "Add a new task")
    list := flag.Bool("list", false, "List all tasks")
    done := flag.Int("done", 0, "Mark a task as done by its ID")
    remove := flag.Int("remove", 0, "Remove a task by its ID")
    flag.Parse()

    
    // Task manager instance   
    var tm TaskManager
    err := tm.LoadFromFile(taskFile)
    if err != nil {
        fmt.Println("Error loading tasks:", err)
        os.Exit(1)
    }

    // Command-line actions
    switch {
    case *add != "":
        tm.AddTask(*add)
    case *list:
        tm.ListTasks()
    case *done > 0:
        tm.MarkTaskDone(*done)
    case *remove > 0:
        tm.RemoveTask(*remove)
    default:
        fmt.Println("No command provided. Use -h for help.")
    }

    
    // Save tasks to file
    err = tm.SaveToFile(taskFile)
    if err != nil {
        fmt.Println("Error saving tasks:", err)
        os.Exit(1)
    }
}
