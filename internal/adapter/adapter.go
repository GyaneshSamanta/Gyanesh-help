package adapter

// OSAdapter abstracts all platform-specific behaviour behind a single interface.
type OSAdapter interface {
	// Package management
	PackageManagerName() string
	InstallPackage(pkg string, args []string) error
	UninstallPackage(pkg string) error
	IsPackageInstalled(pkg string) bool
	PackageVersion(pkg string) (string, error)

	// Lock detection
	LockPaths() []string
	IsLocked() (bool, string, error)

	// Process management
	SuspendProcess(pid int) error
	ResumeProcess(pid int) error
	KillProcess(pid int) error
	RunElevated(cmd string, args []string) error

	// System info
	HomeDir() string
	ConfigDir() string
	OSName() string
	OSDistro() string
	HasGPU() bool
	ShellConfigPaths() []string
}

// contains checks if a string slice contains an item (shared helper).
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
