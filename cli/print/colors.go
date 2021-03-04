package print

import (
	"github.com/fatih/color"
)

var (
	// Warning printer
	Warning = color.New(color.Bold, color.FgYellow).SprintFunc()
	// Warningf printer with string formatter
	Warningf = color.New(color.Bold, color.FgYellow).SprintfFunc()
	// Critical printer
	Critical = color.New(color.Bold, color.FgRed).SprintFunc()
	// Criticalf printer with string formatter
	Criticalf = color.New(color.Bold, color.FgRed).SprintfFunc()
	// Notice printer
	Notice = color.New(color.Bold, color.FgGreen).SprintFunc()
	// Noticef printer with string formatter
	Noticef = color.New(color.Bold, color.FgGreen).SprintfFunc()
	// Bold printer
	Bold = color.New(color.Bold).SprintFunc()
	// Boldf printer with string formatter
	Boldf = color.New(color.Bold).SprintfFunc()
)
