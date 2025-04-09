package datepick

import (
	"time"

	"github.com/charmbracelet/gum/style"
)

// Options are the customization options for the date.
type Options struct {
	Prompt          string        `help:"Prompt to display" default:"> " env:"DATEPICK_PROMPT"`
	PromptStyle     style.Styles  `embed:"" prefix:"prompt." envprefix:"DATEPICK_PROMPT_"`
	CursorTextStyle style.Styles  `embed:"" prefix:"cursor." set:"defaultForeground=212" set:"defaultUnderline=true" envprefix:"DATEPICK_CURSOR_"` //nolint:staticcheck
	Value           string        `help:"Initial value in ISO 8601 format, e.g. 2023-11-28" default:""`
	Header          string        `help:"Header value" default:"" env:"DATEPICK_HEADER"`
	HeaderStyle     style.Styles  `embed:"" prefix:"header." set:"defaultForeground=240" envprefix:"DATEPICK_HEADER_"`
	Timeout         time.Duration `help:"Timeout until input aborts" default:"0" env:"DATEPICK_TIMEOUT"`
}
