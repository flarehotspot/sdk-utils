#!/usr/bin/env node

const fs = require('fs-extra');
const path = require('path');
const ROOT_DIR = path.join(__dirname, '..');
const cleanupDirs = ['.tmp', '.cache/assets', 'public'];

module.exports = (async () => {
  for (let dir of cleanupDirs) {
    dir = path.join(ROOT_DIR, dir);
    if (fs.existsSync(dir)) {
      console.log(`Cleaning up ${dir}...`);
      fs.rmSync(dir, { recursive: true });
    }
  }

  // Remove .app and .so files
  console.log('Removing .app and .so files...');

  const removeFilesWithExtension = (dir, extension) => {
    if (!fs.existsSync(dir)) return;
    const entries = fs.readdirSync(dir);
    entries.forEach((entry) => {
      const fullpath = path.join(dir, entry);
      const stat = fs.statSync(fullpath);
      if (stat.isFile() && entry.endsWith(extension)) {
        fs.rm(fullpath);
      }
      if (stat.isDirectory()) {
        removeFilesWithExtension(fullpath, extension);
      }
    });
  };

  removeFilesWithExtension(path.join(ROOT_DIR, 'core'), '.so');
  removeFilesWithExtension(path.join(ROOT_DIR, 'plugins'), '.so');
  removeFilesWithExtension(path.join(ROOT_DIR, 'main'), '.app');
})();
