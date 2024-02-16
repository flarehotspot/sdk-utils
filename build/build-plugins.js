#!/usr/bin/env node

const execAsync = require('./exec-async');
const fs = require('fs-extra');
const path = require('path');
const buildPlugin = require('./build-plugin');

const ROOT_DIR = path.join(__dirname, '..');

module.exports = (async () => {
  console.log(`Using go from ${(await execAsync('which go')).trim()}...`);

  const pluginsDir = [
    path.join(ROOT_DIR, 'system'),
    path.join(ROOT_DIR, 'plugins')
  ];

  let dirs = [];
  for (let dir of pluginsDir) {
    const entries = await fs.readdir(dir);
    for (let entry of entries) {
      const entryPath = path.join(dir, entry);
      const stat = await fs.stat(entryPath);
      if (stat.isDirectory()) {
        dirs.push(entryPath);
      }
    }
  }

  for (let pluginDir of dirs) {
    await buildPlugin(pluginDir, 'dev');
  }

})();
