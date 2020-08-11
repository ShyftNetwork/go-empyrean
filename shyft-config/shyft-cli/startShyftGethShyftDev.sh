#!/bin/sh
if [ -z "${DBENV}" ]; then
  ./build/bin/geth --allow-insecure-unlock --config ./shyft-config/config.toml --ws --nat=any --mine --minerthreads 4 --targetgaslimit 8000000 --rpc --rpcapi="personal,eth,network,net,web3" --rpcport 8545 --rpcaddr "localhost" --unlock "0x43EC6d0942f7fAeF069F7F63D0384a27f529B062,0x9e602164C5826ebb5A6B68E4AFD9Cd466043dc4A,0x5Bd738164C61FB50eb12E227846CbaeF2dE965Aa,0xC04eE4131895F1d0C294D508AF65D94060AA42BB,0x07D899C4aC0c1725C35C5f816e60273B33a964F7" --password ./shyft-config/unlockPasswords.txt
else
  /bin/geth --allow-insecure-unlock --config ./shyft-config/config.toml --ws --nat=any --mine --minerthreads 4 --targetgaslimit 8000000 --rpc --rpcapi="personal,eth,network,net,web3" --rpcport 8545 --rpcaddr "localhost" --unlock "0x43EC6d0942f7fAeF069F7F63D0384a27f529B062,0x9e602164C5826ebb5A6B68E4AFD9Cd466043dc4A,0x5Bd738164C61FB50eb12E227846CbaeF2dE965Aa,0xC04eE4131895F1d0C294D508AF65D94060AA42BB,0x07D899C4aC0c1725C35C5f816e60273B33a964F7" --password ./shyft-config/unlockPasswords.txt
fi



#./build/bin/geth --allow-insecure-unlock --ws --nat=any --mine --minerthreads 4 --targetgaslimit 8000000 --rpc --rpcapi="personal,eth,network,net,web3" --rpcport 8545 --rpcaddr "localhost" --unlock "0x43EC6d0942f7fAeF069F7F63D0384a27f529B062,0x9e602164C5826ebb5A6B68E4AFD9Cd466043dc4A,0x5Bd738164C61FB50eb12E227846CbaeF2dE965Aa,0xC04eE4131895F1d0C294D508AF65D94060AA42BB,0x07D899C4aC0c1725C35C5f816e60273B33a964F7" --password ./shyft-config/unlockPasswords.txt --datadir shyftData