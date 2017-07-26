'use strict'

/*
 *  mongodb 数据库模型
 */


// 引入mongodb类库
var mongo = require('mongodb')
	, co = require('co');

// mongodb配置
var  MongoAddress = 'mongodb://localhost:27017/netcarshow'
	, table = 'netcarshow'
	, MongoClient = mongo.MongoClient;

MongoClient.connect(MongoAddress, function(err, db){
	
	var collection = db.collection(table);

	collection.find().limit(2).toArray(function(err, docs){

		console.log(docs);
		db.close();
	});
});


