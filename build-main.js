#!/usr/bin/env node

const { execSync } = require('child_process');
const path = require('path');

const WORKDIR = process.cwd();
const mainDir = path.join(WORKDIR, 'main');

try {
  console.log(`Building main.app in ${mainDir}...`);
  execSync(
    `cd ${mainDir} && go build -ldflags="-s -w" -tags="dev" -trimpath -o main.app main.go`
  );
  console.log('Build successful.');
} catch (error) {
  console.error(`Error building main.app: ${error.stderr.toString().trim()}`);
}
