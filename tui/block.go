package tui

type Block struct {
	title       string
	borders     Borders
	titleStyle  Style
	borderStyle Style

	// Padding
	paddingLeft   int
	paddingRight  int
	paddingTop    int
	paddingBottom int
}

func NewBlock() Block {
	return Block{
		title:         "",
		borders:       NoBorders,
		titleStyle:    StyleDefault,
		borderStyle:   StyleDefault,
		paddingTop:    0,
		paddingRight:  0,
		paddingBottom: 0,
		paddingLeft:   0,
	}
}

func (b *Block) Title(title string) {
	b.title = title
}

func (b *Block) Borders(borders Borders) {
	b.borders = borders
}

func (b *Block) BorderStyle(style Style) {
	b.borderStyle = style
}

func (b *Block) TitleStyle(style Style) {
	b.titleStyle = style
}

func (b *Block) Padding(t int, r int, bot int, l int) {
	b.paddingTop = t
	b.paddingRight = r
	b.paddingBottom = bot
	b.paddingLeft = l
}

func (b *Block) Render(buf *OutputBuffer, area Rect) {
	b.renderLeftSide(buf, area)
	b.renderTopSide(buf, area)
	b.renderRightSide(buf, area)
	b.renderBottomSide(buf, area)

	b.renderTopLeftCorner(buf, area)
	b.renderTopRightCorner(buf, area)
	b.renderBottomRightCorner(buf, area)
	b.renderBottomLeftCorner(buf, area)

	b.renderTitle(buf, area)
}

func (b *Block) renderLeftSide(buf *OutputBuffer, area Rect) {
	col := area.Col
	row := area.Row + 1
	r := b.borders.right
	for range area.Height - 2 {
		buf.SetContent(col, row, r, b.borderStyle)
		row++
	}
}

func (b *Block) renderTopSide(buf *OutputBuffer, area Rect) {
	col := area.Col + 1
	row := area.Row
	r := b.borders.top
	for range area.Width - 2 {
		buf.SetContent(col, row, r, b.borderStyle)
		col++
	}
}

func (b *Block) renderRightSide(buf *OutputBuffer, area Rect) {
	col := area.Col + area.Width - 1
	row := area.Row + 1
	r := b.borders.right

	for range area.Height - 2 {
		buf.SetContent(col, row, r, b.borderStyle)
		row++
	}
}

func (b *Block) renderBottomSide(buf *OutputBuffer, area Rect) {
	col := area.Col + 1
	row := area.Row + area.Height - 1
	r := b.borders.bottom

	for range area.Width - 2 {
		buf.SetContent(col, row, r, b.borderStyle)
		col++
	}
}

func (b *Block) renderTopLeftCorner(buf *OutputBuffer, area Rect) {
	buf.SetContent(
		area.Col,
		area.Row,
		b.borders.topLeft,
		b.borderStyle,
	)
}

func (b *Block) renderTopRightCorner(buf *OutputBuffer, area Rect) {
	buf.SetContent(
		area.Col+area.Width-1,
		area.Row,
		b.borders.topRight,
		b.borderStyle,
	)
}

func (b *Block) renderBottomRightCorner(buf *OutputBuffer, area Rect) {
	buf.SetContent(
		area.Col+area.Width-1,
		area.Row+area.Height-1,
		b.borders.bottomRight,
		b.borderStyle,
	)
}

func (b *Block) renderBottomLeftCorner(buf *OutputBuffer, area Rect) {
	buf.SetContent(
		area.Col,
		area.Row+area.Height-1,
		b.borders.bottomLeft,
		b.borderStyle,
	)
}

func (b *Block) renderTitle(buf *OutputBuffer, area Rect) {
	if len(b.title) == 0 {
		return
	}
	col := area.Col + 2
	row := area.Row
	title := " " + b.title + " "
	for _, c := range title {
		buf.SetContent(col, row, c, b.titleStyle)
		col++
	}
}

func (b *Block) InnerArea(area Rect) Rect {
	bWidth := b.borders.width
	bHeight := b.borders.height
	return NewRect(
		area.Col+bWidth+b.paddingLeft,
		area.Row+bHeight+b.paddingTop,
		area.Width-bWidth*2-b.paddingLeft+b.paddingRight,
		area.Height-bHeight*2-b.paddingTop+b.paddingBottom,
	)
}
