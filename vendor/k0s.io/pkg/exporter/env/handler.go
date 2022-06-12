/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package env

import (
	"net/http"
	// "net/http/httputil"
	// "net/url"

	"io"
	"os"
	"strings"
)

func handleEnv(w http.ResponseWriter, r *http.Request) {
	fromEnv := getEnvName(r)
	// log.Println("fromEnv:", fromEnv)
	if fromEnv == "" {
		http.NotFoundHandler().ServeHTTP(w, r)
		return
	}

	envURL := os.Getenv(fromEnv)
	// log.Println("envURL:", envURL)
	if envURL == "" {
		http.NotFoundHandler().ServeHTTP(w, r)
		return
	}

	// io.WriteString(w, envURL)

	/*
	u, err := url.Parse(envURL)
	if err != nil {
		// bad request or internal server error ?
		http.NotFoundHandler().ServeHTTP(w, r)
		return
	}
	*/
	// proxy := httputil.NewSingleHostReverseProxy(u)
	// proxy.ServeHTTP(w, r)
	resp, err := http.Get(envURL)
	if err != nil {
		http.NotFoundHandler().ServeHTTP(w, r)
		return
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}

func NewHandler() http.Handler {
	return http.HandlerFunc(handleEnv)
}

func getEnvName(r *http.Request) string {
	for k, _ := range r.URL.Query() {
		// log.Println("vars:", k)
		if strings.HasPrefix(k, "env.") {
			return strings.TrimPrefix(k, "env.")
		}
	}
	return ""
}
