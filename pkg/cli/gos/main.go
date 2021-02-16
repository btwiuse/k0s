package gos

import (
	"flag"
	"strings"

	"github.com/lunny/tango"
	"github.com/tango-contrib/basicauth"
)

func Run(args []string) error {
	var (
		flagset = flag.NewFlagSet("gos", flag.ContinueOnError)
		dir     = flagset.String("dir", ".", "static dir path")
		listen  = flagset.String("listen", ":8000", "listen port")
		user    = flagset.String("user", "", "basic auth user name")
		pass    = flagset.String("pass", "", "basic auth user password")
		listDir = flagset.Bool("listDir", true, "if list dir files")
		exts    = flagset.String("exts", "", "filtered ext files will be supplied")
	)

	err := flagset.Parse(args)
	if err != nil {
		return err
	}

	t := tango.New()
	if *user != "" {
		t.Use(basicauth.New(*user, *pass))
		t.Logger().Info("Basic auth module loaded")
	}
	var filterExts []string
	if len(*exts) > 0 {
		filterExts = strings.Split(*exts, ",")
	}
	t.Use(tango.Logging())
	t.Use(tango.Static(tango.StaticOptions{
		RootPath:   *dir,
		ListDir:    *listDir,
		FilterExts: filterExts,
	}))

	t.Run(*listen)
	return nil
}
