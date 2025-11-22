## **Building a Simple Ticket Tracker CLI in Go**

### A lightweight, command-line alternative to complex ticketing systems.

Using golang and cobra-cli, I built a simple command-line interface for managing support tickets. Tickets are stored locally in a CSV file.

---

## Introduction

As developers, we often find ourselves juggling multiple tasks, bugs, and feature requests. While tools like Jira, Trello, or GitHub Issues are powerful, sometimes you just need something **simple**, **fast**, and **local** to track your daily work without leaving the terminal.

That's why I built **Ticket CLI**—a simple command-line tool written in Go to track daily tickets and store them in a CSV file. No servers, no databases, just a binary and a text file.

## The Tech Stack

For this project, I choose:
- **[Go](https://go.dev/)**: For its speed, simplicity, and ability to compile into a single binary.
- **[Cobra](https://github.com/spf13/cobra)**: The industry standard for building modern CLI applications in Go. It handles flag parsing, subcommands, and help text generation effortlessly.
- **Standard Library (`encoding/csv`)**: To keep dependencies low, I used Go's built-in CSV support for data persistence.

## How It Works

The project follows a standard Go CLI structure:

```
ticket-cli/
├── cmd/            # Cobra commands (add, list, delete)
├── internal/       # Business logic
│   └── storage/    # CSV handling
└── main.go         # Entry point
```

### 1. The Command Structure

Using Cobra, I defined commands like `add`, `list`, and `delete`. Here's a snippet of how the `add` command handles flags to create a new ticket:

```go
// cmd/add.go
var addCmd = &cobra.Command{
    Use:   "add",
    Short: "Add a new ticket",
    Run: func(cmd *cobra.Command, args []string) {
        // default values logic...

        t := storage.Ticket{
            ID:          fmt.Sprintf("%d", time.Now().UnixNano()),
            Title:       flagTitle,
            Customer:    flagCustomer,
            Priority:    flagPriority,
            Status:      flagStatus,
            Description: flagDescription,
        }

        if err := storage.AppendTicket(t); err != nil {
            fmt.Println("Error saving ticket:", err)
            return
        }
        fmt.Println("Ticket saved with ID:", t.ID)
    },
}
```		
### 2. Data Persistence (The "Database")

Instead of setting up SQLite or a JSON store, I opted for CSV. It's human-readable and easy to debug. The `internal/storage` package handles reading and writing to `tickets.csv`.

```go
// internal/storage/storage.go
func AppendTicket(t Ticket) error {
    // ... (file opening logic)
    w := csv.NewWriter(f)
    defer w.Flush()

    rec := []string{t.ID, t.Date, t.Title, t.Customer, t.Priority, t.Status, t.Description}
    return w.Write(rec)
}
```

## Installation & Usage

You can clone the repo and build it yourself:

```bash
git clone https://github.com/yourusername/ticket-cli
cd ticket-cli
go mod tidy
go build -o ticket-cli .
```

### Adding a Ticket
```bash
./ticket-cli add --title "Fix login bug" --priority high --customer "Acme Corp"
```

### Listing Tickets
```bash
./ticket-cli list
```
### Filter by date
```bash
./ticket-cli list --date 2025-11-15
```
### CSV Storage
A CSV file is created automatically at the Project directory, and it keeps getting updated, once a new ticket is added using the `.ticket-cli add --flags`

### Columns :

```
ID, Date, Title, Customer, Priority, Status, Description
```


## Future Improvements

This is just an MVP. Some ideas for the future include:
- **JSON/SQLite Storage**: For more complex querying.
- **TUI (Text User Interface)**: Using `bubbletea` for an interactive dashboard.
- **Cloud Sync**: Syncing tickets to a Gist or S3 bucket.

## Conclusion

Building CLI tools in Go is a rewarding experience. Cobra makes the interface professional, and Go's standard library handles the rest. If you're looking for a weekend project, try building your own developer tools!

---
