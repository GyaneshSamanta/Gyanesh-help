package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/GyaneshSamanta/gyanesh-help/internal/adapter"
	"github.com/GyaneshSamanta/gyanesh-help/internal/config"
	"github.com/GyaneshSamanta/gyanesh-help/internal/history"
	"github.com/GyaneshSamanta/gyanesh-help/internal/macro"
	_ "github.com/GyaneshSamanta/gyanesh-help/internal/macro/builtins"
	_ "github.com/GyaneshSamanta/gyanesh-help/internal/store/stacks"
	"github.com/GyaneshSamanta/gyanesh-help/internal/ui"
)

var (
	version   string
	buildDate string
	noColor   bool
	verbose   bool
	osAdapter adapter.OSAdapter
)

// SetVersionInfo sets the version info from ldflags.
func SetVersionInfo(v, b string) {
	version = v
	buildDate = b
	rootCmd.Version = v
}

var rootCmd = &cobra.Command{
	Use:   "gyanesh-help",
	Short: "Cross-platform CLI developer utility",
	Long: `gyanesh-help — makes the terminal feel like it already knows what you need.

  Queue management, pause/resume, semantic macros, environment stores,
  smart history, and workspace backup — all offline, all local.`,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if noColor {
			ui.SetColor(false)
		}

		// Load config
		cfg, err := config.Load()
		if err != nil {
			ui.PrintWarning(fmt.Sprintf("Config error: %v (using defaults)", err))
		}
		if cfg != nil && !cfg.UI.Color {
			ui.SetColor(false)
		}

		// Detect OS
		osAdapter = adapter.Detect()

		// Init history DB
		if err := history.InitDB(config.ConfigDir()); err != nil {
			if verbose {
				ui.PrintWarning(fmt.Sprintf("History DB: %v", err))
			}
		}

		// Load user macros
		macroPath := filepath.Join(config.ConfigDir(), "macros.toml")
		macro.LoadUserMacros(macroPath)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		history.Close()
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "Disable colored output")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")

	rootCmd.AddCommand(storeCmd)
	rootCmd.AddCommand(macroCmd)
	rootCmd.AddCommand(explainCmd)
	rootCmd.AddCommand(historyCmd)
	rootCmd.AddCommand(tagCmd)
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(resumeCmd)
	rootCmd.AddCommand(workspaceCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(privacyCmd)
}

// Execute runs the root command.
func Execute() error {
	return rootCmd.Execute()
}

// statusCmd shows system status.
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show gyanesh-help status and session info",
	Run: func(cmd *cobra.Command, args []string) {
		ui.PrintHeader("gyanesh-help Status")
		fmt.Printf("  Version:    %s\n", version)
		fmt.Printf("  Build:      %s\n", buildDate)
		fmt.Printf("  OS:         %s\n", osAdapter.OSName())
		fmt.Printf("  Distro:     %s\n", osAdapter.OSDistro())
		fmt.Printf("  Pkg Manager: %s\n", osAdapter.PackageManagerName())
		fmt.Printf("  Config Dir: %s\n", config.ConfigDir())
		fmt.Printf("  GPU:        %v\n", osAdapter.HasGPU())

		// Show active tag
		sessionFile := filepath.Join(config.ConfigDir(), "session.json")
		if data, err := os.ReadFile(sessionFile); err == nil {
			var session struct {
				ActiveTag string `json:"active_tag"`
			}
			json.Unmarshal(data, &session)
			if session.ActiveTag != "" {
				fmt.Printf("  Active Tag: %s\n", session.ActiveTag)
			}
		}
	},
}

var privacyCmd = &cobra.Command{
	Use:   "privacy",
	Short: "Show privacy and data usage information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
gyanesh-help Privacy Statement
══════════════════════════════════════════════════════

Zero telemetry, analytics, or crash reporting.

All data is stored locally in ~/.gyanesh-help/:
  • history.db  — your command history (SQLite, local only)
  • config.toml — your preferences
  • macros.toml — your custom macros
  • jobs.json   — active/paused job state
  • queue.json  — pending queue entries

Network calls only occur when YOU explicitly invoke:
  1. Package installations (via your OS package manager)
  2. 'gyanesh-help workspace backup/restore' (GitHub API)
  3. Network probes to 1.1.1.1/8.8.8.8 (only during
     managed installations to detect connectivity loss)

Your data never leaves your machine unless you explicitly
run 'workspace backup'. No cloud, no tracking, no surprises.`)
	},
}
