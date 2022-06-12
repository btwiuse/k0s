package bcrypt

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/bcrypt"
)

const (
	// DefaultCost is the default cost for the bcrypt hash
	DefaultCost = 10
)

// BcryptGenerateCmd allows encrypting a password provided through stdin
type BcryptGenerateCmd struct {
	Cost      int `short:"c" long:"cost" default:"10" description:"The cost weight, range of 4-31" value-name:"COST"`
	OutWriter io.Writer
	InReader  io.Reader
}

// Execute performs the Brcrypt operation
func (c *BcryptGenerateCmd) Execute(args []string) error {
	if c.Cost < 4 || c.Cost > 31 {
		return fmt.Errorf("cost %d is outside allowed range (4,31)", c.Cost)
	}

	var in io.Reader
	if c.InReader != nil {
		in = c.InReader
	} else if hasPipedStdin() {
		in = os.Stdin
	} else {
		return fmt.Errorf("you must provide the password through stdin")
	}
	data, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	cb, err := bcrypt.GenerateFromPassword(data, c.Cost)
	if err != nil {
		return fmt.Errorf("error producing bcrypt hash: %v", err)
	}
	_, err = fmt.Fprintf(c.OutWriter, "%s\n", cb)
	return err
}

// NewBcryptGenerateCmd returns a BcryptGenerateCmd with configured defaults
func NewBcryptGenerateCmd() *BcryptGenerateCmd {
	return &BcryptGenerateCmd{Cost: DefaultCost, OutWriter: os.Stdout}
}

func hasPipedStdin() bool {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return true
	}
	return false
}
