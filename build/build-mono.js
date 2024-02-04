#!/usr/bin/env node

const execAsync = require('./exec-async');
const path = require('path');
const buildArgs = require('./build-args.js');
const mainPath = require('./main-path');
const mainDir = path.dirname(mainPath);

const main = async function () {
  await require('./clean-up.js');
  await require('./make-mono.js');
  await require('./make-go.work.js');

  await execAsync(
    `cd ${mainDir} && go build ${buildArgs} -o ${mainPath} main_mono.go`
  );

  console.log(`Built app successfully: ${mainPath}`);
};

module.exports = main();
