# pomodoro
Command line [pomodoro timer](https://en.wikipedia.org/wiki/Pomodoro_Technique), implemented in Go.

## Installation
First install [Go](http://golang.org).

```bash
cd ~
git clone https://github.com/jwatson-CO-edu/pomodoro.git
cd pomodoro
env GOOS=linux GOARCH=amd64 go install
```

## Usage
Usage of pomodoro:

    pomodoro [options] [duration]

Duration defaults to 25 minutes. Durations may be expressed as integer minutes
(e.g. "15") or time with units (e.g. "1m30s" or "90s").

Chimes system bell at the end of the timer, unless -silence is set.

## Screenshots
```bash
$ pomodoro -simple
Start timer for 25m0s.

Countdown: 24:43

$ pomodoro -h
Usage of pomodoro:

    pomodoro [options] [duration]

Duration defaults to 25 minutes. Durations may be expressed as integer minutes
(e.g. "15") or time with units (e.g. "1m30s" or "90s").

Chimes system bell at the end of the timer, unless -silence is set.
  -silence
        Don't ring bell after countdown
  -simple
        Display simple countdown
```

![screenshot](./screenshot.png)

## Recommended helper
[beeep](https://github.com/gen2brain/beeep) can be used to bring up a system alert when pomodoro finishes.
