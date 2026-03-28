package history

import (
	"encoding/json"
	"runtime"
	"time"
)

// Record writes a command execution to the history database.
func Record(cmd string, subCmd string, args []string, exitCode int, durationMs int64, tag string, stack string) error {
	if DB == nil {
		return nil
	}
	argsJSON, _ := json.Marshal(args)
	osLabel := runtime.GOOS
	_, err := DB.Exec(`
		INSERT INTO history (timestamp, command, subcommand, args, exit_code, duration_ms, project_tag, stack, os)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		time.Now().UTC().Format(time.RFC3339),
		cmd, subCmd, string(argsJSON), exitCode, durationMs, tag, stack, osLabel)
	return err
}

// RecordMacro updates macro usage statistics.
func RecordMacro(name string) {
	if DB == nil {
		return
	}
	DB.Exec(`INSERT INTO macro_stats (macro_name, run_count, last_run) VALUES (?, 1, ?)
		ON CONFLICT(macro_name) DO UPDATE SET run_count = run_count + 1, last_run = ?`,
		name, time.Now().UTC().Format(time.RFC3339), time.Now().UTC().Format(time.RFC3339))
}

// RecordStoreAction logs a store install/remove action.
func RecordStoreAction(stack, action, status string, durationMs int64) {
	if DB == nil {
		return
	}
	DB.Exec(`INSERT INTO store_installs (timestamp, stack, action, status, duration_ms) VALUES (?, ?, ?, ?, ?)`,
		time.Now().UTC().Format(time.RFC3339), stack, action, status, durationMs)
}
