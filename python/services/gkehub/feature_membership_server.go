// Copyright 2023 Google LLC. All Rights Reserved.
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
package server

import (
	"context"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	gkehubpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gkehub/gkehub_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub"
)

// FeatureMembershipServer implements the gRPC interface for FeatureMembership.
type FeatureMembershipServer struct{}

// ProtoToFeatureMembershipMeshManagementEnum converts a FeatureMembershipMeshManagementEnum enum from its proto representation.
func ProtoToGkehubFeatureMembershipMeshManagementEnum(e gkehubpb.GkehubFeatureMembershipMeshManagementEnum) *gkehub.FeatureMembershipMeshManagementEnum {
	if e == 0 {
		return nil
	}
	if n, ok := gkehubpb.GkehubFeatureMembershipMeshManagementEnum_name[int32(e)]; ok {
		e := gkehub.FeatureMembershipMeshManagementEnum(n[len("GkehubFeatureMembershipMeshManagementEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureMembershipMeshControlPlaneEnum converts a FeatureMembershipMeshControlPlaneEnum enum from its proto representation.
func ProtoToGkehubFeatureMembershipMeshControlPlaneEnum(e gkehubpb.GkehubFeatureMembershipMeshControlPlaneEnum) *gkehub.FeatureMembershipMeshControlPlaneEnum {
	if e == 0 {
		return nil
	}
	if n, ok := gkehubpb.GkehubFeatureMembershipMeshControlPlaneEnum_name[int32(e)]; ok {
		e := gkehub.FeatureMembershipMeshControlPlaneEnum(n[len("GkehubFeatureMembershipMeshControlPlaneEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum converts a FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum enum from its proto representation.
func ProtoToGkehubFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(e gkehubpb.GkehubFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum) *gkehub.FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := gkehubpb.GkehubFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum_name[int32(e)]; ok {
		e := gkehub.FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(n[len("GkehubFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureMembershipMesh converts a FeatureMembershipMesh object from its proto representation.
func ProtoToGkehubFeatureMembershipMesh(p *gkehubpb.GkehubFeatureMembershipMesh) *gkehub.FeatureMembershipMesh {
	if p == nil {
		return nil
	}
	obj := &gkehub.FeatureMembershipMesh{
		Management:   ProtoToGkehubFeatureMembershipMeshManagementEnum(p.GetManagement()),
		ControlPlane: ProtoToGkehubFeatureMembershipMeshControlPlaneEnum(p.GetControlPlane()),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagement converts a FeatureMembershipConfigmanagement object from its proto representation.
func ProtoToGkehubFeatureMembershipConfigmanagement(p *gkehubpb.GkehubFeatureMembershipConfigmanagement) *gkehub.FeatureMembershipConfigmanagement {
	if p == nil {
		return nil
	}
	obj := &gkehub.FeatureMembershipConfigmanagement{
		ConfigSync:          ProtoToGkehubFeatureMembershipConfigmanagementConfigSync(p.GetConfigSync()),
		PolicyController:    ProtoToGkehubFeatureMembershipConfigmanagementPolicyController(p.GetPolicyController()),
		Binauthz:            ProtoToGkehubFeatureMembershipConfigmanagementBinauthz(p.GetBinauthz()),
		HierarchyController: ProtoToGkehubFeatureMembershipConfigmanagementHierarchyController(p.GetHierarchyController()),
		Version:             dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementConfigSync converts a FeatureMembershipConfigmanagementConfigSync object from its proto representation.
func ProtoToGkehubFeatureMembershipConfigmanagementConfigSync(p *gkehubpb.GkehubFeatureMembershipConfigmanagementConfigSync) *gkehub.FeatureMembershipConfigmanagementConfigSync {
	if p == nil {
		return nil
	}
	obj := &gkehub.FeatureMembershipConfigmanagementConfigSync{
		Git:          ProtoToGkehubFeatureMembershipConfigmanagementConfigSyncGit(p.GetGit()),
		SourceFormat: dcl.StringOrNil(p.GetSourceFormat()),
		PreventDrift: dcl.Bool(p.GetPreventDrift()),
		Oci:          ProtoToGkehubFeatureMembershipConfigmanagementConfigSyncOci(p.GetOci()),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementConfigSyncGit converts a FeatureMembershipConfigmanagementConfigSyncGit object from its proto representation.
func ProtoToGkehubFeatureMembershipConfigmanagementConfigSyncGit(p *gkehubpb.GkehubFeatureMembershipConfigmanagementConfigSyncGit) *gkehub.FeatureMembershipConfigmanagementConfigSyncGit {
	if p == nil {
		return nil
	}
	obj := &gkehub.FeatureMembershipConfigmanagementConfigSyncGit{
		SyncRepo:               dcl.StringOrNil(p.GetSyncRepo()),
		SyncBranch:             dcl.StringOrNil(p.GetSyncBranch()),
		PolicyDir:              dcl.StringOrNil(p.GetPolicyDir()),
		SyncWaitSecs:           dcl.StringOrNil(p.GetSyncWaitSecs()),
		SyncRev:                dcl.StringOrNil(p.GetSyncRev()),
		SecretType:             dcl.StringOrNil(p.GetSecretType()),
		HttpsProxy:             dcl.StringOrNil(p.GetHttpsProxy()),
		GcpServiceAccountEmail: dcl.StringOrNil(p.GetGcpServiceAccountEmail()),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementConfigSyncOci converts a FeatureMembershipConfigmanagementConfigSyncOci object from its proto representation.
func ProtoToGkehubFeatureMembershipConfigmanagementConfigSyncOci(p *gkehubpb.GkehubFeatureMembershipConfigmanagementConfigSyncOci) *gkehub.FeatureMembershipConfigmanagementConfigSyncOci {
	if p == nil {
		return nil
	}
	obj := &gkehub.FeatureMembershipConfigmanagementConfigSyncOci{
		SyncRepo:               dcl.StringOrNil(p.GetSyncRepo()),
		PolicyDir:              dcl.StringOrNil(p.GetPolicyDir()),
		SyncWaitSecs:           dcl.StringOrNil(p.GetSyncWaitSecs()),
		SecretType:             dcl.StringOrNil(p.GetSecretType()),
		GcpServiceAccountEmail: dcl.StringOrNil(p.GetGcpServiceAccountEmail()),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementPolicyController converts a FeatureMembershipConfigmanagementPolicyController object from its proto representation.
func ProtoToGkehubFeatureMembershipConfigmanagementPolicyController(p *gkehubpb.GkehubFeatureMembershipConfigmanagementPolicyController) *gkehub.FeatureMembershipConfigmanagementPolicyController {
	if p == nil {
		return nil
	}
	obj := &gkehub.FeatureMembershipConfigmanagementPolicyController{
		Enabled:                  dcl.Bool(p.GetEnabled()),
		ReferentialRulesEnabled:  dcl.Bool(p.GetReferentialRulesEnabled()),
		LogDeniesEnabled:         dcl.Bool(p.GetLogDeniesEnabled()),
		MutationEnabled:          dcl.Bool(p.GetMutationEnabled()),
		Monitoring:               ProtoToGkehubFeatureMembershipConfigmanagementPolicyControllerMonitoring(p.GetMonitoring()),
		TemplateLibraryInstalled: dcl.Bool(p.GetTemplateLibraryInstalled()),
		AuditIntervalSeconds:     dcl.StringOrNil(p.GetAuditIntervalSeconds()),
	}
	for _, r := range p.GetExemptableNamespaces() {
		obj.ExemptableNamespaces = append(obj.ExemptableNamespaces, r)
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementPolicyControllerMonitoring converts a FeatureMembershipConfigmanagementPolicyControllerMonitoring object from its proto representation.
func ProtoToGkehubFeatureMembershipConfigmanagementPolicyControllerMonitoring(p *gkehubpb.GkehubFeatureMembershipConfigmanagementPolicyControllerMonitoring) *gkehub.FeatureMembershipConfigmanagementPolicyControllerMonitoring {
	if p == nil {
		return nil
	}
	obj := &gkehub.FeatureMembershipConfigmanagementPolicyControllerMonitoring{}
	for _, r := range p.GetBackends() {
		obj.Backends = append(obj.Backends, *ProtoToGkehubFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(r))
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementBinauthz converts a FeatureMembershipConfigmanagementBinauthz object from its proto representation.
func ProtoToGkehubFeatureMembershipConfigmanagementBinauthz(p *gkehubpb.GkehubFeatureMembershipConfigmanagementBinauthz) *gkehub.FeatureMembershipConfigmanagementBinauthz {
	if p == nil {
		return nil
	}
	obj := &gkehub.FeatureMembershipConfigmanagementBinauthz{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementHierarchyController converts a FeatureMembershipConfigmanagementHierarchyController object from its proto representation.
func ProtoToGkehubFeatureMembershipConfigmanagementHierarchyController(p *gkehubpb.GkehubFeatureMembershipConfigmanagementHierarchyController) *gkehub.FeatureMembershipConfigmanagementHierarchyController {
	if p == nil {
		return nil
	}
	obj := &gkehub.FeatureMembershipConfigmanagementHierarchyController{
		Enabled:                         dcl.Bool(p.GetEnabled()),
		EnablePodTreeLabels:             dcl.Bool(p.GetEnablePodTreeLabels()),
		EnableHierarchicalResourceQuota: dcl.Bool(p.GetEnableHierarchicalResourceQuota()),
	}
	return obj
}

// ProtoToFeatureMembership converts a FeatureMembership resource from its proto representation.
func ProtoToFeatureMembership(p *gkehubpb.GkehubFeatureMembership) *gkehub.FeatureMembership {
	obj := &gkehub.FeatureMembership{
		Mesh:             ProtoToGkehubFeatureMembershipMesh(p.GetMesh()),
		Configmanagement: ProtoToGkehubFeatureMembershipConfigmanagement(p.GetConfigmanagement()),
		Project:          dcl.StringOrNil(p.GetProject()),
		Location:         dcl.StringOrNil(p.GetLocation()),
		Feature:          dcl.StringOrNil(p.GetFeature()),
		Membership:       dcl.StringOrNil(p.GetMembership()),
	}
	return obj
}

// FeatureMembershipMeshManagementEnumToProto converts a FeatureMembershipMeshManagementEnum enum to its proto representation.
func GkehubFeatureMembershipMeshManagementEnumToProto(e *gkehub.FeatureMembershipMeshManagementEnum) gkehubpb.GkehubFeatureMembershipMeshManagementEnum {
	if e == nil {
		return gkehubpb.GkehubFeatureMembershipMeshManagementEnum(0)
	}
	if v, ok := gkehubpb.GkehubFeatureMembershipMeshManagementEnum_value["FeatureMembershipMeshManagementEnum"+string(*e)]; ok {
		return gkehubpb.GkehubFeatureMembershipMeshManagementEnum(v)
	}
	return gkehubpb.GkehubFeatureMembershipMeshManagementEnum(0)
}

// FeatureMembershipMeshControlPlaneEnumToProto converts a FeatureMembershipMeshControlPlaneEnum enum to its proto representation.
func GkehubFeatureMembershipMeshControlPlaneEnumToProto(e *gkehub.FeatureMembershipMeshControlPlaneEnum) gkehubpb.GkehubFeatureMembershipMeshControlPlaneEnum {
	if e == nil {
		return gkehubpb.GkehubFeatureMembershipMeshControlPlaneEnum(0)
	}
	if v, ok := gkehubpb.GkehubFeatureMembershipMeshControlPlaneEnum_value["FeatureMembershipMeshControlPlaneEnum"+string(*e)]; ok {
		return gkehubpb.GkehubFeatureMembershipMeshControlPlaneEnum(v)
	}
	return gkehubpb.GkehubFeatureMembershipMeshControlPlaneEnum(0)
}

// FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnumToProto converts a FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum enum to its proto representation.
func GkehubFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnumToProto(e *gkehub.FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum) gkehubpb.GkehubFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum {
	if e == nil {
		return gkehubpb.GkehubFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(0)
	}
	if v, ok := gkehubpb.GkehubFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum_value["FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum"+string(*e)]; ok {
		return gkehubpb.GkehubFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(v)
	}
	return gkehubpb.GkehubFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(0)
}

// FeatureMembershipMeshToProto converts a FeatureMembershipMesh object to its proto representation.
func GkehubFeatureMembershipMeshToProto(o *gkehub.FeatureMembershipMesh) *gkehubpb.GkehubFeatureMembershipMesh {
	if o == nil {
		return nil
	}
	p := &gkehubpb.GkehubFeatureMembershipMesh{}
	p.SetManagement(GkehubFeatureMembershipMeshManagementEnumToProto(o.Management))
	p.SetControlPlane(GkehubFeatureMembershipMeshControlPlaneEnumToProto(o.ControlPlane))
	return p
}

// FeatureMembershipConfigmanagementToProto converts a FeatureMembershipConfigmanagement object to its proto representation.
func GkehubFeatureMembershipConfigmanagementToProto(o *gkehub.FeatureMembershipConfigmanagement) *gkehubpb.GkehubFeatureMembershipConfigmanagement {
	if o == nil {
		return nil
	}
	p := &gkehubpb.GkehubFeatureMembershipConfigmanagement{}
	p.SetConfigSync(GkehubFeatureMembershipConfigmanagementConfigSyncToProto(o.ConfigSync))
	p.SetPolicyController(GkehubFeatureMembershipConfigmanagementPolicyControllerToProto(o.PolicyController))
	p.SetBinauthz(GkehubFeatureMembershipConfigmanagementBinauthzToProto(o.Binauthz))
	p.SetHierarchyController(GkehubFeatureMembershipConfigmanagementHierarchyControllerToProto(o.HierarchyController))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// FeatureMembershipConfigmanagementConfigSyncToProto converts a FeatureMembershipConfigmanagementConfigSync object to its proto representation.
func GkehubFeatureMembershipConfigmanagementConfigSyncToProto(o *gkehub.FeatureMembershipConfigmanagementConfigSync) *gkehubpb.GkehubFeatureMembershipConfigmanagementConfigSync {
	if o == nil {
		return nil
	}
	p := &gkehubpb.GkehubFeatureMembershipConfigmanagementConfigSync{}
	p.SetGit(GkehubFeatureMembershipConfigmanagementConfigSyncGitToProto(o.Git))
	p.SetSourceFormat(dcl.ValueOrEmptyString(o.SourceFormat))
	p.SetPreventDrift(dcl.ValueOrEmptyBool(o.PreventDrift))
	p.SetOci(GkehubFeatureMembershipConfigmanagementConfigSyncOciToProto(o.Oci))
	return p
}

// FeatureMembershipConfigmanagementConfigSyncGitToProto converts a FeatureMembershipConfigmanagementConfigSyncGit object to its proto representation.
func GkehubFeatureMembershipConfigmanagementConfigSyncGitToProto(o *gkehub.FeatureMembershipConfigmanagementConfigSyncGit) *gkehubpb.GkehubFeatureMembershipConfigmanagementConfigSyncGit {
	if o == nil {
		return nil
	}
	p := &gkehubpb.GkehubFeatureMembershipConfigmanagementConfigSyncGit{}
	p.SetSyncRepo(dcl.ValueOrEmptyString(o.SyncRepo))
	p.SetSyncBranch(dcl.ValueOrEmptyString(o.SyncBranch))
	p.SetPolicyDir(dcl.ValueOrEmptyString(o.PolicyDir))
	p.SetSyncWaitSecs(dcl.ValueOrEmptyString(o.SyncWaitSecs))
	p.SetSyncRev(dcl.ValueOrEmptyString(o.SyncRev))
	p.SetSecretType(dcl.ValueOrEmptyString(o.SecretType))
	p.SetHttpsProxy(dcl.ValueOrEmptyString(o.HttpsProxy))
	p.SetGcpServiceAccountEmail(dcl.ValueOrEmptyString(o.GcpServiceAccountEmail))
	return p
}

// FeatureMembershipConfigmanagementConfigSyncOciToProto converts a FeatureMembershipConfigmanagementConfigSyncOci object to its proto representation.
func GkehubFeatureMembershipConfigmanagementConfigSyncOciToProto(o *gkehub.FeatureMembershipConfigmanagementConfigSyncOci) *gkehubpb.GkehubFeatureMembershipConfigmanagementConfigSyncOci {
	if o == nil {
		return nil
	}
	p := &gkehubpb.GkehubFeatureMembershipConfigmanagementConfigSyncOci{}
	p.SetSyncRepo(dcl.ValueOrEmptyString(o.SyncRepo))
	p.SetPolicyDir(dcl.ValueOrEmptyString(o.PolicyDir))
	p.SetSyncWaitSecs(dcl.ValueOrEmptyString(o.SyncWaitSecs))
	p.SetSecretType(dcl.ValueOrEmptyString(o.SecretType))
	p.SetGcpServiceAccountEmail(dcl.ValueOrEmptyString(o.GcpServiceAccountEmail))
	return p
}

// FeatureMembershipConfigmanagementPolicyControllerToProto converts a FeatureMembershipConfigmanagementPolicyController object to its proto representation.
func GkehubFeatureMembershipConfigmanagementPolicyControllerToProto(o *gkehub.FeatureMembershipConfigmanagementPolicyController) *gkehubpb.GkehubFeatureMembershipConfigmanagementPolicyController {
	if o == nil {
		return nil
	}
	p := &gkehubpb.GkehubFeatureMembershipConfigmanagementPolicyController{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	p.SetReferentialRulesEnabled(dcl.ValueOrEmptyBool(o.ReferentialRulesEnabled))
	p.SetLogDeniesEnabled(dcl.ValueOrEmptyBool(o.LogDeniesEnabled))
	p.SetMutationEnabled(dcl.ValueOrEmptyBool(o.MutationEnabled))
	p.SetMonitoring(GkehubFeatureMembershipConfigmanagementPolicyControllerMonitoringToProto(o.Monitoring))
	p.SetTemplateLibraryInstalled(dcl.ValueOrEmptyBool(o.TemplateLibraryInstalled))
	p.SetAuditIntervalSeconds(dcl.ValueOrEmptyString(o.AuditIntervalSeconds))
	sExemptableNamespaces := make([]string, len(o.ExemptableNamespaces))
	for i, r := range o.ExemptableNamespaces {
		sExemptableNamespaces[i] = r
	}
	p.SetExemptableNamespaces(sExemptableNamespaces)
	return p
}

// FeatureMembershipConfigmanagementPolicyControllerMonitoringToProto converts a FeatureMembershipConfigmanagementPolicyControllerMonitoring object to its proto representation.
func GkehubFeatureMembershipConfigmanagementPolicyControllerMonitoringToProto(o *gkehub.FeatureMembershipConfigmanagementPolicyControllerMonitoring) *gkehubpb.GkehubFeatureMembershipConfigmanagementPolicyControllerMonitoring {
	if o == nil {
		return nil
	}
	p := &gkehubpb.GkehubFeatureMembershipConfigmanagementPolicyControllerMonitoring{}
	sBackends := make([]gkehubpb.GkehubFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum, len(o.Backends))
	for i, r := range o.Backends {
		sBackends[i] = gkehubpb.GkehubFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(gkehubpb.GkehubFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum_value[string(r)])
	}
	p.SetBackends(sBackends)
	return p
}

// FeatureMembershipConfigmanagementBinauthzToProto converts a FeatureMembershipConfigmanagementBinauthz object to its proto representation.
func GkehubFeatureMembershipConfigmanagementBinauthzToProto(o *gkehub.FeatureMembershipConfigmanagementBinauthz) *gkehubpb.GkehubFeatureMembershipConfigmanagementBinauthz {
	if o == nil {
		return nil
	}
	p := &gkehubpb.GkehubFeatureMembershipConfigmanagementBinauthz{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	return p
}

// FeatureMembershipConfigmanagementHierarchyControllerToProto converts a FeatureMembershipConfigmanagementHierarchyController object to its proto representation.
func GkehubFeatureMembershipConfigmanagementHierarchyControllerToProto(o *gkehub.FeatureMembershipConfigmanagementHierarchyController) *gkehubpb.GkehubFeatureMembershipConfigmanagementHierarchyController {
	if o == nil {
		return nil
	}
	p := &gkehubpb.GkehubFeatureMembershipConfigmanagementHierarchyController{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	p.SetEnablePodTreeLabels(dcl.ValueOrEmptyBool(o.EnablePodTreeLabels))
	p.SetEnableHierarchicalResourceQuota(dcl.ValueOrEmptyBool(o.EnableHierarchicalResourceQuota))
	return p
}

// FeatureMembershipToProto converts a FeatureMembership resource to its proto representation.
func FeatureMembershipToProto(resource *gkehub.FeatureMembership) *gkehubpb.GkehubFeatureMembership {
	p := &gkehubpb.GkehubFeatureMembership{}
	p.SetMesh(GkehubFeatureMembershipMeshToProto(resource.Mesh))
	p.SetConfigmanagement(GkehubFeatureMembershipConfigmanagementToProto(resource.Configmanagement))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetFeature(dcl.ValueOrEmptyString(resource.Feature))
	p.SetMembership(dcl.ValueOrEmptyString(resource.Membership))

	return p
}

// applyFeatureMembership handles the gRPC request by passing it to the underlying FeatureMembership Apply() method.
func (s *FeatureMembershipServer) applyFeatureMembership(ctx context.Context, c *gkehub.Client, request *gkehubpb.ApplyGkehubFeatureMembershipRequest) (*gkehubpb.GkehubFeatureMembership, error) {
	p := ProtoToFeatureMembership(request.GetResource())
	res, err := c.ApplyFeatureMembership(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FeatureMembershipToProto(res)
	return r, nil
}

// applyGkehubFeatureMembership handles the gRPC request by passing it to the underlying FeatureMembership Apply() method.
func (s *FeatureMembershipServer) ApplyGkehubFeatureMembership(ctx context.Context, request *gkehubpb.ApplyGkehubFeatureMembershipRequest) (*gkehubpb.GkehubFeatureMembership, error) {
	cl, err := createConfigFeatureMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyFeatureMembership(ctx, cl, request)
}

// DeleteFeatureMembership handles the gRPC request by passing it to the underlying FeatureMembership Delete() method.
func (s *FeatureMembershipServer) DeleteGkehubFeatureMembership(ctx context.Context, request *gkehubpb.DeleteGkehubFeatureMembershipRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFeatureMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFeatureMembership(ctx, ProtoToFeatureMembership(request.GetResource()))

}

// ListGkehubFeatureMembership handles the gRPC request by passing it to the underlying FeatureMembershipList() method.
func (s *FeatureMembershipServer) ListGkehubFeatureMembership(ctx context.Context, request *gkehubpb.ListGkehubFeatureMembershipRequest) (*gkehubpb.ListGkehubFeatureMembershipResponse, error) {
	cl, err := createConfigFeatureMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFeatureMembership(ctx, request.GetProject(), request.GetLocation(), request.GetFeature())
	if err != nil {
		return nil, err
	}
	var protos []*gkehubpb.GkehubFeatureMembership
	for _, r := range resources.Items {
		rp := FeatureMembershipToProto(r)
		protos = append(protos, rp)
	}
	p := &gkehubpb.ListGkehubFeatureMembershipResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigFeatureMembership(ctx context.Context, service_account_file string) (*gkehub.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return gkehub.NewClient(conf), nil
}