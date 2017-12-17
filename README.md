## Description

This is an example Ethereum DApp that includes a sample Ethereum smart contract, as well as a simple Server / Frontend for interacting with that contract.

This repo currently powers:

[Biggest G](https://biggestg.com) and [Worlds Greatest Human](https://worldsgreatesthuman.com)

The contract works by storing a single string that users can change by paying a fee. Once the string is changed, the fee to change it again doubles.

## Setup

### Server

1. Cone this repo into your `$GOPATH`
2. Run `glide install`
3. Run `go get -u "golang.org/x/crypto/..."` because glide sucks
4. Run `go build "github.com/richardartoul/worlds" && ./server`

### Ethereum Contract

1. Install truffle `npm install -g truffle`
2. Run `npm install` in the contract directory
3. Run `truffle test` or `truffle compile` in the contract directory