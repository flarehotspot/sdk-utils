#!/usr/bin/env node

process.env.NODE_ENV = 'development';

const path = require('path');
const execAsync = require('./build/exec-async.js');
const buildArgs = require('./build/build-args.js');
const mainGo = path.join(__dirname, 'main/main_mono.go');

(async () => {
  await require('./build/update-version.js');
  await require('./build/clean-up.js');
  await require('./build/make-mono.js');
  await require('./build/make-go.work.js');
  await execAsync(`go run ${buildArgs} ${mainGo}`);
})();
