package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/GyaneshSamanta/gyanesh-help/internal/history"
	"github.com/GyaneshSamanta/gyanesh-help/internal/ui"
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "View command history",
	Run: func(cmd *cobra.Command, args []string) {
		tag, _ := cmd.Flags().GetString("tag")
		search, _ := cmd.Flags().GetString("search")
		since, _ := cmd.Flags().GetString("since")
		failOnly, _ := cmd.Flags().GetBool("failed")
		all, _ := cmd.Flags().GetBool("all")
		exportCSV, _ := cmd.Flags().GetString("export")

		opts := history.QueryOptions{
			Tag:      tag,
			Search:   search,
			FailOnly: failOnly,
			Limit:    20,
		}
		if all {
			opts.Limit = 50000
		}
		if since != "" {
			t, err := time.Parse("2006-01-02", since)
			if err == nil {
				opts.Since = t
			}
		}

		if exportCSV != "" {
			if err := history.ExportCSV(exportCSV, opts); err != nil {
				ui.PrintError(err.Error())
				return
			}
			ui.PrintSuccess(fmt.Sprintf("Exported to %s", exportCSV))
			return
		}

		entries, err := history.Query(opts)
		if err != nil {
			ui.PrintError(err.Error())
			return
		}

		if len(entries) == 0 {
			ui.PrintInfo("No history entries found.")
			return
		}

		headers := []string{"#", "Timestamp", "Command", "Exit", "Duration", "Tag"}
		var rows [][]string
		for _, e := range entries {
			dur := fmt.Sprintf("%dms", e.DurationMs)
			rows = append(rows, []string{
				fmt.Sprintf("%d", e.ID),
				e.Timestamp[:19],
				truncate(e.Command, 40),
				fmt.Sprintf("%d", e.ExitCode),
				dur,
				e.ProjectTag,
			})
		}
		ui.PrintTable(headers, rows)
	},
}

func init() {
	historyCmd.Flags().String("tag", "", "Filter by project tag")
	historyCmd.Flags().String("search", "", "Full-text search")
	historyCmd.Flags().String("since", "", "Filter entries since date (YYYY-MM-DD)")
	historyCmd.Flags().Bool("failed", false, "Show only failed commands")
	historyCmd.Flags().Bool("all", false, "Show all entries")
	historyCmd.Flags().String("export", "", "Export to CSV file path")
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n-3] + "..."
}
