
# Flare Hotspot

Flare Hotpost core repository.

# System Requirements

- Go 1.19.12 (exact version)
- Docker
- Node.js >= 16

# Installation

Clone the project and prepare the development environment.
```sh
git clone git@github.com:flarehotspot/flarehotspot.git
cd flarehotspot
git checkout development
cp go.work.default go.work
git submodule update --init --recursive
```

Pull the latest changes for all the submodules.
```sh
for p in $(ls ./plugins); do cd ./plugins/${p}; git checkout development; git pull; cd ../..; done
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

# Start the server

```sh
docker compose up --build
```
Now you can browse the portal at [http://localhost:3000](http://localhost:3000)

The admin dashboard can be accessed at [http://localhost:3000/admin](http://localhost:3000/admin)

The default admin access is:
```
username: admin
password: admin
```

# Environment variables
```sh
export GOPATH="$HOME/go"
export PATH="$GOPATH/bin:$PATH"
```

# Flare CLI

There are two Flare CLI tools:

Install the `flare` sdk CLI:
```sh
go install ./core/devkit/cli/flare.go
flare --help
```

Install the `flare-internal` CLI:
```sh
go install ./core/internal/cli/flare-internal.go
flare-internal --help
```
