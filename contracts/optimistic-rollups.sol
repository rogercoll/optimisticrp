// SPDX-License-Identifier: MIT

pragma solidity >=0.6.0 <=0.7.3;

import { Lib_MerkleTrie } from "./Lib_MerkleTrie.sol";
import { Lib_BytesUtils } from "./Lib_BytesUtils.sol";
import { Lib_RLPReader } from "./Lib_RLPReader.sol";

contract Optimistic_Rollups {
    
    bytes32 public stateRoot;
    bytes32 public prev_stateRoot;
    address public last_batch_submitter;
    uint256 public immutable lock_time;
    uint256 public immutable required_bond;
    uint256 public last_batch_time;
    mapping(address => address) public aggregators;
    mapping(bytes32 => bool) public valid_stateRoots;
    
    mapping(address => mapping(bytes32 => uint256)) private last_deposits;
    mapping(address => mapping(bytes32 => uint256)) private last_withdraws;
    event New_Deposit(address user, bytes32 stateRoot, uint256 value);
    event New_withdraw(address user, bytes32 stateRoot, uint256 value);
    event Fraud_Proved(address challenger);
    event Invalid_Proof(address challenger);



    constructor(
        uint256 _lock_time,
        uint256 _required_bond
    ) public {
        stateRoot = bytes32(0);
        prev_stateRoot = bytes32(0);
        lock_time = _lock_time;
        last_batch_time = block.timestamp - _lock_time;
        required_bond = _required_bond;
    }
    
    modifier is_aggregator(address user) {
        require(aggregators[user] != address(0), "UNAUTHORIZED_ACCOUNT");
        _;
    }
    
    modifier can_exit_optimism() {
        // Check that enough time has elapsed for potential fraud proofs (10 minutes)
        require (block.timestamp >= last_batch_time + lock_time, "OPTIMISTIC_PERIOD");
        _;
    }
    
    modifier fraud_period() {
        require (block.timestamp < last_batch_time + lock_time, "OPTIMISTIC_PERIOD");
        _;
    }
    
    //Bonds msg.value for msg.sender to become and aggregator
    function bond() external payable {
        require(msg.value >= required_bond, "INSUFFICIENT_BOND");
        aggregators[msg.sender] = msg.sender;
    }
    
    //Deposits funds so they can be used in layer2
    function deposit() external payable can_exit_optimism() {
        //we must verity stateRoot is valid
        last_deposits[msg.sender][stateRoot] += msg.value;
        emit New_Deposit(msg.sender, stateRoot, msg.value);
    }
    
    function withdraw(bytes calldata _key, bytes calldata _value, bytes memory _proof, bytes32 _root) external can_exit_optimism() {
        require(_root == stateRoot, "NOT_VALID_PROOF");
        
        //prevent double withdraw
        require(last_withdraws[msg.sender][stateRoot] == 0, "WITHDRAW_ALREADY_DONE");
        
        require(Lib_MerkleTrie.verifyInclusionProof(_key,_value,_proof,_root) == true, "INVALID_ACCOUNT_PROOF");
        address accAddr = Lib_BytesUtils.toAddress(_key,0);
        
        //Check msg.sender == proof address
        require (accAddr == msg.sender, "INVALID_WITHDRAW_REQUESTER");
        Lib_RLPReader.RLPItem[] memory account = Lib_RLPReader.readList(_value);
        uint256 accBalance = Lib_BytesUtils.toUint256(Lib_RLPReader.readBytes(account[1]));
        last_withdraws[msg.sender][stateRoot] += accBalance;
        msg.sender.transfer(accBalance);
        emit New_withdraw(msg.sender, stateRoot, accBalance);
    }
    
    

    
    
    //[248, 95, 160, 48, 90, 96, 180, 15, 169, 0, 12, 30, 160, 135, 133, 84, 61, 18, 113, 22, 62, 245, 86, 20, 148, 103, 136, 32, 124, 139, 204, 83, 60, 79, 110, 248, 60, 248, 58, 136, 13, 224, 182, 179, 167, 100, 0, 0, 133, 11, 164, 59, 116, 0, 148, 139, 80, 60, 161, 190, 245, 90, 144, 66, 118, 19, 143, 46, 166, 9, 6, 210, 197, 135, 129, 148, 4, 140, 130, 254, 44, 133, 149, 108, 242, 135, 47, 190, 50, 190, 74, 208, 109, 227, 219, 30, 1]
    function newBatch(bytes calldata _batch) external is_aggregator(msg.sender) can_exit_optimism() returns (string memory) {
        //here we should check if fraud proof time has expired
        require(_batch.length > 0, "EMPTY_NEW_BATCH");
        Lib_RLPReader.RLPItem[] memory ls = Lib_RLPReader.readList(_batch);
        Lib_RLPReader.RLPItem memory _stateRoot = ls[0];
        require(stateRoot == abi.decode(Lib_RLPReader.readBytes(_stateRoot), (bytes32)), "INVALID_PREV_STATEROOT");
        prev_stateRoot = stateRoot;
        Lib_RLPReader.RLPItem memory _newstateRoot = ls[1];
        stateRoot = abi.decode(Lib_RLPReader.readBytes(_newstateRoot), (bytes32));
        //At this point we consider the prevBatch as correct
        valid_stateRoots[prev_stateRoot] = true;
        last_batch_submitter = msg.sender;
        last_batch_time = block.timestamp;
    }
    
    //account_proof must contain a proof of the account balance for the previous stateRoot
    function prove_fraud(bytes calldata _key, bytes calldata _value, bytes memory _proof, bytes32 _root, bytes calldata _lastBatch) external fraud_period() {
        //Check proof
        require(_root == prev_stateRoot && stateRoot != prev_stateRoot, "NOT_VALID_PROOF");
        require(Lib_MerkleTrie.verifyInclusionProof(_key,_value,_proof,_root) == true, "INVALID_ACCOUNT_PROOF");
        
        //Extractic account values
        //Research what costs more: keccak256 or toAddress(_key)
        bytes32 accAddr = keccak256(_key); //we will compare with the hash of the bytes representing the account
        Lib_RLPReader.RLPItem[] memory account = Lib_RLPReader.readList(_value);
        uint256 accBalance = Lib_BytesUtils.toUint256(Lib_RLPReader.readBytes(account[1]));
        //uint256 accNonce = Lib_BytesUtils.toUint256(Lib_RLPReader.readBytes(account[0]));
        
        //We must increment and decrease account balance with its last deposits/withdraws as are not contemplated in the account proof
        accBalance += last_deposits[Lib_BytesUtils.toAddress(_key,0)][prev_stateRoot];
        accBalance -= last_withdraws[Lib_BytesUtils.toAddress(_key,0)][prev_stateRoot];

        
        //Now we must verify the value of the account after the applyed batch
        Lib_RLPReader.RLPItem[] memory ls = Lib_RLPReader.readList(_lastBatch);
        Lib_RLPReader.RLPItem[] memory transactions = Lib_RLPReader.readList(ls[2]);
        for (uint256 i = 0; i < transactions.length; i++) {
            Lib_RLPReader.RLPItem[] memory tx_data = Lib_RLPReader.readList(transactions[i]);
            //if is the receipent
            if (keccak256(Lib_RLPReader.readBytes(tx_data[2])) == accAddr) {
                accBalance += Lib_BytesUtils.toUint256(Lib_RLPReader.readBytes(tx_data[1]));
            } else if (keccak256(Lib_RLPReader.readBytes(tx_data[3])) == accAddr) {
                uint256 txValue = abi.decode(Lib_RLPReader.readBytes(tx_data[1]), (uint256));
                if (txValue > accBalance) {
                    emit Fraud_Proved(msg.sender);
                    delete aggregators[last_batch_submitter];
                    stateRoot = prev_stateRoot;
                    msg.sender.transfer(required_bond);
                    return;
                }
                accBalance -= txValue;
            }
        }
        emit Invalid_Proof(msg.sender);

        //if fraud is proved => change to the last apporved stateRoot and reward the prover
    }
    
    function remaining_proof_time() view public returns (uint256) {
        uint256 remaining = (last_batch_time + lock_time) - block.timestamp;
        if (remaining < 0) return 0;
        return remaining;
    }

}

