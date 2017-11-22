#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' SQL事务和ACID属性 '''

"""
SQL的DML语句都工作在一个SQL事务的上下文中。在一个事务中执行的SQL语句是一个工作的逻辑单元，整个
事务可以被提交或回滚，整个过程为原子操作。
SQL的DDL语句（例如：create、drop）不会在事务中工作。
ACID是指原子性（Atomic）、一致性（Consistent）、隔离性（Isolated）和持久性（Durable）。它
们是事务的基本属性，其中每个事务中包括了多个数据库操作。

对于SQLite来说，有4种隔离级别用于定义锁以及事务的本质：
1. isolation_level = None:这是默认的设置，也被称为自动提交（autocommit）模式。在这种模式下，
每个SQL语句都会在执行后直接提交到数据库。它破坏了原子性，而有些奇怪的观点则认为，所有的事务都应当
只包含一个SQL语句。

2. isolation_level = 'DEFERRED': 在这种模式中，事务中锁的添加越晚越好。例如BEGIN语句，并没
有立即获得任何锁。对于其他的读操作（即select语句）可以获得共享锁，写操作将获得保留的锁。然而这样
可以最大化并发，但在多个进程中也会产生死锁。

3. isolation_level = 'EXCLUSIVE':在这种模式中，事务的BEGIN语句会获得一个锁，阻止其他操作的
访问。对于一些链接，在一种特殊的读未提交模式中，忽略锁会导致异常。

在SQL中，需要使用BEGIN TRANSATION和COMMIT TRANSACTION来将括号内的步骤包括在事务中。出现错误
时，需要使用ROLLBACK TRANSACTION语句来进行回滚。在Python中的接口简化了这一点。我们可以执行一个
BEGIN语句。其他语句由sqlite3.Connection对象中的函数来提供，不需要执行SQL语句来终止一个事务
"""