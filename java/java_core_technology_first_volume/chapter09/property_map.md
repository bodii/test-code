#### 属性映射
属性映射(property map)是一个类型非常特殊的映射结构。它有下面3个特性:
1. 键与值都是字符串。
2. 表可以保存到一个文件中，也可以从文件中加载。
3. 使用一个默认的辅助表。
	实现属性映射的Java平台类称为Properties.
	属性映射通常用于程序的特殊配置选项。


#### java.util.Properties
* Properties()
	创建一个空的属性映射。

* Properties(Properties defaults)
	创建一个带有一组默认值的空的属性映射。

* String getProperty(String key)
	获得属性的对应关系；返回与键对应的字符串。如果在映射中不存在，返回默认表中与
	这个键对应的字符串。

* String getProperty(String key, String defaultValue)
	获得在键没有找到时具有的默认值属性；它将返回与键对应的字符串，如果在映射中不
	存在，就返回默认的字符串。

* void load(InputStream in)
	从InputStream加载属性映射。

* void store(OutputStream out, String comentString)
	把属性映射存储到OutputStream.
