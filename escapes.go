package escapes

import (
	"encoding/base64"
	"strconv"
)

// ConsoleDim represents the dimensions of a console in rows and columns.
type ConsoleDim struct {
	Rows int
	Cols int
}

// Common fragments of escape sequences
const (
	Esc = "\u001B["
	Osc = "\u001B]"
	Bel = "\u0007"
)

// Common ANSI escapes sequences. These should be used when the desired action
// is only needed once; otherwise, use the functions (e.g. moving a cursor
// several lines/columns). See: https://docs.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences
const (
	CursorUp       = Esc + "A"
	CursorDown     = Esc + "B"
	CursorForward  = Esc + "C"
	CursorBackward = Esc + "D"
	CursorNextLine = Esc + "E"
	CursorPrevLine = Esc + "F"
	CursorLeft     = Esc + "G"
	CursorTop      = Esc + "d"
	CursorTopLeft  = Esc + "H"
	CursorSave     = Esc + "s"
	CursorRestore  = Esc + "u"

	CursorBlinkEnable  = Esc + "?12h"
	CursorBlinkDisable = Esc + "?12I"
	CursorShow         = Esc + "?25h"
	CursorHide         = Esc + "?25l"

	ScrollUp   = Esc + "S"
	ScrollDown = Esc + "T"

	TextInsertChar = Esc + "@"
	TextDeleteChar = Esc + "P"
	TextEraseChar  = Esc + "X"
	TextInsertLine = Esc + "L"
	TextDeleteLine = Esc + "M"

	EraseRight  = Esc + "K"
	EraseLeft   = Esc + "1K"
	EraseLine   = Esc + "2K"
	EraseDown   = Esc + "J"
	EraseUp     = Esc + "1J"
	EraseScreen = Esc + "2J"

	TextColorBlack         = Esc + "30m"
	TextColorRed           = Esc + "31m"
	TextColorGreen         = Esc + "32m"
	TextColorYellow        = Esc + "33m"
	TextColorBlue          = Esc + "34m"
	TextColorMagenta       = Esc + "35m"
	TextColorCyan          = Esc + "36m"
	TextColorWhite         = Esc + "37m"
	TextColorBrightBlack   = Esc + "30;1m"
	TextColorBrightRed     = Esc + "31;1m"
	TextColorBrightGreen   = Esc + "32;1m"
	TextColorBrightYellow  = Esc + "33;1m"
	TextColorBrightBlue    = Esc + "34;1m"
	TextColorBrightMagenta = Esc + "35;1m"
	TextColorBrightCyan    = Esc + "36;1m"
	TextColorBrightWhite   = Esc + "37;1m"

	BackgroundColorBlack         = Esc + "40m"
	BackgroundColorRed           = Esc + "41m"
	BackgroundColorGreen         = Esc + "42m"
	BackgroundColorYellow        = Esc + "43m"
	BackgroundColorBlue          = Esc + "44m"
	BackgroundColorMagenta       = Esc + "45m"
	BackgroundColorCyan          = Esc + "46m"
	BackgroundColorWhite         = Esc + "47m"
	BackgroundColorBrightBlack   = Esc + "40;1m"
	BackgroundColorBrightRed     = Esc + "41;1m"
	BackgroundColorBrightGreen   = Esc + "42;1m"
	BackgroundColorBrightYellow  = Esc + "43;1m"
	BackgroundColorBrightBlue    = Esc + "44;1m"
	BackgroundColorBrightMagenta = Esc + "45;1m"
	BackgroundColorBrightCyan    = Esc + "46;1m"
	BackgroundColorBrightWhite   = Esc + "47;1m"

	ColorReset = Esc + "0m"

	ClearScreen = "\u001Bc"
)

// CursorPosX returns an escape sequence to move the cursor to an x-coordinate
// (column) at the current y-coordinate (row), where 0 is the leftmost.
func CursorPosX(x int) string {
	return Esc + strconv.Itoa(x+1) + "G"
}

// CursorPosY returns an escape sequence to move the cursor to an y-coordinate
// (row) at the current x-coordinate (column), where 0 is the topmost.
func CursorPosY(y int) string {
	return Esc + strconv.Itoa(y+1) + "d"
}

// CursorPos returns an escape sequence to move the cursor to a coordinate pair,
// where (0, 0) is the origin (top-left corner).
func CursorPos(x, y int) string {
	return Esc + strconv.Itoa(y+1) + ";" + strconv.Itoa(x+1) + "H"
}

// CursorMove returns an escape sequence to move the cursor relative to its
// current position.
func CursorMove(x, y int) string {
	var s string
	if x < 0 {
		s = Esc + strconv.Itoa(-x) + "D"
	} else if x > 0 {
		s = Esc + strconv.Itoa(x) + "C"
	}
	if y < 0 {
		s += Esc + strconv.Itoa(-y) + "A"
	} else if y > 0 {
		s += Esc + strconv.Itoa(y) + "B"
	}
	return s
}

// Scroll returns an escape sequence to scroll the current window. A positive
// number of lines indicates scrolling up, while a negative number of lines
// indicates scrolling down.
func Scroll(n int) string {
	if n > 0 {
		return Esc + strconv.Itoa(n) + "S"
	} else if n < 0 {
		return Esc + strconv.Itoa(-n) + "T"
	} else {
		return ""
	}
}

// TextInsertChars returns an escape sequence to insert spaces to the right of,
// and including, the current cursor position, shifting existing characters to
// the right.
func TextInsertChars(n int) string {
	return Esc + strconv.Itoa(n) + "@"
}

// TextDeleteChars returns an escape sequence to delete characters to the right
// of, and including, the current cursor position, shifting existing characters
// to the left.
func TextDeleteChars(n int) string {
	return Esc + strconv.Itoa(n) + "P"
}

// TextEraseChars returns an escape sequence to insert spaces to the right of,
// and including, the current cursor position, overwriting existing characters
// to the right.
func TextEraseChars(n int) string {
	return Esc + strconv.Itoa(n) + "X"
}

// TextInsertLines returns an escape sequence to insert blank lines below, and
// including the current cursor row, shifting existing lines downwards.
func TextInsertLines(n int) string {
	return Esc + strconv.Itoa(n) + "L"
}

// TextDeleteLines returns an escape sequence to delete the lines below, and
// including, the current cursor row.
func TextDeleteLines(n int) string {
	return Esc + strconv.Itoa(n) + "M"
}

// Link returns an escape sequence to represent linked text.
func Link(url, text string) string {
	return Osc + "8;;" + url + Bel + text + Osc + "8;;" + Bel
}

// Image returns an escape sequence to display an image, preserving the original
// height and width.
func Image(img []byte) string {
	return ImageWidthHeight(img, 0, 0, true)
}

// ImageWidthHeight returns an escape sequence to display an image.
func ImageWidthHeight(img []byte, height, width int, preserveAspectRatio bool) string {
	s := Osc + "1337;File=inline=1"
	if height > 0 {
		s += ";width=" + strconv.Itoa(height)
	}
	if width > 0 {
		s += ";height=" + strconv.Itoa(width)
	}
	if !preserveAspectRatio {
		s += ";preserveAspectRatio=0"
	}

	return s + ":" + base64.StdEncoding.EncodeToString(img) + Bel
}

// SetCwd returns an escape sequence to set the current working directory.
func SetCwd(dir string) string {
	return Osc + "50;CurrentDir=" + dir + Bel
}
