// SPDX-License-Identifier: MIT

pragma solidity >=0.6.0 <=0.7.3;
pragma experimental ABIEncoderV2;

import "./internal-merkle-library.sol";
import "./optimistic-roll-in-compatible.sol";

// TODO: perhaps a owner-controlled sighash whitelist for optimistically performing, and therefore a non-performing exit method
// TODO: consider verifying being view, or some kind of forced "come up for air" mechanism
// TODO: make required bond queryable
// TODO: logic_address cannot be immutable if it will be upgradeable

contract Optimistic_Roll_In {
  event ORI_New_State(address indexed user, bytes32 indexed new_state);
  event ORI_New_Optimistic_State(address indexed user, uint256 indexed block_time);
  event ORI_New_Optimistic_States(address indexed user, uint256 indexed block_time);
  event ORI_Locked(address indexed suspect, address indexed accuser);
  event ORI_Unlocked(address indexed suspect, address indexed accuser, uint256 indexed block_time);
  event ORI_Fraud_Proven(address indexed accuser, address indexed suspect, uint256 indexed transition_index);
  event ORI_Rolled_Back(address indexed user, uint256 indexed tree_size, uint256 indexed block_time);

  Optimistic_Roll_In_Compatible public immutable logic_contract;
  uint256 public immutable lock_time;
  uint256 public immutable required_bond;

  mapping(address => bytes32) public account_states;
  mapping(address => uint256) public balances;
  mapping(address => address) public lockers;
  mapping(address => uint256) public locked_timestamps;
  mapping(address => uint256) public rollback_sizes;

  constructor(
    address _logic_address,
    uint256 _lock_time,
    uint256 _required_bond
  ) public {
    logic_contract = Optimistic_Roll_In_Compatible(_logic_address);
    lock_time = _lock_time;
    required_bond = _required_bond;
  }

  modifier not_initialized(address user) {
    require(account_states[user] == bytes32(0), "ALREADY_INITIALIZED");
    _;
  }

  modifier not_locked(address user) {
    require(lockers[user] == address(0), "ACCOUNT_LOCKED");
    _;
  }

  modifier is_locked(address user) {
    require(lockers[user] != address(0), "ACCOUNT_UNLOCKED");
    _;
  }

  modifier no_rollback_required(address user) {
    require(rollback_sizes[user] == 0, "ROLLBACK_REQUIRED");
    _;
  }

  modifier can_exit_optimism(uint256 last_time) {
    // Check that enough time has elapsed for potential fraud proofs (10 minutes)
    require(last_time + lock_time < block.timestamp, "INSUFFICIENT_TIME");
    _;
  }

  modifier sufficient_calldata(bytes[] calldata call_data) {
    // Check that there is more than 1 call data (if not, user should have called perform_optimistically)
    require(call_data.length > 1, "INSUFFICIENT_CALLDATA");
    _;
  }

  modifier lock_expired(address user) {
    require(locked_timestamps[user] + lock_time <= block.timestamp, "INSUFFICIENT_WINDOW");
    _;
  }

  // Fallback to receive ETH and bond msg.value for msg.sender
  receive() external payable {
    apply_bond(msg.sender);
  }

  // Bonds amount for user, and reverts if resulting balance less than 1 ETH
  function apply_bond(address user) internal returns (uint256 remainder) {
    if (balances[user] == required_bond) return msg.value;

    require(msg.value >= required_bond, "INSUFFICIENT_BOND");

    balances[user] = required_bond;
    return msg.value - required_bond;
  }

  // Bonds msg.value for user
  function bond(address user) external payable {
    require(apply_bond(user) == 0, "INVALID_BOND");
  }

  // Sets user's account state to starting point, and bonds msg.value
  function initialize() external payable not_initialized(msg.sender) {
    address user = msg.sender;
    uint256 remainder = apply_bond(user);

    // call the initializer, passing any remaining amount and get initial state (S_0)
    bytes32 initial_state = logic_contract.initialize_state{ value: remainder }(user);

    // Set account state to combination of empty call data tree, initial state (S_0), and last time of 0 (not in optimism)
    account_states[user] = keccak256(abi.encodePacked(bytes32(0), initial_state, bytes32(0)));

    emit ORI_New_State(user, initial_state);
  }

  // Allows unbonding of ETH if account not locked (clears account state)
  function unbond(address payable destination) external not_locked(msg.sender) {
    address user = msg.sender;
    uint256 amount = balances[user];
    balances[user] = 0;
    account_states[user] = bytes32(0);
    destination.transfer(amount);
  }

  // Withdraws entire balance if no account state, or all minus bond if an account state exists
  function withdraw(address payable destination) external not_locked(msg.sender) {
    address user = msg.sender;
    bool account_exists = account_states[user] != bytes32(0);
    uint256 account_balance = balances[user];
    uint256 amount = account_exists ? account_balance - required_bond : account_balance;

    // Check for underflow
    require(amount <= account_balance, "UNEXPECTED_UNDERFLOW");

    balances[user] = account_balance - amount;
    destination.transfer(amount);
  }

  // Allows unbonding of all user's ETH balance if account has or can have non-optimistic state
  function archive(
    address payable destination,
    bytes32 call_data_root,
    bytes32 state,
    uint256 last_time
  ) external not_locked(msg.sender) can_exit_optimism(last_time) {
    address user = msg.sender;

    // Check that the user it not in an optimistic state, which means that their account state is
    // a call data tree, current state (S_n), and the last block
    require(
      keccak256(abi.encodePacked(call_data_root, state, bytes32(last_time))) == account_states[user],
      "INVALID_ROOTS"
    );

    // Set the account state to an empty call data tree, current state (S_n), and last time 0
    account_states[user] = keccak256(abi.encodePacked(bytes32(0), state, bytes32(0)));

    // Make user their own locker, requiring them to unlock if they wish to unarchive
    lockers[user] = user;

    uint256 amount = balances[user];
    balances[user] = 0;
    destination.transfer(amount);
  }

  // Bonds msg.value for user of an archived account (that they have locked)
  function unarchive() external payable is_locked(msg.sender) {
    address user = msg.sender;

    require(apply_bond(user) == 0, "INVALID_BOND");

    lockers[user] = address(0);
  }

  // Returns true if calling the logic contract with the call data results in new state
  function verify_transition(
    address user,
    bytes calldata call_data,
    bytes32 new_state
  ) internal view returns (bool) {
    // Check that the user is the user extracted from calldata (20 bytes starting after the function signature)
    if (user != abi.decode(call_data[4:], (address))) return false;

    // Compute a new state
    try logic_contract.optimistic_call(call_data) returns (bytes32 state) {
      return state == new_state;
    } catch {
      return false;
    }
  }

  // Calls logic contract and updates account state
  function normal_perform(
    address caller,
    bytes calldata call_data,
    bytes32 call_data_root,
    uint256 last_time
  ) internal not_locked(caller) {
    // Check that the caller is the user extracted from calldata (20 bytes starting after the function signature)
    require(caller == abi.decode(call_data[4:], (address)), "CALLER_USER_MISMATCH");

    // Extract current state (S_n or S_0) from calldata (32 bytes starting after the sig and user)
    bytes32 state = abi.decode(call_data[36:], (bytes32));

    // Check that the user it not in an optimistic state, which means that their account state is
    // a call data tree, current state (S_n or S_0), and the last block is 0
    require(
      keccak256(abi.encodePacked(call_data_root, state, bytes32(last_time))) == account_states[caller],
      "INVALID_ROOTS"
    );

    // Compute a new state (S_n+1 or S_1), reusing the state variable
    state = logic_contract.pessimistic_call{ value: msg.value }(call_data);

    // Set the account state to an empty call data tree, the new state (S_n+1 or S_1), and last time 0
    account_states[caller] = keccak256(abi.encodePacked(bytes32(0), state, bytes32(0)));

    emit ORI_New_State(caller, state);
  }

  // Set the account state to the on-chain computed new state, if the account is not locked
  function perform(bytes calldata call_data) external payable {
    normal_perform(msg.sender, call_data, bytes32(0), 0);
  }

  // Exits optimism by setting the account state to the on-chain computed new state, if the account is not locked
  function perform_and_exit(
    bytes calldata call_data,
    bytes32 call_data_root,
    uint256 last_time
  ) external payable no_rollback_required(msg.sender) can_exit_optimism(last_time) {
    normal_perform(msg.sender, call_data, call_data_root, last_time);
  }

  // Updates account state with optimistic transition data
  function optimistic_perform(
    address caller,
    bytes calldata call_data,
    bytes32 new_state,
    bytes32 call_data_root,
    bytes32[] calldata proof,
    uint256 last_time
  ) internal not_locked(caller) {
    // Extract current state (S_0) from call data (32 bytes starting after sig and user)
    bytes32 state = abi.decode(call_data[36:], (bytes32));

    // Check that the user it not in an optimistic state, which means that their account state is
    // a call data tree, current state (S_0), and the last block
    require(
      keccak256(abi.encodePacked(call_data_root, state, bytes32(last_time))) == account_states[caller],
      "INVALID_ROOTS"
    );

    // Get root of new merkle tree with 1 call data element (CD_0), reusing state as call_data_root
    state = Internal_Merkle_Library.try_append_one_c(call_data_root, call_data, proof);

    // Combine call data root, new state (S_1), and current time into account state
    account_states[caller] = keccak256(abi.encodePacked(state, new_state, bytes32(block.timestamp)));

    emit ORI_New_Optimistic_State(caller, block.timestamp);
  }

  // Enters optimism by updating the account state optimistically with call data and a new state, if the account is not locked
  function perform_optimistically_and_enter(
    bytes calldata call_data,
    bytes32 new_state,
    bytes32[] calldata proof
  ) external {
    optimistic_perform(msg.sender, call_data, new_state, bytes32(0), proof, 0);
  }

  // Updates the account state optimistically with call data and a new state, if the account is not locked
  function perform_optimistically(
    bytes calldata call_data,
    bytes32 new_state,
    bytes32 call_data_root,
    bytes32[] calldata proof,
    uint256 last_time
  ) external {
    optimistic_perform(msg.sender, call_data, new_state, call_data_root, proof, last_time);
  }

  // Updates account state with optimistic batch transition data
  function optimistic_perform_many(
    address user,
    bytes[] calldata call_data,
    bytes32 new_state,
    bytes32 call_data_root,
    bytes32[] calldata proof,
    uint256 last_time
  ) internal not_locked(user) sufficient_calldata(call_data) {
    // Extract current state (S_n) from first call data (32 bytes starting after sig and user)
    bytes32 state = abi.decode(call_data[0][36:], (bytes32));

    // Check that the call data root, current state (S_n), and last time are valid for this user
    require(
      keccak256(abi.encodePacked(call_data_root, state, bytes32(last_time))) == account_states[user],
      "INVALID_ROOTS"
    );

    // Get new merkle root of call data tree, appending several call data (CD_n - CD_n+m), reusing state as call_data_root
    state = Internal_Merkle_Library.try_append_many_c(call_data_root, call_data, proof);

    // Combine call data root, new state (S_n+m), and current time into account state
    account_states[user] = keccak256(abi.encodePacked(state, new_state, bytes32(block.timestamp)));

    emit ORI_New_Optimistic_States(user, block.timestamp);
  }

  // Enters optimism by updating the account state optimistically with several call data and final state, if the account is not locked
  function perform_many_optimistically_and_enter(
    bytes[] calldata call_data,
    bytes32 new_state,
    bytes32[] calldata proof
  ) external {
    optimistic_perform_many(msg.sender, call_data, new_state, bytes32(0), proof, 0);
  }

  // Updates the account state optimistically with several call data and final state, if the account is not locked
  function perform_many_optimistically(
    bytes[] calldata call_data,
    bytes32 new_state,
    bytes32 call_data_root,
    bytes32[] calldata proof,
    uint256 last_time
  ) external {
    optimistic_perform_many(msg.sender, call_data, new_state, call_data_root, proof, last_time);
  }

  // Lock two users (suspect and accuser)
  // Note: accuser and suspect cannot already be locked, but this might have to change so a single accuser isn't overwhelmed with fraud
  function lock(address suspect) external payable not_locked(msg.sender) not_locked(suspect) {
    address accuser = msg.sender;

    // Lock both the accuser and the suspect
    lockers[suspect] = accuser;
    locked_timestamps[suspect] = block.timestamp;
    lockers[accuser] = accuser;
    locked_timestamps[accuser] = block.timestamp;

    // The accuser may be trying to bond at the same time (this also check that have enough bonded)
    require(apply_bond(accuser) == 0, "INVALID_BOND");

    emit ORI_Locked(suspect, accuser);
  }

  // Unlock two users (suspect and accuser)
  function unlock(
    address suspect,
    bytes32 state,
    bytes32 call_data_root,
    uint256 last_time
  ) external is_locked(suspect) lock_expired(suspect) no_rollback_required(suspect) {
    // Check that the call data root, current state (S_n), and last time are valid for this user
    require(
      keccak256(abi.encodePacked(call_data_root, state, bytes32(last_time))) == account_states[suspect],
      "INVALID_ROOTS"
    );

    // Unlock both accounts
    address accuser = lockers[suspect];
    lockers[suspect] = address(0);
    locked_timestamps[suspect] = 0;
    lockers[accuser] = address(0);
    locked_timestamps[accuser] = 0;

    // Give the suspect the accuser's bond for having not proven fraud within a reasonable time frame
    // TODO: consider burning some here to prevent self-reporting breakeven
    uint256 amount = balances[accuser];
    balances[accuser] = 0;
    balances[suspect] += amount;

    // If suspect was not in optimistic state, then just emit event and return
    if (last_time == 0) {
      emit ORI_Unlocked(suspect, accuser, 0);
      return;
    }

    // Combine call data root, current state (S_n), and current time into account state
    // Note: updating last time is important, to prevent suspect blocking fraud proofs by locking themselves
    account_states[suspect] = keccak256(abi.encodePacked(call_data_root, state, bytes32(block.timestamp)));

    emit ORI_Unlocked(suspect, accuser, block.timestamp);
  }

  // Reward accuser for proving fraud in a suspect's transition, and track the expected rolled back account state size
  function prove_fraud(
    address suspect,
    bytes[] calldata call_data,
    bytes32 state,
    bytes32 call_data_root,
    bytes32[] calldata proof,
    uint256 last_time
  ) external is_locked(suspect) {
    address accuser = msg.sender;

    // Only the user that flagged/locked the suspect can prove fraud
    require(lockers[suspect] == accuser, "NOT_LOCKER");

    // Check that the call data root, current state (S_n), and last time are valid for this user
    require(
      keccak256(abi.encodePacked(call_data_root, state, bytes32(last_time))) == account_states[suspect],
      "INVALID_ROOTS"
    );

    // Check that the call data exist
    require(Internal_Merkle_Library.elements_exist_c(call_data_root, call_data, proof), "INVALID_CALLDATA");

    // get the indices of the call data in the call data tre
    uint256[] memory call_data_indices = Internal_Merkle_Library.get_indices(call_data, proof);

    // The transition index is the index of the starting call data of the fraud proof
    uint256 transition_index = call_data_indices[0];

    // If only one call data is provided, the fraud involves the last call data and current state
    if (call_data.length == 1) {
      // Check that the call data index is the last (call data tree size minus 1)
      require(transition_index + 1 == uint256(proof[0]), "INCORRECT_CALLDATA");

      // Fail if the state transition was valid
      require(verify_transition(suspect, call_data[0], state) == false, "VALID_TRANSITION");
    } else {
      // Check that call data provided are consecutive
      require(transition_index + 1 == call_data_indices[1]);

      // Extract new state from second call data (32 bytes starting after the sig and user), reusing state var
      state = abi.decode(call_data[1][36:], (bytes32));

      // Fail if the state transition was valid
      require(verify_transition(suspect, call_data[0], state) == false, "VALID_TRANSITION");
    }

    // Take the suspect's bond and give it to the accuser, reusing last_time var
    // TODO: consider burning some here to prevent self-reporting breakeven
    last_time = balances[suspect];
    balances[suspect] = 0;
    balances[accuser] += last_time;

    // Unlock the accuser's account
    lockers[accuser] = address(0);
    locked_timestamps[accuser] = 0;

    // Set the rollback size to the amount of elements that should be in the call data tree once rolled back
    rollback_sizes[suspect] = transition_index;

    // Set the suspect as the reason for their account's lock
    lockers[suspect] = suspect;
    locked_timestamps[suspect] = 0;

    emit ORI_Fraud_Proven(accuser, suspect, transition_index);
  }

  // Rolls a user back, given the current roots, old roots, and a proof of the optimistic transitions between them
  function rollback(
    bytes32 rolled_back_call_data_root,
    bytes[] calldata rolled_back_call_data,
    bytes32[] calldata roll_back_proof,
    uint256 current_size,
    bytes32 current_size_proof,
    bytes32 call_data_root,
    bytes32 state,
    uint256 last_time
  ) external payable {
    address user = msg.sender;
    uint256 expected_size = rollback_sizes[user];
    require(expected_size != 0, "ROLLBACK_UNNECESSARY");

    // Check that the call data root, current state (S_n), and last time are valid for this user
    require(
      keccak256(abi.encodePacked(call_data_root, state, bytes32(last_time))) == account_states[user],
      "INVALID_ROOTS"
    );

    // Check that the provided size of the current call data tree is correct
    require(Internal_Merkle_Library.verify_size(call_data_root, current_size, current_size_proof), "INVALID_SIZE");

    // Allow incremental roll back by checking that the rolled back call data tree is smaller than the current tree
    uint256 rolled_back_size = uint256(roll_back_proof[0]);
    require(rolled_back_size < current_size, "INSUFFICIENT_ROLLBACK");

    // Check that this is not rolling back too far, though
    require(rolled_back_size >= expected_size, "ROLLBACK_TOO_DEEP");

    // Check that the current call data root is derived by appending the rolled back call data to the rolled back call data root
    require(
      Internal_Merkle_Library.try_append_many_c(rolled_back_call_data_root, rolled_back_call_data, roll_back_proof) ==
        call_data_root,
      "INVALID_ROLLBACK"
    );

    // Extract new state from first rolled back call data (32 bytes starting after the sig and user), reusing state var
    state = abi.decode(rolled_back_call_data[0][36:], (bytes32));

    // Combine rolled back call data root, new state (S_n-m), and current time into account state
    account_states[user] = keccak256(abi.encodePacked(rolled_back_call_data_root, state, bytes32(block.timestamp)));

    // Unlock the user and clear the rollback flag, if roll back is complete
    if (rolled_back_size == expected_size) {
      lockers[user] = address(0);
      rollback_sizes[user] = 0;
    }

    // The user may be trying to bond at the same time (this also check that have enough bonded)
    require(apply_bond(user) == 0, "INVALID_BOND");

    emit ORI_Rolled_Back(user, rolled_back_size, block.timestamp);
  }
}
