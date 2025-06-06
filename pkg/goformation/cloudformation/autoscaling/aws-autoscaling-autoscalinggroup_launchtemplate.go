package autoscaling

import (
	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/policies"
)

// AutoScalingGroup_LaunchTemplate AWS CloudFormation Resource (AWS::AutoScaling::AutoScalingGroup.LaunchTemplate)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplate.html
type AutoScalingGroup_LaunchTemplate struct {

	// LaunchTemplateSpecification AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplate.html#cfn-autoscaling-autoscalinggroup-launchtemplate-launchtemplatespecification
	LaunchTemplateSpecification *AutoScalingGroup_LaunchTemplateSpecification `json:"LaunchTemplateSpecification,omitempty"`

	// Overrides AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplate.html#cfn-autoscaling-autoscalinggroup-launchtemplate-overrides
	Overrides []AutoScalingGroup_LaunchTemplateOverrides `json:"Overrides,omitempty"`

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
func (r *AutoScalingGroup_LaunchTemplate) AWSCloudFormationType() string {
	return "AWS::AutoScaling::AutoScalingGroup.LaunchTemplate"
}
