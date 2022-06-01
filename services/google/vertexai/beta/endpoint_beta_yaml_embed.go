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
// gen_go_data -package beta -var YAML_endpoint blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/vertexai/beta/endpoint.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/vertexai/beta/endpoint.yaml
var YAML_endpoint = []byte("info:\n  title: VertexAI/Endpoint\n  description: The VertexAI Endpoint resource\n  x-dcl-struct-name: Endpoint\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Endpoint\n    parameters:\n    - name: Endpoint\n      required: true\n      description: A full instance of a Endpoint\n  apply:\n    description: The function used to apply information about a Endpoint\n    parameters:\n    - name: Endpoint\n      required: true\n      description: A full instance of a Endpoint\n  delete:\n    description: The function used to delete a Endpoint\n    parameters:\n    - name: Endpoint\n      required: true\n      description: A full instance of a Endpoint\n  deleteAll:\n    description: The function used to delete all Endpoint\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Endpoint\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Endpoint:\n      title: Endpoint\n      x-dcl-id: projects/{{project}}/locations/{{location}}/endpoints/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - displayName\n      - project\n      - location\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. Timestamp when this Endpoint was created.\n          x-kubernetes-immutable: true\n        deployedModels:\n          type: array\n          x-dcl-go-name: DeployedModels\n          readOnly: true\n          description: Output only. The models deployed in this Endpoint. To add or\n            remove DeployedModels use EndpointService.DeployModel and EndpointService.UndeployModel\n            respectively.\n          x-kubernetes-immutable: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: EndpointDeployedModels\n            properties:\n              automaticResources:\n                type: object\n                x-dcl-go-name: AutomaticResources\n                x-dcl-go-type: EndpointDeployedModelsAutomaticResources\n                description: A description of resources that to large degree are decided\n                  by Vertex AI, and require only a modest additional configuration.\n                x-kubernetes-immutable: true\n                x-dcl-conflicts:\n                - dedicatedResources\n                properties:\n                  maxReplicaCount:\n                    type: integer\n                    format: int64\n                    x-dcl-go-name: MaxReplicaCount\n                    description: The maximum number of replicas this DeployedModel\n                      may be deployed on when the traffic against it increases. If\n                      the requested value is too large, the deployment will error,\n                      but if deployment succeeds then the ability to scale the model\n                      to that many replicas is guaranteed (barring service outages).\n                      If traffic against the DeployedModel increases beyond what its\n                      replicas at maximum may handle, a portion of the traffic will\n                      be dropped. If this value is not provided, a no upper bound\n                      for scaling under heavy traffic will be assume, though Vertex\n                      AI may be unable to scale beyond certain replica number.\n                    x-kubernetes-immutable: true\n                  minReplicaCount:\n                    type: integer\n                    format: int64\n                    x-dcl-go-name: MinReplicaCount\n                    description: The minimum number of replicas this DeployedModel\n                      will be always deployed on. If traffic against it increases,\n                      it may dynamically be deployed onto more replicas up to max_replica_count,\n                      and as traffic decreases, some of these extra replicas may be\n                      freed. If the requested value is too large, the deployment will\n                      error.\n                    x-kubernetes-immutable: true\n              createTime:\n                type: string\n                format: date-time\n                x-dcl-go-name: CreateTime\n                readOnly: true\n                description: Output only. Timestamp when the DeployedModel was created.\n                x-kubernetes-immutable: true\n              dedicatedResources:\n                type: object\n                x-dcl-go-name: DedicatedResources\n                x-dcl-go-type: EndpointDeployedModelsDedicatedResources\n                description: A description of resources that are dedicated to the\n                  DeployedModel, and that need a higher degree of manual configuration.\n                x-kubernetes-immutable: true\n                x-dcl-conflicts:\n                - automaticResources\n                properties:\n                  autoscalingMetricSpecs:\n                    type: array\n                    x-dcl-go-name: AutoscalingMetricSpecs\n                    description: The metric specifications that overrides a resource\n                      utilization metric (CPU utilization, accelerator's duty cycle,\n                      and so on) target value (default to 60 if not set). At most\n                      one entry is allowed per metric. If machine_spec.accelerator_count\n                      is above 0, the autoscaling will be based on both CPU utilization\n                      and accelerator's duty cycle metrics and scale up when either\n                      metrics exceeds its target value while scale down if both metrics\n                      are under their target value. The default target value is 60\n                      for both metrics. If machine_spec.accelerator_count is 0, the\n                      autoscaling will be based on CPU utilization metric only with\n                      default target value 60 if not explicitly set. For example,\n                      in the case of Online Prediction, if you want to override target\n                      CPU utilization to 80, you should set autoscaling_metric_specs.metric_name\n                      to `aiplatform.googleapis.com/prediction/online/cpu/utilization`\n                      and autoscaling_metric_specs.target to `80`.\n                    x-kubernetes-immutable: true\n                    x-dcl-send-empty: true\n                    x-dcl-list-type: list\n                    items:\n                      type: object\n                      x-dcl-go-type: EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs\n                      properties:\n                        metricName:\n                          type: string\n                          x-dcl-go-name: MetricName\n                          description: 'The resource metric name. Supported metrics:\n                            * For Online Prediction: * `aiplatform.googleapis.com/prediction/online/accelerator/duty_cycle`\n                            * `aiplatform.googleapis.com/prediction/online/cpu/utilization`'\n                          x-kubernetes-immutable: true\n                        target:\n                          type: integer\n                          format: int64\n                          x-dcl-go-name: Target\n                          description: The target resource utilization in percentage\n                            (1% - 100%) for the given metric; once the real usage\n                            deviates from the target by a certain percentage, the\n                            machine replicas change. The default value is 60 (representing\n                            60%) if not provided.\n                          x-kubernetes-immutable: true\n                  machineSpec:\n                    type: object\n                    x-dcl-go-name: MachineSpec\n                    x-dcl-go-type: EndpointDeployedModelsDedicatedResourcesMachineSpec\n                    description: The specification of a single machine used by the\n                      prediction.\n                    x-kubernetes-immutable: true\n                    properties:\n                      acceleratorCount:\n                        type: integer\n                        format: int64\n                        x-dcl-go-name: AcceleratorCount\n                        description: The number of accelerators to attach to the machine.\n                        x-kubernetes-immutable: true\n                      acceleratorType:\n                        type: string\n                        x-dcl-go-name: AcceleratorType\n                        x-dcl-go-type: EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum\n                        description: 'The type of accelerator(s) that may be attached\n                          to the machine as per accelerator_count. Possible values:\n                          ACCELERATOR_TYPE_UNSPECIFIED, NVIDIA_TESLA_K80, NVIDIA_TESLA_P100,\n                          NVIDIA_TESLA_V100, NVIDIA_TESLA_P4, NVIDIA_TESLA_T4, NVIDIA_TESLA_A100,\n                          TPU_V2, TPU_V3'\n                        x-kubernetes-immutable: true\n                        enum:\n                        - ACCELERATOR_TYPE_UNSPECIFIED\n                        - NVIDIA_TESLA_K80\n                        - NVIDIA_TESLA_P100\n                        - NVIDIA_TESLA_V100\n                        - NVIDIA_TESLA_P4\n                        - NVIDIA_TESLA_T4\n                        - NVIDIA_TESLA_A100\n                        - TPU_V2\n                        - TPU_V3\n                      machineType:\n                        type: string\n                        x-dcl-go-name: MachineType\n                        description: 'The type of the machine. See the [list of machine\n                          types supported for prediction](https://cloud.google.com/vertex-ai/docs/predictions/configure-compute#machine-types)\n                          See the [list of machine types supported for custom training](https://cloud.google.com/vertex-ai/docs/training/configure-compute#machine-types).\n                          For DeployedModel this field is optional, and the default\n                          value is `n1-standard-2`. For BatchPredictionJob or as part\n                          of WorkerPoolSpec this field is required. TODO(rsurowka):\n                          Try to better unify the required vs optional.'\n                        x-kubernetes-immutable: true\n                  maxReplicaCount:\n                    type: integer\n                    format: int64\n                    x-dcl-go-name: MaxReplicaCount\n                    description: The maximum number of replicas this DeployedModel\n                      may be deployed on when the traffic against it increases. If\n                      the requested value is too large, the deployment will error,\n                      but if deployment succeeds then the ability to scale the model\n                      to that many replicas is guaranteed (barring service outages).\n                      If traffic against the DeployedModel increases beyond what its\n                      replicas at maximum may handle, a portion of the traffic will\n                      be dropped. If this value is not provided, will use min_replica_count\n                      as the default value. The value of this field impacts the charge\n                      against Vertex CPU and GPU quotas. Specifically, you will be\n                      charged for max_replica_count * number of cores in the selected\n                      machine type) and (max_replica_count * number of GPUs per replica\n                      in the selected machine type).\n                    x-kubernetes-immutable: true\n                  minReplicaCount:\n                    type: integer\n                    format: int64\n                    x-dcl-go-name: MinReplicaCount\n                    description: The minimum number of machine replicas this DeployedModel\n                      will be always deployed on. This value must be greater than\n                      or equal to 1. If traffic against the DeployedModel increases,\n                      it may dynamically be deployed onto more replicas, and as traffic\n                      decreases, some of these extra replicas may be freed.\n                    x-kubernetes-immutable: true\n              disableContainerLogging:\n                type: boolean\n                x-dcl-go-name: DisableContainerLogging\n                description: For custom-trained Models and AutoML Tabular Models,\n                  the container of the DeployedModel instances will send `stderr`\n                  and `stdout` streams to Stackdriver Logging by default. Please note\n                  that the logs incur cost, which are subject to [Cloud Logging pricing](https://cloud.google.com/stackdriver/pricing).\n                  User can disable container logging by setting this flag to true.\n                x-kubernetes-immutable: true\n              displayName:\n                type: string\n                x-dcl-go-name: DisplayName\n                description: The display name of the DeployedModel. If not provided\n                  upon creation, the Model's display_name is used.\n                x-kubernetes-immutable: true\n              enableAccessLogging:\n                type: boolean\n                x-dcl-go-name: EnableAccessLogging\n                description: These logs are like standard server access logs, containing\n                  information like timestamp and latency for each prediction request.\n                  Note that Stackdriver logs may incur a cost, especially if your\n                  project receives prediction requests at a high queries per second\n                  rate (QPS). Estimate your costs before enabling this option.\n                x-kubernetes-immutable: true\n              enableContainerLogging:\n                type: boolean\n                x-dcl-go-name: EnableContainerLogging\n                description: If true, the container of the DeployedModel instances\n                  will send `stderr` and `stdout` streams to Stackdriver Logging.\n                  Only supported for custom-trained Models and AutoML Tabular Models.\n                x-kubernetes-immutable: true\n              id:\n                type: string\n                x-dcl-go-name: Id\n                description: The ID of the DeployedModel. If not provided upon deployment,\n                  Vertex AI will generate a value for this ID. This value should be\n                  1-10 characters, and valid characters are /[0-9]/.\n                x-kubernetes-immutable: true\n              model:\n                type: string\n                x-dcl-go-name: Model\n                description: The name of the Model that this is the deployment of.\n                  Note that the Model may be in a different location than the DeployedModel's\n                  Endpoint.\n                x-kubernetes-immutable: true\n                x-dcl-references:\n                - resource: Aiplatform/Model\n                  field: selfLink\n              modelVersionId:\n                type: string\n                x-dcl-go-name: ModelVersionId\n                readOnly: true\n                description: Output only. The version ID of the model that is deployed.\n                x-kubernetes-immutable: true\n              privateEndpoints:\n                type: object\n                x-dcl-go-name: PrivateEndpoints\n                x-dcl-go-type: EndpointDeployedModelsPrivateEndpoints\n                readOnly: true\n                description: Output only. Provide paths for users to send predict/explain/health\n                  requests directly to the deployed model services running on Cloud\n                  via private services access. This field is populated if network\n                  is configured.\n                x-kubernetes-immutable: true\n                properties:\n                  explainHttpUri:\n                    type: string\n                    x-dcl-go-name: ExplainHttpUri\n                    readOnly: true\n                    description: Output only. Http(s) path to send explain requests.\n                    x-kubernetes-immutable: true\n                  healthHttpUri:\n                    type: string\n                    x-dcl-go-name: HealthHttpUri\n                    readOnly: true\n                    description: Output only. Http(s) path to send health check requests.\n                    x-kubernetes-immutable: true\n                  predictHttpUri:\n                    type: string\n                    x-dcl-go-name: PredictHttpUri\n                    readOnly: true\n                    description: Output only. Http(s) path to send prediction requests.\n                    x-kubernetes-immutable: true\n                  serviceAttachment:\n                    type: string\n                    x-dcl-go-name: ServiceAttachment\n                    readOnly: true\n                    description: Output only. The name of the service attachment resource.\n                      Populated if private service connect is enabled.\n                    x-kubernetes-immutable: true\n              serviceAccount:\n                type: string\n                x-dcl-go-name: ServiceAccount\n                description: The service account that the DeployedModel's container\n                  runs as. Specify the email address of the service account. If this\n                  service account is not specified, the container runs as a service\n                  account that doesn't have access to the resource project. Users\n                  deploying the Model must have the `iam.serviceAccounts.actAs` permission\n                  on this service account.\n                x-kubernetes-immutable: true\n              sharedResources:\n                type: string\n                x-dcl-go-name: SharedResources\n                description: 'The resource name of the shared DeploymentResourcePool\n                  to deploy on. Format: projects/{project}/locations/{location}/deploymentResourcePools/{deployment_resource_pool}'\n                x-kubernetes-immutable: true\n                x-dcl-references:\n                - resource: Aiplatform/DeploymentResourcePool\n                  field: selfLink\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: The description of the Endpoint.\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: Required. The display name of the Endpoint. The name can be\n            up to 128 characters long and can be consist of any UTF-8 characters.\n        encryptionSpec:\n          type: object\n          x-dcl-go-name: EncryptionSpec\n          x-dcl-go-type: EndpointEncryptionSpec\n          description: Customer-managed encryption key spec for an Endpoint. If set,\n            this Endpoint and all sub-resources of this Endpoint will be secured by\n            this key.\n          x-kubernetes-immutable: true\n          required:\n          - kmsKeyName\n          properties:\n            kmsKeyName:\n              type: string\n              x-dcl-go-name: KmsKeyName\n              description: 'Required. The Cloud KMS resource identifier of the customer\n                managed encryption key used to protect a resource. Has the form: `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.\n                The key needs to be in the same region as where the compute resource\n                is created.'\n              x-kubernetes-immutable: true\n              x-dcl-references:\n              - resource: Cloudkms/CryptoKey\n                field: name\n        etag:\n          type: string\n          x-dcl-go-name: Etag\n          readOnly: true\n          description: Used to perform consistent read-modify-write updates. If not\n            set, a blind \"overwrite\" update happens.\n          x-kubernetes-immutable: true\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: The labels with user-defined metadata to organize your Endpoints.\n            Label keys and values can be no longer than 64 characters (Unicode codepoints),\n            can only contain lowercase letters, numeric characters, underscores and\n            dashes. International characters are allowed. See https://goo.gl/xmQnxf\n            for more information and examples of labels.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        modelDeploymentMonitoringJob:\n          type: string\n          x-dcl-go-name: ModelDeploymentMonitoringJob\n          readOnly: true\n          description: 'Output only. Resource name of the Model Monitoring job associated\n            with this Endpoint if monitoring is enabled by CreateModelDeploymentMonitoringJob.\n            Format: `projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{model_deployment_monitoring_job}`'\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Output only. The resource name of the Endpoint.\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n        network:\n          type: string\n          x-dcl-go-name: Network\n          description: 'The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks)\n            to which the Endpoint should be peered. Private services access must already\n            be configured for the network. If left unspecified, the Endpoint is not\n            peered with any network. Only one of the fields, network or enable_private_service_connect,\n            can be set. [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert):\n            `projects/{project}/global/networks/{network}`. Where `{project}` is a\n            project number, as in `12345`, and `{network}` is network name.'\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Compute/Network\n            field: selfLink\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. Timestamp when this Endpoint was last updated.\n          x-kubernetes-immutable: true\n")

// 23210 bytes
// MD5: c77af7d4cf3820482159279a98691fb3
