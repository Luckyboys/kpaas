// Copyright 2019 Shanghai JingDuo Information Technology co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package task

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	pb "github.com/kpaas-io/kpaas/pkg/deploy/protos"
)

const (
	JointMasterOperation Operation = "join"
	JoinMasterPriority   Priority  = 20
)

type JoinMasterTaskConfig struct {
	operation       Operation
	node            *pb.Node
	masterNodes     []*pb.Node
	clusterConfig   *pb.ClusterConfig
	logFileBasePath string
	priority        int
	parent          string
}

type JoinMasterTask struct {
	Base
	Operation     Operation
	Node          *pb.Node
	MasterNodes   []*pb.Node
	clusterConfig *pb.ClusterConfig
}

func NewJoinMasterTask(taskName string, taskConfig *JoinMasterTaskConfig) (Task, error) {
	var err error
	if taskConfig == nil {
		err = fmt.Errorf("invalid task config: nil")
	} else if taskConfig.node == nil {
		err = fmt.Errorf("invalid task config: node is empty")
	}

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	task := &JoinMasterTask{
		Base: Base{
			Name:              taskName,
			TaskType:          TaskTypeJoinMaster,
			Status:            TaskPending,
			LogFilePath:       GenTaskLogFilePath(taskConfig.logFileBasePath, taskName),
			CreationTimestamp: time.Now(),
			Priority:          taskConfig.priority,
			Parent:            taskConfig.parent,
		},
		Node:          taskConfig.node,
		MasterNodes:   taskConfig.masterNodes,
		clusterConfig: taskConfig.clusterConfig,
		Operation:     JointMasterOperation,
	}

	return task, nil
}
