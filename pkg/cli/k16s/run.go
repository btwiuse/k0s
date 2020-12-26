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

package k16s

import (
	"context"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus"

	"k8s.io/klog/v2"
	"k8s.io/kube-state-metrics/v2/external/store"
	"k8s.io/kube-state-metrics/v2/pkg/allowdenylist"
	"k8s.io/kube-state-metrics/v2/pkg/metricshandler"
	"k8s.io/kube-state-metrics/v2/pkg/options"
)

func Run(args []string) error {
	ctx := context.Background()

	os.Args = append([]string{"k16s"}, args...)
        opts := options.NewOptions()
        opts.AddFlags()
        opts.Parse()

	opts.Kubeconfig      = os.Getenv("KUBECONFIG")
	opts.Port            = 8165
	opts.EnableGZIPEncoding = true

	storeBuilder := store.NewBuilder()

        ksmMetricsRegistry := prometheus.NewRegistry()
        storeBuilder.WithMetrics(ksmMetricsRegistry)

	var resources []string = options.DefaultResources.AsSlice()

	if err := storeBuilder.WithEnabledResources(resources); err != nil {
		klog.Fatalf("Failed to set up resources: %v", err)
	}

	storeBuilder.WithNamespaces(options.DefaultNamespaces)
	storeBuilder.WithAllowDenyList(&allowdenylist.AllowDenyList{})
	storeBuilder.WithGenerateStoreFunc(storeBuilder.DefaultGenerateStoreFunc())

	kubeClient, vpaClient, err := createKubeClient(opts.Apiserver, opts.Kubeconfig)
	if err != nil {
		klog.Fatalf("Failed to create client: %v", err)
	}
	storeBuilder.WithKubeClient(kubeClient)
	storeBuilder.WithVPAClient(vpaClient)
	storeBuilder.WithSharding(opts.Shard, opts.TotalShards)
	storeBuilder.WithAllowLabels(opts.LabelsAllowList)

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

	metricsServer := http.Server{Handler: m}
	metricsServerListenAddress := net.JoinHostPort(opts.Host, strconv.Itoa(opts.Port))
	metricsServerLn, err := net.Listen("tcp", metricsServerListenAddress)
	if err != nil {
		klog.Fatalf("Failed to create MetricsServer Listener: %v", err)
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
