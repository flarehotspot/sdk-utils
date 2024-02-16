#!/usr/bin/env node

const path = require('path');
const fs = require('fs-extra');
const execAsync = require('./exec-async');
const SECRET = process.env.SECRET;
const SYSTEM_DIR = path.join(__dirname, '../system');

const plugins = ['com.flarego.default-theme'];

let systemPlugins = plugins.map((p) => {
  if (SECRET) {
    return `https://oauth2:${SECRET}@github.com/flarehotspot/${p}.git`;
  } else {
    return `git@github.com:flarehotspot/${p}.git`;
  }
});

async function main() {
  await fs.emptyDir(SYSTEM_DIR);
  await Promise.all(
    systemPlugins.map((p) => {
      return execAsync(`git clone ${p}`, { cwd: SYSTEM_DIR });
    })
  );
}

module.exports = main();
