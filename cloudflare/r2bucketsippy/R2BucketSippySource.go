// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package r2bucketsippy


type R2BucketSippySource struct {
	// For `r2_bucket_sippy_provider: aws` or `r2_bucket_sippy_provider: s3`: Access Key ID of an IAM credential (ideally scoped to a single S3 bucket).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/r2_bucket_sippy#access_key_id R2BucketSippy#access_key_id}
	AccessKeyId *string `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	// For `r2_bucket_sippy_provider: azure`: Access key for the Azure Storage account. Mutually exclusive with sasToken.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/r2_bucket_sippy#account_key R2BucketSippy#account_key}
	AccountKey *string `field:"optional" json:"accountKey" yaml:"accountKey"`
	// For `r2_bucket_sippy_provider: azure`: Name of the Azure Storage account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/r2_bucket_sippy#account_name R2BucketSippy#account_name}
	AccountName *string `field:"optional" json:"accountName" yaml:"accountName"`
	// For `r2_bucket_sippy_provider: aws`: Name of the AWS S3 bucket.
	//
	// For `r2_bucket_sippy_provider: gcs`: Name of the GCS bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/r2_bucket_sippy#bucket R2BucketSippy#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// For `r2_bucket_sippy_provider: s3`: URL to the S3-compatible API of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/r2_bucket_sippy#bucket_url R2BucketSippy#bucket_url}
	BucketUrl *string `field:"optional" json:"bucketUrl" yaml:"bucketUrl"`
	// For `r2_bucket_sippy_provider: gcs`: Client email of an IAM credential (ideally scoped to a single GCS bucket).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/r2_bucket_sippy#client_email R2BucketSippy#client_email}
	ClientEmail *string `field:"optional" json:"clientEmail" yaml:"clientEmail"`
	// Available values: "aws", "gcs", "s3", "azure".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/r2_bucket_sippy#cloud_provider R2BucketSippy#cloud_provider}
	CloudProvider *string `field:"optional" json:"cloudProvider" yaml:"cloudProvider"`
	// For `r2_bucket_sippy_provider: azure`: Name of the Azure Blob Storage container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/r2_bucket_sippy#container R2BucketSippy#container}
	Container *string `field:"optional" json:"container" yaml:"container"`
	// For `r2_bucket_sippy_provider: gcs`: Private Key of an IAM credential (ideally scoped to a single GCS bucket).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/r2_bucket_sippy#private_key R2BucketSippy#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// For `r2_bucket_sippy_provider: aws`: AWS region containing the source S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/r2_bucket_sippy#region R2BucketSippy#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// For `r2_bucket_sippy_provider: azure`: Shared Access Signature token for the Azure Storage account. Mutually exclusive with accountKey.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/r2_bucket_sippy#sas_token R2BucketSippy#sas_token}
	SasToken *string `field:"optional" json:"sasToken" yaml:"sasToken"`
	// For `r2_bucket_sippy_provider: aws` or `r2_bucket_sippy_provider: s3`: Secret Access Key of an IAM credential (ideally scoped to a single S3 bucket).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.29.0/docs/resources/r2_bucket_sippy#secret_access_key R2BucketSippy#secret_access_key}
	SecretAccessKey *string `field:"optional" json:"secretAccessKey" yaml:"secretAccessKey"`
}

