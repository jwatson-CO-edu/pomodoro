package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/carlmjohnson/pomodoro/display"
	"github.com/nsf/termbox-go"
)

func fullscreenCountdown( start, finish time.Time, formatter func(time.Duration) string ) {
	// Display the countdown in the terminal

	// Open a termbox display
	err := termbox.Init()
	if err != nil {
		fmt.Fprintln( os.Stderr, "Couldn't open display:", err )
		os.Exit(2)
	}
	defer termbox.Close()

	// Leaks a goroutine
	ticker := time.Tick( 40*time.Millisecond ) // Timer channel
	quit   := make( chan struct{} ) // ---------- Quit  channel

	// Leaks if not quit
	go func() {
		defer close( quit )
		for {
			e := termbox.PollEvent()
			// Quit on any of the common keys for quitting
			if strings.ContainsRune( "CcDdQqXx", e.Ch ) || // [C,D,Q,X]
				e.Key == termbox.KeyCtrlC || // ------------- [Ctrl]+[c]
				e.Key == termbox.KeyCtrlD || // ------------- [Ctrl]+[d]
				e.Key == termbox.KeyCtrlQ || // ------------- [Ctrl]+[q]
				e.Key == termbox.KeyCtrlX {  return  }  // -- [Ctrl]+[x]
		}
	}()

	for render( start, finish, formatter ) {
		select {
			case <-ticker:
			case <-quit:
				termbox.Close()
				os.Exit(1)
				return
		}
	}

}

func render( start, finish time.Time, formatter func(time.Duration) string ) bool {


	now       := time.Now()
	remaining := -now.Sub( finish )
	if remaining < 0 {
		return false
	}

	const timeFmt     = "3:04:05pm"
	screenW, screenH := termbox.Size()
	centerX /*----*/ := screenW / 2
	centerY /*----*/ := screenH / 2

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	startStr := start.Format(timeFmt)
	
	display.Point{
		X:  0, 
		Y:  0,
		Fg: termbox.ColorBlue, 
		Bg: termbox.ColorDefault,
	}.Str("Start")

	display.Point{
		X:  0, 
		Y:  1,
		Fg: termbox.ColorWhite, 
		Bg: termbox.ColorDefault,
	}.Str(startStr)

	nowStr := now.Format( timeFmt )

	display.Point{
		X:  centerX - (len("Now") / 2), 
		Y:  0,
		Fg: termbox.ColorBlue, 
		Bg: termbox.ColorDefault,
	}.Str( "Now" )
	
	display.Point{
		X:  centerX - (len(nowStr) / 2), 
		Y:  1,
		Fg: termbox.ColorWhite, 
		Bg: termbox.ColorDefault,
	}.Str( nowStr )

	finishStr := finish.Format(timeFmt)

	display.Point{
		X:  screenW - len("Finish"), 
		Y:  0,
		Fg: termbox.ColorBlue, 
		Bg: termbox.ColorDefault,
	}.Str( "Finish" )
	
	display.Point{
		X:  screenW - len(finishStr), 
		Y:  1,
		Fg: termbox.ColorWhite, 
		Bg: termbox.ColorDefault,
	}.Str(finishStr)

	remainingStr := formatter(remaining)
	display.Point{
		X:  centerX - (len(remainingStr) * (display.BigCharWidth + 1) / 2), 
		Y:  centerY,
		Fg: termbox.ColorBlue, 
		Bg: termbox.ColorDefault,
	}.BigStr( remainingStr )

	display.Point{
		X:  0, 
		Y:  centerY + 6,
		Fg: termbox.ColorBlue, 
		Bg: termbox.ColorWhite,
	}.ProgressBar( screenW, int( start.Sub( now ) ), int( start.Sub( finish ) ) )

	termbox.Flush()
	return true
}
