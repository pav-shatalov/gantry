package tui

import (
	"github.com/gdamore/tcell/v2"
)

type Block struct {
	title       string
	borders     Borders
	titleStyle  tcell.Style
	borderStyle tcell.Style
}

func NewBlock() Block {
	return Block{
		title:       "",
		borders:     RoundBorders,
		titleStyle:  tcell.StyleDefault,
		borderStyle: tcell.StyleDefault,
	}
}

func (b Block) Title(title string) Block {
	b.title = title
	return b
}

func (b Block) BorderStyle(style tcell.Style) Block {
	b.borderStyle = style
	return b
}

func (b Block) TitleStyle(style tcell.Style) Block {
	b.titleStyle = style
	return b
}

func (b *Block) Render(screen tcell.Screen, area Rect) {
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

func (b *Block) renderLeftSide(screen tcell.Screen, area Rect) {
	col := area.Col
	row := area.Row + 1
	r := b.borders.right
	for range area.Height - 2 {
		screen.SetContent(col, row, r, []rune{}, b.borderStyle)
		row++
	}
}

func (b *Block) renderTopSide(screen tcell.Screen, area Rect) {
	col := area.Col + 1
	row := area.Row
	r := b.borders.top
	for range area.Width - 2 {
		screen.SetContent(col, row, r, []rune{}, b.borderStyle)
		col++
	}
}

func (b *Block) renderRightSide(screen tcell.Screen, area Rect) {
	col := area.Col + area.Width - 1
	row := area.Row + 1
	r := b.borders.right

	for range area.Height - 2 {
		screen.SetContent(col, row, r, []rune{}, b.borderStyle)
		row++
	}
}

func (b *Block) renderBottomSide(screen tcell.Screen, area Rect) {
	col := area.Col + 1
	row := area.Row + area.Height - 1
	r := b.borders.bottom

	for range area.Width - 2 {
		screen.SetContent(col, row, r, []rune{}, b.borderStyle)
		col++
	}
}

func (b *Block) renderTopLeftCorner(screen tcell.Screen, area Rect) {
	screen.SetContent(
		area.Col,
		area.Row,
		b.borders.topLeft,
		[]rune{},
		b.borderStyle,
	)
}

func (b *Block) renderTopRightCorner(screen tcell.Screen, area Rect) {
	screen.SetContent(
		area.Col+area.Width-1,
		area.Row,
		b.borders.topRight,
		[]rune{},
		b.borderStyle,
	)
}

func (b *Block) renderBottomRightCorner(screen tcell.Screen, area Rect) {
	screen.SetContent(
		area.Col+area.Width-1,
		area.Row+area.Height-1,
		b.borders.bottomRight,
		[]rune{},
		b.borderStyle,
	)
}

func (b *Block) renderBottomLeftCorner(screen tcell.Screen, area Rect) {
	screen.SetContent(
		area.Col,
		area.Row+area.Height-1,
		b.borders.bottomLeft,
		[]rune{},
		b.borderStyle,
	)
}

func (b *Block) renderTitle(screen tcell.Screen, area Rect) {
	if len(b.title) == 0 {
		return
	}
	col := area.Col + 2
	row := area.Row
	title := " " + b.title + " "
	for _, c := range title {
		screen.SetContent(col, row, c, []rune{}, b.titleStyle)
		col++
	}
}

func (b *Block) InnerArea(area Rect) Rect {
	return NewRect(area.Col+1, area.Row+1, area.Width-2, area.Height-2)
}
