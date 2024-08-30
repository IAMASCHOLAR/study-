// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/ERC20.sol";

contract Token is ERC20 {
    //is 关键字被用来表示继承关系
    constructor(uint256 initialSupply) ERC20("Token", "Toc"){
       //initialSupply表示初始供应量
        _mint(msg.sender, initialSupply);
        //msg是一个全局变量，提供了当前以太网交易的相关信息，是一个内置对象
    /*msg.sender表示发送当前交易的以太坊地址,
    msg.value：表示与当前交易一起发送的以太币（ether）的数量，单位是wei（1 ether = 10^18 wei）。
    msg.data：表示当前交易的数据部分，可以用来传递给智能合约的函数参数。
    msg.sig：表示当前交易调用的函数签名。
    msg.gas：表示当前交易剩余的gas数量。
    */
    //require 有条件判断的进行回滚，revert无条件回滚
    }
/*
    function _mint(address account, uint256 value) internal {
        if (account == address(0)) {
            revert ERC20InvalidReceiver(address(0));
        }
        _update(address(0), account, value);
    }
    address(0)表示空地址或者无效地址
*/
/*
function _update(address from, address to, uint256 value) internal virtual {
        if (from == address(0)) {
            // Overflow check required: The rest of the code assumes that totalSupply never overflows
            _totalSupply += value;
        } else {
            uint256 fromBalance = _balances[from];
            if (fromBalance < value) {
                revert ERC20InsufficientBalance(from, fromBalance, value);
            }
            unchecked {
                // Overflow not possible: value <= fromBalance <= totalSupply.
                _balances[from] = fromBalance - value;
            }
        }

        if (to == address(0)) {
            unchecked {
                // Overflow not possible: value <= totalSupply or value <= fromBalance <= totalSupply.
                _totalSupply -= value;
            }
        } else {
            unchecked {
                // Overflow not possible: balance + value is at most totalSupply, which we know fits into a uint256.
                _balances[to] += value;
            }
        }

        emit Transfer(from, to, value);
    }
*/

}