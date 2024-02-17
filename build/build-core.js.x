#!/usr/bin/env node

const path = require('path');
const buildPlugin = require('./build-plugin');
const CORE_DIR = path.join(__dirname, '../core');

module.exports = buildPlugin(CORE_DIR, "dev");
