#!/bin/sh
if [ -z "${DBENV}" ]; then
  ./build/bin/geth --config ./shyft-config/config3.toml --disablepg --ws --nat=any --mine --minerthreads 4 --unlock "0xe55f57cac96fd00da6c32c353128c836a58786b0" --password shyft-config/unlockPasswords.txt
else
  /bin/geth --config ./shyft-config/docker-config3.toml --ws --nat=any --mine --minerthreads 4
fi