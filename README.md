PPGo_Api_Demo_Gin
====
什么东西？What?
----
本项目是一个关于gin框架的练习代码。 采用go+mysql实现一个API应用实例。 练习内容包括： 中间件使用、RESTFul路由创建、mysql数据库使用、静态资源加载、页面创建、配置文件读取等等。<br />
Gin是用Golang实现的一种Web框架。基于httprouter，它提供了类似martini但更好性能(路由性能约快40倍)的API服务

有什么价值？
----
1、go get github.com/george518/PPGo_Api_Demo_Gin    
2、创建mysql数据库，并将ppgo_api_demo_gin.sql导入    
3、修改config/config.ini 配置数据库    
4、运行 go build    
5、运行 ./run.sh start|stop    
帮助访问：http://your_host:8000  

安装方法    
----
作用一：可以作为Gin框架的练手项目    
作用二：可以用来快速搭建Gin框架的API应用    
如果感觉项目不错，请赐予star, :)        

API调用示例
----
新增会员 POST<br />
curl -X POST http://127.0.0.1:8000/v0/member?sign=ee14b34513697cd27e0e83e46b084580&ts=1508304821&app_key=1001&method=POST -d "login_name=george518&password=123456"<br />
会员列表 GET<br />
curl -X GET http://127.0.0.1:8000/v0/member?page=1&page_size=4&sign=ee14b34513697cd27e0e83e46b084580&ts=1508304821&app_key=1001&method=GET<br />
会员详情 GET<br />
curl -X GET http://127.0.0.1:8000/v0/member/1?sign=ee14b34513697cd27e0e83e46b084580&ts=1508304821&app_key=1001&method=GET<br />
编辑会员 PUT<br />
curl -X PUT http://127.0.0.1:8000/v0/member/1?sign=ee14b34513697cd27e0e83e46b084580&ts=1508304821&app_key=1001&method=PUT -d "login_name=haodaquan&password=1234"<br />
删除会员 DELETE<br />
curl -X DELETE http://127.0.0.1:8000/v0/member/4?sign=ee14b34513697cd27e0e83e46b084580&ts=1508304821&app_key=1001&method=DELETE<br />

可以在middlewares/auth.go中修改超时验证。
` ``go
//时差两秒返回无权 注意修改
time_check, _ := strconv.Atoi(ts)
if (now - int64(time_check)) > 100000000000 {
    noAuth(c, "Time out")
    return
}
` ``
        
PHP调用接口代码示例
----
` ``php
<?php
/************************************************************
** @Description: PPGo_Api_Demo_Gin API调用DEMO for PHP
** @Author: haodaquan
** @Date:   2017-03-24 14:32:48
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-10-11 13:49:57
*************************************************************/
defined('APP_KEY') OR define("APP_KEY","1001");
defined('APP_SECRET') OR define("APP_SECRET","haodaquan");
$url = "http://127.0.0.1:8000/v0/member/3";
$params["ts"] = time();
$params["method"] = "GET";
$res = http($url,$params,"GET",[],10,true);
print_r($res);
/**
 * [http 调用接口函数]
 * @Date   2016-07-11
 * @Author GeorgeHao
 * @param  string       $url     [接口地址]
 * @param  array        $params  [数组]
 * @param  string       $method  [GET\POST\DELETE\PUT]
 * @param  array        $header  [HTTP头信息]
 * @param  integer      $timeout [超时时间]
 * @param  boolean      $sign    [是否加密]
 * @param  string       $app_key [应用号]
 * @param  string       $app_secret [应用密码]
 * @return [type]                [接口返回数据]
 */
function http($url, $params, $method = 'GET', $header = array(), $timeout = 10,$sign=false,$app_key='',$app_secret='')
{
    $opts = array(
        CURLOPT_TIMEOUT => $timeout,
        CURLOPT_RETURNTRANSFER => 1,
        CURLOPT_SSL_VERIFYPEER => false,
        CURLOPT_SSL_VERIFYHOST => false,
        CURLOPT_HTTPHEADER => $header
    );
    $app_key    = $app_key ? $app_key : APP_KEY;
    $app_secret = $app_secret ? $app_secret : APP_SECRET;
    if($sign)
    {
        $ts = time();
        $check =[
            "app_key=" . $app_key,
            "app_secret=" . $app_secret,
            "method=" . $method,
            "ts=" . $ts];
        sort($check); 
        $url .= '?sign='.md5(join("&", $check)).
                '&ts='.$ts.'&app_key='.$app_key
                .'&method='.$method;
        // print_r($url);
    }
    /* 根据请求类型设置特定参数 */
    switch (strtoupper($method)) {
        case 'GET':
            if($params)
            {
               $opts[CURLOPT_URL] = $url . '?' . http_build_query($params); 
            }else
            {
                $opts[CURLOPT_URL] = $url;
            }
            break;
        case 'POST':
            $params = http_build_query($params);
            $opts[CURLOPT_URL] = $url;
            $opts[CURLOPT_POST] = 1;
            $opts[CURLOPT_POSTFIELDS] = $params;
            break;
        case 'DELETE':
            $opts[CURLOPT_URL] = $url;
            $opts[CURLOPT_HTTPHEADER] = array("X-HTTP-Method-Override: DELETE");
            $opts[CURLOPT_CUSTOMREQUEST] = 'DELETE';
            $opts[CURLOPT_POSTFIELDS] = $params;
            break;
        case 'PUT':
            $opts[CURLOPT_URL] = $url;
            $opts[CURLOPT_POST] = 0;
            $opts[CURLOPT_CUSTOMREQUEST] = 'PUT';
            $opts[CURLOPT_POSTFIELDS] = http_build_query($params);
            break;
        case 'PATCH':
            $opts[CURLOPT_URL] = $url;
            $opts[CURLOPT_POST] = 0;
            $opts[CURLOPT_CUSTOMREQUEST] = 'PATCH';
            $opts[CURLOPT_POSTFIELDS] = http_build_query($params);
            break;
        default:
            throw new Exception('不支持的请求方式！');
    }
  
    /* 初始化并执行curl请求 */
    $ch     = curl_init();
    curl_setopt_array($ch, $opts);
    $data   = curl_exec($ch);
    $error  = curl_error($ch);
    return $data;
}

` ``           


联系我
----
qq:41352963


