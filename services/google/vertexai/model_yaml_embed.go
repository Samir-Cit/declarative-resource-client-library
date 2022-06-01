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
// gen_go_data -package vertexai -var YAML_model blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/vertexai/model.yaml

package vertexai

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/vertexai/model.yaml
var YAML_model = []byte("info:\n  title: VertexAI/Model\n  description: The VertexAI Model resource\n  x-dcl-struct-name: Model\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Model\n    parameters:\n    - name: Model\n      required: true\n      description: A full instance of a Model\n  apply:\n    description: The function used to apply information about a Model\n    parameters:\n    - name: Model\n      required: true\n      description: A full instance of a Model\n  delete:\n    description: The function used to delete a Model\n    parameters:\n    - name: Model\n      required: true\n      description: A full instance of a Model\n  deleteAll:\n    description: The function used to delete all Model\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Model\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Model:\n      title: Model\n      x-dcl-id: projects/{{project}}/locations/{{location}}/models/{{name}}\n      x-dcl-uses-state-hint: true\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - displayName\n      - containerSpec\n      - project\n      - location\n      properties:\n        artifactUri:\n          type: string\n          x-dcl-go-name: ArtifactUri\n          description: Immutable. The path to the directory containing the Model artifact\n            and any of its supporting files. Not present for AutoML Models.\n          x-kubernetes-immutable: true\n        containerSpec:\n          type: object\n          x-dcl-go-name: ContainerSpec\n          x-dcl-go-type: ModelContainerSpec\n          description: Input only. The specification of the container that is to be\n            used when deploying this Model. The specification is ingested upon ModelService.UploadModel,\n            and all binaries it contains are copied and stored internally by Vertex\n            AI. Not present for AutoML Models.\n          x-kubernetes-immutable: true\n          x-dcl-mutable-unreadable: true\n          required:\n          - imageUri\n          properties:\n            args:\n              type: array\n              x-dcl-go-name: Args\n              description: 'Immutable. Specifies arguments for the command that runs\n                when the container starts. This overrides the container''s [`CMD`](https://docs.docker.com/engine/reference/builder/#cmd).\n                Specify this field as an array of executable and arguments, similar\n                to a Docker `CMD`''s \"default parameters\" form. If you don''t specify\n                this field but do specify the command field, then the command from\n                the `command` field runs without any additional arguments. See the\n                [Kubernetes documentation about how the `command` and `args` fields\n                interact with a container''s `ENTRYPOINT` and `CMD`](https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#notes).\n                If you don''t specify this field and don''t specify the `command`\n                field, then the container''s [`ENTRYPOINT`](https://docs.docker.com/engine/reference/builder/#cmd)\n                and `CMD` determine what runs based on their default behavior. See\n                the Docker documentation about [how `CMD` and `ENTRYPOINT` interact](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact).\n                In this field, you can reference [environment variables set by Vertex\n                AI](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables)\n                and environment variables set in the env field. You cannot reference\n                environment variables set in the Docker image. In order for environment\n                variables to be expanded, reference them by using the following syntax:\n                `$(VARIABLE_NAME)` Note that this differs from Bash variable expansion,\n                which does not use parentheses. If a variable cannot be resolved,\n                the reference in the input string is used unchanged. To avoid variable\n                expansion, you can escape this syntax with `$$`; for example: `$$(VARIABLE_NAME)`\n                This field corresponds to the `args` field of the Kubernetes Containers\n                [v1 core API](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).'\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            command:\n              type: array\n              x-dcl-go-name: Command\n              description: 'Immutable. Specifies the command that runs when the container\n                starts. This overrides the container''s [ENTRYPOINT](https://docs.docker.com/engine/reference/builder/#entrypoint).\n                Specify this field as an array of executable and arguments, similar\n                to a Docker `ENTRYPOINT`''s \"exec\" form, not its \"shell\" form. If\n                you do not specify this field, then the container''s `ENTRYPOINT`\n                runs, in conjunction with the args field or the container''s [`CMD`](https://docs.docker.com/engine/reference/builder/#cmd),\n                if either exists. If this field is not specified and the container\n                does not have an `ENTRYPOINT`, then refer to the Docker documentation\n                about [how `CMD` and `ENTRYPOINT` interact](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact).\n                If you specify this field, then you can also specify the `args` field\n                to provide additional arguments for this command. However, if you\n                specify this field, then the container''s `CMD` is ignored. See the\n                [Kubernetes documentation about how the `command` and `args` fields\n                interact with a container''s `ENTRYPOINT` and `CMD`](https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#notes).\n                In this field, you can reference [environment variables set by Vertex\n                AI](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables)\n                and environment variables set in the env field. You cannot reference\n                environment variables set in the Docker image. In order for environment\n                variables to be expanded, reference them by using the following syntax:\n                `$(VARIABLE_NAME)` Note that this differs from Bash variable expansion,\n                which does not use parentheses. If a variable cannot be resolved,\n                the reference in the input string is used unchanged. To avoid variable\n                expansion, you can escape this syntax with `$$`; for example: `$$(VARIABLE_NAME)`\n                This field corresponds to the `command` field of the Kubernetes Containers\n                [v1 core API](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).'\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            env:\n              type: array\n              x-dcl-go-name: Env\n              description: 'Immutable. List of environment variables to set in the\n                container. After the container starts running, code running in the\n                container can read these environment variables. Additionally, the\n                command and args fields can reference these variables. Later entries\n                in this list can also reference earlier entries. For example, the\n                following example sets the variable `VAR_2` to have the value `foo\n                bar`: ```json [ { \"name\": \"VAR_1\", \"value\": \"foo\" }, { \"name\": \"VAR_2\",\n                \"value\": \"$(VAR_1) bar\" } ] ``` If you switch the order of the variables\n                in the example, then the expansion does not occur. This field corresponds\n                to the `env` field of the Kubernetes Containers [v1 core API](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).'\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: ModelContainerSpecEnv\n                required:\n                - name\n                - value\n                properties:\n                  name:\n                    type: string\n                    x-dcl-go-name: Name\n                    description: Required. Name of the environment variable. Must\n                      be a valid C identifier.\n                    x-kubernetes-immutable: true\n                  value:\n                    type: string\n                    x-dcl-go-name: Value\n                    description: 'Required. Variables that reference a $(VAR_NAME)\n                      are expanded using the previous defined environment variables\n                      in the container and any service environment variables. If a\n                      variable cannot be resolved, the reference in the input string\n                      will be unchanged. The $(VAR_NAME) syntax can be escaped with\n                      a double $$, ie: $$(VAR_NAME). Escaped references will never\n                      be expanded, regardless of whether the variable exists or not.'\n                    x-kubernetes-immutable: true\n            healthRoute:\n              type: string\n              x-dcl-go-name: HealthRoute\n              description: 'Immutable. HTTP path on the container to send health checks\n                to. Vertex AI intermittently sends GET requests to this path on the\n                container''s IP address and port to check that the container is healthy.\n                Read more about [health checks](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#health).\n                For example, if you set this field to `/bar`, then Vertex AI intermittently\n                sends a GET request to the `/bar` path on the port of your container\n                specified by the first value of this `ModelContainerSpec`''s ports\n                field. If you don''t specify this field, it defaults to the following\n                value when you deploy this Model to an Endpoint: `/v1/endpoints/ENDPOINT/deployedModels/DEPLOYED_MODEL:predict`\n                The placeholders in this value are replaced as follows: * ENDPOINT:\n                The last segment (following `endpoints/`)of the Endpoint.name][] field\n                of the Endpoint where this Model has been deployed. (Vertex AI makes\n                this value available to your container code as the [`AIP_ENDPOINT_ID`\n                environment variable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).)\n                * DEPLOYED_MODEL: DeployedModel.id of the `DeployedModel`. (Vertex\n                AI makes this value available to your container code as the [`AIP_DEPLOYED_MODEL_ID`\n                environment variable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).)'\n              x-kubernetes-immutable: true\n            imageUri:\n              type: string\n              x-dcl-go-name: ImageUri\n              description: Required. Immutable. URI of the Docker image to be used\n                as the custom container for serving predictions. This URI must identify\n                an image in Artifact Registry or Container Registry. Learn more about\n                the [container publishing requirements](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#publishing),\n                including permissions requirements for the Vertex AI Service Agent.\n                The container image is ingested upon ModelService.UploadModel, stored\n                internally, and this original path is afterwards not used. To learn\n                about the requirements for the Docker image itself, see [Custom container\n                requirements](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#).\n                You can use the URI to one of Vertex AI's [pre-built container images\n                for prediction](https://cloud.google.com/vertex-ai/docs/predictions/pre-built-containers)\n                in this field.\n              x-kubernetes-immutable: true\n            ports:\n              type: array\n              x-dcl-go-name: Ports\n              description: 'Immutable. List of ports to expose from the container.\n                Vertex AI sends any prediction requests that it receives to the first\n                port on this list. Vertex AI also sends [liveness and health checks](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#liveness)\n                to this port. If you do not specify this field, it defaults to following\n                value: ```json [ { \"containerPort\": 8080 } ] ``` Vertex AI does not\n                use ports other than the first one listed. This field corresponds\n                to the `ports` field of the Kubernetes Containers [v1 core API](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).'\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: ModelContainerSpecPorts\n                properties:\n                  containerPort:\n                    type: integer\n                    format: int64\n                    x-dcl-go-name: ContainerPort\n                    description: The number of the port to expose on the pod's IP\n                      address. Must be a valid port number, between 1 and 65535 inclusive.\n                    x-kubernetes-immutable: true\n            predictRoute:\n              type: string\n              x-dcl-go-name: PredictRoute\n              description: 'Immutable. HTTP path on the container to send prediction\n                requests to. Vertex AI forwards requests sent using projects.locations.endpoints.predict\n                to this path on the container''s IP address and port. Vertex AI then\n                returns the container''s response in the API response. For example,\n                if you set this field to `/foo`, then when Vertex AI receives a prediction\n                request, it forwards the request body in a POST request to the `/foo`\n                path on the port of your container specified by the first value of\n                this `ModelContainerSpec`''s ports field. If you don''t specify this\n                field, it defaults to the following value when you deploy this Model\n                to an Endpoint: `/v1/endpoints/ENDPOINT/deployedModels/DEPLOYED_MODEL:predict`\n                The placeholders in this value are replaced as follows: * ENDPOINT:\n                The last segment (following `endpoints/`)of the Endpoint.name][] field\n                of the Endpoint where this Model has been deployed. (Vertex AI makes\n                this value available to your container code as the [`AIP_ENDPOINT_ID`\n                environment variable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).)\n                * DEPLOYED_MODEL: DeployedModel.id of the `DeployedModel`. (Vertex\n                AI makes this value available to your container code as the [`AIP_DEPLOYED_MODEL_ID`\n                environment variable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).)'\n              x-kubernetes-immutable: true\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. Timestamp when this Model was uploaded into Vertex\n            AI.\n          x-kubernetes-immutable: true\n        deployedModels:\n          type: array\n          x-dcl-go-name: DeployedModels\n          readOnly: true\n          description: Output only. The pointers to DeployedModels created from this\n            Model. Note that Model could have been deployed to Endpoints in different\n            Locations.\n          x-kubernetes-immutable: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: ModelDeployedModels\n            properties:\n              deployedModelId:\n                type: string\n                x-dcl-go-name: DeployedModelId\n                description: Immutable. An ID of a DeployedModel in the above Endpoint.\n                x-kubernetes-immutable: true\n              endpoint:\n                type: string\n                x-dcl-go-name: Endpoint\n                description: Immutable. A resource name of an Endpoint.\n                x-kubernetes-immutable: true\n                x-dcl-references:\n                - resource: Aiplatform/Endpoint\n                  field: selfLink\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: The description of the Model.\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: Required. The display name of the Model. The name can be up\n            to 128 characters long and can be consist of any UTF-8 characters.\n        encryptionSpec:\n          type: object\n          x-dcl-go-name: EncryptionSpec\n          x-dcl-go-type: ModelEncryptionSpec\n          description: Customer-managed encryption key spec for a Model. If set, this\n            Model and all sub-resources of this Model will be secured by this key.\n          x-kubernetes-immutable: true\n          required:\n          - kmsKeyName\n          properties:\n            kmsKeyName:\n              type: string\n              x-dcl-go-name: KmsKeyName\n              description: 'Required. The Cloud KMS resource identifier of the customer\n                managed encryption key used to protect a resource. Has the form: `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.\n                The key needs to be in the same region as where the compute resource\n                is created.'\n              x-kubernetes-immutable: true\n              x-dcl-references:\n              - resource: Cloudkms/CryptoKey\n                field: name\n        etag:\n          type: string\n          x-dcl-go-name: Etag\n          readOnly: true\n          description: Used to perform consistent read-modify-write updates. If not\n            set, a blind \"overwrite\" update happens.\n          x-kubernetes-immutable: true\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: The labels with user-defined metadata to organize your Models.\n            Label keys and values can be no longer than 64 characters (Unicode codepoints),\n            can only contain lowercase letters, numeric characters, underscores and\n            dashes. International characters are allowed. See https://goo.gl/xmQnxf\n            for more information and examples of labels.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The resource name of the Model.\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n        originalModelInfo:\n          type: object\n          x-dcl-go-name: OriginalModelInfo\n          x-dcl-go-type: ModelOriginalModelInfo\n          readOnly: true\n          description: Output only. If this Model is a copy of another Model, this\n            contains info about the original.\n          x-kubernetes-immutable: true\n          properties:\n            model:\n              type: string\n              x-dcl-go-name: Model\n              readOnly: true\n              description: 'Output only. The resource name of the Model this Model\n                is a copy of, including the revision. Format: `projects/{project}/locations/{location}/models/{model_id}@{version_id}`'\n              x-kubernetes-immutable: true\n              x-dcl-references:\n              - resource: Aiplatform/Model\n                field: selfLink\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        supportedDeploymentResourcesTypes:\n          type: array\n          x-dcl-go-name: SupportedDeploymentResourcesTypes\n          readOnly: true\n          description: Output only. When this Model is deployed, its prediction resources\n            are described by the `prediction_resources` field of the Endpoint.deployed_models\n            object. Because not all Models support all resource configuration types,\n            the configuration types this Model supports are listed here. If no configuration\n            types are listed, the Model cannot be deployed to an Endpoint and does\n            not support online predictions (PredictionService.Predict or PredictionService.Explain).\n            Such a Model can serve predictions by using a BatchPredictionJob, if it\n            has at least one entry each in supported_input_storage_formats and supported_output_storage_formats.\n          x-kubernetes-immutable: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: ModelSupportedDeploymentResourcesTypesEnum\n            enum:\n            - DEPLOYMENT_RESOURCES_TYPE_UNSPECIFIED\n            - DEDICATED_RESOURCES\n            - AUTOMATIC_RESOURCES\n        supportedExportFormats:\n          type: array\n          x-dcl-go-name: SupportedExportFormats\n          readOnly: true\n          description: Output only. The formats in which this Model may be exported.\n            If empty, this Model is not available for export.\n          x-kubernetes-immutable: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: ModelSupportedExportFormats\n            properties:\n              exportableContents:\n                type: array\n                x-dcl-go-name: ExportableContents\n                readOnly: true\n                description: Output only. The content of this Model that may be exported.\n                x-kubernetes-immutable: true\n                x-dcl-list-type: list\n                items:\n                  type: string\n                  x-dcl-go-type: ModelSupportedExportFormatsExportableContentsEnum\n                  enum:\n                  - EXPORTABLE_CONTENT_UNSPECIFIED\n                  - ARTIFACT\n                  - IMAGE\n              id:\n                type: string\n                x-dcl-go-name: Id\n                readOnly: true\n                description: 'Output only. The ID of the export format. The possible\n                  format IDs are: * `tflite` Used for Android mobile devices. * `edgetpu-tflite`\n                  Used for [Edge TPU](https://cloud.google.com/edge-tpu/) devices.\n                  * `tf-saved-model` A tensorflow model in SavedModel format. * `tf-js`\n                  A [TensorFlow.js](https://www.tensorflow.org/js) model that can\n                  be used in the browser and in Node.js using JavaScript. * `core-ml`\n                  Used for iOS mobile devices. * `custom-trained` A Model that was\n                  uploaded or trained by custom code.'\n                x-kubernetes-immutable: true\n        supportedInputStorageFormats:\n          type: array\n          x-dcl-go-name: SupportedInputStorageFormats\n          readOnly: true\n          description: 'Output only. The formats this Model supports in BatchPredictionJob.input_config.\n            If PredictSchemata.instance_schema_uri exists, the instances should be\n            given as per that schema. The possible formats are: * `jsonl` The JSON\n            Lines format, where each instance is a single line. Uses GcsSource. *\n            `csv` The CSV format, where each instance is a single comma-separated\n            line. The first line in the file is the header, containing comma-separated\n            field names. Uses GcsSource. * `tf-record` The TFRecord format, where\n            each instance is a single record in tfrecord syntax. Uses GcsSource. *\n            `tf-record-gzip` Similar to `tf-record`, but the file is gzipped. Uses\n            GcsSource. * `bigquery` Each instance is a single row in BigQuery. Uses\n            BigQuerySource. * `file-list` Each line of the file is the location of\n            an instance to process, uses `gcs_source` field of the InputConfig object.\n            If this Model doesn''t support any of these formats it means it cannot\n            be used with a BatchPredictionJob. However, if it has supported_deployment_resources_types,\n            it could serve online predictions by using PredictionService.Predict or\n            PredictionService.Explain. TODO(rsurowka): Give a link describing how\n            OpenAPI schema instances are expressed in JSONL and BigQuery. TODO(rsurowka):\n            Should we provide a schema for TFRecord? Or maybe say that at least for\n            now TFRecord input is not supported via schemata (that would also simplify\n            giving them back as part of predictions). TODO(rsurowka): Define CSV format\n            (decide how much we want to support). E.g. no nesting? Or no arrays, or\n            no nested arrays? E.g. https://json-csv.com/ seems to be able to do pretty\n            advanced conversions, but we may decide to make it relatively simple for\n            now.'\n          x-kubernetes-immutable: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n        supportedOutputStorageFormats:\n          type: array\n          x-dcl-go-name: SupportedOutputStorageFormats\n          readOnly: true\n          description: 'Output only. The formats this Model supports in BatchPredictionJob.output_config.\n            If both PredictSchemata.instance_schema_uri and PredictSchemata.prediction_schema_uri\n            exist, the predictions are returned together with their instances. In\n            other words, the prediction has the original instance data first, followed\n            by the actual prediction content (as per the schema). The possible formats\n            are: * `jsonl` The JSON Lines format, where each prediction is a single\n            line. Uses GcsDestination. * `csv` The CSV format, where each prediction\n            is a single comma-separated line. The first line in the file is the header,\n            containing comma-separated field names. Uses GcsDestination. * `bigquery`\n            Each prediction is a single row in a BigQuery table, uses BigQueryDestination\n            . If this Model doesn''t support any of these formats it means it cannot\n            be used with a BatchPredictionJob. However, if it has supported_deployment_resources_types,\n            it could serve online predictions by using PredictionService.Predict or\n            PredictionService.Explain. TODO(rsurowka): Analogous TODOs as for instances\n            field above.'\n          x-kubernetes-immutable: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n        trainingPipeline:\n          type: string\n          x-dcl-go-name: TrainingPipeline\n          readOnly: true\n          description: Output only. The resource name of the TrainingPipeline that\n            uploaded this Model, if any.\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. Timestamp when this Model was most recently updated.\n          x-kubernetes-immutable: true\n        versionCreateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: VersionCreateTime\n          readOnly: true\n          description: Output only. Timestamp when this version was created.\n          x-kubernetes-immutable: true\n        versionDescription:\n          type: string\n          x-dcl-go-name: VersionDescription\n          description: The description of this version.\n          x-kubernetes-immutable: true\n        versionId:\n          type: string\n          x-dcl-go-name: VersionId\n          readOnly: true\n          description: Output only. Immutable. The version ID of the model. A new\n            version is committed when a new model version is uploaded or trained under\n            an existing model id. It is an auto-incrementing decimal number in string\n            representation.\n          x-kubernetes-immutable: true\n        versionUpdateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: VersionUpdateTime\n          readOnly: true\n          description: Output only. Timestamp when this version was most recently\n            updated.\n          x-kubernetes-immutable: true\n")

// 30050 bytes
// MD5: 8542784c84a648c562c9456b9a1f47c5
