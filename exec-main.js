const { spawn } = require('child_process');
const path = require('path');
const isWin = process.platform == 'win32';
const mainFile = isWin ? 'main.exe' : 'main.app';

console.log(`Executing: ${path.join('./main', mainFile)}`);

spawn(
  path.join(process.cwd(), 'main', mainFile),
  { stdio: 'inherit' },
  (err, stdout) => {
    if (err) {
      console.log(err);
    } else {
      console.log(stdout);
    }
  }
);
