<?php

# 策略设计模式的灵活性

# 封装算法
# 设计模式的主要原则之一是封装变化的内容。对于发送到PHP类的不同类型的请求，
# 分别有不同的算法来处理这些MySQL请求，变化的显然是算法。这些变化可能很小，
# 也可能是显著的变化，不过通过使用策略设计模式，我们可以大大简化这个过程。
#
# 算法族
# 一般需要输入、修改、获取和删除数据，这些涉及数据操作的行为就构成了一个“族"
# ,每个族成员可以转换为一个策略。实现这些策略需要不同的算法，还要结合使用
# MySQL命令和PHP mysqli类。把这些操作放在不同的具体类中（所有具体类实现一
# 个公共接口），就成为了策略设计模式的一部分。

# 职责链设计模式
# 职责链设计模式将请求的发送者与接收者分开，这样可以避免请求者与接收者的耦合。
# 另外，这个模式允许将请求沿着一个链传递到多个不同的对象，使这些对象都有机
# 会处理请求。发送者并不需要知道哪一个对象处理这个请求，而对象也不需要知道
# 是谁发送了这个请求。在这二者之间不存在耦合。
#
# 看到职责链模式时，有些人可能会认为“这只不过是一个大的switch语句而已”，从
# 某种程序上看似乎如些，不过实际上并不是这样。这确实会检查一个请求，确实它
# 是否与某个case条件匹配。不过，switch语句是固定的，而职责链会由各个具体处
# 理器定义其后继。基于这种组织，可以把任意多个响应存储在一个MySQL数据库表中，
# 这样一来，多个不同的“咨询台”可以使用相同的表以及相同的PHP职责链。由于每个
# 具体处理器都包含一个方法来指定它自己的后继，Client通过具体处理对象指定后
# 继时也就定义了顺序。
# 另外，由于Client要启动请求链，它可以在开发人员指定的任何位置开始。假设具
# 体请求处理器4、10、15和30可以作为你希望的新咨询台，就可以指定处理器4作为
# 链中的第一个处理器，然后指定处理器10作为它的后继，再指定15为10的后继，而
# 30作为15的后继。所以，职责链具有switch语句不骨的灵活性。
#
# mysql咨询台中的职责链
# 职责链的第一个实现是一个咨询台，用户可以从一系列帮助主题中选择一个要询问
# 的主题。不论是作为PHP设计模式，还是使用Mysql作为一个文本数据存储系统，这
# 都是一个简单的实现，基本说来，在这个实现中，用户将从一组单选钮中选择一个
# 查询，发出请求，然后通过一个响应链搜索。找到正确的响应时，对象从一个MySQL
# 表获取响应并在屏幕上显示。
