// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract ReverseString {

    function reverse(string memory str) external pure returns (string memory) {
        bytes memory strBytes = bytes(str);

        uint len = strBytes.length;
        uint i= 0;
        uint j = len  - 1;
        for(; i< j;) {
            bytes1 temp = strBytes[i];
            strBytes[i] = strBytes[j];
            strBytes[j] = temp;
            i++;
            j--;
        }

        return string(strBytes);
    }
}