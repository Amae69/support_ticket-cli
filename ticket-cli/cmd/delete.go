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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [ticket-id]",
	Short: "Delete a ticket by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		if err := storage.DeleteTicket(id); err != nil {
			fmt.Println("Error deleting ticket:", err)
			os.Exit(1)
		}
		fmt.Println("Ticket deleted successfully")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
