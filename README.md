# task-tracker

> A tiny CLI task tracker in Go for adding, listing, updating, deleting and marking tasks.

**Project**: A simple command-line todo/task tracker that stores tasks in a `tasks.json` file.

**Language**: Go

**Files of interest**:

- `main.go`: CLI entrypoint and command dispatch.
- `handlers.go`: Command handler implementations.
- `os-functions.go` / `crud-functions.go`: Helpers for reading/writing `tasks.json`.
- `tasks.json`: Local JSON file used as the datastore (created automatically when needed).

**Build**:

Run from the project directory:

```powershell
go build -o task-tracker .
```

Or run directly:

```powershell
go run . <command> [args]
```

**Usage**:

```powershell
task-tracker <command> [arguments]
```

Commands supported (examples):

- `add <description>` : Add a new task
  - Example: `task-tracker add "Buy groceries"`
- `update <id> <description>` : Update a task's description
  - Example: `task-tracker update 1 "Buy milk and bread"`
- `delete <id>` : Delete a task by ID
  - Example: `task-tracker delete 1`
- `mark <id> <status>` : Update a task's status (`todo`, `in-progress`, `done`)
  - Example: `task-tracker mark 1 in-progress`
- `list [status]` : List all tasks or filter by status
  - Example: `task-tracker list`
  - Example: `task-tracker list todo`
- `version` : Print the program version

Status values supported:

- `todo`
- `in-progress`
- `done`

**Data format**:

Tasks are stored in `tasks.json` as an array of objects with these fields:

- `id` (int)
- `description` (string)
- `status` (string: `todo`, `in-progress`, `done`)
- `createdAt` (ISO timestamp)
- `updatedAt` (ISO timestamp)

Example `tasks.json` entry:

```json
{
  "id": 1,
  "description": "Buy groceries",
  "status": "todo",
  "createdAt": "2025-11-30T12:34:56Z",
  "updatedAt": "2025-11-30T12:34:56Z"
}
```

**Notes & tips**:

- Run the binary from the project folder so `tasks.json` is created next to it, or provide a path change in code if desired.
- The CLI intentionally keeps behaviour minimal and file-based for simplicity.
