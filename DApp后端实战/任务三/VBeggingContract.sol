// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract VBeggingContract {

    mapping(address => uint) donations;
    event donateEvent(address, uint);
    event withdrawEvent(address, uint, bytes);
    address owner;

    receive () external payable {}
    fallback () external payable {}

    constructor () {
        owner = msg.sender;
    }

    modifier checkOwner {
        require(owner == msg.sender, "you are no owner");
        _;
    }

    function donate () external payable {
        emit donateEvent(msg.sender, msg.value);
        donations[msg.sender] += msg.value;
    } 

    function withdraw () public checkOwner {
        uint balance = address(this).balance;
        require(balance > 0, "contruct amount is 0");
        (bool ok, bytes memory data) = payable(msg.sender).call{value: balance}("");
        if( !ok ) {
            revert("withdraw err");
        }
        emit withdrawEvent(address(this), balance, data);

    }

    function getDonation(address donater) public view returns (uint) {
        return donations[donater];
    }

}