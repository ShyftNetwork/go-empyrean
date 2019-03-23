#!/bin/bash

if [[ $# -lt 1 ]]; then
    echo
    echo Shyft-Geth: No flags detected, see help:
    echo
    echo "   --setup:               Setups postgres and the shyft chain db."
    echo "   --start:               Starts geth."
    echo "   --reset:               Drops postgress and chain db, and reinstantiates both."
    echo "   --js [filename]:       Executes web3 calls with a passed file name. If the file name is sendTransactions.js, $ ./shyft-geth.sh --js sendTransactions"
    echo
    exit 1
fi

illegalCommands=()
while [[ $# -gt 0 ]]
do
key="$1"
case $key in
    --setup)
    sh ./shyft-config/shyft-cli/setup.sh
    shift # past argument
    ;;
    --setup2)
    sh ./shyft-config/shyft-cli/setup2.sh
    shift # past argument
    ;;
    --start)
    sh ./shyft-config/shyft-cli/startShyftGeth.sh
    shift # past argument
    ;;
    --start2)
    sh ./shyft-config/shyft-cli/startShyftGeth2.sh
    shift # past argument
    ;;
    --js)
    sh ./shyft-config/shyft-cli/runJs.sh ./shyft-config/shyft-cli/web3/$2.js
    shift # past argument
    shift # past argument
    ;;
    --reset)
    sh ./shyft-config/shyft-cli/resetShyftGeth.sh
    shift # past argument
    ;;
    --reset2)
    sh ./shyft-config/shyft-cli/resetShyftGeth2.sh
    shift # past argument
    ;;
    *)    # unknown option
    illegalCommands+=("$1") # save it in an array for later
    shift # past argument
    ;;
esac
done

if [[ "${#illegalCommands[@]}" -gt "0" ]]; then
    echo Shyft-Geth: The following commands are not supported: "${illegalCommands[*]}"
fi

