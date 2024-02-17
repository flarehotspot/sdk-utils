#!/usr/bin/env node

const execAsync = require('./exec-async');
const fs = require('fs-extra');
const path = require('path');
const buildArgs = require('./build-args');
const goEnv = require('./go-env');

module.exports = async (pluginDir) => {
  const plugin = path.basename(pluginDir);
  const { GO_BIN } = await goEnv();

  if (await fs.exists(pluginDir)) {
    console.log(`Building plugin ${plugin}...`);

    await execAsync(
      `cd ${pluginDir} && ${GO_BIN} build -buildmode=plugin ${buildArgs} -o plugin.so ./main.go`
    );

    console.log(`Done building plugin: ${plugin}`);
  }
};
