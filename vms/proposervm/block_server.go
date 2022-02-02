// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proposervm

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/proposervm/indexer"
)

var _ indexer.BlockServer = &VM{}

// GetWrappingBlk implements BlockServer interface
func (vm *VM) GetWrappingBlk(blkID ids.ID) (indexer.WrappingBlock, error) {
	vm.ctx.Lock.Lock()
	defer vm.ctx.Lock.Unlock()

	return vm.getPostForkBlock(blkID)
}

// Commit implements BlockServer interface
func (vm *VM) Commit() error {
	vm.ctx.Lock.Lock()
	defer vm.ctx.Lock.Unlock()

	return vm.db.Commit()
}
