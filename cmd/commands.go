package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go_blog_api",
	Short: "simple blog api powered by Go",
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start api server",
	// Long:    "...",
	// Example: "...",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
