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
	title       string
	borders     borders
	titleStyle  tcell.Style
	borderStyle tcell.Style
}

func New() Block {
	return Block{
		title:       "",
		borders:     RoundBorders,
		titleStyle:  tcell.StyleDefault,
		borderStyle: tcell.StyleDefault,
	}
}

func (b *Block) Title(title string) *Block {
	b.title = title
	return b
}

func (b *Block) BorderStyle(style tcell.Style) *Block {
	b.borderStyle = style
	return b
}

func (b *Block) TitleStyle(style tcell.Style) *Block {
	b.titleStyle = style
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
	r := b.borders.right
	for range area.Height - 2 {
		screen.SetContent(col, row, r, []rune{}, b.borderStyle)
		row++
	}
}

func (b *Block) renderTopSide(screen tcell.Screen, area geometry.Rect) {
	col := area.X + 1
	row := area.Y
	r := b.borders.top
	for range area.Width - 2 {
		screen.SetContent(col, row, r, []rune{}, b.borderStyle)
		col++
	}
}

func (b *Block) renderRightSide(screen tcell.Screen, area geometry.Rect) {
	col := area.X + area.Width - 1
	row := area.Y + 1
	r := b.borders.right

	for range area.Height - 2 {
		screen.SetContent(col, row, r, []rune{}, b.borderStyle)
		row++
	}
}

func (b *Block) renderBottomSide(screen tcell.Screen, area geometry.Rect) {
	col := area.X + 1
	row := area.Y + area.Height - 1
	r := b.borders.bottom

	for range area.Width - 2 {
		screen.SetContent(col, row, r, []rune{}, b.borderStyle)
		col++
	}
}

func (b *Block) renderTopLeftCorner(screen tcell.Screen, area geometry.Rect) {
	screen.SetContent(
		area.X,
		area.Y,
		b.borders.topLeft,
		[]rune{},
		b.borderStyle,
	)
}

func (b *Block) renderTopRightCorner(screen tcell.Screen, area geometry.Rect) {
	screen.SetContent(
		area.X+area.Width-1,
		area.Y,
		b.borders.topRight,
		[]rune{},
		b.borderStyle,
	)
}

func (b *Block) renderBottomRightCorner(screen tcell.Screen, area geometry.Rect) {
	screen.SetContent(
		area.X+area.Width-1,
		area.Y+area.Height-1,
		b.borders.bottomRight,
		[]rune{},
		b.borderStyle,
	)
}

func (b *Block) renderBottomLeftCorner(screen tcell.Screen, area geometry.Rect) {
	screen.SetContent(
		area.X,
		area.Y+area.Height-1,
		b.borders.bottomLeft,
		[]rune{},
		b.borderStyle,
	)
}

func (b *Block) renderTitle(screen tcell.Screen, area geometry.Rect) {
	if len(b.title) == 0 {
		return
	}
	col := area.X + 2
	row := area.Y
	title := " " + b.title + " "
	for _, c := range title {
		screen.SetContent(col, row, c, []rune{}, b.titleStyle)
		col++
	}
}

func (b *Block) InnerArea(area geometry.Rect) geometry.Rect {
	return geometry.Rect{
		X:      area.X + 1,
		Y:      area.Y + 1,
		Width:  area.Width - 2,
		Height: area.Height - 2,
	}
}
