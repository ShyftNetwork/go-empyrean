#!/bin/sh
if [ -z "${DBENV}" ]; then
  ./build/bin/geth --config ./shyft-config/config.toml --disablepg --ws --nat=any --mine --minerthreads 4 --unlock "0xf3ba871253bd769b74aff013f52dbaff174bf19a" --password ./shyft-config/unlockPasswords.txt
else
  /bin/geth --config ./shyft-config/docker-config.toml --disablepg --ws --nat=any --mine --minerthreads 4 --unlock "0xf3ba871253bd769b74aff013f52dbaff174bf19a" --password ./shyft-config/unlockPasswords.txt
fi
