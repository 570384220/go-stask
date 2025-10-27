
// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract Voting {
    
    uint voteFlag;
    mapping(uint => mapping(address => uint)) votes;

    function vote(address candidater, uint count) public returns (string memory){
        mapping(address => uint) storage voteInfo = votes[voteFlag];
        voteInfo[candidater] += count;
        return "Voted Success";
    }

    function getVotes(address candidater) public view returns (uint) {
        return votes[voteFlag][candidater];
    }

    function resetVotes() public returns (string memory) {
        voteFlag += 1;
        return "Votes Reset Success";
    }
}

// 0xd839aD9f93B64F3BC3e36D25016976009Abba2FC,1
// 0x0000000000000000000000000000000000001111,2