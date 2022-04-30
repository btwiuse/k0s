package hilighter

import "regexp"

// Colors maps color names to color codes
var Colors = map[string]string{
	"black":         "30",
	"red":           "31",
	"green":         "32",
	"yellow":        "33",
	"blue":          "34",
	"magenta":       "35",
	"cyan":          "36",
	"white":         "37",
	"light_black":   "90",
	"light_red":     "91",
	"light_green":   "92",
	"light_yellow":  "93",
	"light_blue":    "94",
	"light_magenta": "95",
	"light_cyan":    "96",
	"light_white":   "97",

	"on_black":         "40",
	"on_red":           "41",
	"on_green":         "42",
	"on_yellow":        "43",
	"on_blue":          "44",
	"on_magenta":       "45",
	"on_cyan":          "46",
	"on_white":         "47",
	"on_light_black":   "100",
	"on_light_red":     "101",
	"on_light_green":   "102",
	"on_light_yellow":  "103",
	"on_light_blue":    "104",
	"on_light_magenta": "105",
	"on_light_cyan":    "106",
	"on_light_white":   "107",

	"default":    "39",
	"on_default": "49",
}

// Used to disable all color codes
var disable = false

// Cached hex to xterm-256 8-bit mappings
var cachedCodes = map[string]string{}

// Modes maps mode names to mode codes
var Modes = map[string]string{
	"reset":         "0",
	"normal":        "0",
	"bold":          "1",
	"dim":           "2",
	"faint":         "2",
	"italic":        "3",
	"underline":     "4",
	"blink":         "5",
	"blink_slow":    "5",
	"blink_rapid":   "6",
	"inverse":       "7",
	"negative":      "7",
	"swap":          "7",
	"hide":          "8",
	"conceal":       "8",
	"crossed_out":   "9",
	"strikethrough": "9",
	"fraktur":       "20",

	"no_bold":          "21",
	"no_dim":           "22",
	"no_faint":         "22",
	"no_italic":        "23",
	"no_fraktur":       "23",
	"no_underline":     "24",
	"no_blink":         "25",
	"no_blink_slow":    "25",
	"no_blink_rapid":   "26",
	"no_inverse":       "27",
	"no_negative":      "27",
	"no_swap":          "27",
	"no_hide":          "28",
	"no_conceal":       "28",
	"no_crossed_out":   "29",
	"no_strikethrough": "29",
}

// Version is the package version
const Version = "1.10.1"

// Various regular expressions
var allCodes = regexp.MustCompile(`\x1b\[([0-9;]*m|K)`)
var bgCodes = regexp.MustCompile(`\x1b\[(4|10)[0-9;]+m`)
var doubleno = regexp.MustCompile(`no_no_`)
var fgCodes = regexp.MustCompile(`\x1b\[[39][0-9;]+m`)
var hexCode = regexp.MustCompile(`(?i)(on_)?([0-9a-f]{6})`)
var iterate = regexp.MustCompile(
	`(\x1b\[([0-9;]*m|K))*[^\x1b](\x1b\[([0-9;]*m|K))*`,
)
var newline = regexp.MustCompile(`\n`)
var onlyCodes = regexp.MustCompile(
	`(^|\n)(\x1b\[([0-9;]+m|K))+(\n|$)`,
)
var parseHex = regexp.MustCompile(
	`(?i)^#?([0-9a-f]{2})([0-9a-f]{2})([0-9a-f]{2})([0-9a-f]{2})?$`,
)
var wrap = regexp.MustCompile(`wrap(_(\d+))?`)

func init() {
	var key string
	var val string

	// Add all 8-bit colors, fg and bg
	for i := 0; i < 256; i++ {
		key = Sprintf("color_%03d", i)
		val = Sprintf("38;5;%03d", i)
		Colors[key] = val

		key = Sprintf("on_color_%03d", i)
		val = Sprintf("48;5;%03d", i)
		Colors[key] = val
	}
}
