#!/usr/bin/env node

(async () => {
  await require('./build/clean-up.js');
  await require('./build/build-mono.js');
  await require('./build/exec-main.js');
})();
