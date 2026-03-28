# gyanesh-help

> **Cross-platform CLI developer utility** — makes the terminal feel like it already knows what you need.

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat-square&logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)](LICENSE)
[![Binary Size](https://img.shields.io/badge/Binary-~7MB-purple?style=flat-square)]()

---

## What is gyanesh-help?

A **zero-dependency, ultra-lightweight CLI tool** that replaces the friction of:

| Pain Point | What gyanesh-help Does |
|------------|----------------------|
| 🔒 Lock conflicts (`apt`/`dpkg`/`msiexec` busy) | Queues your command and auto-executes when free |
| 📡 Network drops during large installs | Auto-pauses and resumes on connectivity restore |
| 🧠 Unmemorable commands (`git reset --hard HEAD~1`) | 30 semantic macros like `git-undo`, `port-kill` |
| ⏱ Environment setup (2-4 hours) | One-command stores: `gyanesh-help store install mern` |
| 📜 Lost terminal history | Queryable, tagged SQLite history database |
| 💾 Machine migration | Workspace backup & restore via GitHub |

**All offline. All local. Zero telemetry.**

---

## Installation

### Linux / macOS
```bash
curl -fsSL https://raw.githubusercontent.com/GyaneshSamanta/gyanesh-help/main/scripts/install.sh | bash
```

### Windows (PowerShell)
```powershell
iwr https://raw.githubusercontent.com/GyaneshSamanta/gyanesh-help/main/scripts/install.ps1 | iex
```

### Go Install
```bash
go install github.com/GyaneshSamanta/gyanesh-help@latest
```

### Manual
Download the pre-built binary for your platform from [GitHub Releases](https://github.com/GyaneshSamanta/gyanesh-help/releases).

---

## Quick Start

```bash
# Check status
gyanesh-help status

# Install a package (with queue management)
gyanesh-help install vim

# Set up a full MERN stack
gyanesh-help store install mern

# Preview what a store installs
gyanesh-help store preview data-science

# Use semantic macros
gyanesh-help git-undo           # Safe undo last commit
gyanesh-help git-undo --hard    # Destructive undo (asks confirmation)
gyanesh-help port-kill 3000     # Kill process on port 3000
gyanesh-help nuke-node          # Remove node_modules + lock

# Explain any macro before running it
gyanesh-help explain git-undo

# View command history
gyanesh-help history
gyanesh-help history --tag myproject --search "docker"

# Tag your session
gyanesh-help tag set myproject

# Backup workspace to GitHub
gyanesh-help workspace backup

# Restore on a new machine
gyanesh-help workspace restore --repo https://github.com/you/dev-workspace-backup
```

---

## Environment Stores

| Store | Command | What's Included |
|-------|---------|----------------|
| **Data Science** | `store install data-science` | Python, R, JupyterLab, NumPy, Pandas, Matplotlib |
| **Frontend** | `store install frontend` | Node.js, npm, yarn, pnpm, Vite, ESLint, TypeScript |
| **LAMP** | `store install lamp` | Apache, MySQL, PHP, Composer |
| **MERN** | `store install mern` | MongoDB, Express, React, Node.js, PM2 |
| **Backend** | `store install backend` | Docker, PostgreSQL, Redis, HTTPie, Make |
| **AI Dev** | `store install ai-dev` | PyTorch, TensorFlow, HuggingFace, Ollama, CUDA |
| **Claude** | `store install claude` | Anthropic SDK, MCP SDK, promptfoo |

---

## 30 Built-in Macros

| Category | Macros |
|----------|--------|
| **Git** | `git-undo`, `git-clean`, `git-save`, `git-unsave`, `git-whoops`, `git-oops-push`, `git-log-pretty`, `git-branch-clean`, `git-diff-staged` |
| **Network** | `port-kill`, `port-check`, `kill-port`, `ip-info`, `cert-check` |
| **System** | `env-check`, `disk-check`, `process-find`, `hosts-edit`, `path-add` |
| **Filesystem** | `find-big-files`, `find-old-logs`, `nuke-node` |
| **Python** | `pip-freeze-clean`, `venv-create` |
| **Node.js** | `npm-audit-fix` |
| **Docker** | `docker-nuke`, `docker-shell` |
| **SSH** | `ssh-keygen-github` |
| **Workspace** | `backup-now` |

Every macro shows the exact command executed AND a plain-English explanation.

---

## Privacy

- **Zero telemetry, analytics, or crash reporting**
- All data stored locally in `~/.gyanesh-help/`
- No outbound network calls except when you explicitly install packages or backup
- Run `gyanesh-help privacy` for full details

---

## Platform Support

| OS | Package Managers | Architecture |
|----|-----------------|-------------|
| Ubuntu / Debian | `apt` | amd64, arm64 |
| Fedora / RHEL | `dnf` | amd64, arm64 |
| Arch Linux | `pacman` | amd64, arm64 |
| macOS | `brew` | amd64, arm64 (Apple Silicon) |
| Windows 10/11 | `winget`, `choco` | amd64 |

---

## Configuration

Config file: `~/.gyanesh-help/config.toml`

```toml
[core]
lock_poll_interval_secs = 5
lock_timeout_mins = 30

[ui]
color = true
explain_after_macro = true

[network]
probe_host = "1.1.1.1"
```

---

## Building from Source

```bash
git clone https://github.com/GyaneshSamanta/gyanesh-help.git
cd gyanesh-help
go build -ldflags "-s -w" -o gyanesh-help .
```

---

## License

[MIT](LICENSE) — Built by [Gyanesh Samanta](https://github.com/GyaneshSamanta)
