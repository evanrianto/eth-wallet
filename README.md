# Ethereum Wallet Key Pair Generator

This CLI tool generates an Ethereum wallet key pair, including a public key and a private key, using Go. It leverages the `go-ethereum` crypto package for secure cryptographic operations.

## Prerequisites

-   Go (version 1.13 or later recommended)
-   `git` for cloning the repository

## Installation

1. **Clone the Repository**

    First, clone this repository to your local machine using `git`:

    ```bash
    git clone https://github.com/evanrianto/eth-wallet
    cd eth-wallet
    ```

2. **Install Dependencies**

    Run the following command to fetch the necessary Go dependencies, including `go-ethereum`:

    ```bash
    go get -u github.com/ethereum/go-ethereum
    ```

## Usage

To generate a new Ethereum wallet key pair, use the following command from the root of the project directory:

```bash
go run walletgen.go generate
```

This will output a new key pair, printing the private key and public key to the terminal.

### Commands

-   `generate`: Generates a new Ethereum wallet key pair.

## Security Notice

This tool generates key pairs that can be used with Ethereum wallets. **Keep your private key secret and secure.** Exposure of your private key can result in loss of funds and unauthorized access to your wallet.

## Contributing

Contributions to improve the tool or fix issues are welcome. Please follow the standard process: fork the repository, make your changes, and submit a pull request.

<!-- ## License

[Specify the license here, e.g., MIT, Apache 2.0, etc.] -->
