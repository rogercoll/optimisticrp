#!/bin/bash

########################
# include the magic
########################
. demo-magic.sh

# hide the evidence
clear

# A user withdraws funds
p "A user provides a fraud proof to withdraw funds"
pe "go run user/withdraw/main.go"

# hide the evidence
wait
clear
