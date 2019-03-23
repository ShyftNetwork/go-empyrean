#!/bin/sh
if [ -z "${DBENV}" ]; then
  ./build/bin/geth --config ./shyft-config/config2.toml --ws --nat=any --mine --minerthreads 4 --disablepg --unlock "0xeacdd684e99331da8f6b166f5b84f0b0b389bd63" --password ./shyft-config/unlockPasswords.txt
else
  /bin/geth --config ./shyft-config/docker-config2.toml --ws --nat=any --mine --minerthreads 4 --disablepg --unlock "0xeacdd684e99331da8f6b166f5b84f0b0b389bd63" --password ./shyft-config/unlockPasswords.txt
fi
