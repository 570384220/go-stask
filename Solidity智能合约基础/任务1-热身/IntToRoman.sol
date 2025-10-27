// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract IntToRoman {

    mapping(uint => string) private romanNumerals;
    string romanStr;
    event intViewEvent(string roman);

    function intToRoman(string memory str, uint number) public returns (string memory) {
        romanNumerals[1] = "I";
        romanNumerals[4] = "IV";
        romanNumerals[5] = "V";
        romanNumerals[9] = "IV";
        romanNumerals[10] = "X";
        romanNumerals[40] = "XL";
        romanNumerals[50] = "L";
        romanNumerals[90] = "XC";
        romanNumerals[100] = "C";
        romanNumerals[400] = "CD";
        romanNumerals[500] = "D";
        romanNumerals[900] = "CM";
        romanNumerals[1000] = "M";

        if( number >= 1000) {
            (string memory conStr, uint leftNumber)= calculate(number, 1000);
            if(leftNumber > 0) {
                return intToRoman(str, leftNumber);
            } else {
                romanStr = conStr;
                emit intViewEvent(conStr);
                return conStr;
            }
        } else if(number >= 900) {
            (string memory conStr, uint leftNumber)= calculate(number, 900);
            if(leftNumber > 0) {
                return intToRoman(str, leftNumber);
            } else {
                return conStr;
            }
        }else if(number >= 500) {
            (string memory conStr, uint leftNumber)= calculate(number, 500);
            if(leftNumber > 0) {
                return intToRoman(str, leftNumber);
            } else {
                return conStr;
            }
        } else if(number >= 400) {
            (string memory conStr, uint leftNumber)= calculate(number, 400);
            if(leftNumber > 0) {
                return intToRoman(str, leftNumber);
            } else {
                return conStr;
            }
        } else if(number >= 100) {
            (string memory conStr, uint leftNumber)= calculate(number, 100);
            if(leftNumber > 0) {
                return intToRoman(str, leftNumber);
            } else {
                return conStr;
            }
        } else if(number >= 90) {
            (string memory conStr, uint leftNumber)= calculate(number, 90);
            if(leftNumber > 0) {
                return intToRoman(str, leftNumber);
            } else {
                return conStr;
            }
        } else if(number >= 50) {
            (string memory conStr, uint leftNumber)= calculate(number, 50);
            if(leftNumber > 0) {
                return intToRoman(str, leftNumber);
            } else {
                return conStr;
            }
        } else if(number >= 40) {
            (string memory conStr, uint leftNumber)= calculate(number, 40);
            if(leftNumber > 0) {
                return intToRoman(str, leftNumber);
            } else {
                return conStr;
            }
        } else if(number >= 10) {
            (string memory conStr, uint leftNumber)= calculate(number, 10);
            if(leftNumber > 0) {
                return intToRoman(str, leftNumber);
            } else {
                return conStr;
            }
        } else if(number >= 9) {
            (string memory conStr, uint leftNumber)= calculate(number, 9);
            if(leftNumber > 0) {
                return intToRoman(str, leftNumber);
            } else {
                return conStr;
            }
        } else if(number >= 5) {
            (string memory conStr, uint leftNumber)= calculate(number, 5);
            if(leftNumber > 0) {
                return intToRoman(str, leftNumber);
            } else {
                return conStr;
            }
        } else if(number >= 4) {
            (string memory conStr, uint leftNumber)= calculate(number, 4);
            if(leftNumber > 0) {
                return intToRoman(str, leftNumber);
            } else {
                return conStr;
            }
        } else {
            (string memory conStr, uint leftNumber)= calculate(number, 1);
            if(leftNumber > 0) {
                return intToRoman(str, leftNumber);
            } else {
                return conStr;
            }
        }
    }

    function calculate(uint num1, uint num2) public view returns (string memory, uint) {
        uint count = num1 / num2;
        string memory flag = romanNumerals[num2];
        string memory str = "";
        for(uint i=0; i< count; i++) {
            str = string.concat(str, flag);
        }

        uint leftNumber = num1 - count*num2;
        return (str, leftNumber);
    }

}