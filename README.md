# protoc-gen-debug
A utility for debugging protoc plugins

## How to use
First build the app by either `go build .` or `go get github.com/drekle/protoc-gen-debug`.  Once the plugin is built it should be placed somewhere in your path.  When in your path you can invoke this plugin with the desired arguments for another plugin you are building.

For example:
```
protoc --debug_out=. --debug_opt=one=1,two=2 helloworld.proto
```
This will produce a debug.dat file in the current directory.  This is used to simulate the call to potentially another protoc plugin in development with the same arguements and proto file, which can be useful if you would like to debug a protoc plugin in an IDE.

To debug this invocation in an IDE in your desired protoc plugin add the following capabilities to main.

Be able to support an optional flag to pass in a filename:
```
	stdInFile := flag.String("stdinFile", "", "A file to use for stdin")
	flag.Parse()
```

If not set read from standard in, however if set read from file.  Then simply marshal the data read into a `CodeGeneratorRequest` type imported from `github.com/golang/protobuf/protoc-gen-go/plugin`.
```
	req := &plugin.CodeGeneratorRequest{}
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
```

This above example can be found in this same repo, as well as github.com/drekle/protoc-gen-cli.