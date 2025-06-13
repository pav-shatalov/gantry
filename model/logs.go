package model

import (
	"gantry/geometry"
	"regexp"
	"strings"

	"github.com/muesli/reflow/wrap"
)

type LogsViewModel struct {
	Lines  []string
	Scroll int
	area   geometry.Rect
}

func wrappedLines(lines []string, w int) []string {
	re := regexp.MustCompile("[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))")
	logsAsText := strings.Join(lines, "\n")
	asText := re.ReplaceAllString(logsAsText, "")
	return strings.Split(wrap.String(asText, w), "\n")
}

func NewLogsViewModel(logs []string, area geometry.Rect) LogsViewModel {
	lines := wrappedLines(logs, area.Width)
	scroll := max(len(lines)-area.Height, 0)

	return LogsViewModel{
		Lines:  lines,
		Scroll: scroll,
		area:   area,
	}
}

func (m *LogsViewModel) SetLines(lines []string) {
	m.Lines = wrappedLines(lines, m.area.Width)
	scroll := max(len(m.Lines)-m.area.Height, 0)
	m.Scroll = scroll
}
