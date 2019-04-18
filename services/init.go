package services

import (
	"os"
)

func (s *skitService) Init(projectName string) error {
	if err := initGenFolder(projectName); err != nil {
		return err
	}

	if err := initGenFile(projectName); err != nil {
		return err
	}

	return nil
}

func initGenFolder(projectName string) error {
	os.Mkdir(projectName, dirPerm)
	os.Mkdir(projectName+"/toolkit", dirPerm)
	os.Mkdir(projectName+"/cmd", dirPerm)
	os.Mkdir(projectName+"/logger", dirPerm)
	os.Mkdir(projectName+"/migration", dirPerm)
	os.Mkdir(projectName+"/kubernetes", dirPerm)
	os.Mkdir(projectName+"/services", dirPerm)
	return nil
}

func initGenFile(projectName string) error {
	//Kubernetes namespace
	f, err := os.Create(projectName + "/kubernetes/namespace.yml")
	f.Close()
	if err != nil {
		return err
	}

	f, err = os.Create(projectName + "/logger/logger.go")
	_, err = f.WriteString("package logger")
	if err != nil {
		return err
	}
	f.Close()

	f, err = os.Create(projectName + "/toolkit/toolkit.go")
	_, err = f.WriteString("package toolkit")
	if err != nil {
		return err
	}
	f.Close()

	f, err = os.Create(projectName + "/docker-compose.yml")
	if err != nil {
		return err
	}

	f.WriteString(`version: "3.5"`)

	f.Close()

	f, err = os.Create(projectName + "/.gitignore")
	if err != nil {
		return err
	}

	f.WriteString(`**/.env
**/bin
# Created by https://www.gitignore.io/api/go,code
# Edit at https://www.gitignore.io/?templates=go,code

### Code ###
.vscode/*
!.vscode/settings.json
!.vscode/tasks.json
!.vscode/launch.json
!.vscode/extensions.json

### Go ###
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with "go test -c"
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

### Go Patch ###
/vendor/
/Godeps/

# End of https://www.gitignore.io/api/go,code
`)
	f.Close()

	return nil
}
