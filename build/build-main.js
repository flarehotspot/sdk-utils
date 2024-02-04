#!/usr/bin/env node

const { execSync } = require('child_process');
const path = require('path');
const buildArgs = require('./build-args');
const mainPath = require('./main-path');
const mainDir = path.dirname(mainPath);

module.exports = (async () => {
  try {
    console.log(`Building ${mainPath}...`);
    execSync(
      `cd ${mainDir} && go build ${buildArgs} -o ${mainPath} main.go`
    );
    console.log(`Successfully built ${mainPath}.`);
  } catch (error) {
    console.error(`Error building main.app: `, error);
  }
})();
