
# Flare Hotspot

Flare Hotpost core repository.

# System Requirements

- Node.js >= 16
- Go 1.19.12 (exact version)
- [MySQL/MariaDB](https://github.com/adonespitogo/docker-services?tab=readme-ov-file#mariadb-service)

# Installation

Clone the repository and install the dependencies.

```sh
git clone git@github.com:flarehotspot/flarehotspot.git
cd flarehotspot
cp go.work.default go.work
git submodule update --init --recursive
```

Checkout and pull the latest changes for all the submodules

```sh
for p in $(ls ./plugins); do cd ./plugins/${p}; git checkout main; git pull; cd ../..; done
```

Install node modules.

```sh
npm install
```

Unzip the `openwrt-files.zip` file.

```sh
rm -rf ./openwrt-files
unzip openwrt-files.zip -d openwrt-files
```

# Database Connection

Open (or create) the `config/database.json` file and configure the MySQL/MariaDB database connection. Make sure to change the parameters accordingly.

```json
{
    "host": "localhost",
    "username": "root",
    "password": "*****",
    "database": "flarehotspot"
}
```

# Start the server

```sh
make
```
Now you can browse the portal at [http://localhost:3000](http://localhost:3000)

The admin dashboard can be accessed at [http://localhost:3000/admin](http://localhost:3000/admin)

The default admin access is:
```
username: admin
password: admin
```

# Flare CLI

There are two Flare CLI tools:

- `flare` - CLI tool shipped together with the SDK for public use
- `flare-internal` - CLI tool for internal (core) developers use

Install the `flare` sdk CLI:
```sh
$ go install ./core/cli/flare.go
$ flare --help
```

Install the `flare-internal` CLI:
```sh
$ go install ./core/internal/cli/flare-internal.go
$ flare-internal --help
```
