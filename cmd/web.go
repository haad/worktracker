package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	var serverPort int64
	var serverHost string
	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start worktracker webui",
		Long:  `Show worktracker webui to better see stats, project info ...`,
		Run: func(cmd *cobra.Command, args []string) {
			ServerRun(serverHost, serverPort)
		},
	}
	serverCmd.Flags().StringVarP(&serverHost, "host", "H", "localhost", "Server host")
	serverCmd.Flags().Int64VarP(&serverPort, "port", "p", 8080, "Server port")

	rootCmd.AddCommand(serverCmd)
}

func ServerRun(host string, port int64) {
	fmt.Println("Running server on:", host, "while listening on port:", port)
}
