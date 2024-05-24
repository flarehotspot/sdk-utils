
# Flare Hotspot
Flare Hotpost core repository.

# System Requirements
- Docker

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

# Installing Go
```sh
curl -sSL https://github.com/moovweb/gvm/raw/master/binscripts/gvm-installer | bash
gvm install "$(cat .go-version)"
cd ../flarehotspot # load go version
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

# Documentation

Make sure `pipx` is available in your system and install the following packages:

```sh
pipx install mkdocs-material --include-deps
```

Then you can serve the local documentaion server:

```sh
cd flarehotspot
make docs-serve
```

To build the documentation to be uploaded to the docs website:

```sh
make docs-build
```
