
// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract MergesortArray {
    
    function mergeSort(uint[] memory arr1, uint[] memory arr2) public pure returns (uint[] memory) {
        uint[] memory arr3;
        uint count = 0;
        uint i = 0;
        uint j = 0;

        for(; i< arr1.length && j< arr2.length;){
            if(arr1[i] < arr2[j]) {
                arr3[count] = arr1[i];
                count++;
                i++;
            } else {
                arr3[count] = arr2[j];
                count++;
                j++;
            }
        }

        for(; i < arr1.length; i++) {
            arr3[count] = arr1[i];
            i++;
            count++;
        }

        for(; j < arr2.length; j++) {
            arr3[count] = arr2[i];
            j++;
            count++;
        }

        return arr3;
    }
}