package commands

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "go_blog_api",
	Short: "simple blog api powered by Go",
}

func init() {
	rootCmd.AddCommand(serveCmd, migrateCmd, createSuperuserCmd, deleteSuperuserCmd)

	migrateCmd.AddCommand(migrateUpCmd, migrateDownCmd, migrateForceCmd)

	migrateUpCmd.Flags().IntP("steps", "s", -1, "number of steps for up migration (if not set: apply all up migrations)")

	migrateDownCmd.Flags().IntP("steps", "s", -1, "number of steps for down migration (required)")
	migrateDownCmd.MarkFlagRequired("steps")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
