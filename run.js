#!/usr/bin/env node

const fs = require('fs');
const path = require('path');

const WORKDIR = process.cwd();

// Clean up
const cleanupDirs = ['.tmp', '.cache/assets', 'public'];
cleanupDirs.forEach((dir) => {
  const fullPath = path.join(WORKDIR, dir);
  if (fs.existsSync(fullPath)) {
    fs.rmdirSync(fullPath, { recursive: true });
  }
});

// Remove .app and .so files
const removeFilesWithExtension = (dir, extension) => {
  const files = fs.readdirSync(dir);
  files.forEach((file) => {
    if (file.endsWith(extension)) {
      const filePath = path.join(dir, file);
      fs.unlinkSync(filePath);
    }
  });
};

removeFilesWithExtension(WORKDIR, '.app');
removeFilesWithExtension(path.join(WORKDIR, 'plugins'), '.so');

// Build .so files and run
try {
  console.log('Running make-go.work.js...');
  require('./make-go.work.js');
  console.log('Running build-main.js...');
  require('./build-main.js');
  console.log('Running build-plugins.js...');
  require('./build-plugins.js');
  console.log('Running exec-main.js...');
  require('./exec-main.js');
} catch (error) {
  console.error(`Error executing scripts: ${error.message}`);
}
