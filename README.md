# Jim

The jim command line utility enables running long commands with one word, it's basically __alias but better__


# Table of contents

- [Why use `jim`?](#why-use-jim)
- [Installation](#installation)
- [Usage](#usage)

# Why use `jim`?

- The main focus of `jim` is to let you re-use commands that you previously inserted in it. For example:

	```
	jim command
	```

	will launch the command.

- Another use-case is to launch a command with the `--watch` utility that will time the execution of the command and will let you see it using the `--show` utility. 

	```
	jim --watch command
	```

	will launch a command in background and time it.

in future `jim --sync` will let you share commands between multiple devices.

# Installation
## Windows installation

Download 

```
https://github.com/just-hms/jim/releases/latest/download/jim-windows-amd64.tar.gz 
```

and extract it in a folder that is included in the `%PATH%`.


## Linux installation

```sh
$ curl -L https://github.com/just-hms/jim/releases/latest/download/jim-linux-amd64.tar.gz > /tmp/jim.tar.gz
$ sudo mkdir -p /opt/jim && sudo tar -xvf /tmp/jim.tar.gz -C /opt/jim/
$ sudo ln -s /opt/jim/jim /usr/local/bin/jim
```

## Mac-OS installation

```sh
$ curl -L https://github.com/just-hms/jim/releases/latest/download/jim-darwin-amd64.tar.gz > /tmp/jim.tar.gz
$ sudo mkdir -p /opt/jim && sudo tar -xvf /tmp/jim.tar.gz -C /opt/jim/
$ sudo ln -s /opt/jim/jim /usr/local/bin/jim
```

# Usage

type `jim` to check if the installation was completed correctly.

## Actions

### `--add`

add a command

#### syntax

```
jim --add command <value>
```

if no value is specified `jim` will open your default editor and let you insert a file.

### `--clear`

clear all commands

#### syntax

```
jim --clear
```

will remove all commands

### `--help`

`jim` will help you with what you need

#### syntax

```
jim --help <--action>
```

if provided this utility will show more specific help of the action.

### `--ls`

list of all the available commands

#### syntax

```
jim --ls <filter>
```

this command will list all the available commands, filtering them with the provided filter.

### `--mod`

Modify a command 

#### syntax

```
jim --mod command
```

will open the command in your default editor and will let you modify it

### `--rm`

Remove one or more command 

#### syntax

```
jim --rm command_1 command_2 command_3
```

will remove the provided commands


### `--rn`

Rename a command

#### syntax

```
jim --rn command new_name
```

will rename the specified command with the provided `new_name`

### `--run`

Run a command

#### syntax

```
jim <--run> command
```

will run the specified command, `--run` can be omitted

### `--show`

Show a list of all the `--watch` result

#### syntax

```
jim --show <filter>
```

this command will list all the sessions of the commands that were launched by `--watch`, filtering them with the provided filter.


### `--version`

Show the version of the executable

#### syntax

```
jim --version
```

will output the installed `jim` version ex: `v1.0.1`


### `--watch`

Run a command in the background and time it

#### syntax

```
jim --watch command
```

will launch the command in background and add its time of execution in the database, the time will be visible using the `--show` utility

__*User input and output don't work using watch*__

