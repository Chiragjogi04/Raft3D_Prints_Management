# Raft3D

> A distributed 3D printer management system using Raft consensus algorithm for fault-tolerant data persistence and leader election.

---

## 🔧 Prerequisites

- Go 1.20 or later installed on your system
- `curl` for testing the HTTP API

---

## ⚙️ Installation & Build

1. **Install dependencies**

   ```bash
   go mod tidy
   ```

2. **Build the binary**

   ```bash
   go build -o raft3d main.go
   ```

---

## 🏃 Running the Nodes

1. **Start Node 1 (Initial Leader)**

   ```bash
   ./raft3d \
     -node-id node1 \
     -raft-addr 192.168.1.10:5000 \
     -http-addr :8080
   ```

2. **Start Node 2 (Join Node 1)**

   ```bash
   ./raft3d \
     -node-id node2 \
     -raft-addr 192.168.1.11:5000 \
     -http-addr :8081 \
     -join-addr 192.168.1.10:5000
   ```

3. **Start Node 3 (Join Node 1)**

   ```bash
   ./raft3d \
     -node-id node3 \
     -raft-addr 192.168.1.12:5000 \
     -http-addr :8082 \
     -join-addr 192.168.1.10:5000
   ```
---

## Working

1. **Add a Printer**

   ```bash
   curl -X POST \
     -H "Content-Type: application/json" \
     -d '{"id":"p1","company":"Creality","model":"Ender 3"}' \
     http://localhost:8080/api/v1/printers
   ```

2. **List Printers**

   ```bash
   curl http://localhost:8080/api/v1/printers
   ```

3. **Add a Filament**

   ```bash
   curl -X POST \
     -H "Content-Type: application/json" \
     -d '{"id":"f1","type":"PLA","color":"red","total_weight_in_grams":1000,"remaining_weight_in_grams":1000}' \
     http://localhost:8081/api/v1/filaments
   ```

4. **Submit a Print Job**

   ```bash
   curl -X POST \
     -H "Content-Type: application/json" \
     -d '{"id":"j1","printer_id":"p1","filament_id":"f1","filepath":"prints/test.gcode","print_weight_in_grams":50}' \
     http://localhost:8082/api/v1/print_jobs
   ```

5. **Update Print Job Status**

   ```bash
   curl -X POST http://localhost:8082/api/v1/print_jobs/j1?status=Running
   curl -X POST http://localhost:8082/api/v1/print_jobs/j1?status=Done
   ```

6. **List Print Jobs & Filaments**

   ```bash
   curl http://localhost:8082/api/v1/print_jobs
   curl http://localhost:8081/api/v1/filaments
   ```

---
