# Task CLI 📝

[![Go Report Card](https://goreportcard.com/badge/github.com/rjc0des/todo-task-cli)](https://goreportcard.com/report/github.com/rjc0des/todo-task-cli)

Task CLI is a lightweight, offline command-line task manager built with Go. It helps you stay organized without needing an internet connection or a complex database setup.

## Features

- **Simple & Fast:** Minimalist CLI interface.
- **Local Storage:** Tasks are stored securely in a local `tasks.json` file.
- **Smart IDs:** Auto-incrementing task IDs that are never reused.
- **Status Tracking:** Support for `todo`, `in-progress`, and `done` statuses.
- **Cross-Platform:** Runs on Linux, macOS, and Windows.

---

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
    - [Add a Task](#add-a-task)
    - [List Tasks](#list-tasks)
    - [Update Task](#update-task)
    - [Delete Task](#delete-task)
    - [Update Status](#update-status)
- [Development](#-development)
- [Project Structure](#-project-structure)

---

## Installation

### Option 1: Download Binary

1. Go to the [Releases](https://github.com/rjc0des/todo-task-cli/releases) page.
2. Download the binary for your operating system.
3. Make it executable (Linux/macOS) and run it!

### Option 2: Build from Source

Prerequisites: [Go 1.25+](https://go.dev/dl/)

```bash
# Clone the repository
git clone https://github.com/rjc0des/todo-task-cli.git
cd todo-task-cli

# Build the binary
go build -o task-cli
```

### OS Specific Instructions

#### Windows

1. **Download**: Get `task-cli.exe` from releases.
2. **Run via CMD/PowerShell**:
    ```powershell
    .\task-cli.exe list
    ```
3. **Add to PATH** (Optional):
    - Move `task-cli.exe` to a folder (e.g., `C:\Program Files\TaskCLI`).
    - Add that folder to your System PATH environment variable.
    - Now you can run `task-cli list` from anywhere!

#### macOS

1. **Download**: Get the binary from releases.
2. **Permissions**: You might need to allow execution:
    ```bash
    chmod +x task-cli
    ```
3. **Security**: If macOS blocks it ("Unidentified Developer"), Ctrl+Click the file > Open > Open.
4. **Run**:
    ```bash
    ./task-cli list
    ```

---

## Usage

Run `./task-cli` (or just `task-cli` if installed to PATH) followed by a command.

### Add a Task

Create a new task with a description.

```bash
./task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)
```

### List Tasks

View all tasks or filter by status (`todo`, `in-progress`, `done`).

```bash
# List all tasks
./task-cli list

# List only 'todo' tasks
./task-cli list todo

# List only 'done' tasks
./task-cli list done
```

### Update Task

Modify the description of an existing task.

```bash
# Syntax: update <id> <new description>
./task-cli update 1 "Buy organic groceries"
```

### Delete Task

Remove a task permanently.

```bash
# Syntax: delete <id>
./task-cli delete 1
```

### Update Status

Change the status of a task using specific commands.

**Mark as Todo:**

```bash
./task-cli mark-todo 1
```

**Mark as In Progress:**

```bash
./task-cli mark-in-progress 1
```

**Mark as Done:**

```bash
./task-cli mark-done 1
```

---

## Development

### Running Tests

Run the standard Go test suite:

```bash
go test ./...
```

### Running Locally

You can run the application directly without building:

```bash
go run main.go list
```

---

## Project Structure

The project follows a standard Go CLI layout:

- **`cmd/`**: Contains the command logic (add, list, update, delete, etc.).
- **`internal/`**: (If applicable) Internal business logic and data models.
- **`main.go`**: The entry point of the application.
- **`tasks.json`**: The local JSON file where tasks are persisted (created automatically).

---

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feat/amazing-feature`)
3. Commit your Changes (`git commit -m 'feat: Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feat/amazing-feature`)
5. Open a Pull Request
