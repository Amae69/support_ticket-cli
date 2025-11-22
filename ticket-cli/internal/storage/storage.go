package storage

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const fileName = "tickets.csv"

// Ticket is a single ticket record
type Ticket struct {
	ID          string
	Date        string
	Title       string
	Customer    string
	Priority    string
	Status      string
	Description string
}

// getDataFile returns the path for the CSV storage file, which will be in our current directory.
func getDataFile() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(cwd, fileName), nil
}

// ensureHeader ensures the CSV file exists and has a header row
func ensureHeader(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
		w := csv.NewWriter(f)
		defer w.Flush()
		if err := w.Write([]string{"ID", "Date", "Title", "Customer", "Priority", "Status", "Description"}); err != nil {
			return err
		}
	}
	return nil
}

// AppendTicket appends a ticket to the CSV file
func AppendTicket(t Ticket) error {
	path, err := getDataFile()
	if err != nil {
		return err
	}
	if err := ensureHeader(path); err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	rec := []string{t.ID, t.Date, t.Title, t.Customer, t.Priority, t.Status, t.Description}
	if err := w.Write(rec); err != nil {
		return err
	}
	return nil
}

// ReadTickets reads all tickets from CSV and returns them as []Ticket
func ReadTickets() ([]Ticket, error) {
	path, err := getDataFile()
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("no ticket file found at %s (no tickets yet)", path)
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(rows) <= 1 {
		return nil, nil // only header
	}
	var out []Ticket
	for i, row := range rows {
		if i == 0 {
			continue // header
		}
		if len(row) < 7 {
			// skip malformed
			continue
		}
		out = append(out, Ticket{
			ID:          row[0],
			Date:        row[1],
			Title:       row[2],
			Customer:    row[3],
			Priority:    row[4],
			Status:      row[5],
			Description: row[6],
		})
	}
	return out, nil
}

// DeleteTicket deletes a ticket by ID
func DeleteTicket(id string) error {
	tickets, err := ReadTickets()
	if err != nil {
		return err
	}

	var kept []Ticket
	found := false
	for _, t := range tickets {
		if t.ID == id {
			found = true
			continue
		}
		kept = append(kept, t)
	}

	if !found {
		return fmt.Errorf("ticket with ID %s not found", id)
	}

	// Rewrite file
	path, err := getDataFile()
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	// Header
	if err := w.Write([]string{"ID", "Date", "Title", "Customer", "Priority", "Status", "Description"}); err != nil {
		return err
	}

	for _, t := range kept {
		rec := []string{t.ID, t.Date, t.Title, t.Customer, t.Priority, t.Status, t.Description}
		if err := w.Write(rec); err != nil {
			return err
		}
	}

	return nil
}
