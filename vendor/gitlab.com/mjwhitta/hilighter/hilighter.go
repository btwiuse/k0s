package hilighter

import (
	"image/color"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// ColorToTrueColor will convert the color instance to its truecolor
// representation.
func ColorToTrueColor(c color.Color) string {
	var a uint32
	var b uint32
	var g uint32
	var r uint32

	r, g, b, a = c.RGBA()

	if (a != 0) && (a != 0xffff) {
		r = uint32(float64(r*0xffff) / float64(a))
		g = uint32(float64(g*0xffff) / float64(a))
		b = uint32(float64(b*0xffff) / float64(a))
	}

	r >>= 8
	g >>= 8
	b >>= 8
	a >>= 8

	return RGBAToTrueColor(uint8(r), uint8(g), uint8(b), uint8(a))
}

// ColorToXterm256 will convert the color instance to its xterm256
// representation.
func ColorToXterm256(c color.Color) string {
	var a uint32
	var b uint32
	var g uint32
	var r uint32

	r, g, b, a = c.RGBA()

	if (a != 0) && (a != 0xffff) {
		r = uint32(float64(r*0xffff) / float64(a))
		g = uint32(float64(g*0xffff) / float64(a))
		b = uint32(float64(b*0xffff) / float64(a))
	}

	r >>= 8
	g >>= 8
	b >>= 8
	a >>= 8

	return RGBAToXterm256(uint8(r), uint8(g), uint8(b), uint8(a))
}

// Disable will prevent color codes from being used.
func Disable(b bool) {
	disable = b
}

// Hex will take a hex color code and add the appropriate ANSI color
// codes to the provided string.
func Hex(hex string, str string) string {
	switch strings.ToLower(os.Getenv("COLORTERM")) {
	case "truecolor":
		return colorize(HexToTrueColor(hex), str)
	default:
		return colorize(HexToXterm256(hex), str)
	}
}

// HexToTrueColor will convert hex to the 24-bit RGB values.
func HexToTrueColor(hex string) string {
	var hasKey bool
	if _, hasKey = Colors[hex]; hasKey {
		return hex
	}

	var b int
	var g int
	var hexB uint64
	var hexG uint64
	var hexR uint64
	var matches [][]string
	var r int

	r, g, b = 0, 0, 0
	matches = parseHex.FindAllStringSubmatch(hex, -1)
	for _, match := range matches {
		hexB, _ = strconv.ParseUint(match[3], 16, 8)
		b = int(hexB)

		hexG, _ = strconv.ParseUint(match[2], 16, 8)
		g = int(hexG)

		hexR, _ = strconv.ParseUint(match[1], 16, 8)
		r = int(hexR)
	}

	Colors[hex] = Sprintf("38;2;%d;%d;%d", r, g, b)
	Colors["on_"+hex] = Sprintf("48;2;%d;%d;%d", r, g, b)

	return hex
}

// HexToXterm256 will convert hex to xterm-256 8-bit value
// https://stackoverflow.com/questions/11765623/convert-hex-hex_to_256bitto-closest-x11-color-number
func HexToXterm256(hex string) string {
	var hasKey bool
	if _, hasKey = cachedXterm[hex]; hasKey {
		return cachedXterm[hex]
	}

	var average int
	var b int
	var cb int
	var cg int
	var cidx int
	var clrErr float64
	var cr int
	var g int
	var gidx int
	var grayErr float64
	var gv int
	var hexB uint64
	var hexG uint64
	var hexR uint64
	var i2cv []int
	var ib int
	var ig int
	var ir int
	var matches [][]string
	var r int

	// For simplicity, assume RGB space is perceptually uniform.
	// There are 5 places where one of two outputs needs to be
	// chosen when the input is the exact middle:
	// - The r/g/b channels and the gray value:
	//     - the higher value output is chosen
	// - If the gray and color values have same distance from the
	//   input
	//     - color is chosen

	// Note: This is integer math, NOT unsigned integer

	// Calculate the nearest 0-based color index at 16..231
	r, g, b = 0, 0, 0
	matches = parseHex.FindAllStringSubmatch(hex, -1)
	for _, match := range matches {
		hexB, _ = strconv.ParseUint(match[3], 16, 8)
		b = int(hexB)

		hexG, _ = strconv.ParseUint(match[2], 16, 8)
		g = int(hexG)

		hexR, _ = strconv.ParseUint(match[1], 16, 8)
		r = int(hexR)
	}

	// 0..5 each
	ir = (r - 35) / 40
	if r < 48 {
		ir = 0
	} else if r < 115 {
		ir = 1
	}
	ig = (g - 35) / 40
	if g < 48 {
		ig = 0
	} else if g < 115 {
		ig = 1
	}
	ib = (b - 35) / 40
	if b < 48 {
		ib = 0
	} else if b < 115 {
		ib = 1
	}

	// 0..215 lazy evaluation
	cidx = (36 * ir) + (6 * ig) + ib

	// Calculate the nearest 0-based gray index at 232..255
	average = (r + g + b) / 3

	// 0..23
	gidx = (average - 3) / 10
	if average > 238 {
		gidx = 23
	}

	// Calculate the represented colors back from the index
	i2cv = []int{0, 0x5f, 0x87, 0xaf, 0xd7, 0xff}

	// r/g/b 0..255 each
	cr = i2cv[ir]
	cg = i2cv[ig]
	cb = i2cv[ib]

	// same value for r/g/b 0..255
	gv = (10 * gidx) + 8

	// Return the one which is nearer to the original rgb values
	clrErr = math.Pow(float64(int(cr-r)), 2) +
		math.Pow(float64(int(cg-g)), 2) +
		math.Pow(float64(int(cb-b)), 2)
	grayErr = math.Pow(float64(int(gv-r)), 2) +
		math.Pow(float64(int(gv-g)), 2) +
		math.Pow(float64(int(gv-b)), 2)

	if clrErr <= grayErr {
		cachedXterm[hex] = Sprintf("color_%03d", cidx+16)
	} else {
		cachedXterm[hex] = Sprintf("color_%03d", gidx+232)
	}

	return cachedXterm[hex]
}

// Hilight will add the appropriate ANSI code to the specified
// string.
func Hilight(code string, str string) string {
	var clr string
	var hasKey bool
	var match []string
	var matches [][]string
	var width int

	// Call the appropriate function
	if _, hasKey = Colors[code]; hasKey {
		return colorize(code, str)
	} else if _, hasKey = Modes[code]; hasKey {
		return modify(code, str)
	} else {
		switch code {
		case "on_rainbow":
			return OnRainbow(str)
		case "plain":
			return Plain(str)
		case "rainbow":
			return Rainbow(str)
		default:
			// Check if hex color code
			matches = hexCode.FindAllStringSubmatch(code, -1)
			if len(matches) > 0 {
				match = matches[0]

				switch strings.ToLower(os.Getenv("COLORTERM")) {
				case "truecolor":
					clr = HexToTrueColor(match[2])
				default:
					clr = HexToXterm256(match[2])
				}

				if strings.HasPrefix(code, "on_") {
					clr = "on_" + clr
				}

				return colorize(clr, str)
			}

			// Check if wrap
			matches = wrap.FindAllStringSubmatch(code, -1)
			if len(matches) > 0 {
				match = matches[0]

				// Determine wrap width, default to 80
				width = 80
				if len(match) == 3 && len(match[2]) > 0 {
					width, _ = strconv.Atoi(match[2])
				}

				return Wrap(width, str)
			}

			// Otherwise panic
			// panic(
			//     fmt.Errorf(
			//         "hilighter: invalid color or mode: %s",
			//         code,
			//     ),
			// )

			// Scratch that, just do nothing
			return str
		}
	}
}

// Hilights will add the appropriate ANSI codes to the specified
// string.
func Hilights(codes []string, str string) string {
	// Apply all specified color codes
	for _, code := range codes {
		str = Hilight(code, str)
	}

	return str
}

// IsDisabled will return whether or not color codes are disabled.
// Will always return true if GOOS is windows.
func IsDisabled() bool {
	if runtime.GOOS == "windows" {
		return true
	}

	return disable
}

// OnHex will take a hex color code and add the appropriate ANSI color
// codes to the provided string.
func OnHex(hex string, str string) string {
	switch strings.ToLower(os.Getenv("COLORTERM")) {
	case "truecolor":
		return colorize("on_"+HexToTrueColor(hex), str)
	default:
		return colorize("on_"+HexToXterm256(hex), str)
	}
}

// OnRainbow will add multiple ANSI color codes to a string for a
// rainbow effect.
func OnRainbow(str string) string {
	if IsDisabled() {
		// Return the string w/o any color codes
		return Plain(str)
	}

	var chars []string
	var code string
	var colors []int
	var end string
	var line []string
	var lines []string
	var out []string

	// Strip all other bg color codes and split on newline
	lines = strings.Split(plainBg(str), "\n")

	// Loop thru lines and apply bg color codes
	colors = rainbowColors()
	end = "\x1b[" + Colors["on_default"] + "m"
	for i := range lines {
		chars = iterate.FindAllString(lines[i], -1)
		line = []string{}

		// Loop thru non-color-code bytes and apply on_rainbow
		for idx, char := range chars {
			code = strconv.Itoa(colors[idx%len(colors)] + 10)
			line = append(line, "\x1b["+code+"m"+char)
		}

		// Put line back together again, ensure on_default code at end
		out = append(out, strings.Join(line, "")+end)
	}

	// Put lines back together, and remove color codes if the line
	// only contains color codes
	return onlyCodes.ReplaceAllString(strings.Join(out, "\n"), "$1$4")
}

// Plain will strip all ANSI color codes from a string.
func Plain(str string) string {
	// Strip all color codes
	return allCodes.ReplaceAllString(str, "")
}

// Rainbow will add multiple ANSI color codes to a string for a
// rainbow effect.
func Rainbow(str string) string {
	if IsDisabled() {
		// Return the string w/o any color codes
		return Plain(str)
	}

	var chars []string
	var code string
	var colors []int
	var line []string
	var lines []string
	var out []string

	// Strip all other fg color codes and split on newline
	lines = strings.Split(plainFg(str), "\n")

	// Loop thru lines and apply fg color codes
	colors = rainbowColors()
	for i := range lines {
		chars = iterate.FindAllString(lines[i], -1)
		line = []string{}

		// Loop thru non-color-code bytes and apply on_rainbow
		for idx, char := range chars {
			code = strconv.Itoa(colors[idx%len(colors)])
			line = append(line, "\x1b["+code+"m"+char)
		}

		// Put line back together again
		out = append(out, strings.Join(line, ""))
	}

	// Put lines back together, and remove color codes if the line
	// only contains color codes
	return onlyCodes.ReplaceAllString(strings.Join(out, "\n"), "$1$4")
}

// RGBAToTrueColor will convert the RGBA values to their hex
// representation.
func RGBAToTrueColor(r uint8, g uint8, b uint8, a uint8) string {
	return HexToTrueColor(Sprintf("%02x%02x%02x%02x", r, g, b, a))
}

// RGBAToXterm256 will convert the RGBA values to the closest xterm256
// representation.
func RGBAToXterm256(r uint8, g uint8, b uint8, a uint8) string {
	return HexToXterm256(Sprintf("%02x%02x%02x%02x", r, g, b, a))
}

// Sample will return all bg/fg combos of the first 16 8-bit colors.
func Sample() (lines []string) {
	var bg string
	var fg string
	var line string

	for f := 0; f < 16; f++ {
		line = ""

		for b := 0; b < 16; b++ {
			fg = Sprintf("color_%03d", f)
			bg = Sprintf("on_color_%03d", b)
			line += colorize(fg, colorize(bg, " mw "))
		}

		lines = append(lines, line)
	}

	return
}

// Table will return a pretty table of all 8-bit colors.
func Table() (lines []string) {
	var line string

	for i := 0; i < 16; i++ {
		line += Hilightf(
			Sprintf("on_color_%03d", i),
			" %s %s ",
			Blackf("%03d", i),
			Whitef("%03d", i),
		)
		if (i+1)%8 == 0 {
			lines = append(lines, line)
			line = ""
		}
	}

	for i := 16; i < 256; i++ {
		line += Hilightf(
			Sprintf("on_color_%03d", i),
			" %s %s ",
			Blackf("%03d", i),
			Whitef("%03d", i),
		)
		if (i-15)%6 == 0 {
			lines = append(lines, line)
			line = ""
		}
	}

	return
}

// Wrap will wrap a string to the specified width.
func Wrap(width int, str string) string {
	var lc int
	var line = ""
	var lines []string
	var wc int
	var words = strings.Fields(str)

	str = Sprintf(str)

	// Loop thru words
	for _, word := range words {
		lc = len([]rune(Plain(line)))
		wc = len([]rune(Plain(word)))

		if lc == 0 {
			// Can't wrap less than a single word
			line = word
		} else if lc+wc+1 > width {
			// Wrap if line would be longer than width
			lines = append(lines, line)
			line = word
		} else {
			// Append word to line
			line += " " + word
		}
	}

	// Ensure last line is not forgotten
	if len([]rune(line)) != 0 {
		lines = append(lines, line)
	}

	// If original string ended with newline, put it back
	if strings.HasSuffix(str, "\n") {
		lines = append(lines, "")
	}

	// Join lines and return
	return strings.Join(lines, "\n")
}
