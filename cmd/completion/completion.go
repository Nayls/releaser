package completion

import (
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:
Bash:

	$ source <(releaser completion bash)
	
	# To load completions for each session, execute once:
	# Linux:
	$ releaser completion bash > /etc/bash_completion.d/releaser
	# macOS:
	$ releaser completion bash > /usr/local/etc/bash_completion.d/releaser
	
Zsh:

	# If shell completion is not already enabled in your environment,
	# you will need to enable it. You can execute the following once:
	
	$ echo "autoload -U compinit; compinit" >> ~/.zshrc
	
	# To load completions for each session, execute once:
	$ releaser completion zsh > "${fpath[1]}/_releaser
	
	# You will need to start a new shell for this setup to take effect.
	
Fish:

	$ releaser completion fish | source
	
	# To load completions for each session, execute once:
	$ releaser completion fish > ~/.config/fish/completions/releaser.fish
	
PowerShell:

	PS> releaser completion powershell | Out-String | Invoke-Expression
	
	# To load completions for every new session, run:
	PS> releaser completion powershell > releaser.ps1
	# and source this file from your Powershell profile.
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}

func GetCompletionCmd() *cobra.Command {
	return completionCmd
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {}
