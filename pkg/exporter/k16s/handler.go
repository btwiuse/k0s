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
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"

	vpaclientset "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/client/clientset/versioned"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
	"k8s.io/kube-state-metrics/v2/external/store"
	"k8s.io/kube-state-metrics/v2/pkg/allowdenylist"
	"k8s.io/kube-state-metrics/v2/pkg/metricshandler"
	"k8s.io/kube-state-metrics/v2/pkg/options"
)

func opt(args []string) *options.Options{
        osargs := os.Args
        os.Args = append([]string{"k16s"}, args...)
        opts := options.NewOptions()
        opts.AddFlags()
        opts.Parse()

        opts.Kubeconfig      = os.Getenv("KUBECONFIG")
        opts.EnableGZIPEncoding = true

        os.Args = osargs
        return opts
}

func buildStore(opts *options.Options, kubeClient clientset.Interface, vpaClient vpaclientset.Interface) *store.Builder {
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

	storeBuilder.WithKubeClient(kubeClient)
	storeBuilder.WithVPAClient(vpaClient)
	storeBuilder.WithSharding(opts.Shard, opts.TotalShards)
	storeBuilder.WithAllowLabels(opts.LabelsAllowList)
	return storeBuilder
}

func NewHandler() http.Handler {
	ctx := context.Background()

	opts := opt(nil)

	kubeClient, vpaClient, err := createKubeClient(opts.Apiserver, opts.Kubeconfig)
	if err != nil {
		klog.Errorf("Failed to create client: %v", err)
		return http.NotFoundHandler()
	}

	storeBuilder := buildStore(opts, kubeClient, vpaClient)

	m := metricshandler.New(
		opts,
		kubeClient,
		storeBuilder,
		opts.EnableGZIPEncoding,
	)

	go m.Run(ctx)

	return m
}
