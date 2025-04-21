package api

import (
    "encoding/json"
    "net/http"
    "raft3d/models"
    "raft3d/raft"
)

func RegisterHandlers(mux *http.ServeMux, rn *raft.RaftNode) {
    mux.HandleFunc("/api/v1/printers", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "POST":
            var p models.Printer
            json.NewDecoder(r.Body).Decode(&p)
            cmd := raft.Command{Op: "add_printer", Value: json.RawMessage(marshal(p))}
            rn.Apply(cmd)
            json.NewEncoder(w).Encode(p)
        case "GET":
            json.NewEncoder(w).Encode(rn.FSM.Printers)
        }
    })

    mux.HandleFunc("/api/v1/filaments", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "POST":
            var f models.Filament
            json.NewDecoder(r.Body).Decode(&f)
            cmd := raft.Command{Op: "add_filament", Value: json.RawMessage(marshal(f))}
            rn.Apply(cmd)
            json.NewEncoder(w).Encode(f)
        case "GET":
            json.NewEncoder(w).Encode(rn.FSM.Filaments)
        }
    })

    mux.HandleFunc("/api/v1/print_jobs", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "POST":
            var pj models.PrintJob
            json.NewDecoder(r.Body).Decode(&pj)
            pj.Status = "Queued"
            cmd := raft.Command{Op: "add_print_job", Value: json.RawMessage(marshal(pj))}
            rn.Apply(cmd)
            json.NewEncoder(w).Encode(pj)
        case "GET":
            json.NewEncoder(w).Encode(rn.FSM.PrintJobs)
        }
    })

    mux.HandleFunc("/api/v1/print_jobs/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            id := r.URL.Path[len("/api/v1/print_jobs/"):]
            status := r.URL.Query().Get("status")
            if pj, ok := rn.FSM.PrintJobs[id]; ok {
                pj.Status = status
                cmd := raft.Command{Op: "update_print_job", Value: json.RawMessage(marshal(pj))}
                rn.Apply(cmd)
                json.NewEncoder(w).Encode(pj)
            } else {
                http.Error(w, "Print job not found", http.StatusNotFound)
            }
        }
    })
}

func marshal(v interface{}) []byte {
    data, _ := json.Marshal(v)
    return data
}