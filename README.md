# OJ - Output JSON

OJ, short for Output JSON, is a powerful and flexible Command Line Interface (CLI) utility designed to convert the output of various Linux commands into a structured JSON format. This utility is particularly useful for system administrators, developers, and DevOps professionals who often need to parse and manipulate command-line output for scripting, automation, or monitoring purposes.

## Features

- Supports a wide range of Linux commands.
- Handles different output formats using various parsers.
- Extensible design allows for adding custom parsers.
- Converts command-line output to JSON.

## Usage

OJ can be used with any command that produces output. Simply pipe the command's output into OJ and specify the appropriate parser with the `-p` flag. If no parser is specified, OJ uses a default parser.

```bash
command | oj -p parser
```

In the following example, the `KeyValueParser` (`kv`) is used to parse the output of the `ls -l` command.

```bash
ls -l | oj -p kv
```

And to parse the output of the `printenv` command, you can use the following command:

```bash
printenv | oj
```

## Available Parsers

OJ comes with a set of built-in parsers:

- `default`: Parses key-value pairs separated by an equals sign.
- `kv`: Parses key-value pairs separated by a custom separator.
- `csv`: Converts CSV data into JSON.
- `yaml`: Converts YAML data into JSON.
- `xml`: Converts XML data into JSON.
- `log`: Parses common log formats.

## Building

OJ requires the Go compiler version 1.17 or later. You can use the provided `Makefile` to compile the project:

```bash
make build
```

This will produce an executable named `oj` in the `bin` directory.
