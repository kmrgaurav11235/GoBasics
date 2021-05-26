package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com") // GET request to google.com

	if err != nil {
		fmt.Println("[ERROR]:", err)
		os.Exit(1)
	}

	/*
		type Response is a struct. We want to print its body which is of type:
		`Body io.ReadCloser`

		ReadCloser is an interface:
		```
		type ReadCloser interface {
			Reader
			Closer
		}
		```

		Reader:
		```
		type Reader interface {
			Read(p []byte) (n int, err error)
		}
		```

		Closer:
		```
		type Closer interface {
			Close() error
		}
		```

		So, the implmentation of body must have these 2 functions:
		1. Read(p []byte) (n int, err error): It pushes the data that was read inside the byte slice
			provided and returns the number of bytes read and any error encountered. We will use this
			to get the response body.
		2. Close() error: Basic Close method.
	*/

	byteSlice := make([]byte, 99999)
	/* The Read function is not set-up to re-size the slice if the slice is already full. It reads data
	into the byte slice untill the slice is full. So, if we initialize a slice with zero elements, the
	read function will observe that the slice is already full and quit. This is why we create a big byte
	slice that the data can fit into.
	*/

	resp.Body.Read(byteSlice)
	fmt.Println(string(byteSlice))
}
