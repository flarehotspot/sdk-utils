#!/usr/bin/env node

const path = require('path');
const execAsync = require('./exec-async');
const buildArgs = require('./build-args');
const mainPath = require('./main-path');
const goEnv = require('./go-env');
const mainDir = path.dirname(mainPath);

module.exports = (async () => {
  const { GO_BIN } = await goEnv();
  try {
    console.log(`Building ${mainPath}...`);
    await execAsync(
      `cd ${mainDir} && ${GO_BIN} build ${buildArgs} -o ${mainPath} main.go`
    );
    console.log(`Successfully built ${mainPath}.`);
  } catch (error) {
    console.error(`Error building main.app: `, error);
  }
})();
