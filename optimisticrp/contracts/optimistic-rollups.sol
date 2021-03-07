// SPDX-License-Identifier: MIT

pragma solidity >=0.6.0 <=0.7.3;

import "solidity-rlp/contracts/RLPReader.sol";

contract Optimistic_Rollups {
    
    bytes32 stateRoot;
    address to;
    string stringTest;
    // optional way to attach library functions to these data types.
    using RLPReader for RLPReader.RLPItem;
    using RLPReader for RLPReader.Iterator;
    using RLPReader for bytes;
    
    // lets assume that rlpBytes is an encoding of [[1, "nested"], 2, 0x<Address>]
    function someFunctionThatTakesAnEncodedItem(bytes memory rlpBytes) public {
        RLPReader.RLPItem[] memory ls = rlpBytes.toRlpItem().toList(); // must convert to an rlpItem first!

        RLPReader.RLPItem memory item = ls[0]; // the encoding of [1, "nested"].
        item.toList()[0].toUint(); // 1
        string(item.toList()[1].toBytes()); // "nested"

        ls[1].toUint(); // 2
        ls[2].toAddress(); // 0x<Address>
    }
    
    //[139, 104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100]
    function testRLP (bytes memory rlpBytes) public {
        RLPReader.RLPItem memory ls = rlpBytes.toRlpItem();
        stringTest = string(ls.toBytes());
    }
    
    
    //[248, 95, 160, 48, 90, 96, 180, 15, 169, 0, 12, 30, 160, 135, 133, 84, 61, 18, 113, 22, 62, 245, 86, 20, 148, 103, 136, 32, 124, 139, 204, 83, 60, 79, 110, 248, 60, 248, 58, 136, 13, 224, 182, 179, 167, 100, 0, 0, 133, 11, 164, 59, 116, 0, 148, 139, 80, 60, 161, 190, 245, 90, 144, 66, 118, 19, 143, 46, 166, 9, 6, 210, 197, 135, 129, 148, 4, 140, 130, 254, 44, 133, 149, 108, 242, 135, 47, 190, 50, 190, 74, 208, 109, 227, 219, 30, 1]
    function newBatch(bytes calldata _batch) external returns (string memory) {
        RLPReader.RLPItem[] memory ls = _batch.toRlpItem().toList();
        RLPReader.RLPItem memory _stateRoot = ls[0];
        stateRoot = abi.decode(_stateRoot.toBytes(), (bytes32));
        RLPReader.RLPItem[] memory transactions = ls[1].toList();
        RLPReader.RLPItem[] memory tx_data = transactions[0].toList();
        to = tx_data[2].toAddress();
    }
    
    //[160,48,90,96,180,15,169,0,12,30,160,135,133,84,61,18,113,22,62,245,86,20,148,103,136,32,124,139,204,83,60,79,110]
    function readHashRLP(bytes memory _hash) external returns (string memory) {
        RLPReader.RLPItem memory item = _hash.toRlpItem();
        stateRoot = abi.decode(item.toBytes(), (bytes32));
    }
    
    function readHash(bytes calldata _hash) external returns (string memory) {
        stateRoot = abi.decode(_hash[:32], (bytes32));
    }
    
    function getStateRoot() public view returns (bytes32){
        return stateRoot;
    }
    
    function getToAddress() public view returns (address){
        return to;
    }

    //bytes is a shorthand for byte[]
    function calldataTest(bytes calldata _example) external returns (string memory) {
        if (_example.length == 0) {
            stringTest = "Empty";
        } else {
            stringTest = "NotEmpty";
        }
        // cannot modify or return _exampleString
    }
    
    
    
    function getMessage() public view returns (string memory){
                    return stringTest;
    }
    
    

}
