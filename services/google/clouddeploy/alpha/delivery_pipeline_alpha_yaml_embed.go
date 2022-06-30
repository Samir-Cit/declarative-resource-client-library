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
// gen_go_data -package alpha -var YAML_delivery_pipeline blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/clouddeploy/alpha/delivery_pipeline.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/clouddeploy/alpha/delivery_pipeline.yaml
var YAML_delivery_pipeline = []byte("info:\n  title: Clouddeploy/DeliveryPipeline\n  description: The Cloud Deploy `DeliveryPipeline` resource\n  x-dcl-struct-name: DeliveryPipeline\n  x-dcl-has-iam: false\n  x-dcl-ref:\n    text: REST API\n    url: https://cloud.google.com/deploy/docs/api/reference/rest/v1/projects.locations.deliveryPipelines\npaths:\n  get:\n    description: The function used to get information about a DeliveryPipeline\n    parameters:\n    - name: DeliveryPipeline\n      required: true\n      description: A full instance of a DeliveryPipeline\n  apply:\n    description: The function used to apply information about a DeliveryPipeline\n    parameters:\n    - name: DeliveryPipeline\n      required: true\n      description: A full instance of a DeliveryPipeline\n  delete:\n    description: The function used to delete a DeliveryPipeline\n    parameters:\n    - name: DeliveryPipeline\n      required: true\n      description: A full instance of a DeliveryPipeline\n  deleteAll:\n    description: The function used to delete all DeliveryPipeline\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many DeliveryPipeline\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    DeliveryPipeline:\n      title: DeliveryPipeline\n      x-dcl-id: projects/{{project}}/locations/{{location}}/deliveryPipelines/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - project\n      - location\n      properties:\n        annotations:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Annotations\n          description: User annotations. These attributes can only be set and used\n            by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations\n            for more details such as format and size limitations.\n        condition:\n          type: object\n          x-dcl-go-name: Condition\n          x-dcl-go-type: DeliveryPipelineCondition\n          readOnly: true\n          description: Output only. Information around the state of the Delivery Pipeline.\n          properties:\n            pipelineReadyCondition:\n              type: object\n              x-dcl-go-name: PipelineReadyCondition\n              x-dcl-go-type: DeliveryPipelineConditionPipelineReadyCondition\n              description: Details around the Pipeline's overall status.\n              properties:\n                status:\n                  type: boolean\n                  x-dcl-go-name: Status\n                  description: True if the Pipeline is in a valid state. Otherwise\n                    at least one condition in `PipelineCondition` is in an invalid\n                    state. Iterate over those conditions and see which condition(s)\n                    has status = false to find out what is wrong with the Pipeline.\n                updateTime:\n                  type: string\n                  format: date-time\n                  x-dcl-go-name: UpdateTime\n                  description: Last time the condition was updated.\n            targetsPresentCondition:\n              type: object\n              x-dcl-go-name: TargetsPresentCondition\n              x-dcl-go-type: DeliveryPipelineConditionTargetsPresentCondition\n              description: Detalis around targets enumerated in the pipeline.\n              properties:\n                missingTargets:\n                  type: array\n                  x-dcl-go-name: MissingTargets\n                  description: The list of Target names that are missing. For example,\n                    projects/{project_id}/locations/{location_name}/targets/{target_name}.\n                  x-dcl-send-empty: true\n                  x-dcl-list-type: list\n                  items:\n                    type: string\n                    x-dcl-go-type: string\n                    x-dcl-references:\n                    - resource: Clouddeploy/Target\n                      field: selfLink\n                status:\n                  type: boolean\n                  x-dcl-go-name: Status\n                  description: True if there aren't any missing Targets.\n                updateTime:\n                  type: string\n                  format: date-time\n                  x-dcl-go-name: UpdateTime\n                  description: Last time the condition was updated.\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. Time at which the pipeline was created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Description of the `DeliveryPipeline`. Max length is 255 characters.\n        etag:\n          type: string\n          x-dcl-go-name: Etag\n          readOnly: true\n          description: This checksum is computed by the server based on the value\n            of other fields, and may be sent on update and delete requests to ensure\n            the client has an up-to-date value before proceeding.\n          x-kubernetes-immutable: true\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: 'Labels are attributes that can be set and used by both the\n            user and by Google Cloud Deploy. Labels must meet the following constraints:\n            * Keys and values can contain only lowercase letters, numeric characters,\n            underscores, and dashes. * All characters must use UTF-8 encoding, and\n            international characters are allowed. * Keys must start with a lowercase\n            letter or international character. * Each resource is limited to a maximum\n            of 64 labels. Both keys and values are additionally constrained to be\n            <= 128 bytes.'\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Name of the `DeliveryPipeline`. Format is [a-z][a-z0-9\\-]{0,62}.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        serialPipeline:\n          type: object\n          x-dcl-go-name: SerialPipeline\n          x-dcl-go-type: DeliveryPipelineSerialPipeline\n          description: SerialPipeline defines a sequential set of stages for a `DeliveryPipeline`.\n          properties:\n            stages:\n              type: array\n              x-dcl-go-name: Stages\n              description: Each stage specifies configuration for a `Target`. The\n                ordering of this list defines the promotion flow.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: DeliveryPipelineSerialPipelineStages\n                properties:\n                  profiles:\n                    type: array\n                    x-dcl-go-name: Profiles\n                    description: Skaffold profiles to use when rendering the manifest\n                      for this stage's `Target`.\n                    x-dcl-send-empty: true\n                    x-dcl-list-type: list\n                    items:\n                      type: string\n                      x-dcl-go-type: string\n                  targetId:\n                    type: string\n                    x-dcl-go-name: TargetId\n                    description: The target_id to which this stage points. This field\n                      refers exclusively to the last segment of a target name. For\n                      example, this field would just be `my-target` (rather than `projects/project/locations/location/targets/my-target`).\n                      The location of the `Target` is inferred to be the same as the\n                      location of the `DeliveryPipeline` that contains this `Stage`.\n        suspended:\n          type: boolean\n          x-dcl-go-name: Suspended\n          description: When suspended, no new releases or rollouts can be created,\n            but in-progress ones will complete.\n        uid:\n          type: string\n          x-dcl-go-name: Uid\n          readOnly: true\n          description: Output only. Unique identifier of the `DeliveryPipeline`.\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. Most recent time at which the pipeline was updated.\n          x-kubernetes-immutable: true\n")

// 9321 bytes
// MD5: 9759dcd55badfe29a6320979c4f29508
