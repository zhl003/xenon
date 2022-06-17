/*
 * Xenon
 *
 * Copyright 2018 The Xenon Authors.
 * Code is licensed under the GPLv3.
 *
 */

package raft

import (
	"strconv"

	"github.com/radondb/xenon/src/model"
)

// RaftRPC tuple.
type RaftRPC struct {
	raft *Raft
}

// Ping rpc.
// send MsgRaftPing
func (r *RaftRPC) Ping(req *model.RaftRPCRequest, rsp *model.RaftRPCResponse) error {
	ret, err := r.raft.send(MsgRaftPing, req, r.raft.getHeartbeatTimeout())
	if err != nil {
		return err
	}
	*rsp = *ret.(*model.RaftRPCResponse)
	return nil
}

// Heartbeat rpc.
func (r *RaftRPC) Heartbeat(req *model.RaftRPCRequest, rsp *model.RaftRPCResponse) error {
	ret, err := r.raft.send(MsgRaftHeartbeat, req, r.raft.getHeartbeatTimeout())
	if err != nil {
		return err
	}
	*rsp = *ret.(*model.RaftRPCResponse)
	return nil
}

// RequestVote rpc.
func (r *RaftRPC) RequestVote(req *model.RaftRPCRequest, rsp *model.RaftRPCResponse) error {
	ret, err := r.raft.send(MsgRaftRequestVote, req, r.raft.getHeartbeatTimeout())
	if err != nil {
		return err
	}
	*rsp = *ret.(*model.RaftRPCResponse)
	return nil
}

// Status rpc.
func (r *RaftRPC) Status(req *model.RaftStatusRPCRequest, rsp *model.RaftStatusRPCResponse) error {
	rsp.RetCode = model.OK
	rsp.State = r.raft.GetState().String()
	rsp.Stats = r.raft.getStats()
	rsp.IdleCount, _ = strconv.ParseUint(strconv.Itoa(len(r.raft.getIdlePeers())), 10, 64)
	return nil
}

// EnablePurgeBinlog rpc.
func (r *RaftRPC) EnablePurgeBinlog(req *model.RaftStatusRPCRequest, rsp *model.RaftStatusRPCResponse) error {
	r.raft.SetSkipPurgeBinlog(false)
	return nil
}

// DisablePurgeBinlog rpc.
func (r *RaftRPC) DisablePurgeBinlog(req *model.RaftStatusRPCRequest, rsp *model.RaftStatusRPCResponse) error {
	r.raft.SetSkipPurgeBinlog(true)
	return nil
}

// EnableCheckSemiSync rpc.
func (r *RaftRPC) EnableCheckSemiSync(req *model.RaftStatusRPCRequest, rsp *model.RaftStatusRPCResponse) error {
	r.raft.SetSkipCheckSemiSync(false)
	return nil
}

// DisableCheckSemiSync rpc.
func (r *RaftRPC) DisableCheckSemiSync(req *model.RaftStatusRPCRequest, rsp *model.RaftStatusRPCResponse) error {
	r.raft.SetSkipCheckSemiSync(true)
	return nil
}
