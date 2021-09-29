package main

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
	"github.com/mstrangfeld/open-credit-score-chain/backend/contracts/authority"
	"github.com/rs/zerolog/log"
)

const (
	privateKeyEnvVar         = "PRIVATE_KEY"
	ethereumConnectionEnvVar = "ETHEREUM_CONNECTION"
	portEnvVar               = "PORT"
)

var (
	ehtClient         *ethclient.Client    // The Ethereum client connection
	authorityInstance *authority.Authority // The contract instance
	privateKey        *ecdsa.PrivateKey    // The private key of the node
	thisAddress       common.Address       // The address of the node
)

// load the private key and the deriving address from the environment variable
func loadPrivateKey() {
	privHex := os.Getenv(privateKeyEnvVar)
	if privHex == "" {
		log.Fatal().Msg("Private key not set")
	}
	priv, err := crypto.HexToECDSA(privHex)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load private key")
	}
	privateKey = priv
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal().Msg("error casting public key to ECDSA")
	}
	thisAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
}

// create the Ethereum client connection
func createEthereumConnection() {
	connStr := os.Getenv(ethereumConnectionEnvVar)
	if connStr == "" {
		log.Fatal().Msg("Ethereum connection not set")
	}
	ehtc, err := ethclient.Dial(connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to Ethereum")
	}
	ehtClient = ehtc
}

// deploy the authority contract
func deployAuthority() {
	auth := createTransactionOptions(context.Background(), 300000)
	address, tx, instance, err := authority.DeployAuthority(auth, ehtClient)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to deploy contract")
	}
	log.Info().Msgf("Contract address: %s", address.Hex())
	log.Info().Msgf("Contract tx: %s", tx.Hash().Hex())
	authorityInstance = instance
}

func createTransactionOptions(ctx context.Context, gaslimit int) *bind.TransactOpts {
	nonce, err := ehtClient.PendingNonceAt(ctx, thisAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get nonce")
	}
	gasPrice, err := ehtClient.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get gas price")
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(gaslimit)
	auth.GasPrice = gasPrice

	return auth
}

func main() {
	loadPrivateKey()
	createEthereumConnection()
	deployAuthority()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")

	dbApi := api.Group("/db")
	dbApi.Post("/create", func(c *fiber.Ctx) error {
		taxID := c.FormValue("tax_id")
		userAddress := c.FormValue("address")
		addr := common.HexToAddress(userAddress)
		auth := createTransactionOptions(c.Context(), 300000)
		_, err := authorityInstance.CreateDB(auth, taxID, addr)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.SendStatus(200)
	})

	recordApi := api.Group("/record")
	recordApi.Get("/create", func(c *fiber.Ctx) error {
		return nil
	})
	recordApi.Get("/invalidate", func(c *fiber.Ctx) error {
		return nil
	})

	zkApi := api.Group("/zk")
	zkApi.Get("/create", func(c *fiber.Ctx) error {
		return nil
	})
	zkApi.Get("/verify", func(c *fiber.Ctx) error {
		return nil
	})

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal().Err(err).Msg("Server failed")
	}
}
