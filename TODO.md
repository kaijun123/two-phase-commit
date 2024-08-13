Done:
1. Scaffold
  - Create a txt file that contains all the ip:ports for the participants
    - will be used by the coordinator to ping the participants
  - write a bash script to start all of the participants by reading the bash script
  - get the participants to listen to the packets

2. Message tracking for the coordinator
  - Step 1: Request
  - Step 2: Commit

3. protobuf encoding for the messages

TODO:
1. Check behaviour of participant during pause
  - Need another port for the client to pause and read
  - Need 2 goroutines, one for coordinator and one for client
  - Need mutex for state. Might context switch while processing one packet and then suddenly process the other packet
  - Check for possible areas of deadlock
2. Improve client test coverage
3. Fault tolerance:
  - Read, and pause apis
  - If one node dies, and then comes back, it will ask for the latest values from the other participants
  - Check for partitioning. Do not allow for reads if the participant is part of a minority partitioni