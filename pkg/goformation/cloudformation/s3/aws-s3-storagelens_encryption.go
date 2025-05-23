package s3

import (
	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/policies"
)

// StorageLens_Encryption AWS CloudFormation Resource (AWS::S3::StorageLens.Encryption)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-encryption.html
type StorageLens_Encryption struct {

	// SSEKMS AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-encryption.html#cfn-s3-storagelens-encryption-ssekms
	SSEKMS *StorageLens_SSEKMS `json:"SSEKMS,omitempty"`

	// SSES3 AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-encryption.html#cfn-s3-storagelens-encryption-sses3
	SSES3 interface{} `json:"SSES3,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *StorageLens_Encryption) AWSCloudFormationType() string {
	return "AWS::S3::StorageLens.Encryption"
}
