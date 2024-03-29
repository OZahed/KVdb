package cmd

import (
	"KVdb/api/v1"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts the server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listen and serve On port: 8080")
		mux := api.GetMux()
		log.Fatal(http.ListenAndServe(":8080", mux))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
