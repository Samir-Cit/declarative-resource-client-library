// Copyright 2022 Google LLC. All Rights Reserved.
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
// gen_go_data -package alpha -var YAML_connector blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/vpcaccess/alpha/connector.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/vpcaccess/alpha/connector.yaml
var YAML_connector = []byte("info:\n  title: VPCAccess/Connector\n  description: The VPCAccess Connector resource\n  x-dcl-struct-name: Connector\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Connector\n    parameters:\n    - name: Connector\n      required: true\n      description: A full instance of a Connector\n  apply:\n    description: The function used to apply information about a Connector\n    parameters:\n    - name: Connector\n      required: true\n      description: A full instance of a Connector\n  delete:\n    description: The function used to delete a Connector\n    parameters:\n    - name: Connector\n      required: true\n      description: A full instance of a Connector\n  deleteAll:\n    description: The function used to delete all Connector\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Connector\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Connector:\n      title: Connector\n      x-dcl-id: projects/{{project}}/locations/{{location}}/connectors/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - project\n      - location\n      properties:\n        connectedProjects:\n          type: array\n          x-dcl-go-name: ConnectedProjects\n          readOnly: true\n          description: Output only. List of projects using the connector.\n          x-kubernetes-immutable: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n        ipCidrRange:\n          type: string\n          x-dcl-go-name: IPCidrRange\n          description: 'The range of internal addresses that follows RFC 4632 notation.\n            Example: `10.132.0.0/28`.'\n          x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        machineType:\n          type: string\n          x-dcl-go-name: MachineType\n          description: Machine type of VM Instance underlying connector. Default is\n            e2-micro\n          x-kubernetes-immutable: true\n        maxInstances:\n          type: integer\n          format: int64\n          x-dcl-go-name: MaxInstances\n          description: Maximum value of instances in autoscaling group underlying\n            the connector.\n          x-kubernetes-immutable: true\n        maxThroughput:\n          type: integer\n          format: int64\n          x-dcl-go-name: MaxThroughput\n          description: Maximum throughput of the connector in Mbps. Default is 200,\n            max is 1000.\n          x-kubernetes-immutable: true\n        minInstances:\n          type: integer\n          format: int64\n          x-dcl-go-name: MinInstances\n          description: Minimum value of instances in autoscaling group underlying\n            the connector.\n          x-kubernetes-immutable: true\n        minThroughput:\n          type: integer\n          format: int64\n          x-dcl-go-name: MinThroughput\n          description: Minimum throughput of the connector in Mbps. Default and min\n            is 200.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The resource name in the format `projects/*/locations/*/connectors/*`.\n          x-kubernetes-immutable: true\n        network:\n          type: string\n          x-dcl-go-name: Network\n          description: Name of a VPC network.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Compute/Network\n            field: name\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: ConnectorStateEnum\n          readOnly: true\n          description: 'Output only. State of the VPC access connector. Possible values:\n            STATE_UNSPECIFIED, READY, CREATING, DELETING, ERROR, UPDATING'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - READY\n          - CREATING\n          - DELETING\n          - ERROR\n          - UPDATING\n        subnet:\n          type: object\n          x-dcl-go-name: Subnet\n          x-dcl-go-type: ConnectorSubnet\n          description: The subnet in which to house the VPC Access Connector.\n          x-kubernetes-immutable: true\n          properties:\n            name:\n              type: string\n              x-dcl-go-name: Name\n              description: Subnet name (relative, not fully qualified). E.g. if the\n                full subnet selfLink is https://compute.googleapis.com/compute/v1/projects/{project}/regions/{region}/subnetworks/{subnetName}\n                the correct input for this field would be {subnetName}\n              x-kubernetes-immutable: true\n              x-dcl-references:\n              - resource: Compute/Subnetwork\n                field: name\n                parent: true\n            projectId:\n              type: string\n              x-dcl-go-name: ProjectId\n              description: Project in which the subnet exists. If not set, this project\n                is assumed to be the project for which the connector create request\n                was issued.\n              x-kubernetes-immutable: true\n              x-dcl-references:\n              - resource: Cloudresourcemanager/Project\n                field: name\n")

// 6093 bytes
// MD5: 8053a113f347496e5084a618bc280ddd
