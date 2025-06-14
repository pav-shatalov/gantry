package tui

type Borders uint8

const (
	TopBorder    Borders = 0b0001
	RightBorder          = 0b0010
	BottomBorder         = 0b0100
	LeftBorder           = 0b1000
)

const NoBorders = 0b0000
const AllBorders = 0b1111

func (b *Borders) has(side Borders) bool {
	return *b&side != 0
}

type BorderType struct {
	topLeft     rune
	top         rune
	topRight    rune
	right       rune
	bottomRight rune
	bottom      rune
	bottomLeft  rune
	left        rune
	width       int
	height      int
}

var SquareBordersType = BorderType{
	topLeft:     '┌',
	top:         '─',
	topRight:    '┐',
	right:       '│',
	bottomRight: '┘',
	bottom:      '─',
	bottomLeft:  '└',
	left:        '│',
}

var RoundBordersType = BorderType{
	topLeft:     '╭',
	top:         '─',
	topRight:    '╮',
	right:       '│',
	bottomRight: '╯',
	bottom:      '─',
	bottomLeft:  '╰',
	left:        '│',
}
