# Raft3D Prints Management using RAFT Consensus Algorithm

---

## 🚀 How to Run the Project

### 1. Install Dependencies  
```bash
go mod tidy
2. Build the Project
go build -o raft3d main.go
3. Run Raft3D Nodes
Raft3D operates on a Raft cluster with multiple nodes. Below are the steps to start three nodes:

🟢 Node 1 (Initial Node)

./raft3d -node-id node1 -raft-addr 192.168.1.10:5000 -http-addr :8080
🟡 Node 2 (Joining Node)

./raft3d -node-id node2 -raft-addr 192.168.1.11:5000 -http-addr :8080 -join-addr 192.168.1.10:5000
🔵 Node 3 (Joining Node)

./raft3d -node-id node3 -raft-addr 192.168.1.12:5000 -http-addr :8080 -join-addr 192.168.1.10:5000
