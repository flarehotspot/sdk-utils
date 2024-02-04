#!/usr/bin/env node

const execAsync = require('./exec-async');
const fs = require('fs-extra');
const path = require('path');

module.exports = async (pluginDir, tags) => {
  const plugin = path.basename(pluginDir);

  if (await fs.exists(pluginDir)) {
    console.log(`Building plugin ${plugin}...`);

    const buildtags = tags ? `-tags="${tags}"` : '';

    try {
      await execAsync(
        `cd ${pluginDir} && go build -buildmode=plugin ${buildtags} -ldflags="-s -w" -trimpath -o plugin.so ./main.go`
      );
      console.log(`Done building plugin: ${plugin}`);
    } catch (error) {
      console.error(`Error building plugin ${plugin}: `, error);
    }
  }
};
