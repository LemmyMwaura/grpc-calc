# Simple gRPC Service

This is a basic implementation of a calculator service using gRPC in Go. It provides simple arithmetic operations like addition, subtraction, multiplication, and division.

## Requirements

- Go (1.14 or higher)
- Protocol Buffers v3
- gRPC Go

## Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/LemmyMwaura/grpc-calc.git
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

## Usage

1. Start the server:

   ```bash
   make run-server
   ```

2. Run the grpcui client to make calculations. You can either use the provided client or write your own.

   ```bash
   make run-client
   ```
