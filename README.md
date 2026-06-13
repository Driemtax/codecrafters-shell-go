# Go Custom Shell

*This project is a custom shell implementation written in Go, originally created as part of the [CodeCrafters "Build Your Own Shell" Challenge](https://app.codecrafters.io/courses/shell/overview).*

## Overview

This application is a lightweight, interactive command-line interface (CLI) shell. It parses user input, handles complex argument formatting (including single quotes, double quotes, and escape characters), and executes both built-in commands and external system programs. You can find the source code for the integrated calculator here: [Calculator](https://github.com/Driemtax/Calculator). It uses recursive-descent parsing to handle operator precedences. 

### Core Features
*   **Custom Input Parsing:** Accurately interprets spaces, single quotes (`'`), double quotes (`"`), and backslash escapes (`\`) to group arguments properly.
*   **Built-in Commands:** Essential shell utilities implemented natively within the Go application.
*   **External Command Execution:** Seamlessly finds and runs executables located in your system's `PATH` (e.g., `git`, `ls`, `cat`).
*   **External Package Integration:** Demonstrates how to embed external Go modules as native shell applications (see the `calc` command).

---

## Command Overview

| Command | Description |
| :--- | :--- |
| `echo` | Prints the provided arguments to the standard output. |
| `exit` | Gracefully exits the shell. |
| `type` | Identifies whether a command is a shell builtin or an external executable. |
| `pwd` | Prints the absolute path of the current working directory. |
| `cd` | Changes the current working directory. |
| `calc` | Launches an interactive, built-in scientific calculator. |
| *(External)* | Any valid executable found in the system `PATH` will be executed. |

---

## Command Documentation

### `echo [args...]`
Prints the given arguments to the console, separated by spaces, followed by a newline. It respects string literals wrapped in single or double quotes, allowing you to echo strings with multiple consecutive spaces or special characters.

### `exit`
Terminates the shell session and returns a status code of `0` to the operating system.

### `type [command]`
Inspects the given command string and reports how the shell will interpret it. 
*   If it is a builtin (like `cd` or `pwd`), it outputs: `[command] is a shell builtin`.
*   If it is an external program found in the system's `PATH`, it outputs the absolute path to the executable: `[command] is /path/to/executable`.
*   If it cannot be found, it reports: `[command]: not found`.

### `pwd`
Outputs the absolute path of the current working directory to the console.

### `cd [path]`
Changes the shell's current working directory to the specified `path`. 
*   Supports absolute paths (e.g., `/usr/bin`).
*   Supports relative paths (e.g., `./dir` or `../`).
*   Supports the tilde character (`~`) to navigate directly to the current user's home directory.
*   If the directory does not exist, it will print an error: `cd: [path]: No such file or directory`.

### `calc`
Enters an interactive calculator sub-shell. This mode allows you to evaluate mathematical expressions in real-time. Type `help` while inside the calculator for a full list of operations, or `exit` to return to the main shell.
**Supported Operations:**
*   Arithmetic: `+`, `-`, `*`, `/`, and parentheses `()` for grouping.
*   Trigonometry: `sin(x)`, `cos(x)`, `tan(x)` (expects angles in radians).
*   Constants: `pi`.

---

## External Package Integration (Example: Calculator)

A key feature of this shell is its modularity and ability to integrate external Go packages to power built-in commands. 

The `calc` command is powered by an external module. Instead of writing the math evaluation logic directly inside the shell's source code, the shell imports a dedicated evaluation engine:

```go
import "github.com/driemtax/Calculator/pkg/calculator"
```

This demonstrates how you can easily expand the shell's capabilities by pulling in specialized repositories.
