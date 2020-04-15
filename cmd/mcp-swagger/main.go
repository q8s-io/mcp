package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"k8s.io/klog"

	"github.com/q8s-io/mcp/pkg/server"
)

var (
	file string
)

func main() {
	flag.StringVar(&file, "file", "api/openapi-spec/swagger.json", "File to store swagger json.")
	flag.Parse()

	container := server.NewContainer()

	config := restfulspec.Config{
		WebServices: container.RegisteredWebServices(),
		APIPath:     "/apidocs.json",
	}
	swagger := restfulspec.BuildSwagger(config)

	b, err := json.Marshal(swagger)
	if err != nil {
		klog.Errorf("error to marshal json for swagger, %v", err)
		return
	}

	var out bytes.Buffer
	if err := json.Indent(&out, b, "", "\t"); err != nil {
		klog.Errorf("error to indent json, %v", err)
		return
	}

	if err := ioutil.WriteFile(file, out.Bytes(), os.ModePerm); err != nil {
		klog.Errorf("error to write file to %s, %v", file, err)
	}
}
