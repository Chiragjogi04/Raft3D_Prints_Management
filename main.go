package main

import (
    "log"
    "net/http"
    "os"
    "path/filepath"
    "raft3d/api"
    "raft3d/raft"
)

func main() {

    nodeID := "node1" 
    raftDir := filepath.Join("./raft-data", nodeID)
    os.MkdirAll(raftDir, 0700)

    raftNode, err := raft.NewRaftNode(nodeID, raftDir, "localhost:5000")
    if err != nil {
        log.Fatalf("Failed to initialize Raft: %v", err)
    }

    // API setup
    mux := http.NewServeMux()
    api.RegisterHandlers(mux, raftNode)

    // Start server
    log.Println("Starting Raft3D server on :8080")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}