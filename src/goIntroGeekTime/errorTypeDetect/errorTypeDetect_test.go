package errorTypeDetect

import (
	"fmt"
	"testing"

	openapiError "github.com/go-openapi/errors"
	genericError "github.com/pkg/errors"
)

var openApiErrorTimeOut = openapiError.New(500, "Operation Timeout")
var openApiErrorNotFound = openapiError.New(404, "Not Found")
var genericErrorTimeout = genericError.New("Operation Timeout")
var genericErrorNotFound = genericError.New("Not Found")

func openApidFindTimeOut() error {
	return openApiErrorTimeOut
}

func openApiFindNotFound() error {
	return openApiErrorNotFound
}

func genericFindTimeOut() error {
	return genericErrorTimeout
}

func genericFindNotFound() error {
	return genericErrorNotFound
}

func TestErrorTypeDetect(t *testing.T) {
	if err := openApidFindTimeOut(); err != nil {
		if err == openApiErrorTimeOut {
			fmt.Printf("openApi Timeout : %v\n", err)
		}
	}
	if err := openApiFindNotFound(); err != nil {
		if err == openApiErrorNotFound {
			fmt.Printf("openApi Not found : %v\n", err)
		}
	}

	if err := genericFindNotFound(); err != nil {
		if err == genericErrorNotFound {
			fmt.Printf("generic Not Found : %v\n", err)
		}
	}

	if err := genericFindTimeOut(); err != nil {
		if err == genericErrorTimeout {
			fmt.Printf("generic Timeout : %v\n", err)
		}
	}

}
