package releaser_test

import (
	"testing"

	"github.com/Nayls/releaser/cmd/completion"
	"github.com/Nayls/releaser/cmd/generate"
	"github.com/Nayls/releaser/internal/cli"
	"github.com/Nayls/releaser/internal/test"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func buildTestCmd() *cobra.Command {
	cmd := cli.GetRootCmd()

	// Add generate command
	cmd.AddCommand(generate.GetGenerateCmd())

	// Add completion command
	cmd.AddCommand(completion.GetCompletionCmd())

	return cmd
}

func Test_SingleCommandWithoutSubcommand(t *testing.T) {
	out, err := test.ExecuteCommand(buildTestCmd(), "")
	if err != nil {
		assert.Error(t, err)
	}

	assert.NotContains(t, out, `Error:`)

	assert.Contains(t, out, `Application for releaser another host`)

	assert.Contains(t, out, `Usage:`)
	assert.Contains(t, out, `Available Commands:`)
	assert.Contains(t, out, `Flags:`)
}

func Test_NegativeCallSingleCommandWithSubcommand(t *testing.T) {
	out, err := test.ExecuteCommand(buildTestCmd(), "fail")
	if err != nil {
		assert.Error(t, err)
	}

	assert.Contains(t, out, `Error: unknown command "fail" for "releaser"`)

	assert.NotContains(t, out, `Usage:`)
	assert.NotContains(t, out, `Available Commands:`)
	assert.NotContains(t, out, `Flags:`)

	assert.Contains(t, out, `Run 'releaser --help' for usage.`)
}

func Test_NegativeCallSingleCommandWithLongFailFlag(t *testing.T) {
	out, err := test.ExecuteCommand(buildTestCmd(), "--fail")
	if err != nil {
		assert.Error(t, err)
	}

	assert.Contains(t, out, `Error: unknown flag: --fail`)

	assert.Contains(t, out, `Usage:`)
	assert.Contains(t, out, `Available Commands:`)
	assert.Contains(t, out, `Flags:`)
}
