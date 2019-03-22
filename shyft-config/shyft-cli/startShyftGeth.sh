#!/bin/sh
if [ -z "${DBENV}" ]; then
  ./build/bin/geth --config ./shyft-config/config.toml --disablepg --ws --nat=any --mine --minerthreads 4 --unlock "0xc7b55a90e22fa04534170d06e70843ca4f2c5164" --password shyft-config/unlockPasswords.txt
else
  /bin/geth --config ./shyft-config/docker-config.toml --ws --nat=any --mine --minerthreads 4
fi
