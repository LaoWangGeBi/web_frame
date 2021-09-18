var fs = require('fs');

// 非异步
function ReadFileSync(src){
  return fs.readFileSync(src);
}

// 非异步
function ReadFile(src,callback){
  fs.readFile(src,function(err,data){
    callback(err,data)
  });
}

module.exports = {
  ReadFileSync ,
  ReadFile
}