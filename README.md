
# Flare Hotspot
Flare Hotpost core repository.

# System Requirements
- Make
- Docker
- Go

# Installation

Clone the project and prepare the development environment.
```sh
git clone git@github.com:flarehotspot/flarehotspot.git
cd flarehotspot
git checkout development
cp go.work.default go.work
```

Add these to your `/etc/hosts` file:

```sh
127.0.0.1    www.flarehotspot-dev.com
127.0.0.1    superuser.flarehotspot-dev.com
127.0.0.1    developer.flarehotspot-dev.com
127.0.0.1    pgadmin4.flarehotspot-dev.com
127.0.0.1    mailcatcher.flarehotspot-dev.com
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
git clone https://github.com/go-nv/goenv.git ~/.goenv
echo 'export GOENV_ROOT="$HOME/.goenv"' >> ~/.bash_profile
echo 'export PATH="$GOENV_ROOT/bin:$PATH"' >> ~/.bash_profile
echo 'eval "$(goenv init -)"' >> ~/.bash_profile
echo 'export PATH="$GOROOT/bin:$PATH"' >> ~/.bash_profile
echo 'export PATH="$PATH:$GOPATH/bin"' >> ~/.bash_profile
exec $SHELL
goenv install $(cat .go-version)
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

The database can be managed at [http://pgadmin4.flarehotspot-dev.com](http://pgadmin4.flarehotspot-dev.com)

# Flare CLI

Install the `flare` CLI:
```sh
go run ./core/cmd/build-cli/main.go
./bin/flare -h
```

# Documentation

Make sure `pipx` is available in your system and install the following packages:

```sh
pipx install mkdocs-material --include-deps
```

Then you can serve the local documentation server:

```sh
cd flarehotspot
make docs-serve
```

To build the documentation to be uploaded to the docs website:

```sh
make docs-build
```

---

# Steps in implementing git subtree for `go-utils`

## Split the utils to a `git subtree`.

```sh
git subtree split --prefix sdk/utils -b go-utils
```

This will create a new branch called `go-utils` which can be pushed to a git repo.

## Add the necessary `go.mod` file for making the `go-utils` a standalone library.

Example:
```go
module github.com/flarehotspot/go-utils

go 1.22.0
```

## Add the remote url of `flarehotspot/go-utils`

```sh
git remote add go-utils git@github.com:flarehotspot/go-utils.git
```

## Push the `go-utils` branch to a remote git repo.
```sh
git push go-utils go-utils:main
```

# Pushing changes to `go-utils`

```sh
# command guide
# git subtree push --prefix <utils dir name> <go-utils remote name or url> <desired local branch to push>
# don't worry, this will only push the changes inside the `utils` and not the entire local branch

# actual command
git subtree push --prefix sdk/utils go-utils development # or your desired local branch e.g. feat/utils-subtree
```

# Persist changes

For the changes to persist in other codebases that uses the go library, head over to the github or even to the local cloned repo of `go-utils` and create a git tag.

```sh
git checkout go-utils
git tag vx.x.x # creates a tag to the latest commit of the current branch
git push --tags # pushes the created tag
```

Then, update the `go-utils` library by specifying the version of the newly pushed tag.
```sh
go get -u github.com/flarehotspot/go-utils@vx.x.x
```

## Building `dev-kit`

Change owner to `$USER` the `flarehotspot` dir first:
```shell
sudo chown -R $USER <flarehotspot dir>
```

then, `make devkit`.

# Optimizing `gopls`

Make sure to exclude the following directories from `gopls` LSP:

- `.tmp`
