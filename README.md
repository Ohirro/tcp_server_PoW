# tcp_server_PoW

## Proof-of-Work TCP Test Task

This project is a proof-of-concept implementation of a TCP server protected from DDoS attacks using a Proof-of-Work (PoW) challenge-response protocol. The server issues a PoW challenge to each connecting client, and once the client successfully solves the challenge, the server responds with a random "Word of Wisdom" quote.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Architecture and PoW Algorithm](#architecture-and-pow-algorithm)
- [Project Structure](#project-structure)
- [Environment Configuration](#environment-configuration)
- [Usage](#usage)
 - [Makefile Commands](#makefile-commands)
 - [Docker](#docker)
- [Notes](#notes)
- [License](#license)

## Overview

The task was to design and implement a TCP server that:
- Is protected against DDoS attacks using a Proof-of-Work (PoW) challenge-response protocol.
- Explains the choice of the PoW algorithm.
- Sends a random quote from a collection ("Word of Wisdom") after successful PoW verification.
- Provides Dockerfiles for both the server and the client that solves the PoW challenge.

## Features

- **Challenge-Response PoW:** On each TCP connection, the server generates a random challenge and a difficulty level (e.g., a required number of leading zeros in the hash). The client must compute a nonce such that `SHA256(challenge + nonce)` meets the difficulty requirement.
- **Word of Wisdom:** Once the PoW is successfully solved, the server sends a random quote.
- **Dockerized:** Both server and client come with Dockerfiles to simplify deployment and testing.
- **Simple Makefile:** A top-level Makefile is provided to build, run, and clean the project components.

## Architecture and PoW Algorithm

### PoW Challenge-Response Protocol

1. **Connection:** The client connects to the TCP server.
2. **Challenge Issuance:** The server sends a challenge string along with a difficulty level in the format `challenge:difficulty\n`.
3. **Solving PoW:** The client computes a nonce (e.g., by iterating over integer values) such that `SHA256(challenge + nonce)` has a hexadecimal representation that starts with the required number of zeros (determined by the difficulty).
4. **Verification:** The client sends the found nonce back to the server. The server verifies the solution.
5. **Response:** If the nonce is valid, the server sends a random quote. Otherwise, it responds with an error message.

### Choice of PoW Algorithm

We use SHA256 for our PoW:
- **Security:** SHA256 is a widely trusted cryptographic hash function.
- **Simplicity:** It allows a straightforward implementation where the server only needs to check a hash prefix.
- **Performance:** The computational cost can be tuned via the difficulty parameter, making it simple to adjust the workload for clients.

## Project Structure

The repository is organized into two main folders: `server` and `client`. Each component is self-contained with its own configuration and build files.

```
tcp_server_PoW/
├── client/
│   ├── .env              # Client environment configuration (example file provided)
│   ├── go.mod            # Go module file for the client
│   ├── main.go           # Entry point for the client application
│   └── app/
│       ├── app.go        # Client business logic (PoW solving, etc.)
│       └── config.go     # Client configuration loader
├── server/
│   ├── .env              # Server environment configuration (example file provided)
│   ├── go.mod            # Go module file for the server
│   ├── main.go           # Entry point for the server application
│   └── app/
│       ├── app.go        # Server business logic (challenge generation, PoW verification, etc.)
│       └── config.go     # Server configuration loader
├── Makefile              # Top-level Makefile to build/run/clean both client and server
└── README.md             # This README file
```

> **Note:** The actual `.env` files with sensitive or production values are listed in `.gitignore`. Example configuration files are provided for reference.

## Environment Configuration

Both the server and client require an environment configuration file to run. Example files are provided in each folder:

### Server `.env` (located in `server/.env`)

```dotenv
# Port for the TCP server to listen on.
SERVER_PORT=9000
```

### Client `.env` (located in `client/.env`)

```dotenv
# Address (host:port) of the TCP server.
SERVER_ADDRESS=localhost:9000
```

Make sure that when you run the binaries, the working directory is set appropriately so that the `.env` file is loaded. If you use the provided Makefile or Dockerfiles, the working directory is managed accordingly.

## Usage

### Makefile Commands

A top-level Makefile is provided with the following commands:

- **Build everything (both server and client):**
  ```bash
  make
  ```
- **Build and run the server:**
  ```bash
  make run-server
  ```
- **Build and run the client:**
  ```bash
  make run-client
  ```
- **Clean binaries:**
  ```bash
  make clean
  ```

> **Tip:** The Makefile changes directories for building and running, ensuring that the respective `.env` files are correctly loaded.

### Docker

Dockerfiles are provided in both the `server` and `client` directories. To build and run using Docker:

1. **Build Docker Images:**

   For the server:
   ```bash
   cd server
   docker build -t word-of-wisdom-server .
   ```

   For the client:
   ```bash
   cd client
   docker build -t word-of-wisdom-client .
   ```

2. **Run the Containers:**

   Start the server (exposing port 9000):
   ```bash
   docker run -p 9000:9000 word-of-wisdom-server
   ```

   In another terminal, run the client:
   ```bash
   docker run --rm word-of-wisdom-client
   ```

## Notes

- **One-Shot Client:** The client is designed as a one-shot application. It connects to the server, solves the PoW challenge, prints the received quote, and then exits.
- **Environment Files:** The `.env` files provided in the `client` and `server` folders are examples. The actual `.env` files (with sensitive configurations) are excluded via `.gitignore`.
- **Extensibility:** Both the server and client code are structured in a modular way. You can easily add features like improved error handling, logging, or interactive modes.

## License

[MIT License](LICENSE)

