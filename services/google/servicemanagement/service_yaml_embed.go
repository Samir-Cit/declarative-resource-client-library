// Copyright 2021 Google LLC. All Rights Reserved.
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
// GENERATED BY gen_go_data.go
// gen_go_data -package servicemanagement -var YAML_service blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/servicemanagement/service.yaml

package servicemanagement

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/servicemanagement/service.yaml
var YAML_service = []byte("info:\n  title: Servicemanagement/Service\n  description: DCL Specification for the Servicemanagement Service resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Service\n    parameters:\n    - name: Service\n      required: true\n      description: A full instance of a Service\n  apply:\n    description: The function used to apply information about a Service\n    parameters:\n    - name: Service\n      required: true\n      description: A full instance of a Service\n  delete:\n    description: The function used to delete a Service\n    parameters:\n    - name: Service\n      required: true\n      description: A full instance of a Service\n  deleteAll:\n    description: The function used to delete all Service\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Service\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Service:\n      title: Service\n      x-dcl-id: services/{{name}}\n      x-dcl-parent-container: project\n      type: object\n      properties:\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The name of the service. See the overview for naming requirements.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: ID of the project that produces and owns this service.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n")

// 1701 bytes
// MD5: bb2ddc50385698009242060fdf340169