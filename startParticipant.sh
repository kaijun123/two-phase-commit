#!/bin/bash

cd "./participant"

trap 'kill 0' SIGINT;
for n in {8081..8085};
do
  echo $n
  go run ./main.go -port=$n &
done
wait;

# CURR_DIR=$(pwd)
# INFILE="${CURR_DIR}/participants.txt"

# while read -r line
# do
#   echo "inside while loop"

#   # Print the current line (for verification)
#   echo "Processing: $line"

#   IFS=':' read -r hostname port <<< "$line"

#   echo "hostname:${hostname}"
#   echo "port:${port}"

#   # # Run the Go program with the current line as the PORT argument
#   # go run ./participant/main.go "$line"

# done < $INFILE