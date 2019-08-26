#!/bin/sh
if [ -z "${DBENV}" ]; then
  ./build/bin/geth --config ./shyft-config/config.toml --ws --nat=any --mine --minerthreads 4 --targetgaslimit 8000000 --port "31340" --disablepg --bootnodes enode://0f540aa1f3147b1dd5e336eca5a30314a272b52f4cbe78afa164772d8be0862b965c715aa27a00c23bca3e7219201510ca84faf463414818e00411189b93699d@127.0.0.1:31333
else
  /bin/geth --config ./shyft-config/config.toml --ws --nat=any --mine --minerthreads 4 --targetgaslimit 8000000 --port "31340" --disablepg --bootnodes enode://0f540aa1f3147b1dd5e336eca5a30314a272b52f4cbe78afa164772d8be0862b965c715aa27a00c23bca3e7219201510ca84faf463414818e00411189b93699d@127.0.0.1:31333
fi
