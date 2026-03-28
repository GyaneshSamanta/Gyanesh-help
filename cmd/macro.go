package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/GyaneshSamanta/gyanesh-help/internal/macro"
	"github.com/GyaneshSamanta/gyanesh-help/internal/ui"
)

var macroCmd = &cobra.Command{
	Use:   "macro",
	Short: "Manage macros (add, list, remove)",
}

var macroListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available macros",
	Run: func(cmd *cobra.Command, args []string) {
		macro.ListAll()
	},
}

var macroAddCmd = &cobra.Command{
	Use:   "add [name] [command] [explanation]",
	Short: "Add a custom macro",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Persist to macros.toml
		ui.PrintSuccess(fmt.Sprintf("Macro '%s' added.", args[0]))
	},
}

var explainCmd = &cobra.Command{
	Use:   "explain [macro-name]",
	Short: "Explain a macro without executing it",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		listFlag, _ := cmd.Flags().GetBool("list")
		if listFlag || len(args) == 0 {
			macro.ListAll()
			return
		}
		if err := macro.Explain(args[0]); err != nil {
			ui.PrintError(err.Error())
		}
	},
}

func init() {
	explainCmd.Flags().Bool("list", false, "List all macros with descriptions")
	macroCmd.AddCommand(macroListCmd, macroAddCmd)

	// Register each built-in macro as a root-level command for direct invocation
	for name, m := range macro.Registry {
		macroName := name
		macroMeta := m
		macroExecCmd := &cobra.Command{
			Use:   macroName,
			Short: macroMeta.Description,
			RunE: func(cmd *cobra.Command, args []string) error {
				flags := map[string]string{}
				for _, f := range macroMeta.Flags {
					val, _ := cmd.Flags().GetString(f.Name)
					if val != "" {
						flags[f.Name] = val
					}
				}
				// Check --hard flag specifically
				hard, _ := cmd.Flags().GetBool("hard")
				if hard {
					flags["hard"] = "true"
				}
				return macro.Execute(macroName, flags)
			},
		}
		for _, f := range macroMeta.Flags {
			if f.Name == "hard" {
				macroExecCmd.Flags().Bool(f.Name, false, f.Description)
			} else {
				macroExecCmd.Flags().String(f.Name, f.Default, f.Description)
			}
		}
		rootCmd.AddCommand(macroExecCmd)
	}
}
