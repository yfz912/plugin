// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

import (
	"github.com/33cn/chain33/common"
	"github.com/33cn/chain33/types"
	pt "github.com/33cn/plugin/plugin/dapp/paracross/types"
	"github.com/pkg/errors"
)

func (e *Paracross) Exec_Commit(payload *pt.ParacrossCommitAction, tx *types.Transaction, index int) (*types.Receipt, error) {
	a := newAction(e, tx)
	receipt, err := a.Commit(payload)
	if err != nil {
		clog.Error("Paracross commit failed", "error", err, "hash", common.Bytes2Hex(tx.Hash()))
		return nil, errors.Cause(err)
	}
	return receipt, nil
}

func (e *Paracross) Exec_AssetTransfer(payload *types.AssetsTransfer, tx *types.Transaction, index int) (*types.Receipt, error) {
	clog.Debug("Paracross.Exec", "transfer", "")
	_, err := e.checkTxGroup(tx, index)
	if err != nil {
		clog.Error("ParacrossActionAssetTransfer", "get tx group failed", err, "hash", common.Bytes2Hex(tx.Hash()))
		return nil, err
	}
	a := newAction(e, tx)
	receipt, err := a.AssetTransfer(payload)
	if err != nil {
		clog.Error("Paracross AssetTransfer failed", "error", err, "hash", common.Bytes2Hex(tx.Hash()))
		return nil, errors.Cause(err)
	}
	return receipt, nil
}

func (e *Paracross) Exec_AssetWithdraw(payload *types.AssetsWithdraw, tx *types.Transaction, index int) (*types.Receipt, error) {
	clog.Debug("Paracross.Exec", "withdraw", "")
	_, err := e.checkTxGroup(tx, index)
	if err != nil {
		clog.Error("ParacrossActionAssetWithdraw", "get tx group failed", err, "hash", common.Bytes2Hex(tx.Hash()))
		return nil, err
	}
	a := newAction(e, tx)
	receipt, err := a.AssetWithdraw(payload)
	if err != nil {
		clog.Error("ParacrossActionAssetWithdraw failed", "error", err, "hash", common.Bytes2Hex(tx.Hash()))
		return nil, errors.Cause(err)
	}
	return receipt, nil
}

func (e *Paracross) Exec_Miner(payload *pt.ParacrossMinerAction, tx *types.Transaction, index int) (*types.Receipt, error) {
	if index != 0 {
		return nil, pt.ErrParaMinerBaseIndex
	}
	if !types.IsPara() {
		return nil, types.ErrNotSupport
	}
	a := newAction(e, tx)
	return a.Miner(payload)
}

func (e *Paracross) Exec_Transfer(payload *types.AssetsTransfer, tx *types.Transaction, index int) (*types.Receipt, error) {
	a := newAction(e, tx)
	return a.Transfer(payload, tx, index)
}

func (e *Paracross) Exec_Withdraw(payload *types.AssetsWithdraw, tx *types.Transaction, index int) (*types.Receipt, error) {
	a := newAction(e, tx)
	return a.Withdraw(payload, tx, index)
}

func (e *Paracross) Exec_TransferToExec(payload *types.AssetsTransferToExec, tx *types.Transaction, index int) (*types.Receipt, error) {
	a := newAction(e, tx)
	return a.TransferToExec(payload, tx, index)
}
