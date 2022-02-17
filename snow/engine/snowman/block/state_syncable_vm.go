// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/common"
)

// Snowman-VMs implementing state sync, need to be able to link a state summary
// to the block associated with it. This is achieved by structuring Summary.Key
// as the following DefaultSummaryKey/ProposerSummaryKey. Note that these structures
// do not reduce keys expressiveness since DefaultSummaryKey.ContentHash is
// totally defined by the Snowman-VM.

const StateSyncDefaultKeysVersion = 0

// DefaultSummaryKey is primarily used by platform and contract VM
// (ProposerVM only needs to track Default to Proposer summary keys mapping).
// Key is composed associating:
//     blkID of block associated with the Summary
//     hash of Summary content, which allows validating content-key relationship.
type DefaultSummaryKey struct {
	BlkID       ids.ID `serialize:"true"`
	ContentHash []byte `serialize:"true"`
}

// ProposerSummaryKey is used by ProposerVM.
// Key is composed associating:
// proposer block ID of the block wrapping InnerKey.BlkID block
// InnerKey as defined above.
type ProposerSummaryKey struct {
	ProBlkID ids.ID            `serialize:"true"`
	InnerKey DefaultSummaryKey `serialize:"true"`
}

type StateSyncableVM interface {
	common.StateSyncableVM

	// At the end of StateSync process, VM will have rebuilt the state of its blockchain
	// up to a given height. However the block associated with that height may be not known
	// to the VM yet. GetLastSummaryBlockID allows retrival of this block from network
	GetLastSummaryBlockID() (ids.ID, error)

	// SetLastSummaryBlock pass to VM the network-retrieved block associated with its last state summary
	SetLastSummaryBlock([]byte) error
}