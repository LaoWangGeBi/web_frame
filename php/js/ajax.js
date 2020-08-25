
function AjaxGo(obj) {
    let _this = this;
    let _opts = _this.opts = {
        url:obj.url||'', // 地址
        type:obj.type||'POST', // 提交方式post get(默认位post)
        async:obj.async||true, // 是否异步(默认位异步)
        form:obj.form||true, // 是否使用form提交(默认位为否)
        data:obj.data||null, // 要提交的数据
		datatype:obj.datatype||'JSON', // 要提交的数据格式是否为json(默认位为否)
		success:obj.success||function(e,data){}, // 请求完成
    }
    let _tmp;
    let _xhr; //创建请求对象对象
    if (window.XMLHttpRequest){// code for IE7+, Firefox, Chrome, Opera, Safari
        _xhr = new XMLHttpRequest();
    }else{// code for IE6, IE5
        _xhr = new ActiveXObject("Microsoft.XMLHTTP");
    }

    //  监控加载
    _xhr.process = function(e){
        console.log('process');
    }
    //  请求失败
    _xhr.onerror = function(e){

    }
    // 请求成功
    _xhr.onload = function(e){
        if(_xhr.status == 200){ //请求成功

        }else if(_xhr.status == 400){ //内部服务器错误
			console.error('内部服务器错误');
        }
        if(_xhr.readyState == 4){
            //console.log(JSON.parse(_xhr.responseText));
			_opts.success(_this,_xhr.responseText);
        }

    }
    //  状态监控
    _xhr.onreadystatechange = function(e){
        if(_xhr.readyState == 1){ //uninitialized - 还未开始载入

        }else if(_xhr.readyState == 2){ //loading - 载入中


        }else if(_xhr.readyState == 3){ //interactive - 已加载，文档与用户可以开始交互

        }else if(_xhr.readyState == 4){ //complete - 载入完成

        }
        //console.log('onreadystatechange');

    }

    if(_opts.type.toUpperCase() == 'POST'){ // POST
        if(_opts.data){
            _tmp = {};
            if(typeof(_opts.data) == 'string'){
                _opts.data = _opts.data.split('&');
            }
            for(v of _opts.data){
                v = v.split('=');
                _tmp[v[0]] = v[1];
            }
            _opts.data = JSON.stringify(_tmp);
            _tmp = null;
        }
    }

    if(_opts.type.toUpperCase() == 'GET'){
        _tmp = _opts.url.split('?');
        if(_tmp.length > 2){
            _tmp[1] = _tmp[1][_tmp[1].length-1]=='&'?_tmp[1].substr(0,_tmp[1].length-2):_tmp[1];
			_opts.url+= '?'+_tmp[1]
			_opts.url+= _opts.data?'&'+_opts.data:'';
        }else{
			_opts.url+= _opts.data?'?'+_opts.data:'';
        }
        _tmp = _opts.data = null;
    }

    _xhr.open(_opts.type,_opts.url,_opts.async);
    if(_opts.form){
        _xhr.setRequestHeader("Content-Type","application/json");
        //_xhr.setRequestHeader("Content-Type","application/x-www-form-urlencoded;charset=UTF-8");
        // if(_opts.data && !(_opts.data instanceof FormData)){
        //     _tmp = new FormData();
        //     if(typeof(_opts.data) == 'string'){
        //         _opts.data = _opts.data.split('&');
        //     }
        //     for(v of _opts.data){
        //         v = v.split('=');
        //         _tmp.append(v[0],v[1]);
        //     }
        //     _opts.data = _tmp;
        // }
    }else{
        _xhr.setRequestHeader("Content-Type","application/json");
    }
    _xhr.send(_opts.data);

}