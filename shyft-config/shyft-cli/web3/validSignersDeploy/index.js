var Web3 = require('web3')
var fs = require('fs')
var path = require('path')
var solc = require('solc')

var validSignersFile = fs.readFileSync('./validSigners.sol')

var input = {
    language: 'Solidity',
    sources: {
        'ValidSigners.sol': {
            content: validSignersFile.toString()
        }
    },
    settings: {
        outputSelection: {
            '*': {
                '*': [ '*' ]
            }
        }
    }
}

var output = JSON.parse(solc.compile(JSON.stringify(input)))
 
console.log(output)

var abiB
var code

for (var contractName in output.contracts['ValidSigners.sol']) {
    console.log(contractName + ': ' + output.contracts['ValidSigners.sol'][contractName].evm.bytecode.object)
    code = output.contracts['ValidSigners.sol'][contractName].evm.bytecode.object
    var abi = output.contracts['ValidSigners.sol'][contractName].abi
    abiB = abi
    console.log(JSON.stringify(abi))
}

let web3 = new Web3();
web3.setProvider(new web3.providers.HttpProvider('http://localhost:8545'));

var account = web3.eth.accounts[0]

var validSignersContract = web3.eth.contract(abiB)

let contract = validSignersContract.new({from: account, gas: 1000000, data: '0x' + code});

