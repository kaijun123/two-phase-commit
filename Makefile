protobuf:
	@echo "Compiling Protobuf"
	protoc -I=./proto --go_out=./proto ./proto/coordinator.proto
	protoc -I=./proto --go_out=./proto ./proto/participant.proto

start-participant:
	@echo "Starting participants"
	bash startParticipant.sh

start-client:
	@echo "Starting client"
	cd ./client && go run main.go

start-coordinator:
	@echo "Starting client"
	cd ./coordinator && go run main.go