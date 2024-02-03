#!/usr/bin/env node

const { execSync } = require('child_process');
const fs = require('fs');
const path = require('path');
const isDocker = fs.existsSync('/.dockerenv');

const WORKDIR = process.cwd();

console.log(`Using go from ${execSync('which go').toString().trim()}...`);

const vendorDir = path.join(WORKDIR, 'vendor');

fs.readdirSync(vendorDir).forEach((plugin) => {
  const pluginDir = path.join(WORKDIR, 'plugins', plugin);
  const vendorPluginDir = path.join(WORKDIR, 'vendor', plugin);

  if (fs.existsSync(pluginDir) && fs.existsSync(vendorPluginDir) && isDocker) {
    fs.rmSync(vendorPluginDir, { recursive: true });
  }

  if (fs.existsSync(pluginDir)) {
    console.log(`Building plugin ${plugin}...`);

    try {
      execSync(
        `cd ${pluginDir} && go build -buildmode=plugin -ldflags="-s -w" -trimpath -o plugin.so ./main.go`
      );
      fs.copyFileSync(
        path.join(pluginDir, 'main.go'),
        path.join(vendorPluginDir, 'main.go')
      );
      console.log(`Done building plugin: ${plugin}`);
    } catch (error) {
      console.error(`Error building plugin ${plugin}: `, error);
    }
  }
});
