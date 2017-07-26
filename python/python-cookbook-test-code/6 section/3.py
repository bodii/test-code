#！/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic:第六章：数据编码和处理
Description: 这一章主要讨论使用Python 处理各种不同方式编码的数据，比如CSV 文件，
JSON，XML 和二进制包装记录。和数据结构那一章不同的是，这章不会讨论特殊的
算法问题，而是关注于怎样获取和存储这些格式的数据。

Title:          解析简单的XML 数据
Issue:你想从一个简单的XML 文档中提取数据。
Answers:可以使用xml.etree.ElementTree 模块从简单的XML 文档中提取数据。为了演
示，假设你想解析Planet Python 上的RSS 源。
"""

from urllib.request import urlopen
from xml.etree.ElementTree import parse

u = urlopen('http://planet.python.org/rss20.xml')
doc = parse(u)

for item in doc.iterfind('channel/item'):
    title = item.findtext('title')
    date = item.findtext('pubDate')
    link = item.findtext('link')

    print(title)
    print(date)
    print(link)
    print()

'''
Filipe Saraiva: if (LaKademy 2016) goto Rio de Janeiro
Tue, 24 May 2016 23:46:58 +0000
http://blog.filipesaraiva.info/?p=1825

Carl Chenet: Tweet your database with db2twitter
Tue, 24 May 2016 22:00:51 +0000


Mahmoud Hashemi: Managing Python Ecosystems
Tue, 24 May 2016 17:00:00 +0000
http://sedimental.org/managing_python_ecosystems.html

Real Python: Getting Started with the Slack API using Python and Flask
Tue, 24 May 2016 13:18:21 +0000
https://realpython.com/blog/python/getting-started-with-the-slack-api-using-python-and-flask/

Mike Driscoll: Python 101: An Intro to Benchmarking your code
Tue, 24 May 2016 12:30:14 +0000
http://www.blog.pythonlibrary.org/2016/05/24/python-101-an-intro-to-benchmarking-your-code/

Ian Ozsvald: PyDataLondon 2016 Conference Write-up
Tue, 24 May 2016 09:42:16 +0000
http://ianozsvald.com/2016/05/10/pydatalondon-2016-conference-write-up/

EuroPython: EuroPython 2016 Keynote: Rachel Willmer
Tue, 24 May 2016 08:28:41 +0000
http://blog.europython.eu/post/144849205747

Talk Python to Me: #60 Scaling Python to 1000's of cores with Ufora
Tue, 24 May 2016 08:00:00 +0000
https://talkpython.fm/episodes/show/60/scaling-python-to-1000-s-of-cores-with-ufora

Davy Wybiral: Quantum Circuit Simulator
Mon, 23 May 2016 18:47:00 +0000
http://davywybiral.blogspot.com/2012/12/quantum-circuit-simulator.html

Kushal Das: dgplug summer training student Tosin Damilare James Animashaun
Mon, 23 May 2016 17:45:00 +0000
https://kushaldas.in/posts/dgplug-summer-training-student-tosin-damilare-james-animashaun.html

Trey Hunner: Weekly Python Chat: Live From PyCon
Mon, 23 May 2016 16:00:00 +0000
https://treyhunner.com/2016/05/weekly-python-chat-live-from-pycon/

PyCharm: PyCharm 2016.1.4 RC is Available
Mon, 23 May 2016 13:04:27 +0000
http://feedproxy.google.com/~r/Pycharm/~3/uGoIuIgcb9w/

Doug Hellmann: bisect — Maintain Lists in Sorted Order — PyMOTW 3
Mon, 23 May 2016 13:00:11 +0000
http://feeds.doughellmann.com/~r/doughellmann/python/~3/knL3DA02yVI/

Mike Driscoll: PyDev of the Week: Wesley Chun
Mon, 23 May 2016 12:30:57 +0000
http://www.blog.pythonlibrary.org/2016/05/23/pydev-of-the-week-wesley-chun/

Graeme Cross: PyCon AU 2016 registrations are open!
Sun, 22 May 2016 23:26:34 +0000
http://www.curiousvenn.com/2016/05/pycon-au-2016-registrations-are-open/

Evennia: Evennia 0.6 !
Sun, 22 May 2016 23:17:13 +0000
http://evennia.blogspot.com/2016/05/evennia-06.html

Richard Jones: PyCon Australia 2016: Registration Opens!
Sun, 22 May 2016 22:52:24 +0000
http://www.mechanicalcat.net/richard/log/Python/PyCon_Australia_2016__Registration_Opens

Mikko Ohtamaa: Python standard logging pattern
Sun, 22 May 2016 19:30:26 +0000
https://opensourcehacker.com/2016/05/22/python-standard-logging-pattern/

Programando Ciência: Complex scatter plots on Python [PART III] – Inserting labels into elements and defining more than one legend
Sun, 22 May 2016 18:32:47 +0000


Weekly Python Chat: PyCon Sprints
Sun, 22 May 2016 18:00:00 +0000
http://ccst.io/e/sprints

Abu Ashraf Masnun: Building a Facebook Messenger Bot with Python
Sun, 22 May 2016 16:44:28 +0000
http://masnun.com/2016/05/22/building-a-facebook-messenger-bot-with-python.html

PyCon Australia: Tickets now on sale!
Sun, 22 May 2016 13:28:12 +0000
http://2016.pycon-au.org/media/news/39

BangPypers: Talks About Mesh Networks - 2016 May Meetup
Sat, 21 May 2016 17:31:00 +0000
http://bangalore.python.org.in/blog/2016/05/21/talks/

Podcast.__init__: Episode 58 - Wagtail with Tom Dyson
Sat, 21 May 2016 16:23:05 +0000
http://podcastinit.podbean.com/e/episode-58-wagtail-with-tom-dyson/

EuroPython: PyData EuroPython 2016
Sat, 21 May 2016 10:03:46 +0000
http://blog.europython.eu/post/144694765622
'''

"""
很显然，如果你想做进一步的处理，你需要替换print() 语句来完成其他有趣的
事。
在很多应用程序中处理XML 编码格式的数据是很常见的。不仅因为XML 在
Internet 上面已经被广泛应用于数据交换，同时它也是一种存储应用程序数据的常用格
式(比如字处理，音乐库等)。接下来的讨论会先假定读者已经对XML 基础比较熟悉
了。
在很多情况下，当使用XML 来仅仅存储数据的时候，对应的文档结构非常紧凑并
且直观。例如，上面例子中的RSS 订阅源类似于下面的格式：
<?xml version="1.0"?>
<rss version="2.0" xmlns:dc="http://purl.org/dc/elements/1.1/">
<channel>
<title>Planet Python</title>
<link>http://planet.python.org/</link>
<language>en</language>
<description>Planet Python - http://planet.python.org/</description>
<item>
<title>Steve Holden: Python for Data Analysis</title>
<guid>http://holdenweb.blogspot.com/...-data-analysis.html</guid>
<link>http://holdenweb.blogspot.com/...-data-analysis.html</link>
<description>...</description>
<pubDate>Mon, 19 Nov 2012 02:13:51 +0000</pubDate>
</item>
<item>
<title>Vasudev Ram: The Python Data model (for v2 and v3)</title>
<guid>http://jugad2.blogspot.com/...-data-model.html</guid>
<link>http://jugad2.blogspot.com/...-data-model.html</link>
<description>...</description>
<pubDate>Sun, 18 Nov 2012 22:06:47 +0000</pubDate>
</item>
<item>
<title>Python Diary: Been playing around with Object Databases</title>
<guid>http://www.pythondiary.com/...-object-databases.html</guid>
<link>http://www.pythondiary.com/...-object-databases.html</link>
<description>...</description>
<pubDate>Sun, 18 Nov 2012 20:40:29 +0000</pubDate>
</item>
...
</channel>
</rss>
xml.etree.ElementTree.parse() 函数解析整个XML 文档并将其转换成一个文档
对象。然后，你就能使用find() 、iterfind() 和findtext() 等方法来搜索特定的
XML 元素了。这些函数的参数就是某个指定的标签名，例如channel/item 或title
。
每次指定某个标签时，你需要遍历整个文档结构。每次搜索操作会从一个起始元
素开始进行。同样，每次操作所指定的标签名也是起始元素的相对路径。例如，执行
doc.iterfind('channel/item') 来搜索所有在channel 元素下面的item 元素。doc
代表文档的最顶层(也就是第一级的rss 元素)。然后接下来的调用item.findtext()
会从已找到的item 元素位置开始搜索。
ElementTree 模块中的每个元素有一些重要的属性和方法，在解析的时候非常有
用。tag 属性包含了标签的名字，text 属性包含了内部的文本，而get() 方法能获取
属性值。例如：
"""

print(doc)
'''
<xml.etree.ElementTree.ElementTree object at 0x00000000010D2CC0>
'''

e = doc.find('channel/title')
print(e)
# <Element 'title' at 0x0000000002CD9E58>
print(e.tag)
# title
print(e.text)
# Planet Python
print(e.get('some_attribute'))
# None

"""
有一点要强调的是xml.etree.ElementTree 并不是XML 解析的唯一方法。对于
更高级的应用程序，你需要考虑使用lxml 。它使用了和ElementTree 同样的编程接
口，因此上面的例子同样也适用于lxml。你只需要将刚开始的import 语句换成from
lxml.etree import parse 就行了。lxml 完全遵循XML 标准，并且速度也非常快，
同时还支持验证，XSLT，和XPath 等特性。
"""