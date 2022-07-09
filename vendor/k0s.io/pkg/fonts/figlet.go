package fonts

import (
	"bytes"
	"log"
	"os"
	"strings"

	"github.com/lukesampson/figlet/figletlib"
	isatty "github.com/mattn/go-isatty"
	"k0s.io/pkg/console"
)

func Figlet(str string) {
	var (
		std, _     = Fonts["/standard.flf"]
		width  int = 80
	)
	if isatty.IsTerminal(os.Stdin.Fd()) {
		var (
			term    = console.Current()
			size, _ = term.Size()
		)
		width = int(size.Width)
	}
	stdFont, err := figletlib.ReadFontFromBytes([]byte(std))
	if err != nil {
		log.Println("Failed to load standard figlet font:", err)
		log.Println(str)
		return
	}
	buf := bytes.NewBuffer(nil)
	figletlib.FPrintMsg(buf, str, stdFont, width, stdFont.Settings(), "left")
	lines := strings.Split(buf.String(), "\n")
	for _, line := range lines {
		log.Println(line)
	}
}
