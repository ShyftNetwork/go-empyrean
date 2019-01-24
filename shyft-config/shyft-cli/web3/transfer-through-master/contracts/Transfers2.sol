pragma solidity ^0.4.24;

contract Transfers2 {
	function transfer(address _addr) public payable {
		_addr.transfer(msg.value);
	}
}