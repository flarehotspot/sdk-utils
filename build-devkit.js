#!/usr/bin/env node
const fs = require('fs-extra');
const { exec, spawn } = require('child_process');
const { promisify } = require('util');
const execAsync = promisify(exec);

const CORE_VERSION = require('./core/plugin.json').version;
const DOCKER_IMAGE = 'devkit:latest';
const TMP_CONTAINER = 'devkit-tmp';
const CORE_SO = '/root/core.so';
const OUTFILE = 'devkit-extras/core/core.so';
const RELEASE_DIR = 'devkit-release/devkit-' + CORE_VERSION;
const DEVKIT_FILES = [
  './main',
  './core/go.mod',
  './core/go.sum',
  './core/sdk',
  './core/resources',
  './core/plugin.json',
  './install-go.js',
  './make-go.work.js',
  './build-main.js',
  './build-plugins.js',
  './go-version',
  './run.js',
  './exec-main.js',
  './package.json',
  './package-lock.json',
  './plugins/com.adopisoft.basic-flare-theme'
];

async function copyDevkitFiles() {
  for (const file of DEVKIT_FILES) {
    console.log(`Copying ${file} -> ${RELEASE_DIR}/${file}...`);
    await fs.copy(file, `${RELEASE_DIR}/${file}`);
  }
}

async function defaultConfigs() {
  await fs.mkdir(`${RELEASE_DIR}/config`, { recursive: true });
  await fs.copy('./config/.defaults/', `${RELEASE_DIR}/config/.defaults/`);

  // Generate application.json with random secret
  const secret = (await execAsync('openssl rand -hex 16')).stdout.trim();
  await fs.writeFile(
    `${RELEASE_DIR}/config/application.json`,
    `{
    "secret": "${secret}",
    "lang": "en"
}`
  );

  console.log('Created config/application.json:');
  console.log(
    await fs.readFile(`${RELEASE_DIR}/config/application.json`, 'utf-8')
  );
}

async function copyExtrasFiles() {
  console.log('Copying devkit files...');
  await fs.copy('./devkit-extras/', RELEASE_DIR);
}

async function prepare() {
  try {
    await fs.rm(RELEASE_DIR, { recursive: true });
  } catch (err) {
    // Ignore error if directory doesn't exist
  }

  await fs.mkdir(`${RELEASE_DIR}/plugins`, { recursive: true });

  try {
    await execAsync(`docker rm -f "${TMP_CONTAINER}"`);
  } catch (err) {
    // Ignore error if container doesn't exist
  }
}

async function buildDockerImage() {
  return new Promise((resolve, reject) => {
    // await execAsync(`docker build --progress=plain -t "${DOCKER_IMAGE}" .`);
    spawn('docker', ['build', '--progress=plain', '-t', DOCKER_IMAGE, '.'], {
      stdio: 'inherit'
    })
      .on('exit', (code) => {
        if (code === 0) {
          resolve();
        } else {
          reject(new Error('Docker build failed'));
        }
      })
      .on('error', reject);
  });
}

async function copyCoreSo() {
  const containerId = (
    await execAsync(`docker create --name ${TMP_CONTAINER} ${DOCKER_IMAGE}`)
  ).stdout.trim();
  await execAsync(`docker cp ${containerId}:${CORE_SO} ${OUTFILE}`);
  await execAsync(`docker rm ${TMP_CONTAINER}`);
}

async function main() {
  await prepare();
  await buildDockerImage();
  await copyCoreSo();
  await defaultConfigs();
  await copyDevkitFiles();
  await copyExtrasFiles();
}

main().catch((error) => console.error(error));
