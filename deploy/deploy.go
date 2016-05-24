package deploy

import ()

type Deploy interface {
	GetEnvironment() (err error)
	CheckDependencies() (dependencies []string, command string, err error)
	InstallDependency(depend string) (err error)
}
