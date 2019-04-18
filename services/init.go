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

	return nil
}
