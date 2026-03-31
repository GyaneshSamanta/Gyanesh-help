package cmd

import (
	"github.com/spf13/cobra"

	"github.com/GyaneshSamanta/cue/internal/doctor"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check system health and auto-fix common issues",
	Long: `Run a comprehensive suite of health checks to ensure your 
development environment is correctly configured. Checks:
  • Core tool installation (git, node, docker, etc.)
  • Environment PATH configuration
  • SSH key security and strength
  • Git credential management
  • Shared environment store health`,
	Run: func(cmd *cobra.Command, args []string) {
		doctor.RunDoctor(osAdapter)
	},
}

var doctorFixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Attempt to auto-fix identified health issues",
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		issues := doctor.RunDoctor(osAdapter)
		doctor.Fix(osAdapter, issues, all)
	},
}

func init() {
	doctorFixCmd.Flags().Bool("all", false, "Auto-fix all issues without prompting")
	doctorCmd.AddCommand(doctorFixCmd)
}
