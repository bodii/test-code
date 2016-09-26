/**
 *下载 www.netcarshow.com 网站的汽车图片(1600)
 *按汽车分类下载并保存
 */

    // 依赖模块
var fs = require('fs')
    , mkdirp = require('mkdirp')
    , st = require('superagent')
    , request = require('request')
    , async = require('async')
    , cheerio = require('cheerio')
    , eventproxy = require('eventproxy');

var ep = new eventproxy();

// 目标网址
var url = 'https://www.netcarshow.com';

// 本地存储目录
var inputs_dir = '' // 自定义下载的根目录
    , carName = []  // 车辆名称
    , carHref = []  //
    , modelName = []  // 车款名称
    , modelHref = []
    , dir = inputs_dir || './Download';  // 如果未自定义则用当前目录下的 Download 作为根目录



// 判断要保存图片的文件夹是否存在 否：创建目录
mkdir(dir, '保存的根目录创建成功！');

// 调用方法
//getIndex(dir, url);

//console.log(carName);



// 创建目录文件方法
function mkdir(dir, message) {
    var dir = dir || this.dir;
    if(typeof dir != 'string') return ;

    if(!fs.existsSync(dir)){
        mkdirp(dir, function (err) {
            err ?
                console.log(err)
                :
                (message ? console.log(dir, message) : false);
        });
    }
}

// 发送首页请求
function getIndex(dir, url) {

    var options = {
        url: url,
        headers: {
            'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36'
        }
    };

    var aa = {};
    function callback(error, response, body, result) {
        if (!error && response.statusCode == 200) {
            var $ = cheerio.load(body)
              	, names = []
                , hrefs = [];
            $('div.Ll').find('ul').each(function (i) {
                $('div.Ll').find('ul').eq(i).children('li').each(function (){
                    var title = $(this).children('a').attr('title');
                    var href = $(this).children('a').attr('href');
                    names.push(title);
                    hrefs.push(href);

                    //创建车辆品牌目录
                    var newDir = dir + '/' + title;
                    //mkdir(newDir);
                });
            });
            aa.name = names;
            aa.href = hrefs;
        }
    }


    var a = request(options, callback);
    console.log(a)

}






// 发送车辆品牌页请求
function getCar(url, root_url, dir) {
    var options = {
        url: url,
        headers: {
            'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36'
        }
    };

    var regExp = /\['(.*?)'\]/;

    function callback(error, response, body) {
        if (!error && response.statusCode == 200) {
            var $ = cheerio.load(body);
            $('div.Ll').find('ul').children('li').each(function () {
                var title = $(this).children('a').attr('title');
                var href = $(this).children('a').attr('href');
                var filename = $(this).children('a').attr('onmouseover');
                filename = regExp.exec(filename.toString());
                filename = filename[1];
                var newDir = dir + '/' + title;
                //创建车辆品牌目录
                mkdir(newDir);
                //请求车款页
                getModel(root_url + href, root_url, newDir, filename);
            });
        }
    }

    request(options, callback);
}

// 发送车款页请求
function getModel(url, dir, root_url, filename) {
    var options = {
        url: url,
        headers: {
            'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36'
        }
    };

    function callback(error, response, body) {
        if (!error && response.statusCode == 200) {
            var $ = cheerio.load(body);
            $('div#thpg1').find('div.Pi[itemprop=associatedMedia]').each(function (i) {
                // var title = $(this).find('.tss').children('a').attr('title');
                // var href = $(this).find('.tss').children('a[itemprop=image]').attr('href');
                // var file_name = filename + '-1600-0' + (i +1) +'.jpg';
                //请求车款页
                console.log(file_name);
                //download(root_url + href, dir, file_name);
            });
        }
    }

    request(options, callback);
}

// 下载图片
function download(url, dir, filename) {
    writestream = fs.createWriteStream(dir + '/' + filename);

    http.get(url, function (res) {
        res.pipe(writestream);
    });

    writestream.on('finish', function () {
        console.log(filename);
    });
}
