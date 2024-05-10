// Code generated by "internal/generate/listpages/main.go -AWSSDKVersion=2 -ListOps=ListActivatedRulesInRuleGroup,ListByteMatchSets,ListRegexMatchSets,ListRegexPatternSets,ListRuleGroups -Paginator=NextMarker"; DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func listActivatedRulesInRuleGroupPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListActivatedRulesInRuleGroupInput, fn func(*wafregional.ListActivatedRulesInRuleGroupOutput, bool) bool) error {
	for {
		output, err := conn.ListActivatedRulesInRuleGroup(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listByteMatchSetsPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListByteMatchSetsInput, fn func(*wafregional.ListByteMatchSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListByteMatchSets(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listRegexMatchSetsPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListRegexMatchSetsInput, fn func(*wafregional.ListRegexMatchSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListRegexMatchSets(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listRegexPatternSetsPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListRegexPatternSetsInput, fn func(*wafregional.ListRegexPatternSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListRegexPatternSets(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
func listRuleGroupsPages(ctx context.Context, conn *wafregional.Client, input *wafregional.ListRuleGroupsInput, fn func(*wafregional.ListRuleGroupsOutput, bool) bool) error {
	for {
		output, err := conn.ListRuleGroups(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}
