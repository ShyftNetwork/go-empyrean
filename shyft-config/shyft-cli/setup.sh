#!/bin/sh
set +e
dropdb 'shyftdb'
sh ./shyft-config/shyft-cli/resetShyftGeth.sh &&                     # Reset geth data - Remove pg and chain data
sh ./shyft-config/shyft-cli/initShyftGeth.sh                         # Init Shyft Geth
