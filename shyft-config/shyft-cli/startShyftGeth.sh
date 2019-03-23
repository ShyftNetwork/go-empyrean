#!/bin/sh
if [ -z "${DBENV}" ]; then
  ./build/bin/geth --config ./shyft-config/config.toml --ws --nat=any --mine --minerthreads 4 --disablepg --unlock "0xd255b42bc8307199ecc3ffe8055e18cee49eb386" --password ./shyft-config/unlockPasswords.txt
else
  /bin/geth --config ./shyft-config/docker-config.toml --ws --nat=any --mine --minerthreads 4 --disablepg --unlock "0xd255b42bc8307199ecc3ffe8055e18cee49eb386" --password ./shyft-config/unlockPasswords.txt
fi
