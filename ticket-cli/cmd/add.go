/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/Amae69/ticket-cli/internal/storage"
)

var (
	flagTitle       string
	flagCustomer    string
	flagPriority    string
	flagStatus      string
	flagDescription string
	flagDate        string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new ticket",
	Run: func(cmd *cobra.Command, args []string) {
		// default values
		if flagDate == "" {
			flagDate = time.Now().Format("2006-01-02")
		}
		if flagPriority == "" {
			flagPriority = "medium"
		}
		if flagStatus == "" {
			flagStatus = "new"
		}

		id := fmt.Sprintf("%d", time.Now().UnixNano())

		t := storage.Ticket{
			ID:          id,
			Date:        flagDate,
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

		fmt.Println("Ticket saved with ID:", id)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVar(&flagTitle, "title", "", "Ticket title (required)")
	addCmd.Flags().StringVar(&flagCustomer, "customer", "", "Customer or reporter")
	addCmd.Flags().StringVar(&flagPriority, "priority", "", "Priority: low|medium|high")
	addCmd.Flags().StringVar(&flagStatus, "status", "", "Status: new|open|closed")
	addCmd.Flags().StringVar(&flagDescription, "description", "", "Longer description")
	addCmd.Flags().StringVar(&flagDate, "date", "", "Ticket date YYYY-MM-DD (defaults to today)")

	addCmd.MarkFlagRequired("title")
}
