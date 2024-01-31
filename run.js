const { execFile, spawn } = require('child_process');
const path = require('path');
const isWin = process.platform == 'win32';
const prod = process.env.NODE_ENV === 'production';
const buildTags = 'mono' + (!prod ? ' dev' : '');

if (isWin) {
  execFile(
    path.join(process.cwd(), 'main/main.exe'),
    { shell: true },
    (err, stdout) => {
      if (err) {
        console.log(err);
      } else {
        console.log(stdout);
      }
    }
  );
} else {
  spawn('go', ['run', '--tags', buildTags, 'main/main_mono.go'], {
    stdio: 'inherit'
  });
}
