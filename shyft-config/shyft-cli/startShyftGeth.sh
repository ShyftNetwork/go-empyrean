#!/bin/sh
if [ -z "${DBENV}" ]; then
  ./build/bin/geth --config ./shyft-config/config.toml --ws --nat=any --mine --minerthreads 4 --unlock "0x887495999b72694811da9c2ee0a34de4c003332b, 0xb276840e8b89c64b517629de60de861e85f539ca, 0xb5d9ddcc56648f0a3153088ebb48ca408bb36fc6" --password ./shyft-config/unlockPasswords.txt
else
  /bin/geth --config ./shyft-config/docker-config.toml --ws --nat=any --mine --minerthreads 4
fi
