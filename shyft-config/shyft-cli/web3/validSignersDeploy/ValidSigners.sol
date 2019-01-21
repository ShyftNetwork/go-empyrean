pragma solidity ^0.5.1;

contract ValidSigners {

    mapping (address => bool) signers;
    address public owner;

    constructor() public {
      owner = msg.sender;
    }
    
    function addValidSigner(address _new_signer) public {
        if(msg.sender == owner) {
            signers[_new_signer] = true;
        } else {
            revert();
        }
    }
    
    function removeSigner(address _signer) public {
        if(msg.sender == owner) {
            signers[_signer] = false;
        } else {
            revert();
        }
    }

    function isValidSigner(address _signer) public view returns (bool result) {
        return signers[_signer];
    }
}
