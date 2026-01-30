# Task CLI 📝

Task CLI is a lightweight command-line task manager built with Go.

It supports:

- Adding tasks
- Listing tasks (with filters)
- Marking tasks as done
- Editing task text
- Deleting tasks
- JSON persistence with auto-increment IDs
- createdAt and updatedAt timestamps

---

## Features

- Simple CLI interface
- Tasks stored locally in tasks.json
- Auto-increment task IDs (never reused)
- Status support: todo, doing, done
- Cross-platform builds (Linux, macOS, Windows)
- GitHub Releases automation

---

## Installation

### Option 1: Download from GitHub Releases

1. Go to the Releases page of this repository
2. Download the binary for your OS
3. Run it from terminal

Example (Linux/macOS):

```bash
chmod +x task-cli
./task-cli list
```

---

### Option 2: Build from Source

Clone the repo:

```bash
git clone https://github.com/rjc0des/todo-task-cli.git
cd task-cli
```

Build:

```bash
go build -o task-cli
```

Run:

```bash
./task-cli list
```

---

## ⚡ Usage

Add a task:

```bash
task-cli add "Learn Go CLI"
```

List all tasks:

```bash
task-cli list
```

List only TODO tasks:

```bash
task-cli list todo
```

Mark task as done:

```bash
task-cli done 1
```

Edit task text:

```bash
task-cli update 1 "Learn Go CLI properly"
```

Delete a task:

```bash
task-cli delete 1
```

---

## Data Storage

All tasks are stored locally in tasks.json.

Example structure:

```json
{
	"lastId": 2,
	"tasks": [
		{
			"id": 1,
			"text": "Learn Go",
			"status": "todo",
			"createdAt": "...",
			"updatedAt": "..."
		}
	]
}
```
