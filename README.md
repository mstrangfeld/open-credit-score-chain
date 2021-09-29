# Open-Credit-Score-Chain

The Open-Credit-Score-Chain (OCSC) is a proof of concept for an open and decentralized credit scoring system.
Users can see each entry of their credit score records and can choose what they want to share with whom.
The data to calculate the score from is stored encrypted on a blockchain, thus only the user can read the content.
If a company wants to know the credit score of a person, the person can provide a zero-knowledge proof (zk-SNARK) of their score.
This way no unnecessary information is leaked, the user has full transparancy of their records and companies can nevertheless trust the scoring system.

> :warning: **This software is just a proof-of-concept**
> It was developed during a 24h hackathon. Especially the cryptographic modules should not be used in any real-world application!

## Concept

### Why Ethereum?
We don't want to use more energy than necessary for storing the records and calculating the credit score of a person.
Although Ethereum currently uses the proof-of-work algorithm that wastes a lot of energy, it will soon switch to a more resource-friendly proof-of-stake algorithm.
On top of that the Ethereum ecosystem is very mature which makes it a perfect candidate for this task.

### Future
At the moment we are using a dedicated backend to talk to the Ethereum smart contracts and to create the zk-SNARKs.
This is not optimal because users should never have to trust an other entity other than themselves.
In the future the whole interaction with the blockchain should eventually move to the browser where the user has full control over the environment.
However, zk-SNARKs can be computationally very expensive. For this reason there could be a client-side verifiable homomorphic re-encryption of the data, so the server can create a zk-SNARK without ever seeing the un-encrypted data.

## Requirements

When you are running on *NixOS* or you have *Nix* installed you can just run:

```sh
$ nix develop # In case you have nix flakes enabled
$ nix-shell # Otherwise
```

Otherwise you will need the following software installed:
+ ganache (Local ethereum network)
+ go 1.16
+ solc
+ go-ethereum

## Getting started

### Starting the local Ethereum blockchain

```sh
$ ganache-cli -d
```

### Compiling and deploying the smart contracts to the chain

```sh
$ make build-contracts
$ make deploy-contracts
```

### Generate Go bindings for the backend

```sh
$ make gen-go-bindings
```

## License

The Open-Credit-Score-Chain project is licensed under the Apache License, Version 2.0. See LICENSE for details.
