#!/usr/bin/env node

const { execSync } = require('child_process');
const fs = require('fs');
const path = require('path');
const https = require('https');

const WORKDIR = process.cwd();
const CACHE_PATH = path.join(WORKDIR, '.cache');

const GOOS = execSync('go env GOOS').toString().trim();
const GOARCH = execSync('go env GOARCH').toString().trim();
const GO_VERSION = fs
  .readFileSync(path.join(WORKDIR, 'go-version'), 'utf-8')
  .trim();
const GO_TAR = `go${GO_VERSION}.${GOOS}-${GOARCH}.tar.gz`;
const GO_SRC = `https://go.dev/dl/${GO_TAR}`;
const GO_CUSTOM_PATH = process.env.GO_CUSTOM_PATH || path.join(WORKDIR, 'go');
const DL_PATH = path.join(CACHE_PATH, 'downloads', GO_TAR);

console.log(`GOOS: ${GOOS}`);
console.log(`GOARCH: ${GOARCH}`);
console.log(`GO_CUSTOM_PATH: ${GO_CUSTOM_PATH}`);

const usage = () => {
  console.log(`
To use the installed go binary, add these lines to your .bashrc or .zshrc file:
      export PATH="${path.join(GO_CUSTOM_PATH, 'go', 'bin')}:$PATH"
      export PATH="${path.join(GO_CUSTOM_PATH, 'bin')}:$PATH"
      export GOROOT="${path.join(GO_CUSTOM_PATH, 'go')}"
      export GOPATH="${path.dirname(GO_CUSTOM_PATH)}/go"
`);
};

const downloadGo = (GO_SRC_URL, DOWNLOAD_PATH, callback) => {
  fs.mkdirSync(path.dirname(DOWNLOAD_PATH), { recursive: true });
  const file = fs.createWriteStream(DOWNLOAD_PATH);
  const request = https.get(GO_SRC_URL, function (response) {
    if (response.headers.location) {
      console.log('Redirecting to', response.headers.location);
      return downloadGo(response.headers.location, DOWNLOAD_PATH, callback);
    }

    response.pipe(file);
    // after download completed close filestream
    file.on('finish', () => {
      file.close();
      console.log('Download Completed');
      console.log(fs.statSync(DOWNLOAD_PATH));
      callback();
    });

    file.on('error', (err) => {
      fs.unlink(DOWNLOAD_PATH, () => {
        if (callback) callback(err.message);
      });
    });
  });

  request.on('error', (err) => {
    fs.unlink(DOWNLOAD_PATH, () => {
      if (callback) callback(err.message);
    });
  });
};

if (
  fs.existsSync(path.join(GO_CUSTOM_PATH, 'go-version')) &&
  GO_VERSION ===
    fs.readFileSync(path.join(GO_CUSTOM_PATH, 'go-version'), 'utf-8').trim()
) {
  console.log('Go is already installed');
  usage();
  process.exit(0);
} else {
  console.log(`Downloading ${GO_SRC}...`);
  downloadGo(GO_SRC, DL_PATH, (error) => {
    if (error) {
      console.error(error);
      process.exit(1);
    }
    // extract tarfile to GO_CUSTOM_PATH
    console.log(`Extracting ${DL_PATH} to ${GO_CUSTOM_PATH}`);
    fs.mkdirSync(GO_CUSTOM_PATH, { recursive: true });
    execSync(`tar -C ${GO_CUSTOM_PATH} -xzf ${DL_PATH}`);
    fs.renameSync(`${GO_CUSTOM_PATH}/go`, `go.old`);
    fs.rmdirSync(GO_CUSTOM_PATH, { recursive: true });
    fs.renameSync(`go.old`, GO_CUSTOM_PATH);
    console.log(`Installed Go ${GO_VERSION} to ${GO_CUSTOM_PATH}`);
    usage();
  });
}

process.chdir(WORKDIR);
