#!/usr/bin/env node

const execAsync = require('./exec-async');
const fs = require('fs-extra');
const path = require('path');
const buildPlugin = require('./build-plugin');

const ROOT_DIR = path.join(__dirname, '..');

module.exports = (async () => {
  console.log(`Using go from ${(await execAsync('which go')).trim()}...`);

  const pluginsDir = path.join(ROOT_DIR, 'plugins');
  const dirs = await fs.promises.readdir(pluginsDir);

  for (let plugin of dirs) {
    const pluginDir = path.join(pluginsDir, plugin);
    await buildPlugin(pluginDir, "dev");
  }
})();
