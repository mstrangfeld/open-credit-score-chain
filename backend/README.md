# Backend

The backend implements the necessary logic to create encrypted records, decrypt encrypted records and create/verify zk-SNARKs.

## Run the backend

For development purposes you can run the backend using the `go run` command.
```sh
$ go run .
```

## Environment Variables

Following environment variables need to be set:

| Name | Description |
| ---- | ----------- |
| PORT | The port the service needs to listen |
| PRIVATE_KEY | The private key of the authority to be able to create the authority smart contract and new user databases |
| ETHEREUM_CONNECTION | The URL to connect to the Ethereum network |

If you are using *direnv* all those variables will be set automatically to the local testing defaults.

## API

### `GET /api/db/create`

Create a new database for an user

### `GET /api/record/create`

Create a new record

### `GET /api/record/invalidate`

Invalidate a record

### `GET /api/zk/create`

Create a zk-SNARK

### `GET /api/zk/verify`

Verify a zk-SNARK

## Run tests

```sh
$ go test ./...
```
