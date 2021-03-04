package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionsCmd = &cobra.Command{
	Use:                   "completions [bash|zsh|fish]",
	DisableFlagsInUseLine: true,
	Short:                 "Generate shell completions",
	Long: `To load completions:

Bash:

  $ source <(hub-cli completions bash)

  # To load completions for each session, execute once:
  # Linux:
  $ hub-cli completions bash > /etc/bash_completion.d/hub-cli
  # macOS:
  $ hub-cli completions bash > /usr/local/etc/bash_completion.d/hub-cli

Zsh:

  # If shell completions is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ hub-cli completions zsh > "${fpath[1]}/_hub-cli"

  # You will need to start a new shell for this setup to take effect.

fish:

  $ hub-cli completions fish | source

  # To load completions for each session, execute once:
  $ hub-cli completions fish > ~/.config/fish/completions/hub-cli.fish
`,
	ValidArgs: []string{"bash", "zsh", "fish"},
	Args:      cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionsCmd)
}
