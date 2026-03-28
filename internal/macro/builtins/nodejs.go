package builtins

import "github.com/GyaneshSamanta/gyanesh-help/internal/macro"

func registerNodejsMacros() {
	reg(&macro.Macro{
		Name: "npm-audit-fix", Category: "nodejs",
		Description: "Auto-fix npm audit vulnerabilities",
		Commands:    []macro.Step{{OS: "all", Command: "npm audit fix --force"}},
		Explanation: `
✔ Done. npm vulnerabilities were auto-fixed.
─────────────────────────────────────────────────────
--force may have upgraded or downgraded some packages.
Run 'npm test' to verify nothing is broken.
Review changes: git diff package.json package-lock.json
─────────────────────────────────────────────────────`,
		BuiltIn: true,
	})
}
