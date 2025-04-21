Setup and Running the Project

Step 1: Install Dependencies
Run the following command to install the necessary Go dependencies:

go mod tidy
Step 2: Build the Project
Compile the project using the following command:

go build -o raft3d main.go
This will create an executable file named raft3d.

Step 3: Run the Raft3D Nodes
To set up a Raft cluster with three nodes, follow these steps:

Node 1 (Initial Node):

./raft3d -node-id node1 -raft-addr 192.168.1.10:5000 -http-addr :8080
Node 2 (Joining Node):

./raft3d -node-id node2 -raft-addr 192.168.1.11:5000 -http-addr :8080 -join-addr 192.168.1.10:5000
Node 3 (Joining Node):

./raft3d -node-id node3 -raft-addr 192.168.1.12:5000 -http-addr :8080 -join-addr 192.168.1.10:5000
Step 4: Interact with the API
After setting up the nodes, you can interact with the Raft3D API using curl:

Add a Printer:

curl -X POST -H "Content-Type: application/json" -d '{"id":"p1","company":"Creality","model":"Ender 3"}' http://localhost:8080/api/v1/printers
View All Printers:

curl http://localhost:8080/api/v1/printers
Add a Filament:

curl -X POST -H "Content-Type: application/json" -d '{"id":"f1","type":"PLA","color":"red","total_weight_in_grams":1000,"remaining_weight_in_grams":1000}' http://localhost:8080/api/v1/filaments
Add a Print Job:

curl -X POST -H "Content-Type: application/json" -d '{"id":"j1","printer_id":"p1","filament_id":"f1","filepath":"prints/test.gcode","print_weight_in_grams":50}' http://localhost:8080/api/v1/print_jobs
Update Print Job Status to "Running":

curl -X POST http://localhost:8080/api/v1/print_jobs/j1?status=Running
Update Print Job Status to "Done":

curl -X POST http://localhost:8080/api/v1/print_jobs/j1?status=Done
View All Print Jobs:

curl http://localhost:8080/api/v1/print_jobs
View All Filaments:

curl http://localhost:8080/api/v1/filaments
Step 5: Verify and Monitor
Ensure that the Raft nodes are synchronized.
Use curl commands to monitor and interact with printers, filaments, and print jobs.
