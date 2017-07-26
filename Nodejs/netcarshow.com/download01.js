/**
 *下载 www.netcarshow.com 网站的汽车图片(1600)
 *按汽车分类下载并保存
 */

// 依赖模块
var fs = require('fs')
    , mkdirp = require('mkdirp')
    , request = require('request')
    , async = require('async')
    , cheerio = require('cheerio')
	, mongodb = require('mongodb')
	, eventproxy = require('eventproxy')

// 模块变量的初始化 
var MongoClient = mongodb.MongoClient
	, ep = new eventproxy()

//{{{
var url = 'https://www.netcarshow.com' // 目标网址
	, mongoUrl = 'mongodb://localhost:27017/netcarshow' //数据库配置 数据库类型：//数据库地址:端口号/数据库名
	, brandTable = 'car_brand' // 品牌表名
	, seriesTable = 'car_series' // 车系表名

// 本地存储目录
var inputs_dir = '' // 自定义下载的根目录
    , dir = inputs_dir || './Download';  // 如果未自定义则用当前目录下的 Download 作为根目录
//}}}


// 判断要保存图片的文件夹是否存在 否：创建目录
mkdir(dir, '保存的根目录创建成功！');

// 调用方法
getIndex(dir, url);
// 调用方法
getCar();

// 创建目录文件方法
function mkdir(path, message) {//{{{
    
    if(typeof path != 'string') return ;

    if(!fs.existsSync(path)){
        mkdirp(path, function (err) {
            if(!err && message){
				console.log('\n', path, message);
			}else if(err){
                console.log(err);
			}
        });
    }else{

		if(path === dir){
                (message ? console.log('\n', path, message) : console.log('\n', path, '下载根目录已创建！'));
		}
	}

}//}}}

// 发送首页请求
function getIndex(dir, url) {
	
	console.log('\n', '开始@首页-车辆品牌数据 > 爬取中......');

    var options = {
        url: url,
        headers: {
            'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36'
        }
    };

	var local_href  = url;


	// 首页数据爬取处理函数
    var getIndexInfos = function (error, response, body, result) {
        if (!error && response.statusCode == 200) {
            var $ = cheerio.load(body)
              	, cars = [];
            $('div.Ll').find('ul').each(function (i) {
                $('div.Ll').find('ul').eq(i).children('li').each(function (){
                    var title = $(this).children('a').attr('title');
                    var href = $(this).children('a').attr('href');

                    //创建车辆品牌目录
                    var newDir = dir + '/' + title;
                    mkdir(newDir);
					var carInfo = {};
					carInfo.name = title;
					carInfo.href = href;
					carInfo.hrefs = local_href + href; 
					carInfo.dir = newDir + '/';
					carInfo.dirs = __dirname + '/' + newDir.replace(/\.\//, '') + '/';
					cars.push(carInfo);
                });
            });
			
			var data = cars;

			// 写入'netcarshow'数据库的'car_brand'表
			MongoClient.connect(mongoUrl, function(err, db){

				var brand_col = db.collection(brandTable);
				brand_col.drop(); // 初始化'car_brand'表
				var series_col = db.collection(seriesTable);
				series_col.drop(); // 初始化'car_seies'表
				brand_col.insertMany(data, function(err, result){
					db.close();
				});
			});

        }
    }

    request(options, getIndexInfos);
	console.log('\n', '结束@首页-车辆品牌数据 > 已爬完，并写入car_brand表');
}




// 发送车辆品牌页请求
function getCar() {

	// 查找
	var findDocuments = function (db, callback){
		var brand_col = db.collection(brandTable);
		brand_col.find().toArray(function(err, items){
			items.forEach(function(item){
				var series_url = item.hrefs; 
				var brand_dir = item.dir;
				var brand_name = item.name;
				var root_url = url;
				// 调用请求车系页方法
				callback(series_url, brand_dir, brand_name, root_url);	

			});
		});
	}
	
	// 插入	
	var insertDocuments = function (db, data){
		var series_col = db.collection(seriesTable);
		series_col.insertMany(data, function(err, result){
		});
	}


	console.log('\n', '开始@车系页数据爬取......' , '\n');
			
	// 从'car_brand'表读数据
	MongoClient.connect(mongoUrl, function(err, db){
		findDocuments(db, function (url, dir, brand_name, root_url){
			var options = {
				url: url,
				headers: {
					'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36'
				}
			};

			var regExp = /\['(.*?)'\]/;

			// 获取车型页信息
			var getSeriesInfos = function (error, response, body) {
				if (!error && response.statusCode == 200) {
					var $ = cheerio.load(body);
					var seriess = [];
					$('div.Ll').find('ul').children('li').each(function () {
						var title = $(this).children('a').attr('title');
						var href = $(this).children('a').attr('href');
						var filename = $(this).children('a').attr('onmouseover');
						filename = regExp.exec(filename.toString());
						filename = filename[1];
						var newDir = dir + '/' + title;
						//创建车辆品牌目录
						mkdir(newDir);
						
						seriesInfo = {};
						seriesInfo.brand_name = brand_name;
						seriesInfo.series_name = title;
						seriesInfo.series_href = href;
						seriesInfo.series_hrefs = root_url + href;
						seriesInfo.series_dir = newDir;
						seriesInfo.series_dirs =  __dirname + '/' + newDir.replace(/\.\//, '') + '/';
						seriesInfo.series_img = root_url + '/' + filename;
						seriess.push(seriesInfo);	
					});

					console.log('车辆品牌：' + seriess[0].brand_name + ' 下的车系有：' + seriess.length + ' 个');
					// 插入到series_table表中
					insertDocuments(db, seriess);
				}
			}
			request(options, getSeriesInfos);
				
		});
	});
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

// mongdb操作

