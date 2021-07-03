#!/bin/bash

########################
# include the magic
########################
. demo-magic.sh

# hide the evidence
clear

p "Set the new smart contract address"
p 'sed "s/ContractAddr.=.*$/ContractAddr = \"$1\"/" data.go'
sed -i "s/ContractAddr.=.*$/ContractAddr = \"$1\"/" data.go

# cat out the demo environment
batcat -l Go data.go

# hide the evidence
wait
clear

# Aggregator deposit
p "The aggregator stakes 1.5 Ethers to the SC to start sending batches"
pei "go run aggregator/bond/main.go"

# Aggregator deposit
p "The aggregator deposits funds to the SC to start sending transactions"
pei "go run aggregator/deposit/main.go"

# hide the evidence
wait
clear

# Aggregator sends a batch
p "The aggregator sends a new batch with transactions to random receivers"
pe "go run aggregator/send/main.go"

# hide the evidence
wait
clear

# Challenger checks state 
p "Run a challenger node to verify the Optimistic SM state"
pe "go run challenger/main.go"
