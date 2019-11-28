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

package deploy

import (
	"fmt"

	"github.com/kpaas-io/kpaas/pkg/service/model/api"
	"github.com/kpaas-io/kpaas/pkg/service/model/common"
	"github.com/kpaas-io/kpaas/pkg/service/model/wizard"
)

func convertModelAuthenticationTypeToAPIAuthenticationType(authenticationType wizard.AuthenticationType) api.AuthenticationType {

	switch authenticationType {
	case wizard.AuthenticationTypePassword:
		return api.AuthenticationTypePassword
	case wizard.AuthenticationTypePrivateKey:
		return api.AuthenticationTypePrivateKey
	}

	return api.AuthenticationType(fmt.Sprintf("unknown(%s)", authenticationType))
}

func convertModelLabelToAPILabel(label *wizard.Label) api.Label {

	return api.Label{
		Key:   label.Key,
		Value: label.Value,
	}
}

func convertModelAnnotationToAPIAnnotation(annotation *wizard.Annotation) api.Annotation {

	return api.Annotation{
		Key:   annotation.Key,
		Value: annotation.Value,
	}
}

func convertModelTaintToAPITaint(taint *wizard.Taint) api.Taint {

	return api.Taint{
		Key:    taint.Key,
		Value:  taint.Value,
		Effect: convertModelTaintEffectToAPITaintEffect(taint.Effect),
	}
}

func convertModelTaintEffectToAPITaintEffect(effect wizard.TaintEffect) api.TaintEffect {
	switch effect {
	case wizard.TaintEffectNoExecute:
		return api.TaintEffectNoExecute
	case wizard.TaintEffectNoSchedule:
		return api.TaintEffectNoSchedule
	case wizard.TaintEffectPreferNoSchedule:
		return api.TaintEffectPreferNoSchedule
	}

	return api.TaintEffect(fmt.Sprintf("unknown(%s)", effect))
}

func convertModelCheckResultToAPICheckResult(result wizard.CheckResult) api.CheckResult {

	switch result {
	case wizard.CheckResultNotRunning:
		return api.CheckResultNotRunning
	case wizard.CheckResultChecking:
		return api.CheckResultChecking
	case wizard.CheckResultFailed:
		return api.CheckResultFailed
	case wizard.CheckResultPassed:
		return api.CheckResultPassed
	}

	return api.CheckResult(fmt.Sprintf("unknown(%s)", result))
}

func convertModelErrorToAPIError(detail *common.FailureDetail) *api.Error {

	if detail == nil {
		return nil
	}
	return &api.Error{
		Reason:     detail.Reason,
		Detail:     detail.Detail,
		FixMethods: detail.FixMethods,
		LogId:      detail.LogId,
	}
}

func convertModelDeployClusterStatusToAPIDeployClusterStatus(status wizard.DeployClusterStatus) api.DeployClusterStatus {

	switch status {
	case wizard.DeployClusterStatusNotRunning:
		return api.DeployClusterStatusNotRunning
	case wizard.DeployClusterStatusRunning:
		return api.DeployClusterStatusRunning
	case wizard.DeployClusterStatusSuccessful:
		return api.DeployClusterStatusSuccessful
	case wizard.DeployClusterStatusFailed:
		return api.DeployClusterStatusFailed
	case wizard.DeployClusterStatusWorkedButHaveError:
		return api.DeployClusterStatusWorkedButHaveError
	}

	return api.DeployClusterStatus(fmt.Sprintf("unknown(%s)", status))
}

func convertModelMachineRoleToAPIMachineRole(machineRole wizard.MachineRole) api.MachineRole {

	switch machineRole {
	case wizard.MachineRoleMaster:
		return api.MachineRoleMaster
	case wizard.MachineRoleWorker:
		return api.MachineRoleWorker
	case wizard.MachineRoleEtcd:
		return api.MachineRoleEtcd
	}

	return api.MachineRole(fmt.Sprintf("unknown(%s)", machineRole))
}

func convertModelDeployStatusToAPIDeployStatus(status wizard.DeployStatus) api.DeployStatus {

	switch status {
	case wizard.DeployStatusPending:
		return api.DeployStatusPending
	case wizard.DeployStatusDeploying:
		return api.DeployStatusDeploying
	case wizard.DeployStatusCompleted:
		return api.DeployStatusCompleted
	case wizard.DeployStatusFailed:
		return api.DeployStatusFailed
	case wizard.DeployStatusAborted:
		return api.DeployStatusAborted
	}

	return api.DeployStatus(fmt.Sprintf("unknown(%s)", status))
}

func convertAPIMachineRoleToModelMachineRole(role api.MachineRole) wizard.MachineRole {

	switch role {
	case api.MachineRoleMaster:
		return wizard.MachineRoleMaster
	case api.MachineRoleWorker:
		return wizard.MachineRoleWorker
	case api.MachineRoleEtcd:
		return wizard.MachineRoleEtcd
	}

	return wizard.MachineRole(fmt.Sprintf("unknown(%s)", role))
}

func convertAPITaintEffectToModelTaintEffect(effect api.TaintEffect) wizard.TaintEffect {
	switch effect {
	case api.TaintEffectNoExecute:
		return wizard.TaintEffectNoExecute
	case api.TaintEffectNoSchedule:
		return wizard.TaintEffectNoSchedule
	case api.TaintEffectPreferNoSchedule:
		return wizard.TaintEffectPreferNoSchedule
	}

	return wizard.TaintEffect(fmt.Sprintf("unknown(%s)", effect))
}

func convertAPIAuthenticationTypeToModelAuthenticationType(authenticationType api.AuthenticationType) wizard.AuthenticationType {

	switch authenticationType {
	case api.AuthenticationTypePassword:
		return wizard.AuthenticationTypePassword
	case api.AuthenticationTypePrivateKey:
		return wizard.AuthenticationTypePrivateKey
	}

	return wizard.AuthenticationType(fmt.Sprintf("unknown(%s)", authenticationType))
}

func convertModelNodeToAPINode(node *wizard.Node) *api.NodeData {

	machineRoles := make([]api.MachineRole, 0, len(node.MachineRoles))
	for _, role := range node.MachineRoles {
		machineRoles = append(machineRoles, convertModelMachineRoleToAPIMachineRole(role))
	}

	labels := make([]api.Label, 0, len(node.Labels))
	for _, label := range node.Labels {
		labels = append(labels, convertModelLabelToAPILabel(label))
	}

	taints := make([]api.Taint, 0, len(node.Taints))
	for _, taint := range node.Taints {
		taints = append(taints, convertModelTaintToAPITaint(taint))
	}

	return &api.NodeData{
		NodeBaseData: api.NodeBaseData{
			Name:                node.Name,
			Description:         node.Description,
			MachineRole:         machineRoles,
			Labels:              labels,
			Taints:              taints,
			DockerRootDirectory: node.DockerRootDirectory,
		},
		ConnectionData: api.ConnectionData{
			IP:   node.IP,
			Port: node.Port,
			SSHLoginData: api.SSHLoginData{
				Username:           node.Username,
				AuthenticationType: convertModelAuthenticationTypeToAPIAuthenticationType(node.AuthenticationType),
				PrivateKeyName:     node.PrivateKeyName,
			},
		},
	}
}