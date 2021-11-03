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
// gen_go_data -package alpha -var YAML_attestor blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/binaryauthorization/alpha/attestor.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/binaryauthorization/alpha/attestor.yaml
var YAML_attestor = []byte("info:\n  title: BinaryAuthorization/Attestor\n  description: The BinaryAuthorization Attestor resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Attestor\n    parameters:\n    - name: Attestor\n      required: true\n      description: A full instance of a Attestor\n  apply:\n    description: The function used to apply information about a Attestor\n    parameters:\n    - name: Attestor\n      required: true\n      description: A full instance of a Attestor\n  delete:\n    description: The function used to delete a Attestor\n    parameters:\n    - name: Attestor\n      required: true\n      description: A full instance of a Attestor\n  deleteAll:\n    description: The function used to delete all Attestor\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Attestor\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Attestor:\n      title: Attestor\n      x-dcl-id: projects/{{project}}/attestors/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - project\n      properties:\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. A descriptive comment. This field may be updated.\n            The field may be displayed in chooser dialogs.\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'Required. The resource name, in the format: `projects/*/attestors/*`.\n            This field may not be updated.'\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. Time when the attestor was last updated.\n          x-kubernetes-immutable: true\n        userOwnedDrydockNote:\n          type: object\n          x-dcl-go-name: UserOwnedDrydockNote\n          x-dcl-go-type: AttestorUserOwnedDrydockNote\n          description: This specifies how an attestation will be read, and how it\n            will be used during policy enforcement.\n          required:\n          - noteReference\n          properties:\n            delegationServiceAccountEmail:\n              type: string\n              x-dcl-go-name: DelegationServiceAccountEmail\n              readOnly: true\n              description: Output only. This field will contain the service account\n                email address that this Attestor will use as the principal when querying\n                Container Analysis. Attestor administrators must grant this service\n                account the IAM role needed to read attestations from the in Container\n                Analysis (`containeranalysis.notes.occurrences.viewer`). This email\n                address is fixed for the lifetime of the Attestor, but callers should\n                not make any other assumptions about the service account email; future\n                versions may use an email based on a different naming pattern.\n              x-kubernetes-immutable: true\n            noteReference:\n              type: string\n              x-dcl-go-name: NoteReference\n              description: 'Required. The Drydock resource name of a Attestation.\n                Authority Note, created by the user, in the format: `projects/*/notes/*`.\n                This field may not be updated. An attestation by this attestor is\n                stored as a Grafeas Attestation. Authority Occurrence that names a\n                container image and that links to this Note. Grafeas is an external\n                dependency.'\n              x-kubernetes-immutable: true\n              x-dcl-references:\n              - resource: Containeranalysis/Note\n                field: name\n            publicKeys:\n              type: array\n              x-dcl-go-name: PublicKeys\n              description: Optional. Public keys that verify attestations signed by\n                this attestor. This field may be updated. If this field is non-empty,\n                one of the specified public keys must verify that an attestation was\n                signed by this attestor for the image specified in the admission request.\n                If this field is empty, this attestor always returns that no valid\n                attestations exist.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: AttestorUserOwnedDrydockNotePublicKeys\n                properties:\n                  asciiArmoredPgpPublicKey:\n                    type: string\n                    x-dcl-go-name: AsciiArmoredPgpPublicKey\n                    description: ASCII-armored representation of a PGP public key,\n                      as the entire output by the command `gpg --export --armor foo@example.com`\n                      (either LF or CRLF line endings). When using this field, `id`\n                      should be left blank. The BinAuthz API handlers will calculate\n                      the ID and fill it in automatically. BinAuthz computes this\n                      ID as the OpenPGP RFC4880 V4 fingerprint, represented as upper-case\n                      hex. If `id` is provided by the caller, it will be overwritten\n                      by the API-calculated ID.\n                  comment:\n                    type: string\n                    x-dcl-go-name: Comment\n                    description: Optional. A descriptive comment. This field may be\n                      updated.\n                  id:\n                    type: string\n                    x-dcl-go-name: Id\n                    description: The ID of this public key. Signatures verified by\n                      BinAuthz must include the ID of the public key that can be used\n                      to verify them, and that ID must match the contents of this\n                      field exactly. Additional restrictions on this field can be\n                      imposed based on which public key type is encapsulated. See\n                      the documentation on `public_key` cases below for details.\n                  pkixPublicKey:\n                    type: object\n                    x-dcl-go-name: PkixPublicKey\n                    x-dcl-go-type: AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey\n                    description: 'A raw PKIX SubjectPublicKeyInfo format public key.\n                      NOTE: `id` may be explicitly provided by the caller when using\n                      this type of public key, but it MUST be a valid RFC3986 URI.\n                      If `id` is left blank, a default one will be computed based\n                      on the digest of the DER encoding of the public key.'\n                    properties:\n                      publicKeyPem:\n                        type: string\n                        x-dcl-go-name: PublicKeyPem\n                        description: A PEM-encoded public key, as described in https://tools.ietf.org/html/rfc7468#section-13\n                      signatureAlgorithm:\n                        type: string\n                        x-dcl-go-name: SignatureAlgorithm\n                        x-dcl-go-type: AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum\n                        description: 'The signature algorithm used to verify a message\n                          against a signature using this key. These signature algorithm\n                          must match the structure and any object identifiers encoded\n                          in `public_key_pem` (i.e. this algorithm must match that\n                          of the public key). Possible values: SIGNATURE_ALGORITHM_UNSPECIFIED,\n                          RSA_PSS_2048_SHA256, RSA_PSS_3072_SHA256, RSA_PSS_4096_SHA256,\n                          RSA_PSS_4096_SHA512, RSA_SIGN_PKCS1_2048_SHA256, RSA_SIGN_PKCS1_3072_SHA256,\n                          RSA_SIGN_PKCS1_4096_SHA256, RSA_SIGN_PKCS1_4096_SHA512,\n                          ECDSA_P256_SHA256, EC_SIGN_P256_SHA256, ECDSA_P384_SHA384,\n                          EC_SIGN_P384_SHA384, ECDSA_P521_SHA512, EC_SIGN_P521_SHA512'\n                        enum:\n                        - SIGNATURE_ALGORITHM_UNSPECIFIED\n                        - RSA_PSS_2048_SHA256\n                        - RSA_PSS_3072_SHA256\n                        - RSA_PSS_4096_SHA256\n                        - RSA_PSS_4096_SHA512\n                        - RSA_SIGN_PKCS1_2048_SHA256\n                        - RSA_SIGN_PKCS1_3072_SHA256\n                        - RSA_SIGN_PKCS1_4096_SHA256\n                        - RSA_SIGN_PKCS1_4096_SHA512\n                        - ECDSA_P256_SHA256\n                        - EC_SIGN_P256_SHA256\n                        - ECDSA_P384_SHA384\n                        - EC_SIGN_P384_SHA384\n                        - ECDSA_P521_SHA512\n                        - EC_SIGN_P521_SHA512\n")

// 9379 bytes
// MD5: 504fada1cdbb64e6d769cec20682d2c6
