#!/bin/sh
if [ -z "${DBENV}" ]; then
  ./build/bin/geth --config ./shyft-config/config.toml --ws --nat=any --mine --minerthreads 4 --targetgaslimit 8000000 --disablepg
else
  /bin/geth --config ./shyft-config/config.toml --ws --nat=any --mine --minerthreads 4 --targetgaslimit 8000000 --disablepg
fi
