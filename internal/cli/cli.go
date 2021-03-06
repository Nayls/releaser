package cli

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:     "releaser [command]",
		Short:   "Cli application",
		Long:    `A simple git-based release in a closed segment`,
		Version: "v0.1.0",
		// Run: func(cmd *cobra.Command, args []string) {
		// 	if len(args) == 0 {
		// 		cmd.Help()
		// 		os.Exit(0)
		// 	}
		// },
	}
)

func GetRootCmd() *cobra.Command {
	return rootCmd
}
