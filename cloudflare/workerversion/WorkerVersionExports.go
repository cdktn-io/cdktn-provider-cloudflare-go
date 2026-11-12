// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workerversion


type WorkerVersionExports struct {
	// For `type: worker`: Marks this entry as a Worker entrypoint export.
	//
	// For `type: durable-object`: Marks this entry as a Durable Object export.
	// Available values: "worker", "durable-object".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#type WorkerVersion#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// For `type: worker`: Cache override for this entrypoint. Overrides the Worker's global `cache_options.enabled` for this entrypoint only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#cache WorkerVersion#cache}
	Cache *WorkerVersionExportsCache `field:"optional" json:"cache" yaml:"cache"`
	// For `type: durable-object` and `state: created`: Name of the container (declared in the upload's `metadata.containers`) that backs this Durable Object. When set, the namespace is container-enabled. Valid only on live entries.
	//
	// For `state: expecting-transfer`: Name of the container (declared in the upload's
	// `metadata.containers`) that backs this Durable Object once the
	// transfer settles. Valid only on live entries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#container WorkerVersion#container}
	Container *string `field:"optional" json:"container" yaml:"container"`
	// For `state: renamed`: The destination class name.
	//
	// Must differ from the source class
	// (the map key) and must be declared as a live (`created`) entry
	// in the same `exports` map. Write-only: never present in GET
	// responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#renamed_to WorkerVersion#renamed_to}
	RenamedTo *string `field:"optional" json:"renamedTo" yaml:"renamedTo"`
	// For `state: created`: Live export. May be omitted; defaults to `created`.
	//
	// For `state: deleted`: Tombstone that deletes the namespace.
	//
	// For `state: renamed`: Tombstone that renames the namespace's class.
	//
	// For `state: transferred`: Tombstone that transfers the namespace to another script.
	//
	// For `state: expecting-transfer`: Target side of a two-phase transfer.
	// Available values: "created", "deleted", "renamed", "transferred", "expecting-transfer".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#state WorkerVersion#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// For `type: durable-object` and `state: created`: Durable Object storage backend.
	//
	// `sqlite` is the recommended (and
	// only) backend for new namespaces. `legacy-kv` is accepted only for
	// a class whose namespace already exists as KV-backed; the `exports`
	// flow never provisions a new `legacy-kv` namespace.
	//
	// For `type: durable-object` and `state: expecting-transfer`: Durable Object storage backend. `sqlite` is the recommended (and
	// only) backend for new namespaces. `legacy-kv` is accepted only for
	// a class whose namespace already exists as KV-backed; the `exports`
	// flow never provisions a new `legacy-kv` namespace.
	// Available values: "sqlite", "legacy-kv".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#storage WorkerVersion#storage}
	Storage *string `field:"optional" json:"storage" yaml:"storage"`
	// For `state: expecting-transfer`: The source script name to receive the namespace from.
	//
	// Must be
	// in the same account and dispatch-namespace context. Present on
	// reads for `expecting-transfer` entries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#transfer_from WorkerVersion#transfer_from}
	TransferFrom *string `field:"optional" json:"transferFrom" yaml:"transferFrom"`
	// For `state: transferred`: The destination script name.
	//
	// Must be in the same account and
	// the same dispatch-namespace context (or both non-dispatch).
	// Cross-dispatch-namespace transfers are rejected. Write-only:
	// never present in GET responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/worker_version#transferred_to WorkerVersion#transferred_to}
	TransferredTo *string `field:"optional" json:"transferredTo" yaml:"transferredTo"`
}

