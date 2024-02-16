#!/usr/bin/env node

const execAsync = require('./exec-async');
const mainPath = require('./main-path');

module.exports = (async () => {
  await execAsync(`chmod +x ${mainPath}`);
  await execAsync(mainPath);
})();
