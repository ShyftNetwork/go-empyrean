
## Go Empyrean

[![API Reference](https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667)](https://godoc.org/github.com/ShyftNetwork/go-empyrean)
[![Go Report Card](https://goreportcard.com/badge/github.com/ShyftNetwork/go-empyrean)](https://goreportcard.com/report/github.com/ShyftNetwork/go-empyrean)
[![Build Status](https://travis-ci.org/ShyftNetwork/go-empyrean.svg?branch=development)](https://travis-ci.org/ShyftNetwork/shyft_go-ethereum)
[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/ShyftNetwork/go-empyrean?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

go-empyrean is based on a fork of go-ethereum. Much of the functionality and process for starting go-empyrean is the same as that for a regular ethereum node - as reflected in the notes detailed under the heading Go Ethereum below. Documentation for changes added by Shyft is detailed below.

## Shyft Documentation

https://shyftnetwork.github.io/go-empyrean/#setup

## Install 

To get a go-empyrean node ready please ensure that you have the proper installations listed below.

#### Dependencies
    
 - go 1.10
 - postgres 10
    
To install go please review the installation docs [here](https://golang.org/doc/install), but ensure you download version 1.10. If you would like to install go with a script please check out this repo [here](https://github.com/canha/golang-tools-install-script).
    
To install postgres please review the installation docs [here](https://www.postgresql.org/docs/10/static/tutorial-install.html).

#### Govendor and Packages/Dependencies

> Download Go Vendor

```shell
go get -u github.com/kardianos/govendor
```

> To run govendor globally, have this in your bash_profile file:

```shell
export GOPATH=$HOME/go
export PATH=$PATH:$HOME/go/bin
```

> go_empyrean will need to be cloned to this directory:

```shell
$GOPATH/src/github.com/ShyftNetwork/
```

Geth uses govendor to manage packages/dependencies: [Go Vendor](https://github.com/kardianos/govendor)

This has some more information: [Ethereum Wiki](https://github.com/ShyftNetwork/go-empyrean/wiki/Developers'-Guide)

To add a new dependency, run govendor fetch <import-path> , and commit the changes to git. Then the deps will be accessible on other machines that pull from git.

>GOPATH is not strictly necessary however, for govendor it is much easier to use gopath as go will look for binaries in this directory ($GOPATH/bin). To set up GOPATH, read the govendor section.

## Usage 

Once the dependencies and prerequisites are installed, run

    make geth

or, to build the full suite of utilities:

    make all
    
Upon completion, run the below command to setup the postgres instance and shyft chainDB    
    
	./shyft-config/shyft-geth.sh --setup

#### Running Go-empyrean

	./shyft-config/shyft-geth.sh --start

At this point you should see geth running in the terminal and if you opened your postgres instance you should see data being populated into the tables.

#### SHH/Whisper

The Shyft go-empyrean node, unlike go-ethereum starts the SHH whisper client by default. This is to facilitate broadcast messaging from the Shyft bridge to each of the mining nodes.

To disable the whisper client a startup flag `--disablewhisper` is provided, which must be passed into the command line when starting up geth.


	geth --disablewhisper


To overwrite the default whisper variables, the following flags are also provided:

    --shh.maxmessagesize - sets the maximum message size fir the whisper client (integer) -(default: 1048576)  --shh.maxmessagesize=128
    --shh.pow - the minimum POW accepted for processing whisper messages (float64 - default: 0.2) --shh.pow=0.3
    --shh.restrict-light - restrictions connections between two whisper light clients (boolean - default: true) --shh.restrict-light

To authenticate whisper messages a call is made to a smart contract that has a predetermined address on the blockchain.
Upon starting up a geth node if a user wishes to use this functionality they should ensure 
that the WhisperSignersContract variable in config.toml contains the contract address for authentication of Whisper Signers.

The authentication of WhisperSigner's broadcast messages relies on automatically generated go contract bindings using the 
the abigen cmd line utility. Should the contract be changed or modified these bindings will need to be regenerated.
Steps for regenerating are as follows:

```$xslt
1. Generate the abi for the subject contract and save it at ./generated_bindings/contract_abis/whispersigner_abi.json.

2. Run the following command to regenerate the contract bindings:

./build/bin/abigen  --sol ./shyft-config/shyft-cli/web3/validSignersDeploy/ValidSigners.sol  --pkg shyft_contracts --out generated_bindings/whisper_signer_binding.go

```

It should be noted that the authentication currently relies on a smart contract boolean returning function [isValidSigner(bool)], 
that for a given signature address returns true if the contract or contract owner has a public key matching the signature.

## Running with Docker

Two sets of Docker Images are available for go-empyrean, the Postgresql Database, and the Shyft Block Explorer, which can be used for local development and testnet connection. The development settings are included in docker-compose.yml, the testnet settings are included in docker-compose.production.yml. To launch these containers you will need to have docker-compose installed on your computer. Installation instructions for docker-compose are available [here](https://docs.docker.com/install/).

**To build the images for the first time please run the following command:**

	./shyft-config/shyft-geth.sh --setup 

	docker-compose up --build

Running the above command will build Shyft Block Explorer on port `:3000` which is being served by Shyft Block Explorer API on port `:8080`. 
These are images being pulled from docker hub which are publicly available.

#### Docker PostgreSQL

From your local machine you can view the database by connecting to the database in the container at 
	
	127.0.0.1:8001

Use the following credentials: 
 
>``User: 'postgres'``
 
>``Password: 'docker'``
  
>``Database: 'shyftdb'``
 
#### Docker Block Explorer API
 
To access the shyftBlockExplorer open a browser and visit 

	http://localhost:3000

To rebuild any one of the services- issue the following commands:

Services:

   - ShyftGeth
   - Postgres Instance
   - Shyft Explorer API
   - Shyft Example Explorer UI

	docker-compose up -d --no-deps --build <docker compose file service name> 

> Shyft Block Explorer Api:

	docker-compose up -d --no-deps --build shyft_block_api

> Shyft Block Explorer UI:

	docker-compose up -d --no-deps --build shyft_block_ui

> Removing postgres data and chain data

	./shyft-config/shyft-cli/resetShyftGeth

Blockchain data is persisted to **``./ethash/.ethash and ./shyftData__``**. If you would like to reset the test blockchain you will need to delete the **``__./ethash ./shyftData & ./privatenet__``** directories.

NB: The Shyft Geth docker image size is 1+ GB so make sure you have adequate space on your disk drive/

#### CLI

Run `./shyft-config/shyft-geth.sh` with one of the following flags:

| Command    	| Description |
|:-------------:|-------------|
| **`--setup`** | Setup postgres and the shyft chain db |
| `--start` 	| Starts geth. |
| `--reset` 	| Drops postgres and chain DB, and reinstantiates both. |
| `--js [web3 filename]` | Executes web3 calls with a passed file name. If the file name is `sendTransactions.js`, `./shyft-geth --js sendTransactions`. |

                                                                                                                                                                                                        
## Executables

The go-ethereum project comes with several wrappers/executables found in the `cmd` directory.

| Command    | Description |
|:----------:|-------------|
| **`geth`** | Our main Ethereum CLI client. It is the entry point into the Ethereum network (main-, test- or private net), capable of running as a full node (default), archive node (retaining all historical state) or a light node (retrieving data live). It can be used by other processes as a gateway into the Ethereum network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. `geth --help` and the [CLI Wiki page](https://github.com/ShyftNetwork/go-empyrean/wiki/Command-Line-Options) for command line options. |
| `abigen` | Source code generator to convert Ethereum contract definitions into easy to use, compile-time type-safe Go packages. It operates on plain [Ethereum contract ABIs](https://github.com/ethereum/wiki/wiki/Ethereum-Contract-ABI) with expanded functionality if the contract bytecode is also available. However it also accepts Solidity source files, making development much more streamlined. Please see our [Native DApps](https://github.com/ShyftNetwork/go-empyrean/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts) wiki page for details. |
| `bootnode` | Stripped down version of our Ethereum client implementation that only takes part in the network node discovery protocol, but does not run any of the higher level application protocols. It can be used as a lightweight bootstrap node to aid in finding peers in private networks. |
| `evm` | Developer utility version of the EVM (Ethereum Virtual Machine) that is capable of running bytecode snippets within a configurable environment and execution mode. Its purpose is to allow isolated, fine-grained debugging of EVM opcodes (e.g. `evm --code 60ff60ff --debug`). |
| `gethrpctest` | Developer utility tool to support our [ethereum/rpc-test](https://github.com/ethereum/rpc-tests) test suite which validates baseline conformity to the [Ethereum JSON RPC](https://github.com/ethereum/wiki/wiki/JSON-RPC) specs. Please see the [test suite's readme](https://github.com/ethereum/rpc-tests/blob/master/README.md) for details. |
| `rlpdump` | Developer utility tool to convert binary RLP ([Recursive Length Prefix](https://github.com/ethereum/wiki/wiki/RLP)) dumps (data encoding used by the Ethereum protocol both network as well as consensus wise) to user friendlier hierarchical representation (e.g. `rlpdump --hex CE0183FFFFFFC4C304050583616263`). |
| `swarm`    | Swarm daemon and tools. This is the entrypoint for the Swarm network. `swarm --help` for command line options and subcommands. See [Swarm README](https://github.com/ShyftNetwork/go-empyrean/tree/master/swarm) for more information. |
| `puppeth`    | a CLI wizard that aids in creating a new Ethereum network. |


### Configuration

As an alternative to passing the numerous flags to the `geth` binary, you can also pass a configuration file via:

```
$ geth --config /path/to/your_config.toml
```

To get an idea how the file should look like you can use the `dumpconfig` subcommand to export your existing configuration:

```
$ geth --your-favourite-flags dumpconfig
```

_Note: This works only with geth v1.6.0 and above._


### Programatically interfacing Geth nodes

As a developer, sooner rather than later you'll want to start interacting with Geth and the Ethereum
network via your own programs and not manually through the console. To aid this, Geth has built-in
support for a JSON-RPC based APIs ([standard APIs](https://github.com/ethereum/wiki/wiki/JSON-RPC) and
[Geth specific APIs](https://github.com/empyrean/go-ethereum/wiki/Management-APIs)). These can be
exposed via HTTP, WebSockets and IPC (unix sockets on unix based platforms, and named pipes on Windows).

The IPC interface is enabled by default and exposes all the APIs supported by Geth, whereas the HTTP
and WS interfaces need to manually be enabled and only expose a subset of APIs due to security reasons.
These can be turned on/off and configured as you'd expect.

HTTP based JSON-RPC API options:

- `--rpc` Enable the HTTP-RPC server
- `--rpcaddr` HTTP-RPC server listening interface (default: "localhost")
- `--rpcport` HTTP-RPC server listening port (default: 8545)
- `--rpcapi` API's offered over the HTTP-RPC interface (default: "eth,net,web3")
- `--rpccorsdomain` Comma separated list of domains from which to accept cross origin requests (browser enforced)
- `--ws` Enable the WS-RPC server
- `--wsaddr` WS-RPC server listening interface (default: "localhost")
- `--wsport` WS-RPC server listening port (default: 8546)
- `--wsapi` API's offered over the WS-RPC interface (default: "eth,net,web3")
- `--wsorigins` Origins from which to accept websockets requests
- `--ipcdisable` Disable the IPC-RPC server
- `--ipcapi` API's offered over the IPC-RPC interface (default: "admin,debug,eth,miner,net,personal,shh,txpool,web3")
- `--ipcpath` Filename for IPC socket/pipe within the datadir (explicit paths escape it)

You'll need to use your own programming environments' capabilities (libraries, tools, etc) to connect
via HTTP, WS or IPC to a Geth node configured with the above flags and you'll need to speak [JSON-RPC](https://www.jsonrpc.org/specification)
on all transports. You can reuse the same connection for multiple requests!

**Note: Please understand the security implications of opening up an HTTP/WS based transport before
doing so! Hackers on the internet are actively trying to subvert Ethereum nodes with exposed APIs!
Further, all browser tabs can access locally running webservers, so malicious webpages could try to
subvert locally available APIs!**

#### Creating the rendezvous point

With all nodes that you want to run initialized to the desired genesis state, you'll need to start a
bootstrap node that others can use to find each other in your
network and/or over the internet. The
clean way is to configure and run a dedicated bootnode:

```
$ bootnode --genkey=boot.key
$ bootnode --nodekey=boot.key
```

With the bootnode online, it will display an [`enode` URL](https://github.com/ethereum/wiki/wiki/enode-url-format)
that other nodes can use to connect to it and exchange peer information. Make sure to replace the
displayed IP address information (most probably `[::]`) with your externally accessible IP to get the
actual `enode` URL.

_Note: You could also use a full fledged Geth node as a bootnode, but it's the less recommended way._

#### Starting up your member nodes

With the bootnode operational and externally reachable (you can try `telnet <ip> <port>` to ensure
it's indeed reachable), start every subsequent Geth node pointed to the bootnode for peer discovery
via the `--bootnodes` flag. It will probably also be desirable to keep the data directory of your
private network separated, so do also specify a custom `--datadir` flag.

```
$ geth --datadir=path/to/custom/data/folder --bootnodes=<bootnode-enode-url-from-above>
```

_Note: Since your network will be completely cut off from the main and test networks, you'll also
need to configure a miner to process transactions and create new blocks for you._

#### Running a private miner

Mining on the public Ethereum network is a complex task as it's only feasible using GPUs, requiring
an OpenCL or CUDA enabled `ethminer` instance. For information on such a setup, please consult the
[EtherMining subreddit](https://www.reddit.com/r/EtherMining/) and the [Genoil miner](https://github.com/Genoil/cpp-ethereum)
repository.

In a private network setting however, a single CPU miner instance is more than enough for practical
purposes as it can produce a stable stream of blocks at the correct intervals without needing heavy
resources (consider running on a single thread, no need for multiple ones either). To start a Geth
instance for mining, run it with all your usual flags, extended by:

```
$ geth <usual-flags> --mine --minerthreads=1 --etherbase=0x0000000000000000000000000000000000000000
```

Which will start mining blocks and transactions on a single CPU thread, crediting all proceedings to
the account specified by `--etherbase`. You can further tune the mining by changing the default gas
limit blocks converge to (`--targetgaslimit`) and the price transactions are accepted at (`--gasprice`).

## Contribution

If you'd like to contribute to go-empyrean, please fork, fix, commit and send a pull request against the `development` branch.

We have a list of issues available on github.

For general communication, we communicate on [our gitter channel](https://gitter.im/ShyftNetwork/go-empyrean).

Please make sure your contributions adhere to our coding guidelines:

- Code must adhere to the official Go [formatting](https://golang.org/doc/effective_go.html#formatting) guidelines (i.e. uses [gofmt](https://golang.org/cmd/gofmt/)).
- Code must be documented adhering to the official Go [commentary](https://golang.org/doc/effective_go.html#commentary) guidelines.
- Pull requests need to be based on and opened against the `development` branch.


## License

The go-ethereum library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html), also
included in our repository in the `COPYING.LESSER` file.

The go-ethereum binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also included
in our repository in the `COPYING` file.
