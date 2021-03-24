// pormodoro timer
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	start := time.Now()

	finish, err := waitDuration(start)
	if err != nil {
		flag.Usage()
		os.Exit(2)
	}
	wait := finish.Sub(start)

	formatter := formatSeconds
	switch {
		case wait >= 24*time.Hour:
			formatter = formatDays
		case wait >= time.Hour:
			formatter = formatHours
		case wait >= time.Minute:
			formatter = formatMinutes
	}

	fmt.Printf("Start timer for %s.\n\n", wait)

	if *simple {
		simpleCountdown(finish, formatter)
	} else {
		fullscreenCountdown(start, finish, formatter)
	}

	if !*silence {
		fmt.Println("\a") // \a is the bell literal.
	} else {
		fmt.Println()
	}

	beepErr := beeep.Notify("Title", "Message body", "assets/information.png")
	if beepErr != nil {
		panic(beepErr)
	}
}
