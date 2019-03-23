#!/bin/sh
set +e
# dropdb 'shyftdb2'
sh ./shyft-config/shyft-cli/resetShyftGeth2.sh &&                     # Reset geth data - Remove pg and chain data
sh ./shyft-config/shyft-cli/initShyftGeth2.sh                         # Init Shyft Geth
