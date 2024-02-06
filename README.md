# Flare Hotspot

Flare Hotpost core repository.

# System Requirements
- Node.js >= 16
- Go 1.19.12 (exact version)
- MySQL/MariaDB

# Installation
Clone the repository and install the dependencies.
```sh
git clone git@github.com:flarehotspot/flarehotspot.git
cd flarehotspot
npm install
```

Unzip the `config.zip` file.
```sh
rm -rf ./config
unzip config.zip -d config
```

Unzip the `openwrt-files.zip` file.
```sh
rm -rf ./openwrt-files
unzip openwrt-files.zip -d openwrt-files
```

# Database Connection
Open the `config/database.json` file and configure the database connection.
```json
{
    "host": "localhost",
    "username": "root",
    "password": "*****", // your password
    "database": "flarehotspot"
}
```

# Start the server
```sh
node ./run.dev.js
```
