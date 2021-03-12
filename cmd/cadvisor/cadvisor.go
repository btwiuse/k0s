// Copyright 2014 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	cadvisorhttp "github.com/btwiuse/cadvisor/http"
	"github.com/google/cadvisor/container"
	"github.com/google/cadvisor/manager"
	"github.com/google/cadvisor/metrics"
	"github.com/google/cadvisor/utils/sysfs"
	"github.com/google/cadvisor/version"

	// Register container providers
	_ "github.com/btwiuse/cadvisor/container/install"

	// Register CloudProviders
	_ "github.com/google/cadvisor/utils/cloudinfo/aws"
	_ "github.com/google/cadvisor/utils/cloudinfo/azure"
	_ "github.com/google/cadvisor/utils/cloudinfo/gce"

	"k8s.io/klog/v2"
)

var argIp = flag.String("listen_ip", "", "IP to listen on, defaults to all IPs")
var argPort = flag.Int("port", 8080, "port to listen")

var versionFlag = flag.Bool("version", false, "print cAdvisor version and exit")

var prometheusEndpoint = flag.String("prometheus_endpoint", "/metrics", "Endpoint to expose Prometheus metrics on")

var housekeepingConfig = manager.HouskeepingConfig{
	flag.Duration("max_housekeeping_interval", 60*time.Second, "Largest interval to allow between container housekeepings"),
	flag.Bool("allow_dynamic_housekeeping", true, "Whether to allow the housekeeping interval to be dynamic"),
}

var collectorCert = flag.String("collector_cert", "", "Collector's certificate, exposed to endpoints for certificate based authentication.")
var collectorKey = flag.String("collector_key", "", "Key for the collector's certificate")

var urlBasePrefix = flag.String("url_base_prefix", "", "prefix path that will be prepended to all paths to support some reverse proxies")

var rawCgroupPrefixWhiteList = flag.String("raw_cgroup_prefix_whitelist", "", "A comma-separated list of cgroup path prefix that needs to be collected even when -docker_only is specified")

func initialize() {
	// Default logging verbosity to V(2)
	flag.Set("v", "2")
}

func main() {
	initialize()
	klog.InitFlags(nil)
	defer klog.Flush()
	flag.Parse()

	if *versionFlag {
		fmt.Printf("cAdvisor version %s (%s)\n", version.Info["version"], version.Info["revision"])
		os.Exit(0)
	}

	includedMetrics := container.AllMetrics

	memoryStorage, err := NewMemoryStorage()
	if err != nil {
		klog.Fatalf("Failed to initialize storage driver: %s", err)
	}

	sysFs := sysfs.NewRealSysFs()

	collectorHttpClient := createCollectorHttpClient(*collectorCert, *collectorKey)

	resourceManager, err := manager.New(memoryStorage, sysFs, housekeepingConfig, includedMetrics, &collectorHttpClient, strings.Split(*rawCgroupPrefixWhiteList, ","), "")
	if err != nil {
		klog.Fatalf("Failed to create a manager: %s", err)
	}

	mux := http.NewServeMux()

	// Register all HTTP handlers.
	err = cadvisorhttp.RegisterHandlers(mux, resourceManager, "", "", "", "", *urlBasePrefix)
	if err != nil {
		klog.Fatalf("Failed to register HTTP handlers: %v", err)
	}

	containerLabelFunc := metrics.DefaultContainerLabels

	// Register Prometheus collector to gather information about containers, Go runtime, processes, and machine
	cadvisorhttp.RegisterPrometheusHandler(mux, resourceManager, *prometheusEndpoint, containerLabelFunc, includedMetrics)

	// Start the manager.
	if err := resourceManager.Start(); err != nil {
		klog.Fatalf("Failed to start manager: %v", err)
	}

	// Install signal handler.
	installSignalHandler(resourceManager)

	klog.V(1).Infof("Starting cAdvisor version: %s-%s on port %d", version.Info["version"], version.Info["revision"], *argPort)

	rootMux := http.NewServeMux()
	rootMux.Handle(*urlBasePrefix+"/", http.StripPrefix(*urlBasePrefix, mux))

	addr := fmt.Sprintf("%s:%d", *argIp, *argPort)
	klog.Fatal(http.ListenAndServe(addr, rootMux))
}

func installSignalHandler(containerManager manager.Manager) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received.
	go func() {
		sig := <-c
		if err := containerManager.Stop(); err != nil {
			klog.Errorf("Failed to stop container manager: %v", err)
		}
		klog.Infof("Exiting given signal: %v", sig)
		os.Exit(0)
	}()
}

func createCollectorHttpClient(collectorCert, collectorKey string) http.Client {
	//Enable accessing insecure endpoints. We should be able to access metrics from any endpoint
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	if collectorCert != "" {
		if collectorKey == "" {
			klog.Fatal("The collector_key value must be specified if the collector_cert value is set.")
		}
		cert, err := tls.LoadX509KeyPair(collectorCert, collectorKey)
		if err != nil {
			klog.Fatalf("Failed to use the collector certificate and key: %s", err)
		}

		tlsConfig.Certificates = []tls.Certificate{cert}
		tlsConfig.BuildNameToCertificate()
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	return http.Client{Transport: transport}
}
