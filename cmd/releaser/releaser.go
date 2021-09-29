package releaser

import (
	"github.com/Nayls/releaser/cmd/completion"
	"github.com/Nayls/releaser/cmd/generate"
	"github.com/Nayls/releaser/internal/cli"
	"github.com/spf13/cobra"
)

func InitCobraConfig() *cobra.Command {
	rootCmd := cli.GetRootCmd()

	// Add generate command
	rootCmd.AddCommand(generate.GetGenerateCmd())

	// Add completion command
	rootCmd.AddCommand(completion.GetCompletionCmd())

	return rootCmd
}

func init() {}
