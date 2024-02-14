#!/usr/bin/env node

const fs = require('fs');
const path = require('path');

module.exports = (async () => {
  const ROOT_DIR = path.join(__dirname, '..');
  const GO_VERSION = fs
    .readFileSync(path.join(ROOT_DIR, 'core/go-version'), 'utf-8')
    .trim();

  // Get the first two numbers of the version
  const GO_SHORT_VERSION = GO_VERSION.split('.').slice(0, 2).join('.');

  let GOWORK = `go ${GO_SHORT_VERSION}
use (
    ./core
    ./main`;

  if (fs.existsSync(path.join(ROOT_DIR, 'plugins'))) {
    fs.readdirSync(path.join(ROOT_DIR, 'plugins')).forEach((dir) => {
      const d = path.join(ROOT_DIR, 'plugins', dir);
      if (fs.statSync(d).isDirectory()) {
        if (fs.existsSync(path.join(d, 'plugin.json'))) {
          const basename = path.basename(d);
          GOWORK += `\n\t./plugins/${basename}`;
        }
      }
    });
  }

  GOWORK += `
)`;

  fs.writeFileSync(path.join(ROOT_DIR, 'go.work'), GOWORK);

  console.log('go.work file created.');
})();
