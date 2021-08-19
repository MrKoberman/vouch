// Copyright © 2021 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mock

import (
	"context"

	"github.com/attestantio/go-eth2-client/spec/phase0"
)

// Service is a mock sync committee aggregator.
type Service struct{}

// New creates a new mock sync committee aggregator.
func New() *Service {
	return &Service{}
}

// SetBeaconBlockRoot sets the beacon block root used for a given slot.
func (s *Service) SetBeaconBlockRoot(slot phase0.Slot, root phase0.Root) {
}

// Aggregate carries out aggregation for a slot and committee.
func (s *Service) Aggregate(ctx context.Context, details interface{}) {
}
