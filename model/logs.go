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
	area       tui.Rect
	AutoScroll bool
}

func wrappedLines(lines []string, w int) []string {
	re := regexp.MustCompile("[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))")
	logsAsText := strings.Join(lines, "\n")
	asText := re.ReplaceAllString(logsAsText, "")
	return strings.Split(wrap.String(asText, w), "\n")
}

func NewLogsViewModel(logs []string, area tui.Rect) LogsViewModel {
	// TODO: replace hardcoded padding
	lines := wrappedLines(logs, area.Width)
	scroll := scrollPosition(lines, area)

	return LogsViewModel{
		Lines:      lines,
		Scroll:     scroll,
		area:       area,
		AutoScroll: true,
	}
}

func (m *LogsViewModel) SetLines(lines []string) {
	m.Lines = wrappedLines(lines, m.area.Width)
	if m.AutoScroll {
		m.Scroll = scrollPosition(m.Lines, m.area)
	}
}

func (m *LogsViewModel) Reflow(area tui.Rect) {
	m.area = area
	m.SetLines(m.Logs)
}

func scrollPosition(lines []string, area tui.Rect) int {
	return max(len(lines)-area.Height+10, 0)
}
