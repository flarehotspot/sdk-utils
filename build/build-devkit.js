#!/usr/bin/env node

const fs = require('fs-extra');
const path = require('path');
const execAsync = require('./exec-async');

const CORE_VERSION = require('../core/plugin.json').version;
const ROOT_DIR = path.join(__dirname, '..');
const DOCKER_IMAGE = 'devkit:latest';
const TMP_CONTAINER = 'devkit-tmp';
const CORE_SO = '/plugin.so';
const RELEASE_DIR = path.join(
  ROOT_DIR,
  'devkit-release',
  '/devkit-' + CORE_VERSION
);
const OUTFILE = path.join(ROOT_DIR, 'core/plugin.so');
const DOCKER_FILE = path.join(__dirname, 'Dockerfile');
const DEVKIT_FILES = [
  '../main/go.mod',
  '../main/main.app',
  '../core/go.mod',
  '../core/go.sum',
  '../core/sdk',
  '../core/resources',
  '../core/plugin.json',
  '../core/plugin.so',
  '../core/go-version',
  '../build/build-args.js',
  '../build/exec-async.js',
  '../build/main-path.js',
  '../build/install-go.js',
  '../build/make-go.work.js',
  '../build/build-plugin.js',
  '../build/build-plugins.js',
  '../build/exec-main.js',
  '../run.js',
  '../package.json',
  '../package-lock.json',
  '../plugins/com.adopisoft.basic-flare-theme'
];

async function prepare() {
  await require('./clean-up.js');
  try {
    await fs.rm(RELEASE_DIR, { recursive: true });
  } catch (err) {
    // Ignore error if directory doesn't exist
  }

  await fs.mkdir(`${RELEASE_DIR}/plugins`, { recursive: true });

  try {
    await execAsync(`docker rm ${TMP_CONTAINER}`);
  } catch (err) {
    // Ignore error if container doesn't exist
  }
}

async function copyDevkitFiles() {
  for (const file of DEVKIT_FILES) {
    // remove traling "../"
    const src = path.join(ROOT_DIR, 'build', file);
    const dst = path.join(RELEASE_DIR, file.replace(/^\.\.\//, ''));
    console.log(`Copying ${file} -> ${dst}...`);
    await fs.copy(src, dst);
  }
}

async function defaultConfigs() {
  await fs.mkdir(`${RELEASE_DIR}/config`, { recursive: true });
  await fs.copy(
    path.join(__dirname, '../config/.defaults'),
    `${RELEASE_DIR}/config/.defaults/`
  );

  // Generate application.json with random secret
  const secret = await execAsync('openssl rand -hex 16').then((stdout) =>
    stdout.trim()
  );
  const appcfg = { lang: 'en', secret };
  const appcfgPath = path.join(RELEASE_DIR, 'config/application.json');
  await fs.writeFile(appcfgPath, JSON.stringify(appcfg, null, 2));
}

async function copyExtrasFiles() {
  console.log('Copying devkit-extra files...');
  await fs.copy(path.join(__dirname, 'devkit-extras'), RELEASE_DIR);
}

async function buildCore() {
  await execAsync(
    `cd ${ROOT_DIR} && docker build --progress=plain -t ${DOCKER_IMAGE} -f ${DOCKER_FILE} .`
  );
}

async function copyCoreSo() {
  const containerId = await execAsync(
    `docker create --name ${TMP_CONTAINER} ${DOCKER_IMAGE}`
  ).then((stdout) => stdout.trim());
  await execAsync(`docker cp ${containerId}:${CORE_SO} ${OUTFILE}`);
  await execAsync(`docker rm ${TMP_CONTAINER}`);
}

async function buildMain() {
  await require('./build-main.js');
}

async function zipDevkit() {
  const zipFile = RELEASE_DIR + '.zip';
  console.log(`Zipping ${RELEASE_DIR} -> ${zipFile}...`);
  await execAsync(`cd ${RELEASE_DIR} && zip -r ${zipFile} .`);
}

async function main() {
  await prepare();
  await buildCore();
  await copyCoreSo();
  await defaultConfigs();
  await buildMain();
  await copyDevkitFiles();
  await copyExtrasFiles();
  await zipDevkit();
}

module.exports = main();
