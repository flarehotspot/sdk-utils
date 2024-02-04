#!/usr/bin/env node

const path = require('path');
const isWin = process.platform == 'win32';
const mainFile = isWin ? 'main.exe' : 'main.app';

module.exports = path.join(__dirname, '../main', mainFile)
