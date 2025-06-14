package tui

type Block struct {
	title       string
	borders     Borders
	borderType  BorderType
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
		borderType:    SquareBordersType,
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

func (b *Block) BorderType(t BorderType) {
	b.borderType = t
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
	if b.borders.has(LeftBorder) {
		b.renderLeftSide(buf, area)
	}
	if b.borders.has(TopBorder) {
		b.renderTopSide(buf, area)
	}
	if b.borders.has(RightBorder) {
		b.renderRightSide(buf, area)
	}
	if b.borders.has(BottomBorder) {
		b.renderBottomSide(buf, area)
	}

	if b.borders.has(TopBorder) && b.borders.has(LeftBorder) {
		b.renderTopLeftCorner(buf, area)
	}
	if b.borders.has(TopBorder) && b.borders.has(RightBorder) {
		b.renderTopRightCorner(buf, area)
	}
	if b.borders.has(BottomBorder) && b.borders.has(RightBorder) {
		b.renderBottomRightCorner(buf, area)
	}
	if b.borders.has(BottomBorder) && b.borders.has(LeftBorder) {
		b.renderBottomLeftCorner(buf, area)
	}

	b.renderTitle(buf, area)
}

func (b *Block) renderLeftSide(buf *OutputBuffer, area Rect) {
	col := area.Col
	row := area.Row
	r := b.borderType.right
	for range area.Height {
		buf.SetContent(col, row, r, b.borderStyle)
		row++
	}
}

func (b *Block) renderTopSide(buf *OutputBuffer, area Rect) {
	col := area.Col
	row := area.Row
	r := b.borderType.top
	for range area.Width {
		buf.SetContent(col, row, r, b.borderStyle)
		col++
	}
}

func (b *Block) renderRightSide(buf *OutputBuffer, area Rect) {
	col := area.Col + area.Width - 1
	row := area.Row
	r := b.borderType.right

	for range area.Height {
		buf.SetContent(col, row, r, b.borderStyle)
		row++
	}
}

func (b *Block) renderBottomSide(buf *OutputBuffer, area Rect) {
	col := area.Col
	row := area.Row + area.Height - 1
	r := b.borderType.bottom

	for range area.Width {
		buf.SetContent(col, row, r, b.borderStyle)
		col++
	}
}

func (b *Block) renderTopLeftCorner(buf *OutputBuffer, area Rect) {
	buf.SetContent(
		area.Col,
		area.Row,
		b.borderType.topLeft,
		b.borderStyle,
	)
}

func (b *Block) renderTopRightCorner(buf *OutputBuffer, area Rect) {
	buf.SetContent(
		area.Col+area.Width-1,
		area.Row,
		b.borderType.topRight,
		b.borderStyle,
	)
}

func (b *Block) renderBottomRightCorner(buf *OutputBuffer, area Rect) {
	buf.SetContent(
		area.Col+area.Width-1,
		area.Row+area.Height-1,
		b.borderType.bottomRight,
		b.borderStyle,
	)
}

func (b *Block) renderBottomLeftCorner(buf *OutputBuffer, area Rect) {
	buf.SetContent(
		area.Col,
		area.Row+area.Height-1,
		b.borderType.bottomLeft,
		b.borderStyle,
	)
}

func (b *Block) renderTitle(buf *OutputBuffer, area Rect) {
	if len(b.title) == 0 {
		return
	}
	col := area.Col + 1
	row := area.Row
	title := " " + b.title + " "
	for _, c := range title {
		buf.SetContent(col, row, c, b.titleStyle)
		col++
	}
}

func (b *Block) InnerArea(area Rect) Rect {
	lbWidth := 0
	rbWidth := 0
	if b.borders.has(LeftBorder) {
		lbWidth = 1
	}
	if b.borders.has(RightBorder) {
		rbWidth = 1
	}
	tbHeight := 0
	bbHeight := 0
	if b.borders.has(TopBorder) {
		tbHeight = 1
	}
	if b.borders.has(BottomBorder) {
		bbHeight = 1
	}
	return NewRect(
		area.Col+lbWidth+b.paddingLeft,
		area.Row+tbHeight+b.paddingTop,
		area.Width-lbWidth+rbWidth-b.paddingLeft+b.paddingRight,
		area.Height-tbHeight+bbHeight-b.paddingTop+b.paddingBottom,
	)
}
