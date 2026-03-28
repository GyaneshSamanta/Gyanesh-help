package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/GyaneshSamanta/gyanesh-help/internal/store"
	"github.com/GyaneshSamanta/gyanesh-help/internal/ui"
)

var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Manage environment stores (install, preview, verify, remove)",
	Run: func(cmd *cobra.Command, args []string) {
		ui.PrintHeader("Available Environment Stores")
		headers := []string{"Store", "Description", "Size"}
		var rows [][]string
		for _, s := range store.ListStacks() {
			rows = append(rows, []string{
				s.Name(), s.Description(),
				formatSize(s.EstimatedSizeMB()),
			})
		}
		ui.PrintTable(headers, rows)
		ui.PrintInfo("Use 'gyanesh-help store install <name>' to set up a stack.")
	},
}

var storeInstallCmd = &cobra.Command{
	Use:   "install [stack]",
	Short: "Install an environment store",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		verify, _ := cmd.Flags().GetBool("verify")
		err := store.Install(args[0], osAdapter, store.InstallOpts{Verify: verify})
		if err != nil {
			ui.PrintError(err.Error())
		}
	},
}

var storePreviewCmd = &cobra.Command{
	Use:   "preview [stack]",
	Short: "Preview what a store will install",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := store.Preview(args[0], osAdapter); err != nil {
			ui.PrintError(err.Error())
		}
	},
}

var storeVerifyCmd = &cobra.Command{
	Use:   "verify [stack]",
	Short: "Verify installed store components",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := store.Verify(args[0], osAdapter); err != nil {
			ui.PrintError(err.Error())
		}
	},
}

var storeRemoveCmd = &cobra.Command{
	Use:   "remove [stack]",
	Short: "Remove a store and its components",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		if err := store.Remove(args[0], osAdapter, force); err != nil {
			ui.PrintError(err.Error())
		}
	},
}

func init() {
	storeInstallCmd.Flags().Bool("verify", true, "Run verification after install")
	storeRemoveCmd.Flags().Bool("force", false, "Force remove shared components")
	storeCmd.AddCommand(storeInstallCmd, storePreviewCmd, storeVerifyCmd, storeRemoveCmd)
}

func formatSize(mb int) string {
	if mb >= 1000 {
		return fmt.Sprintf("~%.1f GB", float64(mb)/1000)
	}
	return fmt.Sprintf("~%d MB", mb)
}
