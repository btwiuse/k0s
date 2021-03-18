/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"k0s.io/pkg/kproxy"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubectl/pkg/scheme"
)

func main() {
	var (
		staticDir    = ""
		apiPrefix    = "/"
		staticPrefix = "/static/"
		// filter       = &kproxy.FilterServer{}
		keepalive    = 30 * time.Second

		port = ":8090"

		apiserver  = ""
		kubeconfig = "/home/btwiuse/.kube/config.civo"
	)

	clientConfig, err := createRESTConfig(apiserver, kubeconfig)
	if err != nil {
		log.Fatalln(err)
	}

	handler, err := kproxy.NewHandler(staticDir, apiPrefix, staticPrefix, nil /* filter */ , clientConfig, keepalive)
	if err != nil {
		log.Fatalln(err)
	}

	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("listening on", port)
	err = http.Serve(ln, handler)
	if err != nil {
		log.Fatalln(err)
	}
}

func createRESTConfig(apiserver string, kubeconfig string) (*rest.Config, error) {
	config, err := clientcmd.BuildConfigFromFlags(apiserver, kubeconfig)
	if err != nil {
		return nil, err
	}
	config.GroupVersion = &schema.GroupVersion{Group: "", Version: "v1"}
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.APIPath = "/api"
	return config, nil
}
