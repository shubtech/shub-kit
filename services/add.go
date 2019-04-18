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

	if err := addHTTP(serviceName); err != nil {
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

	os.Mkdir("migration/"+serviceName, dirPerm)

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

	// service.go
	f, err := os.OpenFile("services/"+serviceName+"/"+"service/service.go", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	_, err = f.WriteString(`

type Service struct {}`)

	if err != nil {
		return err
	}

	f.Close()

	// compose.go
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

	// endpoint.go
	endpoint, err := os.OpenFile("services/"+serviceName+"/endpoints/endpoints.go",
		os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer endpoint.Close()

	if _, err := endpoint.WriteString(`

import (
	"github.com/go-kit/kit/endpoint"
)

//Endpoints ...
type Endpoints struct {}

//MakeServerEndpoints create endpoint for service
func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{}
}
`); err != nil {
		return err
	}

	// config.go
	config, err := os.OpenFile("services/"+serviceName+"/config/config.go",
		os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer config.Close()

	config.WriteString(`
	
//Config store serivce's config
type Config struct {}

func init() {}
`)

	if _, err := os.Create("services/" + serviceName + "/Dockerfile"); err != nil {
		return err
	}

	if _, err := os.Create("services/" + serviceName + "/docker-compose.yml"); err != nil {
		return err
	}

	return nil
}

func addHTTP(serviceName string) error {
	f, err := os.Create("services/" + serviceName + "/http/encode.go")
	if err != nil {
		return err
	}

	defer f.Close()
	f.WriteString(`package http

import (
	"context"
	"encoding/json"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
)

// reference @hieunmce, thank to him

// encodeResponse is the common method to encode all response types to the client.
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	// maybe we can be smart here by returning text/json error based on request's
	// content-type header
	encodeJSONError(ctx, err, w)
}

func encodeJSONError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// custom headers
	if headerer, ok := err.(kithttp.Headerer); ok {
		for k, values := range headerer.Headers() {
			for _, v := range values {
				w.Header().Add(k, v)
			}
		}
	}
	code := http.StatusBadRequest
	// custome code
	if sc, ok := err.(kithttp.StatusCoder); ok {
		code = sc.StatusCode()
	}
	w.WriteHeader(code)
	// enforce json response
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
	`)

	return nil
}
