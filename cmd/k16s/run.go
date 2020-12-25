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

package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/klog/v2"

	"k8s.io/kube-state-metrics/v2/external/store"
	"k8s.io/kube-state-metrics/v2/pkg/allowdenylist"
	"k8s.io/kube-state-metrics/v2/pkg/metricshandler"
	"k8s.io/kube-state-metrics/v2/pkg/options"
	"k8s.io/kube-state-metrics/v2/pkg/util/proc"
	"k8s.io/kube-state-metrics/v2/pkg/version"
)

func Run(args []string) error {
	os.Args = append([]string{"k16s"}, args...)
	opts := options.NewOptions()
	opts.AddFlags()

	ctx := context.Background()

	err := opts.Parse()
	if err != nil {
		klog.Fatalf("Error: %s", err)
	}

	if opts.Version {
		fmt.Printf("%#v\n", version.GetVersion())
		os.Exit(0)
	}

	if opts.Help {
		opts.Usage()
		os.Exit(0)
	}
	storeBuilder := store.NewBuilder()

	ksmMetricsRegistry := prometheus.NewRegistry()
	durationVec := promauto.With(ksmMetricsRegistry).NewHistogramVec(
		prometheus.HistogramOpts{
			Name:        "http_request_duration_seconds",
			Help:        "A histogram of requests for kube-state-metrics metrics handler.",
			Buckets:     prometheus.DefBuckets,
			ConstLabels: prometheus.Labels{"handler": "metrics"},
		}, []string{"method"},
	)
	storeBuilder.WithMetrics(ksmMetricsRegistry)

	var resources []string
	if len(opts.Resources) == 0 {
		klog.Info("Using default resources")
		resources = options.DefaultResources.AsSlice()
	} else {
		klog.Infof("Using resources %s", opts.Resources.String())
		resources = opts.Resources.AsSlice()
	}

	if err := storeBuilder.WithEnabledResources(resources); err != nil {
		klog.Fatalf("Failed to set up resources: %v", err)
	}

	if len(opts.Namespaces) == 0 {
		klog.Info("Using all namespace")
		storeBuilder.WithNamespaces(options.DefaultNamespaces)
	} else {
		if opts.Namespaces.IsAllNamespaces() {
			klog.Info("Using all namespace")
		} else {
			klog.Infof("Using %s namespaces", opts.Namespaces)
		}
		storeBuilder.WithNamespaces(opts.Namespaces)
	}

	allowDenyList, err := allowdenylist.New(opts.MetricAllowlist, opts.MetricDenylist)
	if err != nil {
		klog.Fatal(err)
	}

	err = allowDenyList.Parse()
	if err != nil {
		klog.Fatalf("error initializing the allowdeny list : %v", err)
	}

	klog.Infof("metric allow-denylisting: %v", allowDenyList.Status())

	storeBuilder.WithAllowDenyList(allowDenyList)

	storeBuilder.WithGenerateStoreFunc(storeBuilder.DefaultGenerateStoreFunc())

	proc.StartReaper()

	kubeClient, vpaClient, err := createKubeClient(opts.Apiserver, opts.Kubeconfig)
	if err != nil {
		klog.Fatalf("Failed to create client: %v", err)
	}
	storeBuilder.WithKubeClient(kubeClient)
	storeBuilder.WithVPAClient(vpaClient)
	storeBuilder.WithSharding(opts.Shard, opts.TotalShards)
	storeBuilder.WithAllowLabels(opts.LabelsAllowList)

	ksmMetricsRegistry.MustRegister(
		prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
		prometheus.NewGoCollector(),
	)

	var g run.Group

	m := metricshandler.New(
		opts,
		kubeClient,
		storeBuilder,
		opts.EnableGZIPEncoding,
	)
	// Run MetricsHandler
	{
		ctxMetricsHandler, cancel := context.WithCancel(ctx)
		g.Add(func() error {
			return m.Run(ctxMetricsHandler)
		}, func(error) {
			cancel()
		})
	}

	telemetryMux := buildTelemetryServer(ksmMetricsRegistry)
	telemetryServer := http.Server{Handler: telemetryMux}
	telemetryListenAddress := net.JoinHostPort(opts.TelemetryHost, strconv.Itoa(opts.TelemetryPort))
	telemetryLn, err := net.Listen("tcp", telemetryListenAddress)
	if err != nil {
		klog.Fatalf("Failed to create Telemetry Listener: %v", err)
	}
	metricsMux := buildMetricsServer(m, durationVec)
	metricsServer := http.Server{Handler: metricsMux}
	metricsServerListenAddress := net.JoinHostPort(opts.Host, strconv.Itoa(opts.Port))
	metricsServerLn, err := net.Listen("tcp", metricsServerListenAddress)
	if err != nil {
		klog.Fatalf("Failed to create MetricsServer Listener: %v", err)
	}

	// Run Telemetry server
	{
		g.Add(func() error {
			klog.Infof("Starting kube-state-metrics self metrics server: %s", telemetryListenAddress)
			return telemetryServer.Serve(telemetryLn)
		}, func(error) {
			ctxShutDown, cancel := context.WithTimeout(ctx, 3*time.Second)
			defer cancel()
			telemetryServer.Shutdown(ctxShutDown)
		})
	}
	// Run Metrics server
	{
		g.Add(func() error {
			klog.Infof("Starting metrics server: %s", metricsServerListenAddress)
			return metricsServer.Serve(metricsServerLn)
		}, func(error) {
			ctxShutDown, cancel := context.WithTimeout(ctx, 3*time.Second)
			defer cancel()
			metricsServer.Shutdown(ctxShutDown)
		})
	}

	if err := g.Run(); err != nil {
		klog.Fatalf("RunGroup Error: %v", err)
	}
	klog.Info("Exiting")
	return nil
}
