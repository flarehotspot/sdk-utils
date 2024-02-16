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

  const PLUGIN_PATHS = ['system', 'plugins'];

  let GOWORK = `go ${GO_SHORT_VERSION}
use (
    ./core
    ./main`;

  for (const dir of PLUGIN_PATHS) {
    const searchPath = `./${dir}`;
    if (fs.existsSync(searchPath)) {
      fs.readdirSync(searchPath).forEach((pluginDir) => {
        const d = path.join(searchPath, pluginDir);
        if (fs.statSync(d).isDirectory()) {
          if (fs.existsSync(path.join(d, 'plugin.json'))) {
            const basename = path.basename(d);
            GOWORK += `\n\t./${path.join(searchPath, pluginDir)}`;
          }
        }
      });
    }
  }

  GOWORK += `
)`;

  fs.writeFileSync(path.join(ROOT_DIR, 'go.work'), GOWORK);

  console.log('go.work file created.');
})();
