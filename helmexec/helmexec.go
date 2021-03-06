package helmexec

// Interface for executing helm commands
type Interface interface {
	SetExtraArgs(args ...string)
	SetHelmBinary(bin string)

	AddRepo(name, repository, certfile, keyfile, username, password string) error
	UpdateRepo() error
	BuildDeps(chart string) error
	UpdateDeps(chart string) error
	SyncRelease(context HelmContext, name, chart string, flags ...string) error
	DiffRelease(context HelmContext, name, chart string, flags ...string) error
	TemplateRelease(chart string, flags ...string) error
	Fetch(chart string, flags ...string) error
	Lint(chart string, flags ...string) error
	ReleaseStatus(context HelmContext, name string, flags ...string) error
	DeleteRelease(context HelmContext, name string, flags ...string) error
	TestRelease(context HelmContext, name string, flags ...string) error
	List(context HelmContext, filter string, flags ...string) (string, error)
	DecryptSecret(context HelmContext, name string, flags ...string) (string, error)
}
