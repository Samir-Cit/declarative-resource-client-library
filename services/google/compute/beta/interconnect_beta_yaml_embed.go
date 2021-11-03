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
// gen_go_data -package beta -var YAML_interconnect blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/beta/interconnect.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/beta/interconnect.yaml
var YAML_interconnect = []byte("info:\n  title: Compute/Interconnect\n  description: The Compute Interconnect resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Interconnect\n    parameters:\n    - name: Interconnect\n      required: true\n      description: A full instance of a Interconnect\n  apply:\n    description: The function used to apply information about a Interconnect\n    parameters:\n    - name: Interconnect\n      required: true\n      description: A full instance of a Interconnect\n  delete:\n    description: The function used to delete a Interconnect\n    parameters:\n    - name: Interconnect\n      required: true\n      description: A full instance of a Interconnect\n  deleteAll:\n    description: The function used to delete all Interconnect\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Interconnect\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Interconnect:\n      title: Interconnect\n      x-dcl-id: projects/{{project}}/global/interconnects/{{name}}\n      x-dcl-locations:\n      - global\n      x-dcl-parent-container: project\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - location\n      - linkType\n      - interconnectType\n      - customerName\n      - project\n      properties:\n        adminEnabled:\n          type: boolean\n          x-dcl-go-name: AdminEnabled\n          description: Administrative status of the interconnect. When this is set\n            to true, the Interconnect is functional and can carry traffic. When set\n            to false, no packets can be carried over the interconnect and no BGP routes\n            are exchanged over it. By default, the status is set to true.\n        circuitInfos:\n          type: array\n          x-dcl-go-name: CircuitInfos\n          readOnly: true\n          description: A list of CircuitInfo objects, that describe the individual\n            circuits in this LAG.\n          x-kubernetes-immutable: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: InterconnectCircuitInfos\n            properties:\n              customerDemarcId:\n                type: string\n                x-dcl-go-name: CustomerDemarcId\n                description: Customer-side demarc ID for this circuit.\n                x-kubernetes-immutable: true\n              googleCircuitId:\n                type: string\n                x-dcl-go-name: GoogleCircuitId\n                description: Google-assigned unique ID for this circuit. Assigned\n                  at circuit turn-up.\n                x-kubernetes-immutable: true\n              googleDemarcId:\n                type: string\n                x-dcl-go-name: GoogleDemarcId\n                description: Google-side demarc ID for this circuit. Assigned at circuit\n                  turn-up and provided by Google to the customer in the LOA.\n                x-kubernetes-immutable: true\n        customerName:\n          type: string\n          x-dcl-go-name: CustomerName\n          description: Customer name, to put in the Letter of Authorization as the\n            party authorized to request a crossconnect.\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: An optional description of this resource. Provide this property\n            when you create the resource.\n        expectedOutages:\n          type: array\n          x-dcl-go-name: ExpectedOutages\n          readOnly: true\n          description: A list of outages expected for this Interconnect.\n          x-kubernetes-immutable: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: InterconnectExpectedOutages\n            properties:\n              affectedCircuits:\n                type: array\n                x-dcl-go-name: AffectedCircuits\n                description: If issue_type is IT_PARTIAL_OUTAGE, a list of the Google-side\n                  circuit IDs that will be affected.\n                x-kubernetes-immutable: true\n                x-dcl-send-empty: true\n                x-dcl-list-type: list\n                items:\n                  type: string\n                  x-dcl-go-type: string\n              description:\n                type: string\n                x-dcl-go-name: Description\n                description: A description about the purpose of the outage.\n                x-kubernetes-immutable: true\n              endTime:\n                type: integer\n                format: int64\n                x-dcl-go-name: EndTime\n                description: Scheduled end time for the outage (milliseconds since\n                  Unix epoch).\n                x-kubernetes-immutable: true\n              issueType:\n                type: string\n                x-dcl-go-name: IssueType\n                x-dcl-go-type: InterconnectExpectedOutagesIssueTypeEnum\n                description: 'Form this outage is expected to take, which can take\n                  one of the following values: OUTAGE: The Interconnect may be completely\n                  out of service for some or all of the specified window. PARTIAL_OUTAGE:\n                  Some circuits comprising the Interconnect as a whole should remain\n                  up, but with reduced bandwidth.nn Note that the versions of this\n                  enum prefixed with \"IT_\" have been deprecated in favor of the unprefixed\n                  values.'\n                x-kubernetes-immutable: true\n                enum:\n                - OUTAGE\n                - PARTIAL_OUTAGE\n              name:\n                type: string\n                x-dcl-go-name: Name\n                description: Unique identifier for this outage notification.\n                x-kubernetes-immutable: true\n              source:\n                type: string\n                x-dcl-go-name: Source\n                x-dcl-go-type: InterconnectExpectedOutagesSourceEnum\n                description: 'The party that generated this notification, which can\n                  take the following value: GOOGLE: this notification as generated\n                  by Google. Note that the value of NSRC_GOOGLE has been deprecated\n                  in favor of GOOGLE.'\n                x-kubernetes-immutable: true\n                enum:\n                - GOOGLE\n              startTime:\n                type: integer\n                format: int64\n                x-dcl-go-name: StartTime\n                description: Scheduled start time for the outage (milliseconds since\n                  Unix epoch).\n                x-kubernetes-immutable: true\n              state:\n                type: string\n                x-dcl-go-name: State\n                x-dcl-go-type: InterconnectExpectedOutagesStateEnum\n                description: 'State of this notification, which can take one of the\n                  following values: ACTIVE: This outage notification is active. The\n                  event could be in the past, present, or future. See start_time and\n                  end_time for scheduling. CANCELLED: The outage associated with this\n                  notification was cancelled before the outage was due to start. Note\n                  that the versions of this enum prefixed with \"NS_\" have been deprecated\n                  in favor of the unprefixed values.'\n                x-kubernetes-immutable: true\n                enum:\n                - ACTIVE\n                - CANCELLED\n        googleIPAddress:\n          type: string\n          x-dcl-go-name: GoogleIPAddress\n          readOnly: true\n          description: IP address configured on the Google side of the Interconnect\n            link. This can be used only for ping tests.\n          x-kubernetes-immutable: true\n        googleReferenceId:\n          type: string\n          x-dcl-go-name: GoogleReferenceId\n          readOnly: true\n          description: Google reference ID to be used when raising support tickets\n            with Google or otherwise to debug backend connectivity issues.\n          x-kubernetes-immutable: true\n        id:\n          type: integer\n          format: int64\n          x-dcl-go-name: Id\n          readOnly: true\n          description: The unique identifier for the resource. This identifier is\n            defined by the server.\n          x-kubernetes-immutable: true\n        interconnectAttachments:\n          type: array\n          x-dcl-go-name: InterconnectAttachments\n          readOnly: true\n          description: A list of the URLs of all InterconnectAttachments configured\n            to use this Interconnect.\n          x-kubernetes-immutable: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n        interconnectType:\n          type: string\n          x-dcl-go-name: InterconnectType\n          x-dcl-go-type: InterconnectInterconnectTypeEnum\n          description: 'Type of interconnect, which can take one of the following\n            values: PARTNER: A partner-managed interconnection shared between customers\n            though a partner. DEDICATED: A dedicated physical interconnection with\n            the customer.nn Note that a value IT_PRIVATE has been deprecated in favor\n            of DEDICATED.'\n          enum:\n          - IT_PRIVATE\n          - PARTNER\n          - DEDICATED\n        linkType:\n          type: string\n          x-dcl-go-name: LinkType\n          x-dcl-go-type: InterconnectLinkTypeEnum\n          description: 'Type of link requested, which can take one of the following\n            values: LINK_TYPE_ETHERNET_10G_LR: A 10G Ethernet with LR optics LINK_TYPE_ETHERNET_100G_LR:\n            A 100G Ethernet with LR optics.nn Note that this field indicates the speed\n            of each of the links in the bundle, not the speed of the entire bundle.'\n          enum:\n          - LINK_TYPE_ETHERNET_10G_LR\n          - LINK_TYPE_ETHERNET_100G_LR\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: URL of the InterconnectLocation object that represents where\n            this connection is to be provisioned.\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Name of the resource. Provided by the client when the resource\n            is created. The name must be 1-63 characters long, and comply with RFC1035.\n            Specifically, the name must be 1-63 characters long and match the regular\n            expression [a-z]([-a-z0-9]*[a-z0-9])? which means the first character\n            must be a lowercase letter, and all following characters must be a dash,\n            lowercase letter, or digit, except the last character, which cannot be\n            a dash.\n        nocContactEmail:\n          type: string\n          x-dcl-go-name: NocContactEmail\n          description: Email address to contact the customer NOC for operations and\n            maintenance notifications regarding this Interconnect. If specified, this\n            will be used for notifications in addition to all other forms described,\n            such as Stackdriver logs alerting and Cloud Notifications.\n        operationalStatus:\n          type: string\n          x-dcl-go-name: OperationalStatus\n          x-dcl-go-type: InterconnectOperationalStatusEnum\n          readOnly: true\n          description: 'The current status of this Interconnect''s functionality,\n            which can take one of the following values: OS_ACTIVE: A valid Interconnect,\n            which is turned up and is ready to use. Attachments may be provisioned\n            on this Interconnect. OS_UNPROVISIONED: An Interconnect that has not completed\n            turnup. No attachments may be provisioned on this Interconnect. OS_UNDER_MAINTENANCE:\n            An Interconnect that is undergoing internal maintenance. No attachments\n            may be provisioned or updated on this Interconnect.'\n          x-kubernetes-immutable: true\n          enum:\n          - OS_ACTIVE\n          - OS_UNPROVISIONED\n          - OS_UNDER_MAINTENANCE\n        peerIPAddress:\n          type: string\n          x-dcl-go-name: PeerIPAddress\n          readOnly: true\n          description: IP address configured on the customer side of the Interconnect\n            link. The customer should configure this IP address during turnup when\n            prompted by Google NOC. This can be used only for ping tests.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        provisionedLinkCount:\n          type: integer\n          format: int64\n          x-dcl-go-name: ProvisionedLinkCount\n          readOnly: true\n          description: Number of links actually provisioned in this interconnect.\n          x-kubernetes-immutable: true\n        requestedLinkCount:\n          type: integer\n          format: int64\n          x-dcl-go-name: RequestedLinkCount\n          description: Target number of physical links in the link bundle, as requested\n            by the customer.\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: Server-defined URL for the resource.\n          x-kubernetes-immutable: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: InterconnectStateEnum\n          readOnly: true\n          description: 'The current state of Interconnect functionality, which can\n            take one of the following values: ACTIVE: The Interconnect is valid, turned\n            up and ready to use. Attachments may be provisioned on this Interconnect.\n            UNPROVISIONED: The Interconnect has not completed turnup. No attachments\n            may be provisioned on this Interconnect. UNDER_MAINTENANCE: The Interconnect\n            is undergoing internal maintenance. No attachments may be provisioned\n            or updated on this Interconnect. Possible values: DEPRECATED, OBSOLETE,\n            DELETED, ACTIVE'\n          x-kubernetes-immutable: true\n          enum:\n          - DEPRECATED\n          - OBSOLETE\n          - DELETED\n          - ACTIVE\n")

// 14454 bytes
// MD5: 3590bd2c55e3ab4aee79c39c18af36df
