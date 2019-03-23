#!/bin/sh
if [ -z "${DBENV}" ]; then
  ./build/bin/geth --identity "ShyftTestnetNode" --keystore ./ --datadir "./shyftData2" init ./shyft-config/ShyftNetwork.json
else
  if [ -d /go-empyrean/shyftData/geth/chaindata ]; then
    echo "Skipping Genesis Initialization as already completed"
    :
  else
    echo "Initializing Custom Genesis Block"
    /bin/geth --identity "ShyftTestnetNode" --datadir "./shyftData2" init ./shyft-config/ShyftNetwork.json
  fi
fi
