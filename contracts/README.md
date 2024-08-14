## Simple NFT Collection factory contracts

## Features

1. FarawayCollectionFactory: deploy a new FarawayCollection via `createCollection`, emits `CollectionCreated(address collection, string name, string symbol)` event
2. FarawayCollection: Mintable, burnable ERC721 w/ uri storage, to mint a token with given uri use `mint`. Emits `TokenMinted(address collection, address recipient, uint256 tokenId, string tokenUri)` event

## Usage

### Build

```shell
$ forge build
```

### Deploy

```shell
$ forge create src/FarawayCollectionFactory.sol:FarawayCollectionFactory --rpc-url <your_rpc_url> --private-key <your_private_key>
```
