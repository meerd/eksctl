package medialive

import (
	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/policies"
)

// Channel_TeletextDestinationSettings AWS CloudFormation Resource (AWS::MediaLive::Channel.TeletextDestinationSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-teletextdestinationsettings.html
type Channel_TeletextDestinationSettings struct {

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
func (r *Channel_TeletextDestinationSettings) AWSCloudFormationType() string {
	return "AWS::MediaLive::Channel.TeletextDestinationSettings"
}
