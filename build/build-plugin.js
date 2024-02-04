#!/usr/bin/env node

const execAsync = require('./exec-async');
const fs = require('fs-extra');
const path = require('path');
const buildArgs = require('./build-args');

module.exports = async (pluginDir) => {
  const plugin = path.basename(pluginDir);

  if (await fs.exists(pluginDir)) {
    console.log(`Building plugin ${plugin}...`);

    await execAsync(
      `cd ${pluginDir} && go build -buildmode=plugin ${buildArgs} -o plugin.so ./main.go`
    );

    console.log(`Done building plugin: ${plugin}`);
  }
};
