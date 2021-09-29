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

### Starting the backend and deploying the authority smart contract

```sh
$ make gen-go-bindings
$ cd backend && go run .
```

### Helpful resources
+ [Solidity Documentation](https://docs.soliditylang.org/en/latest/)
+ [Introduction to zero-knowledge-proofs](https://zkp.science/)
+ [Vitalik Buterin blog post about zk-SNARKs](https://vitalik.ca/general/2021/01/26/snarks.html)
+ [Golang gnark Documentation](https://docs.gnark.consensys.net/en/latest/)
+ [Go Ethereum Book](https://goethereumbook.org/)

## TODOs

- [x] Create the basic smart contracts (see `contracts/`)
- [x] Deploy the contracts to the blockchain (see `backend/main.go`)
- [x] Encrypt and decrypt records for an Ethereum user (see `backend/records/`)
- [x] Create zk-SNARKs of the credit score (see `backend/zkproof/`)
- [x] Verify zk-SNARKs with the public hashes available (see `backend/zkproof/`)
- [ ] Implement the REST API in the backend
- [ ] Create frontend to interact with the backend
- [ ] Access right checking in the smart contracts
- [ ] Authentication to the backend
- [ ] Secure multiparty implementation of the zk-SNARK to be able to create zk-Proofs without having the server knowing about the plain-text records
- [ ] Modular zk-SNARKs for alternative algorithms (at the moment we just create a sum of all inputs)

## License

The Open-Credit-Score-Chain project is licensed under the Apache License, Version 2.0. See LICENSE for details.
