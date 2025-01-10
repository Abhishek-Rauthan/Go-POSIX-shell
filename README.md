# POSIX Shell in Go

A POSIX-compliant shell implementation written in Go. This project aims to mimic the behavior of standard POSIX shells.

## Features

- POSIX-compliant command execution
- Supports standard shell commands (e.g., `echo`, `cd`, `ls`, etc.)
- Built-in shell commands like `cd`, `exit`, and `pwd`

## Installation

To install and run the POSIX shell, follow the steps below:

### Prerequisites

- Go 1.18+ installed. You can download it from [here](https://golang.org/dl/).

### Clone the repository

```bash
git clone https://github.com/Abhishek-Rauthan/posix-shell-go.git
cd posix-shell-go
```

### Build the shell

Run the following command to build the shell binary:

```bash
make
```

This will produce a binary named `posix-shell` that you can run directly from your terminal.

### Run the shell

After building the shell, you can run it by executing the following command:

```bash
make run
```

This will start the POSIX shell and prompt you to start entering shell commands.

## Usage

Once the shell is running, you can enter commands just like in any standard shell. Some examples:

### Running Commands

```bash
$ echo "Hello, World!"
Hello, World!
```

### Built-in Commands

Some built-in commands include:

- `exit`: Exit the shell.
- `cd <dir>`: Change the current directory.
- `pwd`: Print the current working directory.

```bash
$ cd /home/user
$ pwd
/home/user
$ exit
```

## Video Demonstration

[Watch the video](shell.mkv)

## Contributing

We welcome contributions to improve and extend the functionality of the POSIX shell. If you would like to contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature/your-feature`).
6. Create a pull request.
