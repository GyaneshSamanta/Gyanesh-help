package history

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

// QueryOptions defines history query parameters.
type QueryOptions struct {
	Tag      string
	Search   string
	Since    time.Time
	FailOnly bool
	Limit    int
	Offset   int
}

// HistoryEntry represents one row from the history table.
type HistoryEntry struct {
	ID         int
	Timestamp  string
	Command    string
	ExitCode   int
	DurationMs int64
	ProjectTag string
	Stack      string
}

// Query retrieves history entries matching the given options.
func Query(opts QueryOptions) ([]HistoryEntry, error) {
	if DB == nil {
		return nil, fmt.Errorf("history database not initialized")
	}

	if opts.Limit == 0 {
		opts.Limit = 20
	}

	// Use FTS for search queries
	if opts.Search != "" {
		return ftsQuery(opts)
	}

	query := `SELECT id, timestamp, command, exit_code, duration_ms, COALESCE(project_tag,''), COALESCE(stack,'') FROM history`
	var conditions []string
	var args []interface{}

	if opts.Tag != "" {
		conditions = append(conditions, "project_tag = ?")
		args = append(args, opts.Tag)
	}
	if opts.FailOnly {
		conditions = append(conditions, "exit_code != 0")
	}
	if !opts.Since.IsZero() {
		conditions = append(conditions, "timestamp >= ?")
		args = append(args, opts.Since.UTC().Format(time.RFC3339))
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}
	query += " ORDER BY timestamp DESC LIMIT ? OFFSET ?"
	args = append(args, opts.Limit, opts.Offset)

	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanEntries(rows)
}

func ftsQuery(opts QueryOptions) ([]HistoryEntry, error) {
	rows, err := DB.Query(`
		SELECT h.id, h.timestamp, h.command, h.exit_code, h.duration_ms,
			   COALESCE(h.project_tag,''), COALESCE(h.stack,'')
		FROM history h
		JOIN history_fts f ON h.id = f.rowid
		WHERE history_fts MATCH ?
		ORDER BY rank
		LIMIT ? OFFSET ?`, opts.Search, opts.Limit, opts.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanEntries(rows)
}

func scanEntries(rows interface{ Scan(...interface{}) error; Next() bool }) ([]HistoryEntry, error) {
	var entries []HistoryEntry
	for rows.Next() {
		var e HistoryEntry
		if err := rows.Scan(&e.ID, &e.Timestamp, &e.Command, &e.ExitCode, &e.DurationMs, &e.ProjectTag, &e.Stack); err != nil {
			return nil, err
		}
		entries = append(entries, e)
	}
	return entries, nil
}

// ExportCSV writes history to a CSV file.
func ExportCSV(path string, opts QueryOptions) error {
	opts.Limit = 0 // All entries
	entries, err := Query(QueryOptions{Limit: 50000})
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.Write([]string{"id", "timestamp", "command", "exit_code", "duration_ms", "project_tag", "stack"})
	for _, e := range entries {
		w.Write([]string{
			fmt.Sprintf("%d", e.ID), e.Timestamp, e.Command,
			fmt.Sprintf("%d", e.ExitCode), fmt.Sprintf("%d", e.DurationMs),
			e.ProjectTag, e.Stack,
		})
	}
	w.Flush()
	return w.Error()
}
