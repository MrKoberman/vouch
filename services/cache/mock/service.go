// Copyright © 2022 Attestant Limited.
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
	"errors"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	cache "github.com/attestantio/vouch/services/cache"
)

// Service is a mock.
type Service struct {
	blockRootToSlotMap map[phase0.Root]phase0.Slot
}

// New creates a new mock cache.
func New(blockRootToSlotMap map[phase0.Root]phase0.Slot) cache.Service {
	return &Service{
		blockRootToSlotMap: blockRootToSlotMap,
	}
}

// BlockRootToSlot provides the slot for a given block root.
func (s *Service) BlockRootToSlot(ctx context.Context, root phase0.Root) (phase0.Slot, error) {
	slot, exists := s.blockRootToSlotMap[root]
	if exists {
		return slot, nil
	}
	return 0, errors.New("not found")
}
