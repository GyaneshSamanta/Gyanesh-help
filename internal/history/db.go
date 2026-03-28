package history

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

const schema = `
CREATE TABLE IF NOT EXISTS history (
    id           INTEGER PRIMARY KEY AUTOINCREMENT,
    timestamp    TEXT    NOT NULL,
    command      TEXT    NOT NULL,
    subcommand   TEXT,
    args         TEXT,
    exit_code    INTEGER NOT NULL DEFAULT 0,
    duration_ms  INTEGER NOT NULL DEFAULT 0,
    project_tag  TEXT,
    stack        TEXT,
    os           TEXT    NOT NULL,
    notes        TEXT
);

CREATE VIRTUAL TABLE IF NOT EXISTS history_fts
    USING fts5(command, args, notes, content='history', content_rowid='id');

CREATE TRIGGER IF NOT EXISTS history_ai AFTER INSERT ON history BEGIN
    INSERT INTO history_fts(rowid, command, args, notes)
    VALUES (new.id, new.command, new.args, new.notes);
END;

CREATE TRIGGER IF NOT EXISTS history_ad AFTER DELETE ON history BEGIN
    INSERT INTO history_fts(history_fts, rowid, command, args, notes)
    VALUES ('delete', old.id, old.command, old.args, old.notes);
END;

CREATE TABLE IF NOT EXISTS macro_stats (
    macro_name   TEXT    NOT NULL,
    run_count    INTEGER NOT NULL DEFAULT 0,
    last_run     TEXT,
    PRIMARY KEY (macro_name)
);

CREATE TABLE IF NOT EXISTS store_installs (
    id           INTEGER PRIMARY KEY AUTOINCREMENT,
    timestamp    TEXT    NOT NULL,
    stack        TEXT    NOT NULL,
    action       TEXT    NOT NULL,
    status       TEXT    NOT NULL,
    components   TEXT,
    duration_ms  INTEGER
);

CREATE INDEX IF NOT EXISTS idx_history_timestamp   ON history(timestamp);
CREATE INDEX IF NOT EXISTS idx_history_project_tag ON history(project_tag);
CREATE INDEX IF NOT EXISTS idx_history_exit_code   ON history(exit_code);
CREATE INDEX IF NOT EXISTS idx_history_subcommand  ON history(subcommand);
`

// DB wraps the SQLite connection for history operations.
var DB *sql.DB

// InitDB opens/creates the history database with WAL mode.
func InitDB(configDir string) error {
	dbPath := filepath.Join(configDir, "history.db")
	os.MkdirAll(configDir, 0755)

	var err error
	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}

	// Enable WAL for concurrent access
	DB.Exec("PRAGMA journal_mode=WAL")
	DB.Exec("PRAGMA busy_timeout=5000")

	_, err = DB.Exec(schema)
	return err
}

// Close shuts down the database connection.
func Close() {
	if DB != nil {
		DB.Close()
	}
}
