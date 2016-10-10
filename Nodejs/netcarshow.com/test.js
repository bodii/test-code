var mongo = require('./mongo_model')

//mongo.insert({a:1});
var a = mongo.find({a:{$exists:true}});
console.log(a);
