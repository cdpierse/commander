# Commander

![Commander Cli](https://github.com/cdpierse/commander/assets/commander-example.gif?raw=true)

Commander is a simple CLI designed to provide copilot suggestions for command-line operations. It streamlines the process of generating command suggestions in your terminal based on the query you provide.

Commander is built to use [Ollama](https://ollama.com) and is designed to run all queries locally not requiring any external API calls.

The only model currently supported is Llama3 due to its higher performance with code completion tasks.

## Prerequisites

Before you begin, ensure you have the following installed on your system:

- Go (version 1.22 or higher)
- Ollama installed and configured locally, [click here for details](https://ollama.com)
- Llama3 must be pulled and running in Ollama before using Commander. To pull Llama3, run the following command:

```bash
    ollama pull llama3
```

## Installation

There are two main ways to install Commander: running from source or installing it globally using Go's toolchain.

### Option 1: Running from Source

This method is ideal if you plan to modify the source code or prefer to run the latest version directly.

1. **Clone the Repository**

```bash
git clone https://github.com/cdpierse/commander.git
cd commander
```

2. **Build and Run**

   Use the Go toolchain to build and run the application directly:

   ```bash
   go run main.go -q "your query"
   ```

   For example:

   ```bash
   go run main.go -q "delete all pods in kubernetes in the api namespace"
   ```

### Option 2: Installing Globally

If you prefer to install Commander globally:

1. **Clone the Repository**

   ```bash
   git clone https://github.com/cdpierse/commander.git
   cd commander
   ```

2. **Install the Binary**

   Use Go's native installation command to compile and place the binary in your GOPATH/bin, which should be in your PATH if your Go environment is set up correctly:

   ```bash
   go install .
   ```

3. **Run Commander**

   Now, you can run Commander from anywhere using the command:

   ```bash
   commander -q "your query"
   ```

   For example:

   ```bash
   commander -q "install numpy with pip"
   ```

## Usage

To get command line suggestions, use the following syntax:

```bash
commander -q "your query here"
```

Replace `"your query here"` with the actual command you need assistance with.

## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues to improve the documentation or functionality of Commander.
