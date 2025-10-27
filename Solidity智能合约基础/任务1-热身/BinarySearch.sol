
// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract BinarySearch {

    function binarySearch(int[] memory arr, int target) public pure returns ( int ) {
        uint low = 0;
        uint hight = arr.length - 1;
        
        while(low <= hight) {
            uint mid = uint((low + hight) / 2);
            if(arr[mid] == target) {
                return int(mid);
            } else if(arr[mid] > target) {
                hight = mid - 1;
            } else {
                low = mid + 1;
            }
        }
        return -1;
    }
}