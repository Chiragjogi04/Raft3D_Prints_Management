package raft

import (
    "encoding/json"
    "io"
    "raft3d/models"
    "sync"

    "github.com/hashicorp/raft"
)

type FSM struct {
    mu         sync.Mutex
    Printers   map[string]models.Printer
    Filaments  map[string]models.Filament
    PrintJobs  map[string]models.PrintJob
}

type Command struct {
    Op    string          `json:"op"`
    Value json.RawMessage `json:"value"`
}

func (f *FSM) Apply(log *raft.Log) interface{} {
    f.mu.Lock()
    defer f.mu.Unlock()

    var cmd Command
    if err := json.Unmarshal(log.Data, &cmd); err != nil {
        return err
    }

    switch cmd.Op {
    case "add_printer":
        var p models.Printer
        json.Unmarshal(cmd.Value, &p)
        f.Printers[p.ID] = p
    case "add_filament":
        var fil models.Filament
        json.Unmarshal(cmd.Value, &fil)
        f.Filaments[fil.ID] = fil
    case "add_print_job":
        var pj models.PrintJob
        json.Unmarshal(cmd.Value, &pj)
        f.PrintJobs[pj.ID] = pj
    case "update_print_job":
        var pj models.PrintJob
        json.Unmarshal(cmd.Value, &pj)
        if old, exists := f.PrintJobs[pj.ID]; exists {
            if pj.Status == "Done" && old.Status == "Running" {
                filament := f.Filaments[old.FilamentID]
                filament.RemainingWeightInGrams -= old.PrintWeightInGrams
                f.Filaments[old.FilamentID] = filament
            }
            f.PrintJobs[pj.ID] = pj
        }
    }
    return nil
}

func (f *FSM) Snapshot() (raft.FSMSnapshot, error) {
    f.mu.Lock()
    defer f.mu.Unlock()
    return &Snapshot{State: f}, nil
}

func (f *FSM) Restore(rc io.ReadCloser) error {
    f.mu.Lock()
    defer f.mu.Unlock()
    var snap Snapshot
    if err := json.NewDecoder(rc).Decode(&snap); err != nil {
        return err
    }
    f.Printers = snap.State.Printers
    f.Filaments = snap.State.Filaments
    f.PrintJobs = snap.State.PrintJobs
    return nil
}

type Snapshot struct {
    State *FSM
}

func (s *Snapshot) Persist(sink raft.SnapshotSink) error {
    err := json.NewEncoder(sink).Encode(s)
    if err != nil {
        sink.Cancel()
        return err
    }
    return sink.Close()
}

func (s *Snapshot) Release() {}