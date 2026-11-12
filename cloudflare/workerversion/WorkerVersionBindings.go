// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workerversion


type WorkerVersionBindings struct {
	// For `type: ai` or `type: ai_search` or `type: ai_search_namespace` or `type: messaging` or `type: analytics_engine` or `type: artifacts` or `type: assets` or `type: browser` or `type: d1` or `type: data_blob` or `type: dispatch_namespace` or `type: durable_object_namespace` or `type: hyperdrive` or `type: images` or `type: json` or `type: kv_namespace` or `type: media` or `type: mtls_certificate` or `type: plain_text` or `type: pipelines` or `type: k2` or `type: queue` or `type: ratelimit` or `type: r2_bucket` or `type: secret_text` or `type: send_email` or `type: service` or `type: text_blob` or `type: vectorize` or `type: version_metadata` or `type: secrets_store_secret` or `type: flagship` or `type: secret_key` or `type: workflow` or `type: wasm_module` or `type: vpc_service` or `type: vpc_network` or `type: worker_loader`: A JavaScript variable name for the binding.
	//
	// For `type: inherit`: The name of the inherited binding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#name WorkerVersion#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The kind of resource that the binding provides.
	//
	// Available values: "ai", "ai_search", "ai_search_namespace", "messaging", "analytics_engine", "artifacts", "assets", "browser", "d1", "data_blob", "dispatch_namespace", "durable_object_namespace", "hyperdrive", "inherit", "images", "json", "kv_namespace", "media", "mtls_certificate", "plain_text", "pipelines", "k2", "queue", "ratelimit", "r2_bucket", "secret_text", "send_email", "service", "text_blob", "vectorize", "version_metadata", "secrets_store_secret", "flagship", "secret_key", "workflow", "wasm_module", "vpc_service", "vpc_network", "worker_loader".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#type WorkerVersion#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// For `type: secret_key`: Algorithm-specific key parameters. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#algorithm WorkerVersion#algorithm}
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// For `type: send_email`: List of allowed destination addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#allowed_destination_addresses WorkerVersion#allowed_destination_addresses}
	AllowedDestinationAddresses *[]*string `field:"optional" json:"allowedDestinationAddresses" yaml:"allowedDestinationAddresses"`
	// For `type: send_email`: List of allowed sender addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#allowed_sender_addresses WorkerVersion#allowed_sender_addresses}
	AllowedSenderAddresses *[]*string `field:"optional" json:"allowedSenderAddresses" yaml:"allowedSenderAddresses"`
	// For `type: flagship`: ID of the Flagship app to bind to for feature flag evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#app_id WorkerVersion#app_id}
	AppId *string `field:"optional" json:"appId" yaml:"appId"`
	// For `type: r2_bucket`: R2 bucket to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#bucket_name WorkerVersion#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// For `type: mtls_certificate`: Identifier of the certificate to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#certificate_id WorkerVersion#certificate_id}
	CertificateId *string `field:"optional" json:"certificateId" yaml:"certificateId"`
	// For `type: durable_object_namespace`: The exported class name of the Durable Object.
	//
	// For `type: workflow`: Class name of the Workflow. Should only be provided if the Workflow belongs to this script.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#class_name WorkerVersion#class_name}
	ClassName *string `field:"optional" json:"className" yaml:"className"`
	// For `type: d1`: Identifier of the D1 database to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#database_id WorkerVersion#database_id}
	DatabaseId *string `field:"optional" json:"databaseId" yaml:"databaseId"`
	// For `type: analytics_engine`: The name of the dataset to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#dataset WorkerVersion#dataset}
	Dataset *string `field:"optional" json:"dataset" yaml:"dataset"`
	// For `type: send_email`: Destination address for the email.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#destination_address WorkerVersion#destination_address}
	DestinationAddress *string `field:"optional" json:"destinationAddress" yaml:"destinationAddress"`
	// For `type: durable_object_namespace`: The dispatch namespace the Durable Object script belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#dispatch_namespace WorkerVersion#dispatch_namespace}
	DispatchNamespace *string `field:"optional" json:"dispatchNamespace" yaml:"dispatchNamespace"`
	// For `type: service`: Entrypoint to invoke on the target Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#entrypoint WorkerVersion#entrypoint}
	Entrypoint *string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// For `type: durable_object_namespace`: The environment of the script_name to bind to.
	//
	// For `type: service`: Optional environment if the Worker utilizes one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#environment WorkerVersion#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// For `type: secret_key`: Data format of the key. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format). Available values: "raw", "pkcs8", "spki", "jwk".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#format WorkerVersion#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
	// For `type: d1`: Identifier of the D1 database to bind to.
	//
	// For `type: hyperdrive`: Identifier of the Hyperdrive connection to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#id WorkerVersion#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// For `type: vpc_network`: Enables Gateway identity for the binding.
	//
	// Requires network_id to be "cf1:network" and cannot be combined with tunnel_id or an explicit "cf1:vnet:<UUID>" network_id.
	// Available values: "runtime-email-alpha".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#identity WorkerVersion#identity}
	Identity *string `field:"optional" json:"identity" yaml:"identity"`
	// For `type: vectorize`: Name of the Vectorize index to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#index_name WorkerVersion#index_name}
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// For `type: ai_search`: The user-chosen instance name.
	//
	// Must exist at deploy time. The worker can search, chat, update, and manage items/jobs on this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#instance_name WorkerVersion#instance_name}
	InstanceName *string `field:"optional" json:"instanceName" yaml:"instanceName"`
	// For `type: json`: JSON data to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#json WorkerVersion#json}
	Json *string `field:"optional" json:"json" yaml:"json"`
	// For `type: r2_bucket`: The [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions) of the R2 bucket. Available values: "eu", "fedramp", "fedramp-high", "us".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#jurisdiction WorkerVersion#jurisdiction}
	Jurisdiction *string `field:"optional" json:"jurisdiction" yaml:"jurisdiction"`
	// For `type: secret_key`: Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#key_base64 WorkerVersion#key_base64}
	KeyBase64 *string `field:"optional" json:"keyBase64" yaml:"keyBase64"`
	// For `type: secret_key`: Key data in [JSON Web Key](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#json_web_key) format. Required if `format` is "jwk".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#key_jwk WorkerVersion#key_jwk}
	KeyJwk *string `field:"optional" json:"keyJwk" yaml:"keyJwk"`
	// For `type: ai_search`: The namespace the instance belongs to.
	//
	// Defaults to "default" if omitted. Customers who don't use namespaces can simply omit this field.
	//
	// For `type: ai_search_namespace`: The user-chosen namespace name. Must exist before deploy -- Wrangler handles auto-creation on deploy failure (R2 bucket pattern). The "default" namespace is auto-created by config-api for new accounts. Grants full access (CRUD + search + chat) to all instances within the namespace.
	//
	// For `type: messaging`: The Messaging namespace to bind to.
	//
	// For `type: artifacts`: The Artifacts namespace exposed to the Worker in the Worker's account. Must be 2-63 characters, start with an ASCII alphanumeric character, contain only ASCII alphanumeric characters, dots, underscores, and hyphens, and must not end with a hyphen. The namespace does not need to be created before binding it.
	//
	// For `type: dispatch_namespace`: The name of the dispatch namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#namespace WorkerVersion#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// For `type: durable_object_namespace` or `type: kv_namespace`: Namespace identifier tag.
	//
	// For `type: ratelimit`: Identifier of the rate limit namespace to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#namespace_id WorkerVersion#namespace_id}
	NamespaceId *string `field:"optional" json:"namespaceId" yaml:"namespaceId"`
	// For `type: vpc_network`: Identifier of the network to bind to.
	//
	// Use "cf1:network" to select the account's default virtual network at runtime, or "cf1:vnet:<UUID>" to select an explicit virtual network in the Worker's account. Mutually exclusive with tunnel_id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#network_id WorkerVersion#network_id}
	NetworkId *string `field:"optional" json:"networkId" yaml:"networkId"`
	// For `type: inherit`: The old name of the inherited binding.
	//
	// If set, the binding will be renamed from `old_name` to `name` in the new version. If not set, the binding will keep the same name between versions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#old_name WorkerVersion#old_name}
	OldName *string `field:"optional" json:"oldName" yaml:"oldName"`
	// For `type: dispatch_namespace`: Outbound worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#outbound WorkerVersion#outbound}
	Outbound *WorkerVersionBindingsOutbound `field:"optional" json:"outbound" yaml:"outbound"`
	// For `type: data_blob`: The name of the file containing the data content. Only accepted for `service worker syntax` Workers.
	//
	// For `type: text_blob`: The name of the file containing the text content. Only accepted for `service worker syntax` Workers.
	//
	// For `type: wasm_module`: The name of the file containing the WebAssembly module content. Only accepted for `service worker syntax` Workers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#part WorkerVersion#part}
	Part *string `field:"optional" json:"part" yaml:"part"`
	// For `type: pipelines`: Name of the Pipeline to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#pipeline WorkerVersion#pipeline}
	Pipeline *string `field:"optional" json:"pipeline" yaml:"pipeline"`
	// For `type: queue`: Name of the Queue to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#queue_name WorkerVersion#queue_name}
	QueueName *string `field:"optional" json:"queueName" yaml:"queueName"`
	// For `type: durable_object_namespace`: The script where the Durable Object is defined, if it is external to this Worker.
	//
	// For `type: workflow`: Script name that contains the Workflow. If not provided, defaults to this script name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#script_name WorkerVersion#script_name}
	ScriptName *string `field:"optional" json:"scriptName" yaml:"scriptName"`
	// For `type: secrets_store_secret`: Name of the secret in the store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#secret_name WorkerVersion#secret_name}
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
	// For `type: service`: Name of Worker to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#service WorkerVersion#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// For `type: vpc_service`: Identifier of the VPC service to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#service_id WorkerVersion#service_id}
	ServiceId *string `field:"optional" json:"serviceId" yaml:"serviceId"`
	// For `type: ratelimit`: The rate limit configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#simple WorkerVersion#simple}
	Simple *WorkerVersionBindingsSimple `field:"optional" json:"simple" yaml:"simple"`
	// For `type: secrets_store_secret`: ID of the store containing the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#store_id WorkerVersion#store_id}
	StoreId *string `field:"optional" json:"storeId" yaml:"storeId"`
	// For `type: k2`: ID of a K2 stream owned by the account deploying the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#stream WorkerVersion#stream}
	Stream *string `field:"optional" json:"stream" yaml:"stream"`
	// For `type: plain_text`: The text value to use.
	//
	// For `type: secret_text`: The secret value to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#text WorkerVersion#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
	// For `type: vpc_network`: UUID of the Cloudflare Tunnel to bind to. Mutually exclusive with network_id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#tunnel_id WorkerVersion#tunnel_id}
	TunnelId *string `field:"optional" json:"tunnelId" yaml:"tunnelId"`
	// For `type: secret_key`: Allowed operations with the key. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#usages WorkerVersion#usages}
	Usages *[]*string `field:"optional" json:"usages" yaml:"usages"`
	// For `type: inherit`: Identifier for the version to inherit the binding from.
	//
	// This can be a version ID, or the literal "latest" to inherit from the most recently uploaded version, which may not be the deployed version. Defaults to "latest". Only the Workers API version endpoints under `/accounts/{account_id}/workers/workers/{worker_id}/versions` accept a version ID. Other endpoints, including script upload and script version upload, accept only "latest" and reject a version ID with error 10057.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#version_id WorkerVersion#version_id}
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
	// For `type: workflow`: Name of the Workflow to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#workflow_name WorkerVersion#workflow_name}
	WorkflowName *string `field:"optional" json:"workflowName" yaml:"workflowName"`
}

