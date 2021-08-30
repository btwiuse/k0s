package hilighter

import (
	"regexp"
	"strings"
)

func bgColor(code string, str string) string {
	var colorized string

	// Strip all other bg color codes and don't extend bg color over
	// newlines
	str = newline.ReplaceAllString(
		plainBg(str),
		"\x1b["+Colors["on_default"]+"m\n\x1b["+Colors[code]+"m",
	)

	// Wrap whole thing with specified color code
	colorized = "\x1b[" + Colors[code] + "m" + str +
		"\x1b[" + Colors["on_default"] + "m"

	// Remove color codes, if the line only contains color codes
	return onlyCodes.ReplaceAllString(colorized, "$1$4")
}

func colorize(clr string, str string) string {
	if IsDisabled() {
		// Return the string w/o any color codes
		return Plain(str)
	}

	// Call the appropriate function
	if strings.HasPrefix(clr, "on_") {
		return bgColor(clr, str)
	}
	return fgColor(clr, str)
}

func fgColor(code string, str string) string {
	var colorized string

	// Strip all other fg color codes and don't extend fg color over
	// newlines
	str = newline.ReplaceAllString(
		plainFg(str),
		"\x1b["+Colors["default"]+"m\n\x1b["+Colors[code]+"m",
	)

	// Wrap whole thing with specified color code
	colorized = "\x1b[" + Colors[code] + "m" + str +
		"\x1b[" + Colors["default"] + "m"

	// Remove color codes, if the line only contains color codes
	return onlyCodes.ReplaceAllString(colorized, "$1$4")
}

func modify(mode string, str string) string {
	if IsDisabled() {
		// Return the string w/o any color codes
		return Plain(str)
	}

	var hasKey bool
	var modified string
	var off string
	var opposite string
	var r *regexp.Regexp
	var rm string

	// Reverse mode
	opposite = "no_" + mode
	if strings.HasPrefix(opposite, "no_no_") {
		opposite = doubleno.ReplaceAllString(opposite, "")
	}

	// Store specified mode code for removal
	rm = Modes[mode]

	// Determine the off color code, if it exists
	off = ""
	if _, hasKey = Modes[opposite]; hasKey {
		// Store opposite code for removal
		rm += "|" + Modes[opposite]

		// Save opposite mode code sequence if it starts with "no_"
		if strings.HasPrefix(opposite, "no_") {
			off = "\x1b[" + Modes[opposite] + "m"
		}
	}

	// Remove other occurrences of specified mode and opposite
	r = regexp.MustCompile(`\x1b\[(` + rm + ")m")
	str = r.ReplaceAllString(str, "")

	// Wrap the whole thing with specified color code
	modified = "\x1b[" + Modes[mode] + "m" + str + off

	// Remove color codes, if the line only contains color codes
	return onlyCodes.ReplaceAllString(modified, "$1$4")
}

func plainBg(str string) string {
	return bgCodes.ReplaceAllString(str, "")
}

func plainFg(str string) string {
	return fgCodes.ReplaceAllString(str, "")
}

func rainbowColors() []int {
	// Don't include black, white, light_black, and light_white
	return []int{31, 32, 33, 34, 35, 36, 91, 92, 93, 94, 95, 96}
}
