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
// GENERATED BY gen_go_data.go
// gen_go_data -package alpha -var YAML_connection blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudbuildv2/alpha/connection.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudbuildv2/alpha/connection.yaml
var YAML_connection = []byte("info:\n  title: Cloudbuildv2/Connection\n  description: The Cloudbuildv2 Connection resource\n  x-dcl-struct-name: Connection\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Connection\n    parameters:\n    - name: connection\n      required: true\n      description: A full instance of a Connection\n  apply:\n    description: The function used to apply information about a Connection\n    parameters:\n    - name: connection\n      required: true\n      description: A full instance of a Connection\n  delete:\n    description: The function used to delete a Connection\n    parameters:\n    - name: connection\n      required: true\n      description: A full instance of a Connection\n  deleteAll:\n    description: The function used to delete all Connection\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Connection\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Connection:\n      title: Connection\n      x-dcl-id: projects/{{project}}/locations/{{location}}/connections/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - project\n      - location\n      properties:\n        annotations:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Annotations\n          description: Allows clients to store small amounts of arbitrary data.\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. Server assigned timestamp for when the connection\n            was created.\n          x-kubernetes-immutable: true\n        disabled:\n          type: boolean\n          x-dcl-go-name: Disabled\n          description: If disabled is set to true, functionality is disabled for this\n            connection. Repository based API methods and webhooks processing for repositories\n            in this connection will be disabled.\n        etag:\n          type: string\n          x-dcl-go-name: Etag\n          readOnly: true\n          description: This checksum is computed by the server based on the value\n            of other fields, and may be sent on update and delete requests to ensure\n            the client has an up-to-date value before proceeding.\n          x-kubernetes-immutable: true\n        githubConfig:\n          type: object\n          x-dcl-go-name: GithubConfig\n          x-dcl-go-type: ConnectionGithubConfig\n          description: Configuration for connections to github.com.\n          x-dcl-conflicts:\n          - githubEnterpriseConfig\n          - gitlabConfig\n          properties:\n            appInstallationId:\n              type: integer\n              format: int64\n              x-dcl-go-name: AppInstallationId\n              description: GitHub App installation id.\n            authorizerCredential:\n              type: object\n              x-dcl-go-name: AuthorizerCredential\n              x-dcl-go-type: ConnectionGithubConfigAuthorizerCredential\n              description: OAuth credential of the account that authorized the Cloud\n                Build GitHub App. It is recommended to use a robot account instead\n                of a human user account. The OAuth token must be tied to the Cloud\n                Build GitHub App.\n              properties:\n                oauthTokenSecretVersion:\n                  type: string\n                  x-dcl-go-name: OAuthTokenSecretVersion\n                  description: 'A SecretManager resource containing the OAuth token\n                    that authorizes the Cloud Build connection. Format: `projects/*/secrets/*/versions/*`.'\n                  x-dcl-references:\n                  - resource: Secretmanager/SecretVersion\n                    field: selfLink\n                username:\n                  type: string\n                  x-dcl-go-name: Username\n                  readOnly: true\n                  description: Output only. The username associated to this token.\n        githubEnterpriseConfig:\n          type: object\n          x-dcl-go-name: GithubEnterpriseConfig\n          x-dcl-go-type: ConnectionGithubEnterpriseConfig\n          description: Configuration for connections to an instance of GitHub Enterprise.\n          x-dcl-conflicts:\n          - githubConfig\n          - gitlabConfig\n          required:\n          - hostUri\n          properties:\n            appId:\n              type: integer\n              format: int64\n              x-dcl-go-name: AppId\n              description: Id of the GitHub App created from the manifest.\n            appInstallationId:\n              type: integer\n              format: int64\n              x-dcl-go-name: AppInstallationId\n              description: ID of the installation of the GitHub App.\n            appSlug:\n              type: string\n              x-dcl-go-name: AppSlug\n              description: The URL-friendly name of the GitHub App.\n            hostUri:\n              type: string\n              x-dcl-go-name: HostUri\n              description: Required. The URI of the GitHub Enterprise host this connection\n                is for.\n            privateKeySecretVersion:\n              type: string\n              x-dcl-go-name: PrivateKeySecretVersion\n              description: SecretManager resource containing the private key of the\n                GitHub App, formatted as `projects/*/secrets/*/versions/*`.\n              x-dcl-references:\n              - resource: Secretmanager/SecretVersion\n                field: selfLink\n            serviceDirectoryConfig:\n              type: object\n              x-dcl-go-name: ServiceDirectoryConfig\n              x-dcl-go-type: ConnectionGithubEnterpriseConfigServiceDirectoryConfig\n              description: Configuration for using Service Directory to privately\n                connect to a GitHub Enterprise server. This should only be set if\n                the GitHub Enterprise server is hosted on-premises and not reachable\n                by public internet. If this field is left empty, calls to the GitHub\n                Enterprise server will be made over the public internet.\n              required:\n              - service\n              properties:\n                service:\n                  type: string\n                  x-dcl-go-name: Service\n                  description: 'Required. The Service Directory service name. Format:\n                    projects/{project}/locations/{location}/namespaces/{namespace}/services/{service}.'\n                  x-dcl-references:\n                  - resource: Servicedirectory/Service\n                    field: selfLink\n            sslCa:\n              type: string\n              x-dcl-go-name: SslCa\n              description: SSL certificate to use for requests to GitHub Enterprise.\n            webhookSecretSecretVersion:\n              type: string\n              x-dcl-go-name: WebhookSecretSecretVersion\n              description: SecretManager resource containing the webhook secret of\n                the GitHub App, formatted as `projects/*/secrets/*/versions/*`.\n              x-dcl-references:\n              - resource: Secretmanager/SecretVersion\n                field: selfLink\n        gitlabConfig:\n          type: object\n          x-dcl-go-name: GitlabConfig\n          x-dcl-go-type: ConnectionGitlabConfig\n          description: Configuration for connections to gitlab.com or an instance\n            of GitLab Enterprise.\n          x-dcl-conflicts:\n          - githubConfig\n          - githubEnterpriseConfig\n          required:\n          - webhookSecretSecretVersion\n          - readAuthorizerCredential\n          - authorizerCredential\n          properties:\n            authorizerCredential:\n              type: object\n              x-dcl-go-name: AuthorizerCredential\n              x-dcl-go-type: ConnectionGitlabConfigAuthorizerCredential\n              description: Required. A GitLab personal access token with the `api`\n                scope access.\n              required:\n              - userTokenSecretVersion\n              properties:\n                userTokenSecretVersion:\n                  type: string\n                  x-dcl-go-name: UserTokenSecretVersion\n                  description: 'Required. A SecretManager resource containing the\n                    user token that authorizes the Cloud Build connection. Format:\n                    `projects/*/secrets/*/versions/*`.'\n                  x-dcl-references:\n                  - resource: Secretmanager/SecretVersion\n                    field: selfLink\n                username:\n                  type: string\n                  x-dcl-go-name: Username\n                  readOnly: true\n                  description: Output only. The username associated to this token.\n            hostUri:\n              type: string\n              x-dcl-go-name: HostUri\n              description: The URI of the GitLab Enterprise host this connection is\n                for. If not specified, the default value is https://gitlab.com.\n              x-dcl-server-default: true\n            readAuthorizerCredential:\n              type: object\n              x-dcl-go-name: ReadAuthorizerCredential\n              x-dcl-go-type: ConnectionGitlabConfigReadAuthorizerCredential\n              description: Required. A GitLab personal access token with the minimum\n                `read_api` scope access.\n              required:\n              - userTokenSecretVersion\n              properties:\n                userTokenSecretVersion:\n                  type: string\n                  x-dcl-go-name: UserTokenSecretVersion\n                  description: 'Required. A SecretManager resource containing the\n                    user token that authorizes the Cloud Build connection. Format:\n                    `projects/*/secrets/*/versions/*`.'\n                  x-dcl-references:\n                  - resource: Secretmanager/SecretVersion\n                    field: selfLink\n                username:\n                  type: string\n                  x-dcl-go-name: Username\n                  readOnly: true\n                  description: Output only. The username associated to this token.\n            serverVersion:\n              type: string\n              x-dcl-go-name: ServerVersion\n              readOnly: true\n              description: Output only. Version of the GitLab Enterprise server running\n                on the `host_uri`.\n            serviceDirectoryConfig:\n              type: object\n              x-dcl-go-name: ServiceDirectoryConfig\n              x-dcl-go-type: ConnectionGitlabConfigServiceDirectoryConfig\n              description: Configuration for using Service Directory to privately\n                connect to a GitLab Enterprise server. This should only be set if\n                the GitLab Enterprise server is hosted on-premises and not reachable\n                by public internet. If this field is left empty, calls to the GitLab\n                Enterprise server will be made over the public internet.\n              required:\n              - service\n              properties:\n                service:\n                  type: string\n                  x-dcl-go-name: Service\n                  description: 'Required. The Service Directory service name. Format:\n                    projects/{project}/locations/{location}/namespaces/{namespace}/services/{service}.'\n                  x-dcl-references:\n                  - resource: Servicedirectory/Service\n                    field: selfLink\n            sslCa:\n              type: string\n              x-dcl-go-name: SslCa\n              description: SSL certificate to use for requests to GitLab Enterprise.\n            webhookSecretSecretVersion:\n              type: string\n              x-dcl-go-name: WebhookSecretSecretVersion\n              description: Required. Immutable. SecretManager resource containing\n                the webhook secret of a GitLab Enterprise project, formatted as `projects/*/secrets/*/versions/*`.\n              x-dcl-references:\n              - resource: Secretmanager/SecretVersion\n                field: selfLink\n        installationState:\n          type: object\n          x-dcl-go-name: InstallationState\n          x-dcl-go-type: ConnectionInstallationState\n          readOnly: true\n          description: Output only. Installation state of the Connection.\n          x-kubernetes-immutable: true\n          properties:\n            actionUri:\n              type: string\n              x-dcl-go-name: ActionUri\n              readOnly: true\n              description: Output only. Link to follow for next action. Empty string\n                if the installation is already complete.\n              x-kubernetes-immutable: true\n            message:\n              type: string\n              x-dcl-go-name: Message\n              readOnly: true\n              description: Output only. Message of what the user should do next to\n                continue the installation. Empty string if the installation is already\n                complete.\n              x-kubernetes-immutable: true\n            stage:\n              type: string\n              x-dcl-go-name: Stage\n              x-dcl-go-type: ConnectionInstallationStateStageEnum\n              readOnly: true\n              description: 'Output only. Current step of the installation process.\n                Possible values: STAGE_UNSPECIFIED, PENDING_CREATE_APP, PENDING_USER_OAUTH,\n                PENDING_INSTALL_APP, COMPLETE'\n              x-kubernetes-immutable: true\n              enum:\n              - STAGE_UNSPECIFIED\n              - PENDING_CREATE_APP\n              - PENDING_USER_OAUTH\n              - PENDING_INSTALL_APP\n              - COMPLETE\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Immutable. The resource name of the connection, in the format\n            `projects/{project}/locations/{location}/connections/{connection_id}`.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        reconciling:\n          type: boolean\n          x-dcl-go-name: Reconciling\n          readOnly: true\n          description: Output only. Set to true when the connection is being set up\n            or updated in the background.\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. Server assigned timestamp for when the connection\n            was updated.\n          x-kubernetes-immutable: true\n")

// 15422 bytes
// MD5: 8aeeed1345743eb8568528a5672af403
