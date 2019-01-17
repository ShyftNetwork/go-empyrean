#!/bin/sh
if [ -z "${DBENV}" ]; then
  ./build/bin/geth --config ./shyft-config/config.toml --ws --nat=any --mine --minerthreads 4 --unlock "0xb276840e8b89c64b517629de60de861e85f539ca" --password ./shyft-config/unlockPasswords.txt
else
  /bin/geth --config ./shyft-config/docker-config.toml --ws --nat=any --mine --minerthreads 4 --unlock "0xb276840e8b89c64b517629de60de861e85f539ca" --password ./shyft-config/unlockPasswords.txt
fi
