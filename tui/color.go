package tui

type Color uint64

const (
	ColorReset Color = iota
	ColorBlack
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite

	// Dimmed|bright|dull whatever they call it. I prefer to call them Dimmed
	ColorBlackDimmed
	ColorRedDimmed
	ColorGreenDimmed
	ColorYellowDimmed
	ColorBlueDimmed
	ColorMagentaDimmed
	ColorCyanDimmed
	ColorWhiteDimmed
)
