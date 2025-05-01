This is an implementation of the
["Build Your Own Shell" Challenge](https://app.codecrafters.io/courses/shell/overview).

In this challenge, you'll build your own POSIX compliant shell that's capable of
interpreting shell commands, running external programs and builtin commands like
cd, pwd, echo and more. Along the way, you'll learn about shell command parsing,
REPLs, builtin commands, and more.

**Note**: If you're viewing this repo on GitHub, head over to
[codecrafters.io](https://codecrafters.io) to try the challenge.

# 🐚 GoShell – A Simple Shell in Go

GoShell is a minimalistic, interactive command-line interface (shell) written in Go. It offers basic shell functionalities like echo, cd, pwd, type, and executing external programs. Ideal for learning, experimenting, or as a foundation for your own shell projects!
### ✨ Features

- Interactive Input: Commands are entered and executed directly in the shell.
- Built-in Commands:
   - echo: Prints text to the console.
   - pwd: Shows the current working directory.
   - cd <dir>: Changes the directory (including support for ~).
   - type <command>: Shows how a command is interpreted (builtin, external, or not found).
   - exit: Exits the shell.
- External Programs: Programs in $PATH can be called as usual (e.g., ls, git, cat...).
- Argument Parsing: Supports single quotes, double quotes, and escape characters in the input.
- Error Messages: Clear outputs for errors or unknown commands.

### 🚀 Getting Started
Prerequisites

- Go (at least version 1.16 recommended)

Installation & Start

```bash
git clone https://github.com/Driemtax/codecrafters-shell-go.git
cd goshell
go run main.go
```

### 🛠️ Examples

```shell
$ echo "Hello World!"
Hello World!

$ pwd
/home/user/goshell

$ cd ..
$ pwd
/home/user

$ type echo
echo is a shell builtin

$ type ls
ls is /bin/ls

$ exit
```

### ⚠️ Notes & Tips

- Whitespace & Quotes: The shell supports single and double quotes, as well as escape characters, but the parsing is not perfect and may be limited for complex inputs.
- External Programs: Only programs located in the $PATH can be started.
- Error Handling: Faulty commands or invalid directories are reported informatively.

### 💡 ToDo / Ideas for Extensions

- Better support for pipes (|) and redirections (>, <)
- Improved argument parsing logic
- Support for environment variables
- Autocompletion and history

### 📄 License

MIT License – feel free to use, modify, and contribute!
