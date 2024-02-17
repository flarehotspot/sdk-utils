#!/usr/bin/env node

const path = require('path');
const fs = require('fs-extra');
const goCustomPath = process.env.GO_CUSTOM_PATH;
const goCustomBin = path.join(goCustomPath, 'bin', 'go');
const execAsync = require('./exec-async');

module.exports = async () => {
  const GO_BIN = (await fs.exists(goCustomBin))
    ? goCustomBin
    : await execAsync(`which go`).then((str) => str.trim());
  const GOOS = await execAsync(`${GO_BIN} env GOOS`).then((str) => str.trim());
  const GOARCH = await execAsync(`${GO_BIN} env GOARCH`).then((str) => str.trim());

  const env = { GOOS, GOARCH, GO_BIN };

  console.log('ENV:', env);

  return env;
};
