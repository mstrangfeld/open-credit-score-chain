# Open-Credit-Score-Chain

The Open-Credit-Score-Chain (OCSC) is a proof of concept for an open and decentralized credit scoring system.
Users can see each entry of their credit score records and can choose what they want to share with whom.
The data to calculate the score from is stored encrypted on a blockchain, thus only the user can read the content.
If a company wants to know the credit score of a person, the person can provide a zero-knowledge proof (zk-SNARK) of their score.
This way no unnecessary information is leaked, the user has full transparancy of their records and companies can nevertheless trust the scoring system.

> :warning: **This software is just a proof-of-concept**
> It was developed during a 24h hackathon. Especially the cryptographic modules should not be used in any real-world application.

## Concept

### Why Blockchain?
For a consistent and trusted record of entries 

### Why Ethereum?

### Credit Score DB Smart Contract
The main part of the application is the credit score db, which is implemented as a smart contract.
In this case it is an Ethereum smart contract but it can be migrated to other blockchains if necessary.
Every user that wants to be able 

### What is necessary for a real-world application of this project?


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


## License

The Open-Credit-Score-Chain project is licensed under the Apache License, Version 2.0. See LICENSE for details.
