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
		go checkLink(link) // This will start a new Go Routine to run this function.
		/*
			* The Main Routine will now start new Child Routines to run this function every time the for-loop
				gets here.
			* Unfortunately, this code has an issue. Currently, it is not printing anything. This is because the
				Main Routine is the one that controls when our program exits.
			* So, once the Main Routine runs the for-loop and spawns all the Child Routines, it exits. Meanwhile,
				the Child Routines have not finished their http calls yet. So, the child routines cannot print
				anything to the console.
			* The way to solve this is Channels. Channels are used to communicate between different Go Routines.
			* We will create one Channel which will allow communication between all the different Go Routines.
			* Channels are typed, i.e. the data that we attempt to share between these routines must all be of
				the same type.
		*/
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
