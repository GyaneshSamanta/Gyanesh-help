package builtins

import "github.com/GyaneshSamanta/gyanesh-help/internal/macro"

func registerWorkspaceMacros() {
	reg(&macro.Macro{
		Name: "backup-now", Category: "workspace",
		Description: "Quick alias for workspace backup",
		Commands:    []macro.Step{{OS: "all", Command: "gyanesh-help workspace backup"}},
		Explanation: `
✔ Workspace backup triggered.
─────────────────────────────────────────────────────
This is a shortcut for 'gyanesh-help workspace backup'.
Your shell configs, macros, and store manifests will
be pushed to your private GitHub backup repo.
─────────────────────────────────────────────────────`,
		BuiltIn: true,
	})
}
