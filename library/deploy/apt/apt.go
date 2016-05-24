package apt

import (
	"os"
)

type Deploy struct{}

func (d *Deploy) GetEnvironment() (err error) {
	if _, err = os.Stat("/usr/bin/apt-get"); err == nil {
		return
	}
	return
}

func (d *Deploy) CheckDependencies() (dependencies []string, command string, err error) {
	return
}

func (d *Deploy) InstallDependency(depend string) (err error) {
	return
}
