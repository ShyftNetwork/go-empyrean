var Web3 = require('web3')
var web3 = new Web3(new Web3.providers.HttpProvider("http://127.0.0.1:8545"));
var contract_address = "0x3f69c384ecaf00925ec5310754a1044db618a1c9"

var mytokenContract = web3.eth.contract([{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"inputs":[{"name":"initialSupply","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"}]).at(contract_address)

console.log(mytokenContract.balanceOf('0x887495999b72694811da9c2ee0a34de4c003332b'))

mytokenContract.transfer.sendTransaction('0xb276840e8b89c64b517629de60de861e85f539ca', 3, {from: '0x887495999b72694811da9c2ee0a34de4c003332b', gas: 3000000})
