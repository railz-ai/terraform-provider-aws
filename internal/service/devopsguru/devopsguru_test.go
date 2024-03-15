// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devopsguru_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func TestAccDevOpsGuru_serial(t *testing.T) {
	t.Parallel()

	testCases := map[string]map[string]func(t *testing.T){
		"ResourceCollection": {
			"basic":          testAccResourceCollection_basic,
			"cloudformation": testAccResourceCollection_cloudformation,
			"disappears":     testAccResourceCollection_disappears,
			// "tags":           testAccResourceCollection_tags,
		},
	}

	acctest.RunSerialTests2Levels(t, testCases, 0)
}
