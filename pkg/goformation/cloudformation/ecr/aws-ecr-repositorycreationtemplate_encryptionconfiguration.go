package ecr

import (
	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/types"

	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/policies"
)

// RepositoryCreationTemplate_EncryptionConfiguration AWS CloudFormation Resource (AWS::ECR::RepositoryCreationTemplate.EncryptionConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-repositorycreationtemplate-encryptionconfiguration.html
type RepositoryCreationTemplate_EncryptionConfiguration struct {

	// EncryptionType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-repositorycreationtemplate-encryptionconfiguration.html#cfn-ecr-repositorycreationtemplate-encryptionconfiguration-encryptiontype
	EncryptionType *types.Value `json:"EncryptionType,omitempty"`

	// KmsKey AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-repositorycreationtemplate-encryptionconfiguration.html#cfn-ecr-repositorycreationtemplate-encryptionconfiguration-kmskey
	KmsKey *types.Value `json:"KmsKey,omitempty"`

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
func (r *RepositoryCreationTemplate_EncryptionConfiguration) AWSCloudFormationType() string {
	return "AWS::ECR::RepositoryCreationTemplate.EncryptionConfiguration"
}
