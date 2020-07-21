package gost

import (
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime"

	_ "net/http/pprof"

	"github.com/ginuerzh/gost"
	"github.com/go-log/log"
)

var (
	configureFile string
	baseCfg       = &baseConfig{}
	pprofAddr     string
	pprofEnabled  = os.Getenv("PROFILING") != ""
)

func Main(args []string) {
	gost.SetLogger(&gost.LogLogger{})

	var (
		printVersion bool
		flagset      = flag.NewFlagSet("gost", flag.ContinueOnError)
	)

	flagset.Var(&baseCfg.route.ChainNodes, "F", "forward address, can make a forward chain")
	flagset.Var(&baseCfg.route.ServeNodes, "L", "listen address, can listen on multiple ports (required)")
	flagset.StringVar(&configureFile, "C", "", "configure file")
	flagset.BoolVar(&baseCfg.Debug, "D", false, "enable debug log")
	flagset.BoolVar(&printVersion, "V", false, "print version")
	if pprofEnabled {
		flagset.StringVar(&pprofAddr, "P", ":6060", "profiling HTTP server address")
	}
	flagset.Parse(args)

	if printVersion {
		fmt.Fprintf(os.Stdout, "gost %s (%s %s/%s)\n",
			gost.Version, runtime.Version(), runtime.GOOS, runtime.GOARCH)
		os.Exit(0)
	}

	if configureFile != "" {
		_, err := parseBaseConfig(configureFile)
		if err != nil {
			log.Log(err)
			os.Exit(1)
		}
	}
	if flagset.NFlag() == 0 {
		flagset.PrintDefaults()
		os.Exit(0)
	}
	if pprofEnabled {
		go func() {
			log.Log("profiling server on", pprofAddr)
			log.Log(http.ListenAndServe(pprofAddr, nil))
		}()
	}

	// NOTE: as of 2.6, you can use custom cert/key files to initialize the default certificate.
	tlsConfig, err := tlsConfig(defaultCertFile, defaultKeyFile, "")
	if err != nil {
		// generate random self-signed certificate.
		cert, err := gost.GenCertificate()
		if err != nil {
			log.Log(err)
			os.Exit(1)
		}
		tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
	} else {
		log.Log("load TLS certificate files OK")
	}

	gost.DefaultTLSConfig = tlsConfig

	if err := start(); err != nil {
		log.Log(err)
		os.Exit(1)
	}

	select {}
}

func start() error {
	gost.Debug = baseCfg.Debug

	var routers []router
	rts, err := baseCfg.route.GenRouters()
	if err != nil {
		return err
	}
	routers = append(routers, rts...)

	for _, route := range baseCfg.Routes {
		rts, err := route.GenRouters()
		if err != nil {
			return err
		}
		routers = append(routers, rts...)
	}

	if len(routers) == 0 {
		return errors.New("invalid config")
	}
	for i := range routers {
		fmt.Println(routers[i])
		go routers[i].Serve()
	}

	return nil
}
