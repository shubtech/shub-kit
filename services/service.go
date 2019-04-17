package services

//Service define service interface
type Service interface {
	Init(projectName string) error
	Add(serviceName string) error
	Endpoint(endpointName string) error
}

type skitService struct{}

//MakeService create service
func MakeService() Service {
	return &skitService{}
}
