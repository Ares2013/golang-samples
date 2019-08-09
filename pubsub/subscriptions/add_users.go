// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package subscription

// [START pubsub_set_subscription_policy]

import (
	"context"
	"fmt"

	"cloud.google.com/go/iam"
	"cloud.google.com/go/pubsub"
)

// addUsers adds all IAM users to a subscription
func addUsers(c *pubsub.Client, subName string) error {
	ctx := context.Background()

	sub := c.Subscription(subName)
	policy, err := sub.IAM().Policy(ctx)
	if err != nil {
		return fmt.Errorf("Policy: %v", err)
	}
	// Other valid prefixes are "serviceAccount:", "user:"
	// See the documentation for more values.
	policy.Add(iam.AllUsers, iam.Viewer)
	policy.Add("group:cloud-logs@google.com", iam.Editor)
	if err := sub.IAM().SetPolicy(ctx, policy); err != nil {
		return fmt.Errorf("SetPolicy: %v", err)
	}
	// NOTE: It may be necessary to retry this operation if IAM policies are
	// being modified concurrently. SetPolicy will return an error if the policy
	// was modified since it was retrieved.
	return nil
}

// [END pubsub_set_subscription_policy]
