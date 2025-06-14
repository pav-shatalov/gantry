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
	topLeftCorner     rune
	horizontal        rune
	topRightCorner    rune
	vertical          rune
	bottomRightCorner rune
	bottomLeftCorner  rune
}

var SquareBordersType = BorderType{
	topLeftCorner:     '┌',
	bottomLeftCorner:  '└',
	topRightCorner:    '┐',
	bottomRightCorner: '┘',
	horizontal:        '─',
	vertical:          '│',
}

var RoundBordersType = BorderType{
	topLeftCorner:     '╭',
	bottomLeftCorner:  '╰',
	topRightCorner:    '╮',
	bottomRightCorner: '╯',
	horizontal:        '─',
	vertical:          '│',
}
