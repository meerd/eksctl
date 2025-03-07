package config

import (
	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/policies"
)

// RemediationConfiguration_RemediationParameterValue AWS CloudFormation Resource (AWS::Config::RemediationConfiguration.RemediationParameterValue)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-remediationconfiguration-remediationparametervalue.html
type RemediationConfiguration_RemediationParameterValue struct {

	// ResourceValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-remediationconfiguration-remediationparametervalue.html#cfn-config-remediationconfiguration-remediationparametervalue-resourcevalue
	ResourceValue *RemediationConfiguration_ResourceValue `json:"ResourceValue,omitempty"`

	// StaticValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-remediationconfiguration-remediationparametervalue.html#cfn-config-remediationconfiguration-remediationparametervalue-staticvalue
	StaticValue *RemediationConfiguration_StaticValue `json:"StaticValue,omitempty"`

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
func (r *RemediationConfiguration_RemediationParameterValue) AWSCloudFormationType() string {
	return "AWS::Config::RemediationConfiguration.RemediationParameterValue"
}
