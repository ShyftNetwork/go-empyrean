#!/bin/sh
if [ -z "${DBENV}" ]; then
  ./build/bin/geth sh
else
  /bin/geth --config ./shyft-config/docker-config.toml --ws --nat=any --mine --minerthreads 4
fi
