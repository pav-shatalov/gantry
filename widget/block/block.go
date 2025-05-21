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

func (b *Block) Title(title string) *Block {
	b.title = title
	return b
}

func (b *Block) Render(screen tcell.Screen, area geometry.Rect) {
	b.renderLeftSide(screen, area)
	b.renderTopSide(screen, area)
	b.renderRightSide(screen, area)
	b.renderBottomSide(screen, area)

	b.renderTopLeftCorner(screen, area)
	b.renderTopRightCorner(screen, area)
	b.renderBottomRightCorner(screen, area)
	b.renderBottomLeftCorner(screen, area)

	b.renderTitle(screen, area)
}

func (b *Block) renderLeftSide(screen tcell.Screen, area geometry.Rect) {
	col := area.X
	row := area.Y + 1
	borders := RoundBorders
	r := borders.right
	for range area.Height - 2 {
		screen.SetContent(col, row, r, []rune{}, tcell.StyleDefault)
		row++
	}
}

func (b *Block) renderTopSide(screen tcell.Screen, area geometry.Rect) {
	col := area.X + 1
	row := area.Y
	borders := RoundBorders
	r := borders.top
	for range area.Width - 2 {
		screen.SetContent(col, row, r, []rune{}, tcell.StyleDefault)
		col++
	}
}

func (b *Block) renderRightSide(screen tcell.Screen, area geometry.Rect) {
	col := area.X + area.Width - 1
	row := area.Y + 1
	borders := RoundBorders
	r := borders.right

	for range area.Height - 2 {
		screen.SetContent(col, row, r, []rune{}, tcell.StyleDefault)
		row++
	}
}

func (b *Block) renderBottomSide(screen tcell.Screen, area geometry.Rect) {
	col := area.X + 1
	row := area.Y + area.Height - 1
	borders := RoundBorders
	r := borders.bottom

	for range area.Width - 2 {
		screen.SetContent(col, row, r, []rune{}, tcell.StyleDefault)
		col++
	}
}

func (b *Block) renderTopLeftCorner(screen tcell.Screen, area geometry.Rect) {
	screen.SetContent(area.X, area.Y, RoundBorders.topLeft, []rune{}, tcell.StyleDefault)
}

func (b *Block) renderTopRightCorner(screen tcell.Screen, area geometry.Rect) {
	screen.SetContent(area.X+area.Width-1, area.Y, RoundBorders.topRight, []rune{}, tcell.StyleDefault)
}

func (b *Block) renderBottomRightCorner(screen tcell.Screen, area geometry.Rect) {
	screen.SetContent(area.X+area.Width-1, area.Y+area.Height-1, RoundBorders.bottomRight, []rune{}, tcell.StyleDefault)
}

func (b *Block) renderBottomLeftCorner(screen tcell.Screen, area geometry.Rect) {
	screen.SetContent(area.X, area.Y+area.Height-1, RoundBorders.bottomLeft, []rune{}, tcell.StyleDefault)
}

func (b *Block) renderTitle(screen tcell.Screen, area geometry.Rect) {
	if len(b.title) == 0 {
		return
	}
	col := area.X + 2
	row := area.Y
	title := " " + b.title + " "
	for _, c := range title {
		screen.SetContent(col, row, c, []rune{}, tcell.StyleDefault)
		col++
	}
}
