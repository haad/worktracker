package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ServerPort int64
var ServerHost string

func init() {
	rootCmd.AddCommand(ServerCmd)
	ServerCmd.AddCommand(ProjCreateCmd)

	ServerCmd.Flags().StringVarP(&ServerHost, "host", "H", "localhost", "Server host")
	ServerCmd.Flags().Int64VarP(&ServerPort, "port", "p", 8080, "Server port")
}

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start worktracker webui",
	Long:  `Show worktracker webui to better see stats, project info ...`,
	Run: func(cmd *cobra.Command, args []string) {
		ServerRun(ServerHost, ServerPort)
	},
}

func ServerRun(host string, port int64) {
	fmt.Println("Running server on:", host, "while listening on port:", port)
}
