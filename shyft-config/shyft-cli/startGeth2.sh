#!/bin/sh
if [ -z "${DBENV}" ]; then
  ./build/bin/geth --config ./shyft-config/config2.toml --disablepg --ws --nat=any --mine --minerthreads 4 --unlock "0xc32eceb11a1c719ffd9b8af792f41eadf46b8ffa" --password shyft-config/unlockPasswords.txt
else
  /bin/geth --config ./shyft-config/docker-config2.toml --ws --nat=any --mine --minerthreads 4
fi