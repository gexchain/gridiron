<!--
order: 6
-->

# Snapshots & Archive Nodes

Quickly sync your node with Gridiron using a snapshot or serve queries for prev versions using archive nodes {synopsis}

## List of Snapshots and Archives

Below is a list of publicly available snapshots that you can use to sync with the Gridiron mainnet and
archived [9001-1 mainnet](https://github.com/gridchain/mainnet/tree/main/gridiron_9001-1):

<!-- markdown-link-check-disable -->

### Snapshots

| Name        | URL                                                                     |
| -------------|------------------------------------------------------------------------ |
| `Staketab`   | [github.com/staketab/nginx-cosmos-snap](https://github.com/staketab/nginx-cosmos-snap/blob/main/docs/gridiron.md) |
| `Polkachu`   | [polkachu.com](https://www.polkachu.com/tendermint_snapshots/gridiron)                   |
| `Nodes Guru` | [snapshots.nodes.guru/gridiron_9001-2/](snapshots.nodes.guru/gridiron_9001-2/)                   |
| `Notional`   | [mainnet/pruned/gridiron_9001-2(pebbledb)](https://snapshot.notional.ventures/gridiron/) <br> [mainnet/archive/gridiron_9001-2(pebbledb)](https://snapshot.notional.ventures/gridiron-archive/) <br> [testnet/archive/gridiron_9000-4(pebbledb)](https://snapshot.notional.ventures/gridiron-testnet-archive/)                   |

### Archives
<!-- markdown-link-check-disable -->

| Name           | URL                                                                             |
| ---------------|---------------------------------------------------------------------------------|
| `Nodes Guru`   | [snapshots.nodes.guru/gridiron_9001-1](https://snapshots.nodes.guru/gridiron_9001-1/)                                    |
| `Polkachu`     | [polkachu.com/tendermint_snapshots/gridiron](https://www.polkachu.com/tendermint_snapshots/gridiron)                           |
| `Forbole`      | [bigdipper.live/gridiron_9001-1](https://s3.bigdipper.live.eu-central-1.linodeobjects.com/gridchain_9001-1.tar.lz4) |

To access snapshots and archives, follow the process below (this code snippet is to access a snapshot of the current network, `gridiron_9001-2`, from Nodes Guru):

```bash
cd $HOME/.gridirond/data
wget https://snapshots.nodes.guru/gridiron_9001-2/gridiron_9001-2-410819.tar
tar xf gridiron_9001-2-410819.tar
```

### PebbleDB

To use PebbleDB instead of GoLevelDB when using snapshots from Notional:

Build:

```bash
go mod edit -replace github.com/tendermint/tm-db=github.com/baabeetaa/tm-db@pebble
go mod tidy
go install -tags pebbledb -ldflags "-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=pebbledb" ./...
```

Download snapshot:

```bash
cd $HOME/.gridirond/
URL_SNAPSHOT="https://snapshot.notional.ventures/gridiron/data_20221024_193254.tar.gz"
wget -O - "$URL_SNAPSHOT" |tar -xzf -
```

Start:

Set `db_backend = "pebbledb"` in `config.toml` or start with `--db_backend=pebbledb`

```bash
gridirond start --db_backend=pebbledb
```

**Note**: use this [workaround](https://github.com/notional-labs/cosmosia/blob/main/docs/pebbledb.md) when upgrading a node running PebbleDB.
