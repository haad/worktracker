package cmd

import (
	"fmt"

	"github.com/haad/worktracker/web"
	"github.com/spf13/cobra"
)

func init() {
	var serverAddr string
	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start worktracker webui",
		Long:  `Show worktracker webui to better see stats, project info ...`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running server on:", serverAddr)
			web.StartServer(serverAddr)
		},
	}
	serverCmd.Flags().StringVarP(&serverAddr, "addr", "A", "localhost:8080", "Server host and port")

	rootCmd.AddCommand(serverCmd)
}
