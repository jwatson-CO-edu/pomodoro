package display

import "github.com/nsf/termbox-go" // Terminal UI

const ( // Giant character dimensions
	BigCharWidth  = 3
	BigCharHeight = 5
)

// Dict of large display characters
var bigChars = map[rune][BigCharHeight]string{
	0: {
		"X X",
		" X ",
		"X X",
		" X ",
		"X X",
	},
	' ': {
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
	},
	'0': {
		"XXX",
		"X X",
		"X X",
		"X X",
		"XXX",
	},
	'1': {
		" X ",
		"XX ",
		" X ",
		" X ",
		"XXX",
	},
	'2': {
		"XXX",
		"  X",
		"XXX",
		"X  ",
		"XXX",
	},
	'3': {
		"XXX",
		"  X",
		"XXX",
		"  X",
		"XXX",
	},
	'4': {
		"X X",
		"X X",
		"XXX",
		"  X",
		"  X",
	},
	'5': {
		"XXX",
		"X  ",
		"XXX",
		"  X",
		"XXX",
	},
	'6': {
		"XXX",
		"X  ",
		"XXX",
		"X X",
		"XXX",
	},
	'7': {
		"XXX",
		"  X",
		"  X",
		" X ",
		" X ",
	},
	'8': {
		"XXX",
		"X X",
		"XXX",
		"X X",
		"XXX",
	},
	'9': {
		"XXX",
		"X X",
		"XXX",
		"  X",
		"  X",
	},
	':': {
		"   ",
		" X ",
		"   ",
		" X ",
		"   ",
	},
	'.': {
		"   ",
		"   ",
		"   ",
		" XX",
		" XX",
	},
}

type Point struct {
	// Terminal Coordinate
	X, Y   int
	Fg, Bg termbox.Attribute
}

func (p Point) Char(ch rune) {
	// Giant terminal character
	termbox.SetCell( p.X, p.Y, ch, p.Fg, p.Bg )
}

func (p Point) Str( s string ) {
	// Create `Point`s from a string
	for _, c := range s {
		p.Char(c)
		p.X++
	}
}

func (p Point) Pattern( pattern [BigCharHeight]string ) {
	// Build terminal characters per row of a string
	for y, line := range pattern {
		for x, c := range line {
			q := p
			q.X += x
			q.Y += y
			if c != ' ' {  q.Fg, q.Bg = p.Bg, p.Fg  }
			q.Char(' ')
		}
	}
}

func (p Point) BigChar( ch rune ) {
	// `Pattern` for one char
	pattern, ok := bigChars[ch]
	if !ok {  pattern = bigChars[0]  }
	p.Pattern( pattern )
}

func (p Point) BigStr( s string ) {
	// Collection of `Pattern`s for a char string
	xOffset := p.X
	for i, c := range s {
		p.X = xOffset + i*(BigCharWidth+1)
		p.BigChar(c)
	}
}

func (p Point) ProgressBar( length, cur, total int ) {
	// Graphic representation of time remaining
	divider := (length * cur) / total
	for x := 0; x < length; x++ {
		ch := ' '
		q := p
		q.X += x
		if x == divider {
			ch = '░'
		}
		if x < divider {
			q.Fg, q.Bg = p.Bg, p.Fg
		}
		q.Char(ch)
	}
}
