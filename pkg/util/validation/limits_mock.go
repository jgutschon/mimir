// SPDX-License-Identifier: AGPL-3.0-only
// Provenance-includes-location: https://github.com/cortexproject/cortex/blob/master/pkg/util/validation/limits_test.go
// Provenance-includes-license: Apache-2.0
// Provenance-includes-copyright: The Cortex Authors.

package validation

import "github.com/grafana/dskit/flagext"

// mockTenantLimits exposes per-tenant limits based on a provided map
type mockTenantLimits struct {
	limits map[string]*Limits
}

// NewMockTenantLimits creates a new mockTenantLimits that returns per-tenant limits based on
// the given map
func NewMockTenantLimits(limits map[string]*Limits) TenantLimits {
	return &mockTenantLimits{
		limits: limits,
	}
}

func (l *mockTenantLimits) ByUserID(userID string) *Limits {
	return l.limits[userID]
}

func (l *mockTenantLimits) AllByUserID() map[string]*Limits {
	return l.limits
}

func MockOverrides(customize func(defaults *Limits)) *Overrides {
	defaults := Limits{}
	flagext.DefaultValues(&defaults)
	customize(&defaults)

	overrides, err := NewOverrides(defaults, nil)
	if err != nil {
		// This function is expected to be used only in tests, so we're not afraid of panicking.
		panic(err)
	}

	return overrides
}
