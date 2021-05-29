package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func oldMain() {
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

func oldMain2() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("[ERROR]:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)
	/*
		io.Copy() function copies from Reader (source) to Writer (destination):
		```
		func Copy(dst Writer, src Reader) (written int64, err error)
		```

		The Writer interface has a write() fuction that writes from byte slice p to the underlying data stream.
		```
		type Writer interface {
			Write(p []byte) (n int, err error)
		}
		```
		The output for a Writer can be anything like outgoing http request, text file, image, terminal etc.

		In our code:
		```
		io.Copy(os.Stdout, resp.Body)
		```
		As expected, 'os.Stdout' has a type of '*File'. The 'File' type in Go implements the Writer interface.

		io.Copy	(dst, 													src		)
					▼														▼
				Something that implements 'Writer' interface			Something that implements 'Reader' interface
					▼														▼
				os.Stdout												resp.Body
					▼
				Value of type 'File'
					▼
				File has a function 'Write()'
					▼
				Therefore, it implements the 'Writer' interface
	*/
}

// A new struct
type logWriter struct{}

// Implementing the write() method. logWriter now implements the Writer interface.
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote", len(bs), "bytes.")
	return len(bs), nil
}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("[ERROR]:", err)
		os.Exit(1)
	}

	io.Copy(logWriter{}, resp.Body)
}
