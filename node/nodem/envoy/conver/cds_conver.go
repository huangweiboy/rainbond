// RAINBOND, Application Management Platform
// Copyright (C) 2014-2017 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package conver

import (
	"fmt"
	"strconv"

	"github.com/Sirupsen/logrus"
	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	api_model "github.com/goodrain/rainbond/api/model"
	envoyv2 "github.com/goodrain/rainbond/node/core/envoy/v2"
	corev1 "k8s.io/api/core/v1"
)

//OneNodeCluster conver cluster of on envoy node
func OneNodeCluster(serviceAlias, namespace string, configs *corev1.ConfigMap, services []*corev1.Service) ([]cache.Resource, error) {
	resources, _, err := GetPluginConfigs(configs)
	if err != nil {
		return nil, err
	}
	var clusters []cache.Resource
	if resources.BaseServices != nil && len(resources.BaseServices) > 0 {
		for _, cl := range upstreamClusters(serviceAlias, namespace, resources.BaseServices, services) {
			if err := cl.Validate(); err != nil {
				logrus.Errorf("cluster validate failure %s", err.Error())
			} else {
				clusters = append(clusters, cl)
			}
		}
	}
	if resources.BasePorts != nil && len(resources.BasePorts) > 0 {
		for _, cl := range downstreamClusters(serviceAlias, namespace, resources.BasePorts) {
			if err := cl.Validate(); err != nil {
				logrus.Errorf("cluster validate failure %s", err.Error())
			} else {
				clusters = append(clusters, cl)
			}
		}
	}
	if len(clusters) == 0 {
		logrus.Warningf("configmap name: %s; plugin-config: %s; create clusters zero length", configs.Name, configs.Data["plugin-config"])
	}
	return clusters, nil
}

// upstreamClusters handle upstream app cluster
// handle kubernetes inner service
func upstreamClusters(serviceAlias, namespace string, dependsServices []*api_model.BaseService, services []*corev1.Service) (cdsClusters []*v2.Cluster) {
	var clusterConfig = make(map[string]*api_model.BaseService, len(dependsServices))
	for i, dService := range dependsServices {
		depServiceIndex := fmt.Sprintf("%s_%s_%s_%d", namespace, serviceAlias, dService.DependServiceAlias, dService.Port)
		clusterConfig[depServiceIndex] = dependsServices[i]
	}
	for _, service := range services {
		inner, ok := service.Labels["service_type"]
		destServiceAlias := GetServiceAliasByService(service)
		port := service.Spec.Ports[0]
		if !ok || inner != "inner" {
			continue
		}
		getOptions := func() (d envoyv2.RainbondPluginOptions) {
			relPort, _ := strconv.Atoi(service.Labels["origin_port"])
			if relPort == 0 {
				relPort = int(port.TargetPort.IntVal)
			}
			depServiceIndex := fmt.Sprintf("%s_%s_%s_%d", namespace, serviceAlias, GetServiceAliasByService(service), relPort)
			if _, ok := clusterConfig[depServiceIndex]; ok {
				return envoyv2.GetOptionValues(clusterConfig[depServiceIndex].Options)
			}
			return envoyv2.GetOptionValues(nil)
		}
		var clusterOption envoyv2.ClusterOptions
		clusterOption.Name = fmt.Sprintf("%s_%s_%s_%v", namespace, serviceAlias, GetServiceAliasByService(service), port.Port)
		options := getOptions()
		clusterOption.OutlierDetection = envoyv2.CreatOutlierDetection(options)
		clusterOption.CircuitBreakers = envoyv2.CreateCircuitBreaker(options)
		clusterOption.ServiceName = fmt.Sprintf("%s_%s_%s_%v", namespace, serviceAlias, destServiceAlias, port.Port)
		clusterOption.ClusterType = v2.Cluster_EDS
		clusterOption.HealthyPanicThreshold = options.HealthyPanicThreshold
		cluster := envoyv2.CreateCluster(clusterOption)
		if cluster != nil {
			cdsClusters = append(cdsClusters, cluster)
		}
	}
	return
}

//downstreamClusters handle app self cluster
//only local port
func downstreamClusters(serviceAlias, namespace string, ports []*api_model.BasePort) (cdsClusters []*v2.Cluster) {
	for i := range ports {
		port := ports[i]
		address := envoyv2.CreateSocketAddress(port.Protocol, "127.0.0.1", uint32(port.Port))
		clusterName := fmt.Sprintf("%s_%s_%v", namespace, serviceAlias, port.Port)
		option := envoyv2.GetOptionValues(port.Options)
		cluster := envoyv2.CreateCluster(envoyv2.ClusterOptions{
			Name:                     clusterName,
			ServiceName:              "",
			ClusterType:              v2.Cluster_STATIC,
			CircuitBreakers:          envoyv2.CreateCircuitBreaker(option),
			OutlierDetection:         envoyv2.CreatOutlierDetection(option),
			MaxRequestsPerConnection: option.MaxRequestsPerConnection,
			Hosts:                    []*core.Address{&address},
			HealthyPanicThreshold:    option.HealthyPanicThreshold,
		})
		if cluster != nil {
			cdsClusters = append(cdsClusters, cluster)
		}
	}
	return
}
