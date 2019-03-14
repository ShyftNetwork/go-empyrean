#!/bin/sh
if [ -z "${DBENV}" ]; then
  ./build/bin/geth --config ./shyft-config/config2.toml --ws --nat=any --mine --minerthreads 4 --disablepg --unlock "0x181dE8FB0a836A8e68a4c68bF758B2E0eE2a8145" --password ./shyft-config/unlockPasswords.txt
else
  /bin/geth --config ./shyft-config/docker-config2.toml --ws --nat=any --mine --minerthreads 4 --disablepg --unlock "0x181dE8FB0a836A8e68a4c68bF758B2E0eE2a8145" --password ./shyft-config/unlockPasswords.txt
fi
