// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package r2bucketsippy


type R2BucketSippySource struct {
	// Access Key ID of an IAM credential (ideally scoped to a single S3 bucket).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/r2_bucket_sippy#access_key_id R2BucketSippy#access_key_id}
	AccessKeyId *string `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	// Access key for the Azure Storage account. Mutually exclusive with sasToken.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/r2_bucket_sippy#account_key R2BucketSippy#account_key}
	AccountKey *string `field:"optional" json:"accountKey" yaml:"accountKey"`
	// Name of the Azure Storage account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/r2_bucket_sippy#account_name R2BucketSippy#account_name}
	AccountName *string `field:"optional" json:"accountName" yaml:"accountName"`
	// Name of the AWS S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/r2_bucket_sippy#bucket R2BucketSippy#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// URL to the S3-compatible API of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/r2_bucket_sippy#bucket_url R2BucketSippy#bucket_url}
	BucketUrl *string `field:"optional" json:"bucketUrl" yaml:"bucketUrl"`
	// Client email of an IAM credential (ideally scoped to a single GCS bucket).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/r2_bucket_sippy#client_email R2BucketSippy#client_email}
	ClientEmail *string `field:"optional" json:"clientEmail" yaml:"clientEmail"`
	// Available values: "aws", "gcs", "s3", "azure".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/r2_bucket_sippy#cloud_provider R2BucketSippy#cloud_provider}
	CloudProvider *string `field:"optional" json:"cloudProvider" yaml:"cloudProvider"`
	// Name of the Azure Blob Storage container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/r2_bucket_sippy#container R2BucketSippy#container}
	Container *string `field:"optional" json:"container" yaml:"container"`
	// Private Key of an IAM credential (ideally scoped to a single GCS bucket).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/r2_bucket_sippy#private_key R2BucketSippy#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Name of the AWS availability zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/r2_bucket_sippy#region R2BucketSippy#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Shared Access Signature token for the Azure Storage account. Mutually exclusive with accountKey.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/r2_bucket_sippy#sas_token R2BucketSippy#sas_token}
	SasToken *string `field:"optional" json:"sasToken" yaml:"sasToken"`
	// Secret Access Key of an IAM credential (ideally scoped to a single S3 bucket).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.23.0/docs/resources/r2_bucket_sippy#secret_access_key R2BucketSippy#secret_access_key}
	SecretAccessKey *string `field:"optional" json:"secretAccessKey" yaml:"secretAccessKey"`
}

