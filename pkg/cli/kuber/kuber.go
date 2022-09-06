package kuber

import (
	"fmt"
	"log"
	"net/http"

	"k0s.io"
	"k0s.io/pkg/apiserver"
	"k0s.io/pkg/apiserver/routes"
	"k0s.io/pkg/utils"
	"k8s.io/apimachinery/pkg/version"
)

var Info version.Info = version.Info{
	Major:        "1",
	Minor:        "23",
	GitVersion:   "v1.23.0",
	GitCommit:    "10bca343e85180f668522fe252552da20220cb7a",
	GitTreeState: "clean",
	BuildDate:    "2021-09-23T22:46:18Z",
	GoVersion:    "go1.16.8",
	Compiler:     "gc",
	Platform:     "linux/amd64",
}

func Run([]string) error {
	port := utils.EnvPORT(k0s.HUB_PORT)
	handler := apiserver.DefaultAPIServerHandler()
	routes.Version{Version: &Info}.Install(handler.GoRestfulContainer)
	routes.Index{}.Install(handler, handler.NonGoRestfulMux)
	log.Println(fmt.Sprintf("listening on http://127.0.0.1%s", port))
	return http.ListenAndServe(port, handler)
}
