package main

import (
	"fmt"
	"github.com/xackery/peqrepack/deploy"
	"github.com/xackery/peqrepack/library/deploy/apt"
	"os"
)

var instance deploy.Deploy

func main() {
	var err error

	//Step 1: Check Environment
	fmt.Printf("Determining environment...")
	err = LearnEnvironment()
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}

	//Step 2: Check Dependencies
	fmt.Printf("Checking dependencies...")
	dependencies, command, err := instance.CheckDependencies()
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("done\n")

	//Step 3: Install Dependencies
	if len(dependencies) > 0 {
		fmt.Printf("The following dependencies need to be installed: ")
		dependList := ""
		for _, depend := range dependencies {
			dependList += depend + ", "
		}
		dependList = dependList[0 : len(dependList)-2]
		fmt.Printf("%s\n", dependList)
		fmt.Printf("To install them, this script will execute: '%s'\nContinue? [y]", command)
		for _, depend := range dependencies {
			fmt.Printf("Installing %s", depend)
			err = instance.InstallDependency(depend)
			if err != nil {
				fmt.Printf("error: %s\n", err.Error())
				return
			}
		}

	}

}

//Figure out which environment we're in
func LearnEnvironment() (err error) {
	instance = &apt.Deploy{}
	err = instance.GetEnvironment()
	if err == nil {
		return
	}
	return
}
