package main

import (
	"fmt"
	"os"

	"github.com/shubtech/shub-kit/color"
	"github.com/shubtech/shub-kit/services"
)

func main() {
	if len(os.Args) != 3 {
		usage()
		return
	}

	s := services.MakeService()

	if os.Args[1] == "init" {
		if err := s.Init(os.Args[2]); err != nil {
			panic(err)
		}
	} else if os.Args[1] == "add" {
		if err := s.Add(os.Args[2]); err != nil {
			panic(err)
		}
	} else if os.Args[1] == "endpoint" {
		if err := s.Endpoint(os.Args[2]); err != nil {
			panic(err)
		}
	} else {
		usage()
	}
}

func usage() {
	mess := color.Sprintf(color.InfoColor, "Usages: \n\t",
		color.DebugColor,
		"skit init project_name ",
		color.NormalColor,
		"to initial project structer\n\t",
		color.DebugColor, "skit add service_name ",
		color.NormalColor,
		"add a service to microservice\n\n",
		color.DebugColor,
		"Happy Hacking!",
	)
	fmt.Println(mess)
}
