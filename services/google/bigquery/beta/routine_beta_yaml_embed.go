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
// gen_go_data -package beta -var YAML_routine blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/bigquery/beta/routine.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/bigquery/beta/routine.yaml
var YAML_routine = []byte("info:\n  title: Bigquery/Routine\n  description: The Bigquery Routine resource\n  x-dcl-struct-name: Routine\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Routine\n    parameters:\n    - name: Routine\n      required: true\n      description: A full instance of a Routine\n  apply:\n    description: The function used to apply information about a Routine\n    parameters:\n    - name: Routine\n      required: true\n      description: A full instance of a Routine\n  delete:\n    description: The function used to delete a Routine\n    parameters:\n    - name: Routine\n      required: true\n      description: A full instance of a Routine\n  deleteAll:\n    description: The function used to delete all Routine\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: dataset\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Routine\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: dataset\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    ArgumentsDataType:\n      x-dcl-has-create: false\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      x-dcl-go-name: DataType\n      x-dcl-go-type: RoutineArgumentsDataType\n      description: Required unless argument_kind = ANY_TYPE.\n      x-kubernetes-immutable: true\n      required:\n      - typeKind\n      properties:\n        arrayElementType:\n          $ref: '#/components/schemas/ArgumentsDataType'\n          x-dcl-go-name: ArrayElementType\n          x-kubernetes-immutable: true\n          x-dcl-conflicts:\n          - structType\n        structType:\n          type: object\n          x-dcl-go-name: StructType\n          x-dcl-go-type: RoutineArgumentsDataTypeStructType\n          description: The fields of this struct, in order, if type_kind = \"STRUCT\".\n          x-kubernetes-immutable: true\n          x-dcl-conflicts:\n          - arrayElementType\n          properties:\n            fields:\n              type: array\n              x-dcl-go-name: Fields\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: RoutineArgumentsDataTypeStructTypeFields\n                properties:\n                  name:\n                    type: string\n                    x-dcl-go-name: Name\n                    description: Optional. The name of this field. Can be absent for\n                      struct fields.\n                    x-kubernetes-immutable: true\n                  type:\n                    $ref: '#/components/schemas/ArgumentsDataType'\n                    x-dcl-go-name: Type\n                    x-kubernetes-immutable: true\n        typeKind:\n          type: string\n          x-dcl-go-name: TypeKind\n          x-dcl-go-type: RoutineArgumentsDataTypeTypeKindEnum\n          description: 'Required. The top level type of this field. Can be any standard\n            SQL data type (e.g., \"INT64\", \"DATE\", \"ARRAY\"). Possible values: TYPE_KIND_UNSPECIFIED,\n            INT64, BOOL, FLOAT64, STRING, BYTES, TIMESTAMP, DATE, TIME, DATETIME,\n            INTERVAL, GEOGRAPHY, NUMERIC, BIGNUMERIC, JSON, ARRAY, STRUCT'\n          x-kubernetes-immutable: true\n          enum:\n          - TYPE_KIND_UNSPECIFIED\n          - INT64\n          - BOOL\n          - FLOAT64\n          - STRING\n          - BYTES\n          - TIMESTAMP\n          - DATE\n          - TIME\n          - DATETIME\n          - INTERVAL\n          - GEOGRAPHY\n          - NUMERIC\n          - BIGNUMERIC\n          - JSON\n          - ARRAY\n          - STRUCT\n    Routine:\n      title: Routine\n      x-dcl-id: projects/{{project}}/datasets/{{dataset}}/routines/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - project\n      - dataset\n      - routineType\n      - definitionBody\n      properties:\n        arguments:\n          type: array\n          x-dcl-go-name: Arguments\n          description: Optional.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: RoutineArguments\n            properties:\n              argumentKind:\n                type: string\n                x-dcl-go-name: ArgumentKind\n                x-dcl-go-type: RoutineArgumentsArgumentKindEnum\n                description: 'Optional. Defaults to FIXED_TYPE. Possible values: ARGUMENT_KIND_UNSPECIFIED,\n                  FIXED_TYPE, ANY_TYPE'\n                enum:\n                - ARGUMENT_KIND_UNSPECIFIED\n                - FIXED_TYPE\n                - ANY_TYPE\n              dataType:\n                $ref: '#/components/schemas/ArgumentsDataType'\n                x-dcl-go-name: DataType\n                x-kubernetes-immutable: true\n              mode:\n                type: string\n                x-dcl-go-name: Mode\n                x-dcl-go-type: RoutineArgumentsModeEnum\n                description: 'Optional. Specifies whether the argument is input or\n                  output. Can be set for procedures only. Possible values: MODE_UNSPECIFIED,\n                  IN, OUT, INOUT'\n                enum:\n                - MODE_UNSPECIFIED\n                - IN\n                - OUT\n                - INOUT\n              name:\n                type: string\n                x-dcl-go-name: Name\n                description: Optional. The name of this argument. Can be absent for\n                  function return argument.\n        creationTime:\n          type: integer\n          format: int64\n          x-dcl-go-name: CreationTime\n          readOnly: true\n          description: Output only. The time when this routine was created, in milliseconds\n            since the epoch.\n          x-kubernetes-immutable: true\n        dataset:\n          type: string\n          x-dcl-go-name: Dataset\n          description: Required. The ID of the dataset containing this routine.\n          x-dcl-references:\n          - resource: Bigquery/Dataset\n            field: name\n            parent: true\n        definitionBody:\n          type: string\n          x-dcl-go-name: DefinitionBody\n          description: \"Required. The body of the routine. For functions, this is\n            the expression in the AS clause. If language=SQL, it is the substring\n            inside (but excluding) the parentheses. For example, for the function\n            created with the following statement: `CREATE FUNCTION JoinLines(x string,\n            y string) as (concat(x, \\\"\\n\\\", y))` The definition_body is `concat(x,\n            \\\"\\n\\\", y)` (\\n is not replaced with linebreak). If language=JAVASCRIPT,\n            it is the evaluated string in the AS clause. For example, for the function\n            created with the following statement: `CREATE FUNCTION f() RETURNS STRING\n            LANGUAGE js AS 'return \\\"\\n\\\";\\n'` The definition_body is `return \\\"\\n\\\";\\n`\n            Note that both \\n are replaced with linebreaks.\"\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. The description of the routine if defined.\n        determinismLevel:\n          type: string\n          x-dcl-go-name: DeterminismLevel\n          x-dcl-go-type: RoutineDeterminismLevelEnum\n          description: 'Optional. The determinism level of the JavaScript UDF if defined.\n            Possible values: DETERMINISM_LEVEL_UNSPECIFIED, DETERMINISTIC, NOT_DETERMINISTIC'\n          enum:\n          - DETERMINISM_LEVEL_UNSPECIFIED\n          - DETERMINISTIC\n          - NOT_DETERMINISTIC\n        etag:\n          type: string\n          x-dcl-go-name: Etag\n          readOnly: true\n          description: Output only. A hash of this resource.\n          x-kubernetes-immutable: true\n        importedLibraries:\n          type: array\n          x-dcl-go-name: ImportedLibraries\n          description: Optional. If language = \"JAVASCRIPT\", this field stores the\n            path of the imported JAVASCRIPT libraries.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n        language:\n          type: string\n          x-dcl-go-name: Language\n          x-dcl-go-type: RoutineLanguageEnum\n          description: 'Optional. Defaults to \"SQL\". Possible values: LANGUAGE_UNSPECIFIED,\n            SQL, JAVASCRIPT'\n          enum:\n          - LANGUAGE_UNSPECIFIED\n          - SQL\n          - JAVASCRIPT\n        lastModifiedTime:\n          type: integer\n          format: int64\n          x-dcl-go-name: LastModifiedTime\n          readOnly: true\n          description: Output only. The time when this routine was last modified,\n            in milliseconds since the epoch.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Required. The ID of the routine. The ID must contain only letters\n            (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 256\n            characters.\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: Required. The ID of the project containing this routine.\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        returnType:\n          $ref: '#/components/schemas/ArgumentsDataType'\n          x-dcl-go-name: ReturnType\n        routineType:\n          type: string\n          x-dcl-go-name: RoutineType\n          x-dcl-go-type: RoutineRoutineTypeEnum\n          description: 'Required. The type of routine. Possible values: ROUTINE_TYPE_UNSPECIFIED,\n            SCALAR_FUNCTION, PROCEDURE'\n          enum:\n          - ROUTINE_TYPE_UNSPECIFIED\n          - SCALAR_FUNCTION\n          - PROCEDURE\n        strictMode:\n          type: boolean\n          x-dcl-go-name: StrictMode\n          description: Optional. Can be set for procedures only. If true (default),\n            the definition body will be validated in the creation and the updates\n            of the procedure. For procedures with an argument of ANY TYPE, the definition\n            body validtion is not supported at creation/update time, and thus this\n            field must be set to false explicitly.\n")

// 10606 bytes
// MD5: 19e7c2f83114048d14011b126770df35
