# Getting Started

-----

## Requirements

```bash
docker
make
golang
```

## Clone Repository

Clone and initialize the project with the following commands:

```
git clone git@github.com:flarehotspot/flarehotspot.git
cd flarehotspot
git submodule update --init --recursive
make checkout_main
make pull
```

## Configure Database

To easily install `mariadb` database, install and follow this installation in [this repository](https://github.com/adonespitogo/docker-services).

Create a file `config/database.yml` with the following content:

```
host: localhost
username: root
password: rootpass
database: flarehotspot
```

## Run Application

To run the application, run `make` command. Then visit http://localhost:3000/admin

The default admin credentials are:

```
Username: admin
Password: admin
```
