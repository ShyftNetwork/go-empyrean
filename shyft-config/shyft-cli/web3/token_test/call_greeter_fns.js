var greetings = [
  'HI GREG',
  'THIS IS DUSTIN',
  'YOUR CODE OVERLORD',
  '01010101'
]

sendTxes = async (web3, greeter, proxyGreeter) => {
  var greeterContractAddr = greeter
  var proxyGreeterAddr = proxyGreeter
  var greeterContract = web3.eth
    .contract([
      {
        constant: false,
        inputs: [{ name: '_greeting', type: 'string' }],
        name: 'setGreeting',
        outputs: [],
        payable: false,
        stateMutability: 'nonpayable',
        type: 'function'
      },
      {
        constant: true,
        inputs: [],
        name: 'greet',
        outputs: [{ name: '', type: 'string' }],
        payable: false,
        stateMutability: 'view',
        type: 'function'
      },
      {
        constant: true,
        inputs: [],
        name: 'greeting',
        outputs: [{ name: '', type: 'string' }],
        payable: false,
        stateMutability: 'view',
        type: 'function'
      },
      {
        inputs: [{ name: '_greeting', type: 'string' }],
        payable: false,
        stateMutability: 'nonpayable',
        type: 'constructor'
      }
    ])
    .at(greeterContractAddr)
  var proxygreeterContract = web3.eth
    .contract([
      {
        constant: false,
        inputs: [{ name: '_greeting', type: 'string' }],
        name: 'proxySetGreeting',
        outputs: [],
        payable: false,
        stateMutability: 'nonpayable',
        type: 'function'
      },
      {
        inputs: [{ name: '_address', type: 'address' }],
        payable: false,
        stateMutability: 'nonpayable',
        type: 'constructor'
      }
    ])
    .at(proxyGreeterAddr)
  for (i = 0; i < greetings.length; i++) {
    console.log(greetings[i])
    var res = await proxygreeterContract.proxySetGreeting.sendTransaction(
      greetings[i],
      { from: '0x887495999b72694811da9c2ee0a34de4c003332b', gas: 3000000 }
    )
    console.log(res)
    await sleep(10000)
    console.log(greeterContract.greet())
  }
}

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms))
}

module.exports = { sendTxes: sendTxes }
