# Jim

The jim command line utility enables running long commands with one word, it's basically __alias but better__

<!-- ## Windows installation

use `powershell`

```powershell
	
$ wget -O %TMP%/jim.tar.gz https://github.com/just-hms/jim/releases/download/v0.0.4/jim-windows-amd64.tar.gz
$ mkdir -p %Programfiles%/jim
$ tar -xvf %TMP%/jim.tar.gz -C %Programfiles%/jim
$ setx
 
``` --> 

## Linux installation

```sh
$ curl -L https://github.com/just-hms/jim/releases/download/v0.0.4/jim-linux-amd64.tar.gz > /tmp/jim.tar.gz 
$ sudo mkdir -p /opt/jim
$ sudo tar -xvf /tmp/jim.tar.gz -C /opt/jim/
$ sudo sh -c 'grep -qxF "export PATH=/opt/jim:\$PATH" /etc/profile || echo "export PATH=/opt/jim:\$PATH" >> /etc/profile'
```

## Mac-OS installation

```sh
$ curl -L https://github.com/just-hms/jim/releases/download/v0.0.4/jim-darwin-amd64.tar.gz > /tmp/jim.tar.gz 
$ sudo mkdir -p /opt/jim
$ sudo tar -xvf /tmp/jim.tar.gz -C /opt/jim/
$ sudo sh -c 'grep -qxF "export PATH=/opt/jim:\$PATH" /etc/profile || echo "export PATH=/opt/jim:\$PATH" >> /etc/profile'
```


