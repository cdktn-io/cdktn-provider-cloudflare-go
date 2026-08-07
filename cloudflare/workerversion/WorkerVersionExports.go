// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workerversion


type WorkerVersionExports struct {
	// The kind of export. Available values: "worker", "durable-object".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/worker_version#type WorkerVersion#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Cache override for this entrypoint. It applies only to `type: worker` entries and overrides the Worker's global `cache_options.enabled` for that entrypoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/worker_version#cache WorkerVersion#cache}
	Cache *WorkerVersionExportsCache `field:"optional" json:"cache" yaml:"cache"`
	// Destination class name for a `state: renamed` tombstone.
	//
	// The
	// target must appear as a live (`created`) entry in the same
	// `exports` map. Write-only: never present in GET responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/worker_version#renamed_to WorkerVersion#renamed_to}
	RenamedTo *string `field:"optional" json:"renamedTo" yaml:"renamedTo"`
	// Lifecycle state of the export entry. Defaults to `created` (a normal, live export) when omitted.
	//
	// `deleted`, `renamed`, and `transferred` are tombstones:
	// write-only lifecycle operations that retire, rename, or hand
	// off a provisioned Durable Object namespace. They are applied
	// at upload and are filtered out of GET responses, so a read
	// only ever returns `created` or `expecting-transfer`.
	//
	// `expecting-transfer` is a live export whose data is being
	// received from another script via the two-phase transfer flow;
	// it carries `storage` and `transfer_from`.
	// Available values: "created", "deleted", "renamed", "transferred", "expecting-transfer".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/worker_version#state WorkerVersion#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// Storage backend for a `type: durable-object` export.
	//
	// Required
	// for live Durable Object entries (`created` and
	// `expecting-transfer`). `sqlite` selects SQLite-backed storage;
	// `legacy-kv` selects the legacy key-value storage.
	// Available values: "sqlite", "legacy-kv".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/worker_version#storage WorkerVersion#storage}
	Storage *string `field:"optional" json:"storage" yaml:"storage"`
	// Source script for a `state: expecting-transfer` entry.
	//
	// The
	// namespace on this script is materialised from the source
	// script's data via the pending-transfer flow. Present on reads
	// for `expecting-transfer` entries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/worker_version#transfer_from WorkerVersion#transfer_from}
	TransferFrom *string `field:"optional" json:"transferFrom" yaml:"transferFrom"`
	// Destination script for a `state: transferred` tombstone.
	//
	// Must
	// reference a script in the same account; cross-dispatch-namespace
	// transfers are rejected. Write-only: never present in GET
	// responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/worker_version#transferred_to WorkerVersion#transferred_to}
	TransferredTo *string `field:"optional" json:"transferredTo" yaml:"transferredTo"`
}

