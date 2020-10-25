package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {

	stdInFile := flag.String("stdinFile", "", "A file to use for stdin")
	flag.Parse()

	// os.Stdin will contain data which will unmarshal into the following object:
	// https://godoc.org/github.com/golang/protobuf/protoc-gen-go/plugin#CodeGeneratorRequest
	req := &plugin.CodeGeneratorRequest{}
	resp := &plugin.CodeGeneratorResponse{}

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	if *stdInFile != "" {
		data, err = ioutil.ReadFile(*stdInFile)
		if err != nil {
			panic(err)
		}
	}

	// You must use the requests unmarshal method to handle this type
	if err := proto.Unmarshal(data, req); err != nil {
		panic(err)
	}
	test, err := proto.Marshal(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		errStr := err.Error()
		resp.Error = &errStr
	} else {
		fName := "debug.dat"
		content := fmt.Sprintf("%s", string(test))
		resp.File = append(resp.File, &pluginpb.CodeGeneratorResponse_File{
			Name:    &fName,
			Content: &content,
		})
	}

	marshalled, err := proto.Marshal(resp)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(marshalled)
}
