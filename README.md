<p align="center">
<a href="https://go.dev"><img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="golang"></a>
<a style="margin:0px 1rem;"href="https://github.com/tidwall/buntdb"><img style="height:28px;" src="https://github.com/tidwall/buntdb/raw/master/logo.png" alt="Bunt DB"></a>
<a href="https://github.com/features/actions"><img src="https://img.shields.io/badge/GitHub_Actions-2088FF?style=for-the-badge&logo=github-actions&logoColor=white" alt="Github actions"></a>
</p>
	
# Jim

The `jim` command line utility enables running long commands with one word, it's basically __alias but better__

## Why `jim`?

- The main focus of `jim` is to let you re-use shell commands that you previously inserted in it. For example:

	```
	jim command
	```

	will launch the specified command.

- Another use-case is to launch a command with the `--watch` utility that will time its execution and will let you see it using the `--show` utility. 

	```
	jim --watch command
	```

	will launch the specified command *in background* and time it.

In future `jim --sync` will let you share commands between multiple devices.

## New features 

### Shebang 

if a command value starts with something like

```
#!/usr/bin/env python3

...
```

`jim` will launch the command using `python3`

## Installation
### Windows installation

Download

https://github.com/just-hms/jim/releases/latest/download/jim-windows-amd64.tar.gz 

and extract it in a folder that is included in the `%PATH%`.

#### Linux installation

```sh
$ curl -L https://github.com/just-hms/jim/releases/latest/download/jim-linux-amd64.tar.gz > /tmp/jim.tar.gz
$ sudo mkdir -p /opt/jim && sudo tar -xvf /tmp/jim.tar.gz -C /opt/jim/
$ [ ! -f "/usr/local/bin/jim" ] && sudo ln -s /opt/jim/jim /usr/local/bin/jim
```

#### Mac-OS installation

```sh
$ curl -L https://github.com/just-hms/jim/releases/latest/download/jim-darwin-amd64.tar.gz > /tmp/jim.tar.gz
$ sudo mkdir -p /opt/jim && sudo tar -xvf /tmp/jim.tar.gz -C /opt/jim/
$ [ ! -f "/usr/local/bin/jim" ] && sudo ln -s /opt/jim/jim /usr/local/bin/jim
```

## Usage

Type `jim` to check if the installation was completed correctly.

### Available Actions

#### `--add`

Adds a new command to `jim`

```
$ jim --add command <value>
```

If no value is specified `jim` will open your *default editor* and will let you insert a set of instruction in a temporary file.

<br>

#### `--clear`

Clear all commands

```
$ jim --clear
```

Will remove all commands.

<br>

#### `--help`

`jim` will help you with what you need

```
$ jim --help <--action>
```

If provided, this utility will show more specific help for the action.

<br>

#### `--ls`

List of all the available commands

```
$ jim --ls <filter>
```

Will list all the available commands, filtering them with the provided filter.

<br>

#### `--mod`

Modify a command 

```
$ jim --mod command
```

Will open the command in your default editor and will let you modify it.

<br>

#### `--rm`

Remove one or more command 

```
$ jim --rm command_1 <command_2> ...
```

Will remove the provided commands.

<br>

#### `--rn`

Rename a command

```
$ jim --rn command new_name
```

Will rename the specified command with the provided `new_name`.

<br>

#### `--run`

Run a command

```
$ jim <--run> command
```

Will run the specified command in *your default shell*, `--run` can be omitted.

<br>

#### `--show`

Show a list of all the `--watch` results

```
$ jim --show <filter>
```

Will list all of the commands' sessions. Filtering them with the provided filter.

A session is created when a command is launched with `--watch`.

<br>

#### `--upgrade`

Upgrade `jim`, this action needs to be run with `sudo` on Linux

```
# jim --version
```

If you have installed the last version this action will do nothing

<br>

#### `--version`

Show the version of the executable

```
$ jim --version
```

Will output the installed `jim` version ex: `v1.0.1`.

<br>

#### `--watch`

Run a command in the background and time it

```
$ jim --watch command
```

Will launch the command in background and save its time of execution. The time that the command took to execute will be visible using the `--show` utility.

__*User input and output don't work using `--watch`*__
