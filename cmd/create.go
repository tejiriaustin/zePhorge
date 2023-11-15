package cmd

import "github.com/spf13/cobra"

// apiCmd represents the api command
var createCmd = &cobra.Command{
	Use:   "create-server",
	Short: "creates a new server with the default configurations",
	Run:   create,
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func create(cmd *cobra.Command, args []string) {

}
