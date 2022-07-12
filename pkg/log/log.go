package log

import "log"

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

var Fatalf = log.Fatalf
var Fatalln = log.Fatalln

const Ldate = log.Ldate

type Logger = log.Logger

const Lshortfile = log.Lshortfile
const Ltime = log.Ltime

var New = log.New
var Print = log.Print
var Printf = log.Printf
var Println = log.Println
var SetFlags = log.SetFlags
