<?php
header("Access-Control-Allow-Origin: *");
header("content-type:text/html;charset=utf-8");
$request_type = $_SERVER['REQUEST_METHOD'];

if($_SERVER['REQUEST_URI'] == '/'){
    header("Location: index.html");
    exit();
}

//$filepath = dirname(__FILE__); //想要读取的目录
$filepath = "D:\Work"; //想要读取的目录
//$filepath = dirname(__FILE__);
readFileRecursive($filepath);


function printFile($filepath)
{
	//substr(string,start,length)函数返回字符串的一部分；start规定在字符串的何处开始 ；length规定要返回的字符串长度。默认是直到字符串的结尾。
	//strripos(string,find,start)查找 "php" 在字符串中最后一次出现的位置； find为规定要查找的字符；start可选。规定开始搜索的位置

	//读取文件后缀名
	//$filetype = substr ( $filename, strripos ( $filename, "." ) + 1 );
	//判断是不是以txt结尾并且是文件
	#if ($filetype == "txt" && is_file ( $filepath . "/" . $filename ))
	if ( is_file ( $filepath))
	{
		$filename=iconv("gb2312","utf-8",$filepath);
		echo $filename."内容如下:"."<br/>";
		$fp = fopen ( $filepath, "r" );//打开文件
		#while (! feof ( $f )) //一直输出直到文件结尾
		$i = 1;
		while ($i < 10)
		{
			$line = fgets ( $fp );
			echo $line."<br/>";
			$i = $i +1;
		}
		fclose($fp);
	}
}

function readFileRecursive($filepath)
{
    $arr = array();
    if (is_dir ( $filepath )) //判断是不是目录
	{
        $dirhandle = opendir ( $filepath );//打开文件夹的句柄
		if ($dirhandle)
		{
			//判断是不是有子文件或者文件夹
			while ( ($filename = readdir ( $dirhandle ))!= false )
			{
                if($filename=="." || $filename==".."){
                    continue;
                }
                if(is_dir($filepath.'/'.$filename)){
                    //echo $filename;
                    $arr[] = array('url'=>$filepath."\\".$filename,'name'=>$filename.'');
                }
			}
			closedir ( $dirhandle );
		}
    }
    $result = json_encode(array("code"=>1,"data"=>$arr));
    echo $result;
}

?>
