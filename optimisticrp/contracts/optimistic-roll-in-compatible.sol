// SPDX-License-Identifier: MIT

pragma solidity >=0.6.0 <=0.7.3;
pragma experimental ABIEncoderV2;

interface Optimistic_Roll_In_Compatible {
  // User (address) is a mandatory first field
  function initialize_state(address user) external payable returns (bytes32 initial_state);

  // call_data will be function selector (bytes4), user (non-payable address), current state (bytes32), and abi encoded args
  function optimistic_call(bytes calldata call_data) external view returns (bytes32 new_state);

  // call_data will be function selector (bytes4), user (payable address), current state (bytes32), and abi encoded args
  function pessimistic_call(bytes calldata call_data) external payable returns (bytes32 new_state);
}
