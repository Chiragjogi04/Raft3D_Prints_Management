# Raft3D Prints management using RAFT consensus Algorithm

---

##  How to Run the Project

1. **Install Dependencies**  
   ```properties
   go mod tidy

2. **Build the Project**:
   ```properties
   go build -o raft3d main.go

3. **Run Raft3D Nodes**:
   Raft3D operates on a Raft cluster with multiple nodes. Below are the steps to start three nodes.

   Node 1 (Initial Node):
   ```bash
   ./raft3d -node-id node1 -raft-addr 192.168.1.10:5000 -http-addr :8080

   Node 2 (Joining Node):
   ```bash
   ./raft3d -node-id node2 -raft-addr 192.168.1.11:5000 -http-addr :8080 -join-addr 192.168.1.10:5000
