/*
Copyright 2019 The Unity Scheduler Authors

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

package conf

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.infra.cloudera.com/yunikorn/k8s-shim/pkg/common"
	"time"
)

type SchedulerConf struct {
	ClusterId      string `json:"clusterId"`
	ClusterVersion string `json:"clusterVersion"`
	SchedulerName  string `json:"schedulerName"`
	Interval       int    `json:"schedulingIntervalSecond"`
	KubeConfig     string `json:"absoluteKubeConfigFilePath"`
}

func (conf *SchedulerConf) GetSchedulingInterval() time.Duration {
	return time.Duration(conf.Interval) * time.Second
}

func (conf *SchedulerConf) GetKubeConfigPath() string {
	return conf.KubeConfig
}

func ParseFromCommandline() *SchedulerConf {
	var clusterId *string
	var clusterVersion *string
	var schedulerName *string
	var kubeConfig *string
	var schedulingInterval *int

	kubeConfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	schedulingInterval = flag.Int("interval", 1, "scheduling interval in seconds")
	clusterId = flag.String("clusterid", common.ClusterId, "cluster id")
	clusterVersion = flag.String("clusterversion", common.ClusterVersion, "cluster version")
	schedulerName = flag.String("name", common.SchedulerName, "name of the scheduler")

	flag.Parse()

	return &SchedulerConf{
		ClusterId: *clusterId,
		ClusterVersion: *clusterVersion,
		SchedulerName: *schedulerName,
		Interval: *schedulingInterval,
		KubeConfig: *kubeConfig,
	}
}

func (conf *SchedulerConf) DumpConfiguration() {
	c,_ := json.MarshalIndent(&conf, "", " ")
	glog.V(3).Info(fmt.Sprintf("Scheduler conf: \n %s", string(c)))
}