# Jim

The jim command line utility enables running long commands with one word, it's basically __alias but better__

# Installation
## Windows installation

download 

```
https://github.com/just-hms/jim/releases/latest/download/jim-windows-amd64.tar.gz 
```

and extract it in a folder that is included in the %PATH%


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
