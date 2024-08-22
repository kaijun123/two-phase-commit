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

4. Check behaviour of participant during pause
  - Need another port for the client to pause and read
  - Need 2 goroutines, one for coordinator and one for client
  - Need mutex for state. Might context switch while processing one packet and then suddenly process the other packet
  - Check for possible areas of deadlock

TODO:
1. Code out complete 2PC first
2. Improve client test coverage
3. Fault tolerance:
  - Read, and pause apis
  - If one node dies, and then comes back, it will ask for the latest values from the other participants
  - Check for partitioning. Do not allow for reads if the participant is part of a minority partition


2PC status:
Completed: (from participant POV)
- Normal successful prepare-commit
- Participant is partitioned after the prepare phase, and does not receive commit message
  - participant is now set to timeout if it does not receive a commit message after prepare
TODO:
- Add P2P communication/ Termination protocol when there is a timeout/ coord failure
- Part recovery protocol
- Coord failures - resends last message(?)


