/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Amae69/ticket-cli/internal/storage"
	"github.com/spf13/cobra"
)

var (
	flagListDate string
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tickets (today by default)",
	Run: func(cmd *cobra.Command, args []string) {
		tickets, err := storage.ReadTickets()
		if err != nil {
			fmt.Println("Error reading tickets:", err)
			os.Exit(1)
		}
		if flagListDate != "" {
			// filter by date
			var filtered []storage.Ticket
			for _, t := range tickets {
				if t.Date == flagListDate {
					filtered = append(filtered, t)
				}
			}
			tickets = filtered
		}

		if len(tickets) == 0 {
			fmt.Println("No tickets found")
			return
		}

		// Print a simple table
		fmt.Printf("% -20s % -10s % -30s % -10s % -8s\n", "ID", "Date", "Title", "Customer", "Priority")
		for _, t := range tickets {
			fmt.Printf("% -20s % -10s % -30s % -10s % -8s\n", t.ID, t.Date, t.Title, t.Customer, t.Priority)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVar(&flagListDate, "date", "", "Filter tickets by date YYYY-MM-DD")
}
