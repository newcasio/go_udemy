package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	//https://golang.org/pkg/net/http/#Get
	//returns a response and an error, so here we define these
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// fmt.Println(resp)
	//resp is a pointer to a Response object which is a struct
	//body of Response object is of type ReadCloser which is interface.  Having a interface as the type gives flexibility, will allow the body to be of any type (eg string, ints, structs, etc).  For this example returning the body of a http get, this response could be anything depending on the url we visit.

	//new empty byteSlice, with 99999 empty spaces
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}

	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil

}
