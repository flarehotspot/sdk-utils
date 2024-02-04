#!/usr/bin/env node

(async () => {
  await require('./build/make-go.work.js');
  await require('./build/build-plugins.js');
  await require('./build/exec-main.js');
})();
