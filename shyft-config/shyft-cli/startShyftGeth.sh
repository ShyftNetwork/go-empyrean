#!/bin/sh
if [ -z "${DBENV}" ]; then
  ./build/bin/geth --config ./shyft-config/config.toml --ws --nat=any --mine --minerthreads 4 --disablepg --unlock "0xef0D57593513aAE0F357d7211870Bf6b02cAF6E7" --password ./shyft-config/unlockPasswords.txt
else
  /bin/geth --config ./shyft-config/docker-config.toml --ws --nat=any --mine --minerthreads 4 --disablepg --unlock "0xef0D57593513aAE0F357d7211870Bf6b02cAF6E7" --password ./shyft-config/unlockPasswords.txt
fi
