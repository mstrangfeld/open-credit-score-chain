.PHONY: build-contracts
build-contracts:
	solc --optimize --bin --abi --overwrite ./contracts/Authority.sol -o ./build/contracts
	solc --optimize --bin --abi --overwrite ./contracts/CreditScoreDB.sol -o ./build/contracts

gen-go-bindings: build-contracts
	abigen --sol ./contracts/Authority.sol --pkg=authority --out=./backend/contracts/authority/Authority.go
	abigen --sol ./contracts/CreditScoreDB.sol --pkg=creditscoredb --out=./backend/contracts/creditscoredb/CreditScoreDB.go

.PHONY: deploy-contracts
deploy-contracts:

.PHONY: clean
clean:
	rm -rf build
