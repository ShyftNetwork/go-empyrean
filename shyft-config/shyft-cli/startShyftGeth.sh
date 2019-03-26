#!/bin/sh
if [ -z "${DBENV}" ]; then
  ./build/bin/geth --config ./shyft-config/config.toml --ws --rpc --nat=any --mine --minerthreads 4
else
  /bin/geth --config ./shyft-config/docker-config.toml --ws --rpc --nat=any --mine --minerthreads 4
fi
