package block

import (
	"gantry/geometry"

	"github.com/gdamore/tcell/v2"
)

type borders struct {
	topLeft     rune
	top         rune
	topRight    rune
	right       rune
	bottomRight rune
	bottom      rune
	bottomLeft  rune
	left        rune
}

var SquareBorders = borders{
	topLeft:     '┌',
	top:         '─',
	topRight:    '┐',
	right:       '│',
	bottomRight: '┘',
	bottom:      '─',
	bottomLeft:  '└',
	left:        '│',
}

var RoundBorders = borders{
	topLeft:     '╭',
	top:         '─',
	topRight:    '╮',
	right:       '│',
	bottomRight: '╯',
	bottom:      '─',
	bottomLeft:  '╰',
	left:        '│',
}

type Block struct {
	title string
}

func New() Block {
	return Block{title: ""}
}

func (b *Block) Render(screen tcell.Screen, area geometry.Rect) {
	col := area.X
	row := area.Y
	borders := RoundBorders
	for range area.Height {
		isFirstRow := row == area.Y
		isLastRow := row == area.Y+area.Height-1
		for range area.Width {
			r := ' '
			isLeftEdge := col == area.X
			isRightEdge := col == area.X+area.Width-1
			if isFirstRow {
				r = borders.top
				if isLeftEdge {
					r = borders.topLeft
				}
				if isRightEdge {
					r = borders.topRight
				}
			}

			if isLastRow {
				r = borders.bottom
				if isLeftEdge {
					r = borders.bottomLeft
				}
				if isRightEdge {
					r = borders.bottomRight
				}
			}

			if isLeftEdge && !(isFirstRow || isLastRow) {
				r = borders.left
			}

			if isRightEdge && !(isFirstRow || isLastRow) {
				r = borders.right
			}

			screen.SetContent(col, row, r, []rune{}, tcell.StyleDefault)
			col++
		}
		row++
		col = area.X
	}
}
