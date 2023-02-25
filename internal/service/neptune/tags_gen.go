// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package neptune

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/neptune"
	"github.com/aws/aws-sdk-go/service/neptune/neptuneiface"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

// ListTags lists neptune service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ListTags(ctx context.Context, conn neptuneiface.NeptuneAPI, identifier string) (tftags.KeyValueTags, error) {
	input := &neptune.ListTagsForResourceInput{
		ResourceName: aws.String(identifier),
	}

	output, err := conn.ListTagsForResourceWithContext(ctx, input)

	if err != nil {
		return tftags.New(ctx, nil), err
	}

	return KeyValueTags(ctx, output.TagList), nil
}

// []*SERVICE.Tag handling

// Tags returns neptune service tags.
func Tags(tags tftags.KeyValueTags) []*neptune.Tag {
	result := make([]*neptune.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &neptune.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from neptune service tags.
func KeyValueTags(ctx context.Context, tags []*neptune.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.StringValue(tag.Key)] = tag.Value
	}

	return tftags.New(ctx, m)
}

// UpdateTags updates neptune service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func UpdateTags(ctx context.Context, conn neptuneiface.NeptuneAPI, identifier string, oldTagsMap interface{}, newTagsMap interface{}) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	if removedTags := oldTags.Removed(newTags); len(removedTags) > 0 {
		input := &neptune.RemoveTagsFromResourceInput{
			ResourceName: aws.String(identifier),
			TagKeys:      aws.StringSlice(removedTags.IgnoreAWS().Keys()),
		}

		_, err := conn.RemoveTagsFromResourceWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	if updatedTags := oldTags.Updated(newTags); len(updatedTags) > 0 {
		input := &neptune.AddTagsToResourceInput{
			ResourceName: aws.String(identifier),
			Tags:         Tags(updatedTags.IgnoreAWS()),
		}

		_, err := conn.AddTagsToResourceWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}
