// Package imagebuilder provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version (devel) DO NOT EDIT.
package imagebuilder

// Defines values for Distributions.
const (
	DistributionsCentos8 Distributions = "centos-8"

	DistributionsCentos9 Distributions = "centos-9"

	DistributionsRhel84 Distributions = "rhel-84"

	DistributionsRhel85 Distributions = "rhel-85"
)

// Defines values for ImageStatusStatus.
const (
	ImageStatusStatusBuilding ImageStatusStatus = "building"

	ImageStatusStatusFailure ImageStatusStatus = "failure"

	ImageStatusStatusPending ImageStatusStatus = "pending"

	ImageStatusStatusRegistering ImageStatusStatus = "registering"

	ImageStatusStatusSuccess ImageStatusStatus = "success"

	ImageStatusStatusUploading ImageStatusStatus = "uploading"
)

// Defines values for ImageTypes.
const (
	ImageTypesAmi ImageTypes = "ami"

	ImageTypesAws ImageTypes = "aws"

	ImageTypesAzure ImageTypes = "azure"

	ImageTypesEdgeCommit ImageTypes = "edge-commit"

	ImageTypesEdgeContainer ImageTypes = "edge-container"

	ImageTypesEdgeInstaller ImageTypes = "edge-installer"

	ImageTypesGcp ImageTypes = "gcp"

	ImageTypesGuestImage ImageTypes = "guest-image"

	ImageTypesImageInstaller ImageTypes = "image-installer"

	ImageTypesRhelEdgeCommit ImageTypes = "rhel-edge-commit"

	ImageTypesRhelEdgeInstaller ImageTypes = "rhel-edge-installer"

	ImageTypesVhd ImageTypes = "vhd"

	ImageTypesVsphere ImageTypes = "vsphere"
)

// Defines values for UploadStatusStatus.
const (
	UploadStatusStatusFailure UploadStatusStatus = "failure"

	UploadStatusStatusPending UploadStatusStatus = "pending"

	UploadStatusStatusRunning UploadStatusStatus = "running"

	UploadStatusStatusSuccess UploadStatusStatus = "success"
)

// Defines values for UploadTypes.
const (
	UploadTypesAws UploadTypes = "aws"

	UploadTypesAwsS3 UploadTypes = "aws.s3"

	UploadTypesAzure UploadTypes = "azure"

	UploadTypesGcp UploadTypes = "gcp"
)

// AWSS3UploadRequestOptions defines model for AWSS3UploadRequestOptions.
type AWSS3UploadRequestOptions map[string]interface{}

// AWSS3UploadStatus defines model for AWSS3UploadStatus.
type AWSS3UploadStatus struct {
	Url string `json:"url"`
}

// AWSUploadRequestOptions defines model for AWSUploadRequestOptions.
type AWSUploadRequestOptions struct {
	ShareWithAccounts []string `json:"share_with_accounts"`
}

// AWSUploadStatus defines model for AWSUploadStatus.
type AWSUploadStatus struct {
	Ami    string `json:"ami"`
	Region string `json:"region"`
}

// ArchitectureItem defines model for ArchitectureItem.
type ArchitectureItem struct {
	Arch       string   `json:"arch"`
	ImageTypes []string `json:"image_types"`
}

// Architectures defines model for Architectures.
type Architectures []ArchitectureItem

// AzureUploadRequestOptions defines model for AzureUploadRequestOptions.
type AzureUploadRequestOptions struct {
	// Name of the resource group where the image should be uploaded.
	ResourceGroup string `json:"resource_group"`

	// ID of subscription where the image should be uploaded.
	SubscriptionId string `json:"subscription_id"`

	// ID of the tenant where the image should be uploaded. This link explains how
	// to find it in the Azure Portal:
	// https://docs.microsoft.com/en-us/azure/active-directory/fundamentals/active-directory-how-to-find-tenant
	TenantId string `json:"tenant_id"`
}

// AzureUploadStatus defines model for AzureUploadStatus.
type AzureUploadStatus struct {
	ImageName string `json:"image_name"`
}

// ComposeMetadata defines model for ComposeMetadata.
type ComposeMetadata struct {
	// ID (hash) of the built commit
	OstreeCommit *string `json:"ostree_commit,omitempty"`

	// Package list including NEVRA
	Packages *[]PackageMetadata `json:"packages,omitempty"`
}

// ComposeRequest defines model for ComposeRequest.
type ComposeRequest struct {
	Customizations *Customizations `json:"customizations,omitempty"`
	Distribution   Distributions   `json:"distribution"`
	ImageName      *string         `json:"image_name,omitempty"`
	ImageRequests  []ImageRequest  `json:"image_requests"`
}

// ComposeResponse defines model for ComposeResponse.
type ComposeResponse struct {
	Id string `json:"id"`
}

// ComposeStatus defines model for ComposeStatus.
type ComposeStatus struct {
	ImageStatus ImageStatus `json:"image_status"`
}

// ComposeStatusError defines model for ComposeStatusError.
type ComposeStatusError struct {
	Details *interface{} `json:"details,omitempty"`
	Id      int          `json:"id"`
	Reason  string       `json:"reason"`
}

// ComposesResponse defines model for ComposesResponse.
type ComposesResponse struct {
	Data  []ComposesResponseItem `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
}

// ComposesResponseItem defines model for ComposesResponseItem.
type ComposesResponseItem struct {
	CreatedAt string      `json:"created_at"`
	Id        string      `json:"id"`
	ImageName *string     `json:"image_name,omitempty"`
	Request   interface{} `json:"request"`
}

// Customizations defines model for Customizations.
type Customizations struct {
	Filesystem   *[]Filesystem `json:"filesystem,omitempty"`
	Packages     *[]string     `json:"packages,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
}

// DistributionItem defines model for DistributionItem.
type DistributionItem struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

// Distributions defines model for Distributions.
type Distributions string

// DistributionsResponse defines model for DistributionsResponse.
type DistributionsResponse []DistributionItem

// Filesystem defines model for Filesystem.
type Filesystem struct {
	MinSize    int    `json:"min_size"`
	Mountpoint string `json:"mountpoint"`
}

// GCPUploadRequestOptions defines model for GCPUploadRequestOptions.
type GCPUploadRequestOptions struct {
	// List of valid Google accounts to share the imported Compute Node image with.
	// Each string must contain a specifier of the account type. Valid formats are:
	//   - 'user:{emailid}': An email address that represents a specific
	//     Google account. For example, 'alice@example.com'.
	//   - 'serviceAccount:{emailid}': An email address that represents a
	//     service account. For example, 'my-other-app@appspot.gserviceaccount.com'.
	//   - 'group:{emailid}': An email address that represents a Google group.
	//     For example, 'admins@example.com'.
	//   - 'domain:{domain}': The G Suite domain (primary) that represents all
	//     the users of that domain. For example, 'google.com' or 'example.com'.
	//     If not specified, the imported Compute Node image is not shared with any
	//     account.
	ShareWithAccounts []string `json:"share_with_accounts"`
}

// GCPUploadStatus defines model for GCPUploadStatus.
type GCPUploadStatus struct {
	ImageName string `json:"image_name"`
	ProjectId string `json:"project_id"`
}

// HTTPError defines model for HTTPError.
type HTTPError struct {
	Detail string `json:"detail"`
	Title  string `json:"title"`
}

// HTTPErrorList defines model for HTTPErrorList.
type HTTPErrorList struct {
	Errors []HTTPError `json:"errors"`
}

// ImageRequest defines model for ImageRequest.
type ImageRequest struct {
	Architecture  string        `json:"architecture"`
	ImageType     ImageTypes    `json:"image_type"`
	Ostree        *OSTree       `json:"ostree,omitempty"`
	UploadRequest UploadRequest `json:"upload_request"`
}

// ImageStatus defines model for ImageStatus.
type ImageStatus struct {
	Error        *ComposeStatusError `json:"error,omitempty"`
	Status       ImageStatusStatus   `json:"status"`
	UploadStatus *UploadStatus       `json:"upload_status,omitempty"`
}

// ImageStatusStatus defines model for ImageStatus.Status.
type ImageStatusStatus string

// ImageTypes defines model for ImageTypes.
type ImageTypes string

// OSTree defines model for OSTree.
type OSTree struct {
	Ref *string `json:"ref,omitempty"`
	Url *string `json:"url,omitempty"`
}

// Package defines model for Package.
type Package struct {
	Name    string `json:"name"`
	Summary string `json:"summary"`
}

// PackageMetadata defines model for PackageMetadata.
type PackageMetadata struct {
	Arch      string  `json:"arch"`
	Epoch     *string `json:"epoch,omitempty"`
	Name      string  `json:"name"`
	Release   string  `json:"release"`
	Sigmd5    string  `json:"sigmd5"`
	Signature *string `json:"signature,omitempty"`
	Type      string  `json:"type"`
	Version   string  `json:"version"`
}

// PackagesResponse defines model for PackagesResponse.
type PackagesResponse struct {
	Data  []Package `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
}

// Readiness defines model for Readiness.
type Readiness struct {
	Readiness string `json:"readiness"`
}

// Subscription defines model for Subscription.
type Subscription struct {
	ActivationKey string `json:"activation-key"`
	BaseUrl       string `json:"base-url"`
	Insights      bool   `json:"insights"`
	Organization  int    `json:"organization"`
	ServerUrl     string `json:"server-url"`
}

// UploadRequest defines model for UploadRequest.
type UploadRequest struct {
	Options interface{} `json:"options"`
	Type    UploadTypes `json:"type"`
}

// UploadStatus defines model for UploadStatus.
type UploadStatus struct {
	Options interface{}        `json:"options"`
	Status  UploadStatusStatus `json:"status"`
	Type    UploadTypes        `json:"type"`
}

// UploadStatusStatus defines model for UploadStatus.Status.
type UploadStatusStatus string

// UploadTypes defines model for UploadTypes.
type UploadTypes string

// Version defines model for Version.
type Version struct {
	Version string `json:"version"`
}

// ComposeImageJSONBody defines parameters for ComposeImage.
type ComposeImageJSONBody ComposeRequest

// GetComposesParams defines parameters for GetComposes.
type GetComposesParams struct {
	// max amount of composes, default 100
	Limit *int `json:"limit,omitempty"`

	// composes page offset, default 0
	Offset *int `json:"offset,omitempty"`
}

// GetPackagesParams defines parameters for GetPackages.
type GetPackagesParams struct {
	// distribution to look up packages for
	Distribution Distributions `json:"distribution"`

	// architecture to look up packages for
	Architecture GetPackagesParamsArchitecture `json:"architecture"`

	// packages to look for
	Search string `json:"search"`

	// max amount of packages, default 100
	Limit *int `json:"limit,omitempty"`

	// packages page offset, default 0
	Offset *int `json:"offset,omitempty"`
}

// GetPackagesParamsArchitecture defines parameters for GetPackages.
type GetPackagesParamsArchitecture string

// ComposeImageJSONRequestBody defines body for ComposeImage for application/json ContentType.
type ComposeImageJSONRequestBody ComposeImageJSONBody
