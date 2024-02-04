#!/usr/bin/env node

const { execSync } = require('child_process');
const path = require('path');
const mainPath = require('./main-path');
const mainDir = path.dirname(mainPath);

module.exports = (async () => {
  try {
    console.log(`Building ${mainPath}...`);
    execSync(
      `cd ${mainDir} && go build -ldflags="-s -w" -tags="dev" -trimpath -o ${mainPath} main.go`
    );
    console.log(`Successfully built ${mainPath}!`);
  } catch (error) {
    console.error(`Error building main.app: `, error);
  }
})();
