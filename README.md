# Optimistic Rollups PoC

![Go](https://github.com/rogercoll/optimisticrp/workflows/Go/badge.svg)

## On-chain features

`deposit()` => Any account that want to add funds into the Optimistic layer 2 scheme, thus will be able to transfer ethers to other accounts.

`newBatch(bytes calldata _batch)` => It can only be used by aggregator nodes. A batch will contain the previous stateRoot, the new one and the collected transactions.


## Aggregator rules

1- Before submitting a new batch, the provided state root must contain all the deposits (`deposit()`) account update since the last submitted batch. 
