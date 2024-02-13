+++
title = "Getting Started"
description = "Series of tutorials to get you started with using the Flare Hotspot SDK."
date = 2021-05-01T08:00:00+00:00
updated = 2021-05-01T08:00:00+00:00
draft = false
weight = 100
sort_by = "weight"
template = "docs/page.html"

[extra]
lead = ""
toc = true
top = false
+++

# Getting Started

## Overview

The plugin SDK is a set of tools and libraries that allow you to build applications that interact with the [Flare Hotspot System](https://www.flarehotspot.com). It is developed using the [Go Programming Language](https://go.dev). Go is a simple but fast and efficient programming language suitable for building system applications while also being easy to learn for beginners. This series of tutorials will guide you through the process of building a simple plugin using the Flare Hotspot SDK.

Througout this tutorial, we assume you are using a Debian based operating system like [Ubuntu](https://ubuntu.com/) or some of its [derivatives](https://en.wikipedia.org/wiki/Category:Ubuntu_derivatives). Although this works on any other linux distro,
we'll just stick to one that's more familiar for most users.

---

# Prerequisites

Before we can proceed with installation, we must first install some development tools needed to compile and run the system. Below are the needed tools.

- Node.js >= 18 (18 recommended)
- Go 1.19.12 (must be exact version)
- MariaDB (or MySQL)

The instructions for installing these tools are provided in the next section of this page.

## System Packages
Before we proceed, let's make sure all the needed system tools are available in our system.
```sh
sudo apt update
sudo apt install -y curl wget git build-essential
```

We also need to remove existing installation of Node.js, Go, MySQL and MariaDB if any.
```sh
sudo apt remove -y --purge nodejs npm golang-go mysql-server mariadb-server
```

## Installing Node.js

Node.js is needed to execute scripts to compile and run the system. Let's install Node.js using a PPA maintained by NodeSource:

```sh
cd ~
curl -sL https://deb.nodesource.com/setup_18.x -o nodesource_setup.sh
sudo bash nodesource_setup.sh
```

Then install Node.js:

```sh
sudo apt update
sudo apt install -y nodejs
```

To verify that Node.js is installed in your machine, type:

```sh
node -v
# v18.0.0
```

It should output the installed version of Node.js in your system.

## Installing Go

We need to install a specific version of Go which is `1.19.12`.

```sh
cd ~
wget https://go.dev/dl/go1.19.12.linux-amd64.tar.gz
sudo rm -rf /usr/local/go # remove any previously installed go
sudo tar -C /usr/local -xzf go1.19.12.linux-amd64.tar.gz
```

We need to add the go path to our `$PATH` environment variable:
```sh
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
source ~/.bashrc
```

Verify the installed Go version. The result should be `go version go1.19.12 linux/amd64`.
```sh
go version
# go version go1.19.12 linux/amd64
```

## Installing MariaDB
Install the `mariadb-server` package:
```sh
sudo apt update
sudo apt install mariadb-server
sudo mysql_secure_installation
```

You will then be asked to enter the current password of the root mariadb user. Just PRESS **ENTER**.
```
NOTE: RUNNING ALL PARTS OF THIS SCRIPT IS RECOMMENDED FOR ALL MariaDB
      SERVERS IN PRODUCTION USE!  PLEASE READ EACH STEP CAREFULLY!

In order to log into MariaDB to secure it, you'll need the current
password for the root user.  If you've just installed MariaDB, and
you haven't set the root password yet, the password will be blank,
so you should just press enter here.

Enter current password for root (enter for none):
```

You will then be asked to switch the root authentication using unix_socket. Just type **n**.
```
Setting the root password or using the unix_socket ensures that nobody
can log into the MariaDB root user without the proper authorisation.

You already have your root account protected, so you can safely answer 'n'.

Switch to unix_socket authentication [Y/n] n
```

You will then be asked to change the MariaDB root password. Just type **n** since this is just for development setup and convenience is a priority over security.
```
OK, successfully used password, moving on...

Setting the root password ensures that nobody can log into the MariaDB
root user without the proper authorisation.

Set root password? [Y/n] n
```

From here on, you can press **Y** and then ENTER to accept the defaults for all the subsequent questions.

After the installation process, let's create a new database user. Type:
```sh
sudo mariadb
```
Once you're inside MariaDB console, type the command below to create a new user with root privileges. Make sure to set the `password` as you like.
```sql
MariaDB [(none)]> GRANT ALL ON *.* TO 'admin'@'localhost' IDENTIFIED BY 'password' WITH GRANT OPTION;
```

Then fluash the privilges so that it persists in the current session.
```sql
MariaDB [(none)]> FLUSH PRIVILEGES;
```

Let's create our development database.
```sql
MariaDB [(none)]> CREATE DATABASE flarehotspot_dev;
```

Once you're done, exit the MariaDB console:
```
MariaDB [(none)]> exit
```

Our database should now be ready. To test the connection to your database, let's enter the MariaDB console using our newly created "admin" user and connect to our "flarehotspot_dev" database.
```sh
mariadb -h localhost -u admin -ppassword flarehotspot_dev
```

You should now be in the Mariadb console. Now exit the console:
```sql
MariaDB [flarehotspot_dev]> exit
```

---

# Installing the SDK

## Download SDK
To install the plugin SDK, download the latest **devkit-x.x.xzip** file from [sdk-releases](https://github.com/flarehotspot/sdk-releases/releases) repository.
After downloading, extract the zip file to your desired location.
```sh
# replace ~/Downloads/devkit-x.x.x.zip with the path to the downloaded zip file
unzip ~/Downloads/devkit-x.x.x.zip -d ~/Documents/flare-devkit
cd ~/Documents/flare-devkit
```

Below is the directory structure of the zip file:
```
|- config
|- core
|- main
|- mock-files
|- plugins
    |-- com.flarego.default-theme
    |-- com.flarego.sample-plugin
```

Notice the `com.flarego.sample-plugin` directory. We are going to use this sample plugin as the base of our project.

## Configure Project
Our application needs to connect to the database we [created](#installing-mariadb) earlier, open the file `config/database.json` and set the database connection settings.
```json
{
    "host":     "localhost",
    "username": "admin",
    "password": "password",
    "database": "flarehotspot_dev"
}
```

## Start Application
Before starting the application, make sure to install the needed node modules.
```sh
npm install
```

Then start the application by running the `npm start` command.
```sh
npm start
```
You can now browse the application in [http://localhost:3000](http://localhost:3000)

The admin dashboard can be accessed in [http://localhost:3000/admin](http://localhost:3000/admin)

The default login for the admin dashboard is:
```
username: admin
password: admin
```
