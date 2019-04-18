package services

import (
	"os"
)

func (s *skitService) Add(serviceName string) error {
	if err := addGenFolder(serviceName); err != nil {
		return err
	}

	if err := addGenFile(serviceName); err != nil {
		return err
	}

	return nil
}

func addGenFolder(serviceName string) error {
	os.Mkdir("cmd/"+serviceName, dirPerm)
	os.Mkdir("cmd/"+serviceName+"/server", dirPerm)
	servicePath := "services/" + serviceName
	os.Mkdir(servicePath, dirPerm)

	for _, f := range []string{"config", "db", "endpoints", "http", "model", "service", "util"} {
		os.Mkdir("services/"+serviceName+"/"+f, dirPerm)
	}

	return nil
}

func addGenFile(serviceName string) error {
	m, err := os.Create("cmd/" + serviceName + "/server/main.go")
	if err != nil {
		return err
	}
	m.WriteString("package main\n\nfunc main() {\n\n}")
	m.Close()

	for _, f := range []string{"config", "db", "endpoints", "http", "model", "service", "util"} {
		ff, err := os.Create("services/" + serviceName + "/" + f + "/" + f + ".go")
		if err != nil {
			return err
		}

		ff.WriteString("package " + f)
		ff.Close()
	}

	compose, err := os.Create("services/" + serviceName + "/service/compose.go")
	if err != nil {
		return err
	}
	compose.WriteString("package compose")
	composeSource := `

import (
"reflect"
)

//Compose applies middlewares to Service
func Compose(s interface{}, mws ...interface{}) interface{} {
	for i := len(mws) - 1; i >= 0; i-- {
		vv := reflect.ValueOf(mws[i]).Call([]reflect.Value{reflect.ValueOf(s)})
		s = vv[0].Interface()
	}
	return s
}`
	compose.WriteString(composeSource)
	compose.Close()

	if _, err := os.Create("services/" + serviceName + "/Dockerfile"); err != nil {
		return err
	}

	if _, err := os.Create("services/" + serviceName + "/docker-compose.yml"); err != nil {
		return err
	}

	return nil
}
