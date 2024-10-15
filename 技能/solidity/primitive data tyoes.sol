//boolean,uint256,int256,address
// SPDX-License-Identifier:MIT
pragma solidity ^0.8.26;
contract Primitives {
    bool public boo = true ;
    uint256 public u = 123;
    int256 public minInt = type(int256).min;
    address public addr = 0XCA;
    
}