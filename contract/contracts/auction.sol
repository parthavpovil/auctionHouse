// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract AuctionVault {
    address public owner;

    // Track balances for each user
    mapping(address => uint256) public balances;

    // Events for transparency
    event Deposited(address indexed user, uint256 amount);
    event Withdrawn(address indexed user, uint256 amount);

    modifier onlyOwner() {
        require(msg.sender == owner, "Not authorized");
        _;
    }

    constructor() {
        owner = msg.sender;
    }

    // User deposits ETH
    function deposit() external payable {
        require(msg.value > 0, "Deposit must be greater than 0");
        balances[msg.sender] += msg.value;
        emit Deposited(msg.sender, msg.value);
    }

    // Owner approves and releases funds to a user
    function approveWithdrawal(address user, uint256 amount) external onlyOwner {
        require(balances[user] >= amount, "Insufficient balance");
        balances[user] -= amount;

        (bool success, ) = user.call{value: amount}("");
        require(success, "Transfer failed");

        emit Withdrawn(user, amount);
    }

    // Owner can check contract balance
    function contractBalance() external view returns (uint256) {
        return address(this).balance;
    }
}
