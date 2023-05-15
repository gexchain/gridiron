#!/bin/bash

KEY="dev0"
CHAINID="gridiron_9000-1"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t gridiron-datadir.XXXXX)

echo "create and add new keys"
./gridirond keys add $KEY --home $DATA_DIR --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init Gridiron with moniker=$MONIKER and chain-id=$CHAINID"
./gridirond init $MONIKER --chain-id $CHAINID --home $DATA_DIR
echo "prepare genesis: Allocate genesis accounts"
./gridirond add-genesis-account \
"$(./gridirond keys show $KEY -a --home $DATA_DIR --keyring-backend test)" 1000000000000000000afury,1000000000000000000stake \
--home $DATA_DIR --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./gridirond gentx $KEY 1000000000000000000stake --keyring-backend test --home $DATA_DIR --keyring-backend test --chain-id $CHAINID
echo "prepare genesis: Collect genesis tx"
./gridirond collect-gentxs --home $DATA_DIR
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./gridirond validate-genesis --home $DATA_DIR

echo "starting gridiron node $i in background ..."
./gridirond start --pruning=nothing --rpc.unsafe \
--keyring-backend test --home $DATA_DIR \
>$DATA_DIR/node.log 2>&1 & disown

echo "started gridiron node"
tail -f /dev/null