package model

import (
	"gantry/tui"
	"regexp"
	"strings"

	"github.com/muesli/reflow/wrap"
)

type LogsViewModel struct {
	Logs       []string
	Lines      []string
	Scroll     int
	Area       tui.Rect
	AutoScroll bool
}

func wrappedLines(lines []string, w int) []string {
	asString := strings.Join(lines, "\n")
	re := regexp.MustCompile("[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))")
	withoutAnsi := re.ReplaceAllString(asString, "")
	wrapped := wrap.String(withoutAnsi, w-4) // borders + padding
	return strings.Split(wrapped, "\n")
}

func NewLogsViewModel(logs []string, area tui.Rect) LogsViewModel {
	lines := wrappedLines(logs, area.Width)
	scroll := scrollPosition(lines, area)

	return LogsViewModel{
		Lines:      lines,
		Scroll:     scroll,
		Area:       area,
		AutoScroll: true,
	}
}

func (m *LogsViewModel) SetLines(lines []string) {
	m.Lines = wrappedLines(lines, m.Area.Width)
	if m.AutoScroll {
		m.Scroll = scrollPosition(m.Lines, m.Area)
	}
}

func (m *LogsViewModel) Reflow(area tui.Rect) {
	m.Area = area
	m.SetLines(m.Logs)
}

func scrollPosition(lines []string, area tui.Rect) int {
	return max(len(lines)-area.Height+10, 0)
}
