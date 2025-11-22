package storage

import (
	"os"
	"testing"
)

func TestAppendAndReadTickets(t *testing.T) {
	// Setup temporary file

	// Backup original fileName if necessary, but here we can just swap the variable if it were exported.
	// Since fileName is a private const, we can't change it easily without refactoring.
	// However, the `getDataFile` uses `os.Getwd()`.
	// A better approach for testing would be to make the filename configurable or run in a temp dir.
	// For now, let's try to run in a temp directory.

	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "ticket-cli-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Change working directory to temp dir so `getDataFile` uses it
	oldWd, _ := os.Getwd()
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(oldWd)

	// Test Append
	ticket := Ticket{
		ID:          "1",
		Date:        "2025-01-01",
		Title:       "Test Ticket",
		Customer:    "Test Customer",
		Priority:    "High",
		Status:      "New",
		Description: "Test Description",
	}

	if err := AppendTicket(ticket); err != nil {
		t.Fatalf("AppendTicket failed: %v", err)
	}

	// Test Read
	tickets, err := ReadTickets()
	if err != nil {
		t.Fatalf("ReadTickets failed: %v", err)
	}

	if len(tickets) != 1 {
		t.Fatalf("Expected 1 ticket, got %d", len(tickets))
	}

	if tickets[0].ID != ticket.ID {
		t.Errorf("Expected ID %s, got %s", ticket.ID, tickets[0].ID)
	}

	// Test Delete
	if err := DeleteTicket(ticket.ID); err != nil {
		t.Fatalf("DeleteTicket failed: %v", err)
	}

	tickets, err = ReadTickets()
	if err != nil {
		t.Fatalf("ReadTickets after delete failed: %v", err)
	}

	if len(tickets) != 0 {
		t.Fatalf("Expected 0 tickets after delete, got %d", len(tickets))
	}
}
