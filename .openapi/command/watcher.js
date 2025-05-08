const chokidar = require('chokidar');
const exec = require('child_process').exec;

const watcher = chokidar.watch('/openapi', {
  persistent: true,
  usePolling: true,
  interval: 1000,
  ignored: /\.output\/*/,
});

const stdout = (error, stdout, stderr) => {
  if(stdout){
    console.log(stdout);
  }
  if(stderr){
    console.log('ERROR: ' + stderr);
  }
  if (error !== null) {
    console.log('Exec error: ' + error);
  }
}

exec('mkdir /api/openapi/');

const run = () => {
  console.log('start building with openapi.yaml');
  exec('swagger-cli bundle -o /openapi/.output/openapi.yaml -t yaml /openapi/openapi.yaml', stdout);
  exec('~/go/bin/oapi-codegen -config /openapi/config/go/config.yaml -templates /openapi/config/go/templates /openapi/.output/openapi.yaml', stdout);
//   exec('openapi-generator-cli generate -i /openapi/.output/openapi.yaml -g typescript-fetch -o /app/application/src/api.gen.d -t /openapi/config/typescript/templates', stdout);
//   exec('openapi-generator-cli generate -i /openapi/.output/openapi.yaml -g typescript-fetch -o /app/webmaster/api.gen.d -t /openapi/config/typescript/templates', stdout);
  console.log('finished building with openapi.yaml');
}

watcher.on('ready', function () {
  watcher.on('change', async () => {
    run()
  });
});
run()