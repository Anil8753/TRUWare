const path = require('path');
const cp = require('child_process');

const cwd = __dirname;
const networkPath = path.join(cwd, '../', 'network');

function ab2str(buf) {
    return String.fromCharCode.apply(null, new Uint8Array(buf));
}

process.chdir(networkPath)

if (process.argv[1] == 'up') {
    const buf = cp.execSync('"./minifab" up');
    console.log(ab2str(buf));
} else {
    const buf = cp.execSync('"./minifab" down');
    console.log(ab2str(buf));   
}
