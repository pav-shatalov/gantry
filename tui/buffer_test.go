package tui_test

import (
	"gantry/geometry"
	"gantry/tui"
	"testing"
)

func TestBuffer(t *testing.T) {
	buf := tui.NewBuffer(10, 1)
	actualW := buf.Width()
	expectedW := 10
	if actualW != expectedW {
		t.Errorf("Actual width: %d, Expected widht: %d", actualW, expectedW)
	}
}

func TestMerging(t *testing.T) {
	buf := tui.NewBuffer(100, 20)
	subBuf := tui.NewBuffer(10, 10)
	subBuf.SetContent(0, 0, '1', tui.StyleDefault)

	buf.FillFrom(&subBuf, geometry.Rect{Col: 0, Row: 0, Width: 10, Height: 10})

	cell := buf.GetCell(0, 0)

	if cell.R != '1' {
		t.Errorf("Expected %s, got %s", string('1'), string(cell.R))
	}
}
