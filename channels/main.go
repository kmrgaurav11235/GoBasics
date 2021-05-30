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

	c := make(chan string) // channel of type string

	for _, link := range links {
		// go checkLink(link) // This will start a new Go Routine to run this function.
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
		go checkLink(link, c) // now, we pass in the channel as well
	}
	// Unfortunately, the old program was not concurrent. This is because receiving a message over a channel is a
	// Blocking operation. So, we need to move the `fmt.Println` statement outside the loop. Otherwise the code
	// inside the for-loop will be executed sequentially.

	// Receiving and printing data from channel.
	// We need to run `fmt.Println` statement 5 times to capture response from the 5 http calls.
	// for i := 0; i < len(links); i++ {

	// Infinite loop
	// for {
	// 	go checkLink(<-c, c)
	// }

	// Alternative loop syntax
	// The for-loop below says: Wait for the channel to return a value. Assign that value to l. Then run the 'checkLink'
	// go Routine for it.
	for l := range c {
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link // Pushing the same link back into the channel so that we can repeat the check for the same URL
	} else {
		fmt.Println(link, "is up.")
		c <- link
	}

	/*
		Sending data with Channels:
		* ```channel <- 5``` : Send the value '5' into the channel.
		* ```myNumber <- channel```: Wait for a value to be sent into the channel. When we get one, assign the value
			to variable 'myNumber'.
		* ```fmt.Println(<- channel)```: Wait for a value to be sent into the channel. When we get one, print the value
			to the console.
	*/
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
