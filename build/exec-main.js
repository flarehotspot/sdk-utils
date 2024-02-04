#!/usr/bin/env node

const { exec } = require('child_process');
const mainPath = require('./main-path');

module.exports = (async () => {
  const proc = exec(mainPath);
  proc.stdout.pipe(process.stdout);
  proc.stderr.pipe(process.stderr);
})();
