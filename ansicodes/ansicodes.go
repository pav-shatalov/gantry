package ansicodes

import "fmt"

const Escape = "\033"
const CSI = Escape + "["

const Reset = CSI + "0m"
const Bold = CSI + "1m"
const Underline = CSI + "4m"
const Reversed = CSI + "7m"

// Text colors
const Black = CSI + "30m"
const Red = CSI + "31m"
const Green = CSI + "32m"
const Yellow = CSI + "33m"
const Blue = CSI + "34m"
const Magenta = CSI + "35m"
const Cyan = CSI + "36m"
const White = CSI + "37m"

// Background colors
const BgBlack = CSI + "40m"
const BgRed = CSI + "41m"
const BgGreen = CSI + "42m"
const BgYellow = CSI + "43m"
const BgBlue = CSI + "44m"
const BgMagenta = CSI + "45m"
const BgCyan = CSI + "46m"
const BgWhite = CSI + "47m"

// Cursor movement
func CursorUp(n int) string {
	return CSI + fmt.Sprintf("%dA", n)
}
func CursorDown(n int) string {
	return CSI + fmt.Sprintf("%dB", n)
}
func CursorForward(n int) string {
	return CSI + fmt.Sprintf("%dC", n)
}
func CursorBackward(n int) string {
	return CSI + fmt.Sprintf("%dD", n)
}

func CursorMoveAt(row int, col int) string {
	return CSI + fmt.Sprintf("%d;%dH", row, col)
}

const CursorHide = CSI + "?25l"
const CursorShow = CSI + "?25h"

// TrueColor support
func SetForegroundColor(r, g, b int) string {
	return CSI + fmt.Sprintf("38;2;%d;%d;%dm", r, g, b)
}
func SetBackgroundColor(r, g, b int) string {
	return CSI + fmt.Sprintf("48;2;%d;%d;%dm", r, g, b)
}

const ClearScreen = CSI + "2J"
const ClearLine = CSI + "2K"

const EnterAlternateScreen = CSI + "?1049h"
const ExitAlternateScreen = CSI + "?1049l"
