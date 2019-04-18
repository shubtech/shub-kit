package services

import "fmt"

func (s *skitService) Endpoint(endpointName string) error {
	fmt.Println("endpoint")
	return nil
}
