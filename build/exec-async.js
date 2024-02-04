#!/usr/bin/env node

const { exec } = require('child_process');

module.exports = async (cmd) => {
  console.log(`Executing: ${cmd}`)
  return await new Promise((resolve, reject) => {
    const proc = exec(cmd, (err, stdout, _) => {
      if (err) {
        reject(err);
        return;
      }
      resolve(stdout);
    });

    proc.stdout.pipe(process.stdout);
    proc.stderr.pipe(process.stderr);
  });
};
