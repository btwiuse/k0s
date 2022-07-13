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
	"github.com/pkg/errors"
	vpaclientset "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/client/clientset/versioned"
	clientset "k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/btwiuse/k16s/v2/pkg/version"
	"k8s.io/klog/v2"
)

func createKubeClient(apiserver string, kubeconfig string) (clientset.Interface, vpaclientset.Interface, error) {
	config, err := clientcmd.BuildConfigFromFlags(apiserver, kubeconfig)
	if err != nil {
		return nil, nil, err
	}

	config.UserAgent = version.GetVersion().String()
	config.AcceptContentTypes = "application/vnd.kubernetes.protobuf,application/json"
	config.ContentType = "application/vnd.kubernetes.protobuf"

	kubeClient, err := clientset.NewForConfig(config)
	if err != nil {
		return nil, nil, err
	}

	vpaClient, err := vpaclientset.NewForConfig(config)
	if err != nil {
		return nil, nil, err
	}
	// Informers don't seem to do a good job logging error messages when it
	// can't reach the server, making debugging hard. This makes it easier to
	// figure out if apiserver is configured incorrectly.
	klog.Infof("Testing communication with server")
	v, err := kubeClient.Discovery().ServerVersion()
	if err != nil {
		return nil, nil, errors.Wrap(err, "error while trying to communicate with apiserver")
	}
	klog.Infof("Running with Kubernetes cluster version: v%s.%s. git version: %s. git tree state: %s. commit: %s. platform: %s",
		v.Major, v.Minor, v.GitVersion, v.GitTreeState, v.GitCommit, v.Platform)
	klog.Infof("Communication with server successful")

	return kubeClient, vpaClient, nil
}
