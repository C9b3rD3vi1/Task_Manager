# Task_Manager

A simple design and implementation for a CLI Task Manager in Go,


# CLI Task Manager

A simple command-line tool written in Go to manage daily tasks. This tool allows you to add, list, mark as done, and remove tasks, all from your terminal. Tasks are saved to a JSON file, so your list persists between sessions.

## Features

- **Add Task**: Add a new task with the `-add` flag.
- **List Tasks**: List all tasks with the `-list` flag.
- **Mark Task as Done**: Mark a task as completed with the `-done` flag followed by the task ID.
- **Remove Task**: Remove a task with the `-remove` flag followed by the task ID.
- **Save and Load**: Tasks are automatically saved to and loaded from a `tasks.json` file.

## Installation

1. **Clone the repository**:

    ```bash
    git clone https://github.com/C9b3rD3vi1/Task_Manager.git
    cd cli-task-manager
    ```

2. **Build the project**:

    ```bash
    go build -o task-manager
    ```

3. **Run the tool**:

    After building, you can run the tool using:

    ```bash
    ./task-manager -add "Buy groceries"
    ./task-manager -list
    ./task-manager -done 1
    ./task-manager -remove 1
    ```

## Usage

### Add a Task

```bash
./task-manager -add "Your task description"


List All Tasks

./task-manager -list


Mark a Task as Done

./task-manager -done [task_id]


Remove a Task

./task-manager -remove [task_id]


Help
You can get help by running:

./task-manager -h


## Example

./task-manager -add "Complete Go project"
./task-manager -add "Read a book"
./task-manager -list
./task-manager -done 1
./task-manager -remove 2
./task-manager -list


## Contribution
Contributions are welcome! If you have suggestions, ideas, or want to report a bug, please open an issue or submit a pull request.

## License
This project is licensed under the MIT !License. See the LICENSE file for details.