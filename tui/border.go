package tui

type Borders struct {
	topLeft     rune
	top         rune
	topRight    rune
	right       rune
	bottomRight rune
	bottom      rune
	bottomLeft  rune
	left        rune
}

var SquareBorders = Borders{
	topLeft:     '┌',
	top:         '─',
	topRight:    '┐',
	right:       '│',
	bottomRight: '┘',
	bottom:      '─',
	bottomLeft:  '└',
	left:        '│',
}

var RoundBorders = Borders{
	topLeft:     '╭',
	top:         '─',
	topRight:    '╮',
	right:       '│',
	bottomRight: '╯',
	bottom:      '─',
	bottomLeft:  '╰',
	left:        '│',
}
