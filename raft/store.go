package raft

import (
    "encoding/json"
    "os"
    "path/filepath"
    "time"

    "raft3d/models" 
    "github.com/hashicorp/raft"
    "github.com/hashicorp/raft-boltdb"
)

type RaftNode struct {
    Raft    *raft.Raft
    FSM     *FSM
    RaftDir string
}

func NewRaftNode(nodeID, raftDir, bindAddr string) (*RaftNode, error) {
    config := raft.DefaultConfig()
    config.LocalID = raft.ServerID(nodeID)

    fsm := &FSM{
        Printers:  make(map[string]models.Printer),
        Filaments: make(map[string]models.Filament),
        PrintJobs: make(map[string]models.PrintJob),
    }

    logStore, err := raftboltdb.NewBoltStore(filepath.Join(raftDir, "raft-log.db"))
    if err != nil {
        return nil, err
    }
    stableStore, err := raftboltdb.NewBoltStore(filepath.Join(raftDir, "raft-stable.db"))
    if err != nil {
        return nil, err
    }
    snapshotStore, err := raft.NewFileSnapshotStore(raftDir, 1, os.Stderr)
    if err != nil {
        return nil, err
    }

    transport, err := raft.NewTCPTransport(bindAddr, nil, 3, 10*time.Second, os.Stderr)
    if err != nil {
        return nil, err
    }

    r, err := raft.NewRaft(config, fsm, logStore, stableStore, snapshotStore, transport)
    if err != nil {
        return nil, err
    }

    r.BootstrapCluster(raft.Configuration{
        Servers: []raft.Server{
            {
                ID:      config.LocalID,
                Address: transport.LocalAddr(),
            },
        },
    })

    return &RaftNode{Raft: r, FSM: fsm, RaftDir: raftDir}, nil
}

func (rn *RaftNode) Apply(cmd Command) error {
    data, _ := json.Marshal(cmd)
    future := rn.Raft.Apply(data, 10*time.Second)
    return future.Error()
}