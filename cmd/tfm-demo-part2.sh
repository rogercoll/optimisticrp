#!/bin/bash

########################
# include the magic
########################
. demo-magic.sh

# hide the evidence
clear

# Aggregator sends a batch
p "The aggregator sends an invalid batch within the Fraud period, his account does not have enough funds"
pe "go run aggregator/send2/main.go"

# hide the evidence
wait
clear

# Aggregator sends a batch
p "The aggregator sends an invalid batch, his account does not have enough funds"
pe "go run aggregator/send2/main.go"

# hide the evidence
wait
clear

# Aggregator sends a batch
p "The challenger checks the on-chain state"
pe "go run challenger/main.go"
