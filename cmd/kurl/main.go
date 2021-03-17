// Copyright 2016 Alexander Brand
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/spf13/pflag"

	rest "k8s.io/client-go/rest"
	api "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
)

var server = pflag.StringP("server", "s", "http://127.0.0.1:8080", "address of the K8s API Server")
var proxyMode = pflag.Bool("proxy", false, "start in proxy mode")
var proxyPort = pflag.String("proxy-port", "9090", "set the port when running in proxy mode")

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	pflag.Usage = usage
	pflag.Parse()

	if *proxyMode {
		// Don't use proxy set in http_proxy env var
		noProxyTransport := &http.Transport{
			Proxy: nil,
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			// Find the request pod's IP
			podIP, err := getPodIP(r.Host)
			if err != nil {
				log.Printf("error getting pod IP: %v", err)
				return
			}
			if podIP == "" {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprint(w, "Pod not found\n")
				return
			}

			// Reconstruct URL with pod IP and create reverse proxy
			u := r.URL
			u.Scheme = "http"
			u.Host = podIP

			proxy := httputil.NewSingleHostReverseProxy(u)
			proxy.Transport = noProxyTransport
			proxy.ServeHTTP(w, r)
		})

		fmt.Println("Starting proxy on port", *proxyPort)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *proxyPort), nil))
	}

	// Not using proxy mode, find pod IP and issue GET request
	args := pflag.Args()
	if len(args) != 1 {
		usage()
	}
	podName := args[0]

	podIP, err := getPodIP(podName)
	if err != nil {
		log.Fatalf("error getting pod IP: %v", err)
	}
	if podIP == "" {
		log.Fatal("pod not found")
	}

	resp, err := http.Get("http://" + podIP)
	if err != nil {
		log.Fatalf("error connecting to pod: %v\n", err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatalf("error writing response to STDOUT: %v\n", err)
	}
}

func usage() {
	fmt.Print("kurl: curl for Kubernetes\n\n")
	fmt.Println("Usage: kurl POD_NAME")
	pflag.PrintDefaults()
	os.Exit(1)
}

func getPodIP(podName string) (string, error) {
	c, err := createKubeClient(*server, "")
	if err != nil {
		return "", fmt.Errorf("error building client: %v", err)
	}

	res := c.Get().Resource("pods").Timeout(5 * time.Second).Do(context.Background())
	obj, err := res.Get()
	if err != nil {
		return "", fmt.Errorf("error getting object from response: %v\n", err)
	}
	list, ok := obj.(*api.PodList)
	if !ok {
		return "", fmt.Errorf("got something other than a pod list")
	}

	var podIP string
	for _, p := range list.Items {
		if strings.HasPrefix(p.Name, podName) {
			podIP = p.Status.PodIP
		}
	}
	if podIP == "" {
		return "", nil
	}

	log.Println(podName, podIP)
	return podIP, nil
}

func createKubeClient(apiserver string, kubeconfig string) (*rest.RESTClient, error) {
        config, err := clientcmd.BuildConfigFromFlags(apiserver, kubeconfig)
        if err != nil {
                return nil, err
        }
	config.GroupVersion = &schema.GroupVersion{Group: "", Version: "v1"}
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.APIPath = "/api"

	c, err := rest.UnversionedRESTClientFor(config)
	if err != nil {
		return nil, err
	}
	return c, nil
}
