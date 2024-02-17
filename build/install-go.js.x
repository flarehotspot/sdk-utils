#!/usr/bin/env node

const fs = require('fs-extra');
const path = require('path');
const https = require('https');
const execAsync = require('./exec-async');
const goEnv = require('./go-env');

module.exports = (async () => {
  const ROOT_DIR = path.join(__dirname, '..');
  const CACHE_PATH = path.join(ROOT_DIR, '.cache');
  const EXTRACT_PATH = path.join(CACHE_PATH, 'extracted/go');
  const GO_VERSION_PATH = path.join(ROOT_DIR, 'core/go-version');

  const { GOOS, GOARCH } = await goEnv();
  const GO_VERSION = fs.readFileSync(GO_VERSION_PATH, 'utf-8').trim();
  const GO_TAR = `go${GO_VERSION}.${GOOS}-${GOARCH}.tar.gz`;
  const GO_SRC = `https://go.dev/dl/${GO_TAR}`;
  const GO_CUSTOM_PATH =
    process.env.GO_CUSTOM_PATH || path.join(ROOT_DIR, 'go');
  const DL_PATH = path.join(CACHE_PATH, 'downloads', GO_TAR);

  console.log(`GO_CUSTOM_PATH: ${GO_CUSTOM_PATH}`);

  const usage = () => {
    console.log(`
To use the installed go binary, add these lines to your .bashrc or .zshrc file:
      export PATH="${path.join(GO_CUSTOM_PATH, 'bin')}:$PATH"
      export PATH="$HOME/go/bin:$PATH"
      export GOROOT="${path.join(GO_CUSTOM_PATH)}"
`);
  };

  const downloadGo = async (GO_SRC_URL, DOWNLOAD_PATH) => {
    await new Promise((resolve, reject) => {
      fs.emptyDirSync(path.dirname(DOWNLOAD_PATH));
      const file = fs.createWriteStream(DOWNLOAD_PATH);
      const request = https.get(GO_SRC_URL, function (response) {
        if (response.headers.location) {
          console.log('Redirecting to', response.headers.location);
          const dl = downloadGo(response.headers.location, DOWNLOAD_PATH);
          dl.then(resolve);
          dl.catch(reject);
          return;
        }

        response.pipe(file);
        // after download completed close filestream
        file.on('finish', () => {
          file.close();
          console.log('Download Completed');
          console.log(fs.statSync(DOWNLOAD_PATH));
          resolve();
        });

        file.on('error', (err) => {
          fs.rmSync(DOWNLOAD_PATH);
          reject(err);
        });
      });

      request.on('error', (err) => {
        fs.rmSync(DOWNLOAD_PATH);
        reject(err);
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
    await downloadGo(GO_SRC, DL_PATH);

    console.log(`Cleaning up ${GO_CUSTOM_PATH}`);
    if (await fs.exists(GO_CUSTOM_PATH))
      await fs.rm(GO_CUSTOM_PATH, { recursive: true });

    console.log(`Extracting ${DL_PATH} to ${EXTRACT_PATH}`);
    await fs.emptyDir(EXTRACT_PATH);
    await execAsync(`tar -C ${EXTRACT_PATH} -xzf ${DL_PATH}`);
    await fs.copy(`${EXTRACT_PATH}/go`, GO_CUSTOM_PATH);
    await fs.rm(EXTRACT_PATH, { recursive: true });
    console.log(`Installed Go ${GO_VERSION} to ${GO_CUSTOM_PATH}`);
    usage();
    process.exit(0);
  }
})();
