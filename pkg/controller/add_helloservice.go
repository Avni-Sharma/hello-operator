package controller

import (
	"github.com/Avni-Sharma/hello-operator/pkg/controller/helloservice"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, helloservice.Add)
}
