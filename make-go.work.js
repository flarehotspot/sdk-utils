const fs = require('fs');
const path = require('path');

const WORKDIR = process.cwd();
const GO_VERSION = fs.readFileSync(path.join(WORKDIR, 'go-version'), 'utf-8').trim();

// Get the first two numbers of the version
const GO_SHORT_VERSION = GO_VERSION.split('.').slice(0, 2).join('.');

let GOWORK = `go ${GO_SHORT_VERSION}
use (
    ./core
    ./main`;

if (fs.existsSync(path.join(WORKDIR, 'plugins'))) {
    fs.readdirSync(path.join(WORKDIR, 'plugins')).forEach((dir) => {
        const d = path.join(WORKDIR, 'plugins', dir);
        if (fs.statSync(d).isDirectory()) {
            if (!fs.existsSync(path.join(d, 'package.yml'))) {
                GOWORK += `
                ${d}`;
            }
        }
    });
}

GOWORK += `
)`;

fs.writeFileSync(path.join(WORKDIR, 'go.work'), GOWORK);

console.log('go.work file created.');
