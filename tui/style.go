package tui

type Style struct {
	bg Color
	fg Color
}

var StyleDefault Style

func (s Style) Fg(c Color) Style {
	s.fg = c
	return s
}

func (s Style) Bg(c Color) Style {
	s.bg = c
	return s
}
