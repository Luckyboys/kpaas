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

package server

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"

	pb "github.com/kpaas-io/kpaas/pkg/deploy/protos"
	"github.com/kpaas-io/kpaas/pkg/deploy/task"
)

func (c *controller) getCheckNodesLog(aTask task.Task, nodeName string) (*pb.GetCheckNodesLogReply, error) {
	if aTask == nil {
		return nil, fmt.Errorf("Task is nil")
	}

	// Collect all log files into a buffer and return its bytes to the caller.
	// Notes, this implementation is based on the assumption that the log files' size will not be huge (less than
	// several hundred MBs). If they are huge, we need to consider to send back the log via stream (gRPC support this).
	var buf bytes.Buffer

	// Get all actions of the check nodes task
	actions := task.GetAllActions(aTask)
	for _, act := range actions {
		node := act.GetNode()
		if node == nil || node.GetName() != nodeName {
			continue
		}
		logFilePath := act.GetLogFilePath()
		if logFilePath == "" {
			continue
		}
		logFile, err := os.Open(logFilePath)
		if err != nil {
			logrus.Errorf("Open log file %q failed: %v", logFilePath, err)
			return nil, err
		}
		defer logFile.Close()

		if _, err = io.Copy(&buf, logFile); err != nil {
			logrus.Warnf("Copy log file %q failed: %v", logFilePath, err)
			continue
		}

		// Append two blank lines between each log files
		buf.WriteString("\n\n")
	}

	result := &pb.GetCheckNodesLogReply{
		Log: buf.Bytes(),
	}

	logrus.Debugf("Finish getCheckNodesLog, log length: %d", len(result.Log))

	return result, nil
}
