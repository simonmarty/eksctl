package fargate

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/kris-nova/logger"
)

// DeleteProfile drains and delete the Fargate profile with the provided name.
func (c *Client) DeleteProfile(ctx context.Context, name string, waitForDeletion bool) error {
	if name == "" {
		return errors.New("invalid Fargate profile name: empty")
	}
	out, err := c.api.DeleteFargateProfile(ctx, deleteRequest(c.clusterName, name))
	logger.Debug("Fargate profile: delete request: received: %#v", out)
	if err != nil {
		return fmt.Errorf("failed to delete Fargate profile %q: %w", name, err)
	}
	if waitForDeletion {
		return c.waitForDeletion(ctx, name)
	}

	profiles, err := c.api.ListFargateProfiles(ctx, &eks.ListFargateProfilesInput{
		ClusterName: &c.clusterName,
	})

	if err != nil {
		return err
	}

	//If waitForDeletion is false then the profile might still exist until deletion finishes
	if len(profiles.FargateProfileNames) == 0 ||
		(len(profiles.FargateProfileNames) == 1 && profiles.FargateProfileNames[0] == name) {
		stack, err := c.stackManager.GetFargateStack(ctx)
		if err != nil {
			logger.Debug("failed to fetch fargate stack to delete, skipping deletion")
			return nil
		}

		if stack == nil {
			logger.Info("no fargate stack to delete")
		} else {
			logger.Info("deleting unused fargate role")
			_, err = c.stackManager.DeleteStackBySpec(ctx, stack)
			return err
		}
	}

	return nil
}

func (c *Client) waitForDeletion(ctx context.Context, name string) error {
	// Clone this client's policy to ensure this method is re-entrant/thread-safe:
	retryPolicy := c.retryPolicy.Clone()
	for !retryPolicy.Done() {
		names, err := c.ListProfiles(ctx)
		if err != nil {
			return err
		}
		if !contains(names, name) {
			return nil
		}

		timer := time.NewTimer(retryPolicy.Duration())
		select {
		case <-timer.C:

		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return fmt.Errorf("timed out while waiting for Fargate profile %q's deletion", name)
}

func deleteRequest(clusterName string, profileName string) *eks.DeleteFargateProfileInput {
	request := &eks.DeleteFargateProfileInput{
		ClusterName:        &clusterName,
		FargateProfileName: &profileName,
	}
	logger.Debug("Fargate profile: delete request: sending: %#v", request)
	return request
}

func contains(array []string, target string) bool {
	for _, value := range array {
		if value == target {
			return true
		}
	}
	return false
}
