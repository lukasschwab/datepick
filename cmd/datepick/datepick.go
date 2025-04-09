package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/lukasschwab/datepick"
)

func main() {
	options := &datepick.Options{}
	ctx := kong.Parse(
		options,
		kong.Description("A date picker for the terminal."),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: false,
		}),
		kong.Vars{
			"defaultHeight":           "0",
			"defaultWidth":            "0",
			"defaultAlign":            "left",
			"defaultBorder":           "none",
			"defaultBorderForeground": "",
			"defaultBorderBackground": "",
			"defaultBackground":       "",
			"defaultForeground":       "",
			"defaultMargin":           "0 0",
			"defaultPadding":          "0 0",
			"defaultUnderline":        "false",
			"defaultBold":             "false",
			"defaultFaint":            "false",
			"defaultItalic":           "false",
			"defaultStrikethrough":    "false",
		},
	)
	if err := ctx.Run(); err != nil {
		if errors.Is(err, datepick.ErrAborted) {
			os.Exit(130)
		}
		fmt.Println(err)
		os.Exit(1)
	}
}
