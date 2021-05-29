package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://github.com",
		"http://golang.org",
		"http://colorhunt.co",
	}

	for _, link := range links {
		checkLink(link) // starts a new Go Routine to run this function.
	}
}

func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
	} else {
		fmt.Println(link, "is up.")
	}
}

/*
* By default, Go tries to use one core.

				--------------------------
						One CPU Core
				--------------------------
							▼
	----------------------------------------------------
						Go Scheduler
	----------------------------------------------------
	Go Scheduler runs one routine until it is finished
	or makes a Blocking call (like an http request).
	----------------------------------------------------
		▼					▼					▼
	----------------------------------------------------
	Go Routine 1	|	Go Routine 2	|	Go Routine 3
	----------------------------------------------------

* You change the default setting and ask Go to use multiple cores.

	----------------------------------------------------
	CPU Core 1		|	CPU Core  2		|	CPU Core 3
	----------------------------------------------------
		▼					▼					▼
	----------------------------------------------------
						Go Scheduler
	----------------------------------------------------
	Go Scheduler runs one routine until it is finished
	or makes a Blocking call (like an http request).
	----------------------------------------------------
		▼					▼					▼
	----------------------------------------------------
	Go Routine 1	|	Go Routine 2	|	Go Routine 3
	----------------------------------------------------

* "Concurrency is not parallelism." - You would hear this quote a lot.
* Concurrency: We have have multiple threads executing code. If one thread blocks, another one
	is picked up and worked on.
* Parallelism: Multiple threads executed at the exact same time. Required multiple CPUs.
*/
