var config = require('./config.json');
var readFile = require('./readFile.js');
var http = require('http');
var { URL , Url} = require('url');

http.createServer((request, response) => {
  if(request.url == '/favicon.ico'){
    return false;
  }
  //获取返回的url对象的query属性值
  var myUrl = new URL(request.url, config.hostname);
  var searchParams = myUrl.searchParams;
  console.log('-----------');
  console.log(myUrl);
  console.log('-----------');

  // var data = readFile.ReadFileSync('config.json');
  // // 发送相应数据 "Hello Word"
  // response.end(data.toString());
  // readFile.ReadFile('config.json',function(err,data){
  //   if (err) {
  //     response.write(err.toString());
  //   } else {
  //     response.write(data.toString());
  //   }
  // });

  // 发送 HTTP 头部
  // HTTP 状态值 200 : ok
  // 内容类型：text/plain
  response.writeHead(200, {'Content-Type':'text/plain'});
  response.write(`\nmethod=${request.method}\n`);
  response.write(`url=${request.url}\n`);
  response.write(`id=${searchParams.get('id')}\n`);
  response.write(`name=${searchParams.get('id')}\n`);
  response.end();
}).listen(config.port);

console.log(`Server running at http://127.0.0.1:${config.port}/`);