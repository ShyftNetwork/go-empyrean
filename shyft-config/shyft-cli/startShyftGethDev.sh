#!/bin/sh
if [ -z "${DBENV}" ]; then
  ./build/bin/geth --config ./shyft-config/config.toml --ws --nat=any --mine --minerthreads 4 --targetgaslimit 8000000 --rpc --rpcapi="personal,eth,network,net,web3" --rpcport 8545 --rpcaddr "localhost"
else
  /bin/geth --config ./shyft-config/config.toml --ws --nat=any --mine --minerthreads 4 --targetgaslimit 8000000 --rpc --rpcapi="personal,eth,network,net,web3" --rpcport 8545 --rpcaddr "localhost"
fi
