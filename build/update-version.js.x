#!/usr/bin/env node

const fs = require('fs-extra');
const path = require('path');
const PACKAGE_JSON_PATH = path.join(__dirname, '../package.json');
const coreVersion = require('./core-version');

async function main() {
  const CORE_VERSION = await coreVersion();
  console.log(`Upading package.json version: ${CORE_VERSION}`);
  const packageJson = JSON.parse(await fs.readFile(PACKAGE_JSON_PATH, 'utf8'));
  packageJson.version = await coreVersion();
  await fs.writeFile(PACKAGE_JSON_PATH, JSON.stringify(packageJson, null, 2));
}

module.exports = main();
