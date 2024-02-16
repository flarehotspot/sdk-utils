#!/usr/bin/env node

process.env.NODE_ENV = 'production';
process.env.DEVKIT_BUILD = 'true';

const fs = require('fs-extra');
const path = require('path');
const execAsync = require('./exec-async');
const goEnv = require('./go-env');
const coreVersion = require('./core-version');

async function main() {
  const { GOARCH } = await goEnv();
  const CORE_VERSION = await coreVersion();
  const ROOT_DIR = path.join(__dirname, '..');
  const RELEASE_DIR = path.join(
    ROOT_DIR,
    'devkit-release',
    `devkit-${CORE_VERSION}-${GOARCH}`
  );
  const DEVKIT_FILES = [
    '../main/go.mod',
    '../main/main.app',
    '../core/plugin.so',
    '../core/go.mod',
    '../core/go.sum',
    '../core/plugin.json',
    '../core/sdk',
    '../core/resources',
    '../core/go-version',
    '../build/build-args.js',
    '../build/exec-async.js',
    '../build/main-path.js',
    '../build/install-go.js',
    '../build/make-go.work.js',
    '../build/build-plugin.js',
    '../build/build-plugins.js',
    '../build/exec-main.js',
    '../build/go-env.js',
    '../build/setup_nodejs_16.x',
    '../run.js',
    '../package.json',
    '../package-lock.json',
    '../system'
  ];

  async function prepare() {
    await require('./clean-up.js');
    try {
      await fs.rm(RELEASE_DIR, { recursive: true });
    } catch (err) {
      // Ignore error if directory doesn't exist
    }

    await fs.mkdir(`${RELEASE_DIR}/plugins`, { recursive: true });
    await fs.mkdir(`${RELEASE_DIR}/system`, { recursive: true });
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

  async function copyConfigs() {
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

  await prepare();
  await require('./update-version.js');
  await require('./make-go.work.js');
  await require('./build-main.js');
  await require('./build-core.js');
  await copyConfigs();
  await copyDevkitFiles();
  await copyExtrasFiles();
}

module.exports = main();
