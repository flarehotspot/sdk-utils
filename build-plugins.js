#!/usr/bin/env node

const { execSync, spawn } = require('child_process');
const fs = require('fs');
const path = require('path');

const WORKDIR = process.cwd();

console.log(`Using go from ${execSync('which go').toString().trim()}...`);

const pluginsDir = path.join(WORKDIR, 'plugins');

(async () => {
  const dirs = await fs.promises.readdir(pluginsDir);

  dirs.forEach(async (plugin) => {
    const pluginDir = path.join(WORKDIR, 'plugins', plugin);
    if (fs.existsSync(pluginDir)) {
      console.log(`Building plugin ${plugin}...`);
      try {
        await new Promise((resolve, reject) => {
          spawn(
            'go',
            [
              'build',
              '-buildmode',
              'plugin',
              '-ldflags',
              '-s -w',
              '-trimpath',
              '-o',
              'plugin.so',
              './main.go'
            ],
            { stdio: 'inherit', cwd: pluginDir }
          )
            .on('close', (code) => {
              if (code === 0) {
                resolve();
              } else {
                reject(new Error(`Failed to build plugin ${plugin}`));
              }
            })
            .on('error', (error) => {
              reject(error);
            });
        });
        console.log(`Done building plugin: ${plugin}`);
      } catch (error) {
        console.error(`Error building plugin ${plugin}: `, error);
      }
    }
  });
})();
