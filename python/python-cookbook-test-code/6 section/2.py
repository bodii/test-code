#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第六章：数据编码和处理
Desc: 这一章主要讨论使用 Python 处理各种不同方式编码的数据，比如 CSV 文件，
JSON，XML 和二进制包装记录。和数据结构那一章不同的是，这章不会讨论特殊的
算法问题，而是关注于怎样获取和存储这些格式的数据。

Title;      读写 JSON 数据
Issue: 你想读写 JSON(JavaScript Object Notation) 编码格式的数据。
Answer: json 模块提供了一种很简单的方式来编码和解码 JSON 数据。其中两个主要的函
数是 json.dumps() 和 json.loads() ，要比其他序列化函数库如 pickle 的接口少得
多。
"""

import json

data = {
    'name' : 'ACME',
    'shares' : 100,
    'price' : 542.23
}

json_str  = json.dumps(data)
print(json_str)
# {"price": 542.23, "shares": 100, "name": "ACME"}

"""
下面演示如何将一个 JSON 编码的字符串转换回一个 Python 数据结构：
"""
data = json.loads(json_str)
print(data)
# {'name': 'ACME', 'shares': 100, 'price': 542.23}

"""
如果你要处理的是文件而不是字符串，你可以使用 json.dump() 和 json.load()
来编码和解码 JSON 数据。
"""
with open('../test file/data.json', 'w', newline='') as f:
    json.dump(data, f)

with open('../test file/data.json','r') as f:
    data = json.load(f)
    print(data)
# {'price': 542.23, 'name': 'ACME', 'shares': 100}

"""
JSON 编码支持的基本数据类型为 None ， bool ， int ， float 和 str ，以及包
含这些类型数据的 lists，tuples 和 dictionaries。对于 dictionaries，keys 需要是字符串
类型 (字典中任何非字符串类型的 key 在编码时会先转换为字符串)。为了遵循 JSON
规范，你应该只编码 Python 的 lists 和 dictionaries。而且，在 web 应用程序中，顶层
对象被编码为一个字典是一个标准做法。
JSON 编码的格式对于 Python 语法而已几乎是完全一样的，除了一些小的差异之
外。比如，True 会被映射为 true，False 被映射为 false，而 None 会被映射为 null。下
面是一个例子，演示了编码后的字符串效果：
"""
print(json.dumps(False))
d = {
    'a' : True,
    'b' : 'Hello',
    'c' : 'None'
}
print(json.dumps(d))
# {"a": true, "c": "None", "b": "Hello"}
# with open('../test file/data2.json', 'w') as f:
#     f.write(json.dumps(d))


'''
如果你试着去检查 JSON 解码后的数据，你通常很难通过简单的打印来确定它的
结构，特别是当数据的嵌套结构层次很深或者包含大量的字段时。为了解决这个问题，
可以考虑使用 pprint 模块的 pprint() 函数来代替普通的 print() 函数。它会按照
key 的字母顺序并以一种更加美观的方式输出。下面是一个演示如何漂亮的打印输出
Twitter 上搜索结果的例子
'''
from urllib.request import urlopen

#u = urlopen('http://search.twitter.com/search.json?q=python&rpp=5')
#resp = json.loads(u.read().decode('utf-8'))
# from ppint import pprint
# pprint(resp)

'''
{'completed_in': 0.074,
'max_id': 264043230692245504,
'max_id_str': '264043230692245504',
'next_page': '?page=2&max_id=264043230692245504&q=python&rpp=5',
'page': 1,
'query': 'python',
'refresh_url': '?since_id=264043230692245504&q=python',
'results': [{'created_at': 'Thu, 01 Nov 2012 16:36:26 +0000',
    'from_user': ...
    },
    {'created_at': 'Thu, 01 Nov 2012 16:36:14 +0000',
    'from_user': ...
    },
    {'created_at': 'Thu, 01 Nov 2012 16:36:13 +0000',
    'from_user': ...
    },
    {'created_at': 'Thu, 01 Nov 2012 16:36:07 +0000',
    'from_user': ...
    }
    {'created_at': 'Thu, 01 Nov 2012 16:36:04 +0000',
    'from_user': ...
}],
'results_per_page': 5,
'since_id': 0,
'since_id_str': '0'}
'''

"""
一般来讲，JSON 解码会根据提供的数据创建 dicts 或 lists。如果你想要创建其他
类型的对象，可以给 json.loads() 传递 object pairs hook 或 object hook 参数。例如，
下面是演示如何解码 JSON 数据并在一个 OrderedDict 中保留其顺序的例子
"""
s = '{"name" : "ACME","shares":50, "price":490.1}'
from collections import OrderedDict
data = json.loads(s, object_pairs_hook=OrderedDict)
print(data)
# OrderedDict([('name', 'ACME'), ('shares', 50), ('price', 490.1)])

"""
下面是如何将一个 JSON 字典转换为一个 Python 对象例子
"""
class JSONObject:
    def __init__(self, d):
        self.__dict__ = d

datas = json.loads(s, object_hook=JSONObject)
print(datas.name)
# ACME
print(datas.shares)
# 50
print(datas.price)
# 490.1

"""
最后一个例子中，JSON 解码后的字典作为一个单个参数传递给 init () 。然
后，你就可以随心所欲的使用它了，比如作为一个实例字典来直接使用它。
在编码 JSON 的时候，还有一些选项很有用。如果你想获得漂亮的格式化字符串后
输出，可以使用 json.dumps() 的 indent 参数。它会使得输出和 pprint() 函数效果类
似。
"""
print(json.dumps(data))
# {"name": "ACME", "shares": 50, "price": 490.1}
print(json.dumps(data, indent=4))
'''
{
    "name": "ACME",
    "shares": 50,
    "price": 490.1
}
'''


"""
对象实例通常并不是 JSON 可序列化的。
"""
class Point:
    def __init__(self, x, y):
        self.x = x
        self.y = y

p = Point(2, 3)
#print(json.dumps(p))

"""
如果你想序列化对象实例，你可以提供一个函数，它的输入是一个实例，返回一个
可序列化的字典。
"""
def serialize_instance(obj):
    d = {'__classname__' : type(obj).__name__}
    d.update(vars(obj))
    return d

classes = {
    'Point'  : Point
}

def unserialize_object(d):
    clsname = d.pop('__classname__', None)
    if clsname:
        cls = classes[clsname]
        obj = cls.__new__(cls)
        for key, value in d.items():
            setattr(obj, key, value)
            return obj
    else:
        return d

p = Point(2, 3)
s = json.dumps(p, default=serialize_instance)
print(s)
# {"y": 3, "__classname__": "Point", "x": 2}
a = json.loads(s, object_hook = unserialize_object)
print(a)
# <__main__.Point object at 0x00000195A10AB160>
print(a.x)
# 2

"""
json 模块还有很多其他选项来控制更低级别的数字、特殊值如 NaN 等的解析。可
以参考官方文档获取更多细节。
"""