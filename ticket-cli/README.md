# Ticket CLI

A simple command-line interface for managing support tickets, written in Go. Tickets are stored locally in a CSV file.

## Installation

1.  Clone the repository:
    ```bash
    git clone https://github.com/Amae69/ticket-cli.git
    cd ticket-cli
    ```
2.  Build the application:
    ```bash
    go build -o ticket-cli
    ```

## Usage

### Add a Ticket

Add a new ticket with a title, priority, and other optional details.

```bash
./ticket-cli add --title "Login issue" --priority high --customer "Acme Corp"
```

**Flags:**
-   `--title`: Ticket title (required)
-   `--priority`: Priority (low, medium, high)
-   `--customer`: Customer name
-   `--status`: Status (new, open, closed)
-   `--description`: Detailed description
-   `--date`: Date (YYYY-MM-DD), defaults to today

### List Tickets

List all tickets, optionally filtered by date.

```bash
./ticket-cli list
./ticket-cli list --date 2025-10-27
```

### Delete a Ticket

Delete a ticket by its ID.

```bash
./ticket-cli delete <ticket-id>
```
