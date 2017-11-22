#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 用web应用程序框架实现REST '''


''' 定义进程 '''

import multiprocessing

class Simulation( multiprocessing.Process ):
    def __init__( self, setup_queue, result_queue ):
        self.setup_queue = setup_queue
        self.result_queue = result_queue
        super().__init__()

    def run( self ):
        """ waits for a termaination """
        print( self.__class__.__name__, 'start' )
        item = self.setup_queue.get()
        while item != (None, None):
            table, player = item
            self.sim = Simulate( table, player, samples=1 )
            results = list( self.sim )
            self.result_queue.put( (table, player, results[0]) )
            item = self.setup_queue.get()
        print( self.__class__.__name__, 'finish' )


class Simulate:
    def __init__( self, table, player, samples ):
        pass
    
    def __iter__( self ):
        yields summaries

class Summarize( multiprocessing.Process ):
    def __init__( self, queue ):
        self.queue = queue
        super().__init__()

    def run( self ):
        """waits for a termaination"""
        print( self.__class__.__name__, 'start' )
        count = 0
        item = self.queue.get()

        while item != (None, None, None):
            print( item )
            count += 1
            item = self.queue.get()
        print( self.__class__.__name__, 'finish', count )



# 创建multiprocessing.Quue或者它的某个子类的实例,两个定义处理管道的队列
setup_q = multiprocessing.SimpleQueue()
result_q = multiprocessing.SimpleQueue()

result = Summarize( result_q )
result.start()

# 创建4个并发的模拟进程
simulators = []
for i in range(4):
    sim = Simulation( setup_q, result_q )
    sim.start()
    simulators.append( sim )

"""
这4个并发的进程会竞争工作资源，每个都会试图从等待请求的队列中获取下一个请求。一旦这
4个模拟器满载，未处理的请求会开始填充队列。一旦队列和进都进入等待。驱动者函数就可以
开始将请求存入setup_q队列中。
"""

# 以下代码中的循环会生成大量的请求
table = Table( decks = 6, limi=50, dealer=Hit7(), split=ReSplit(), payout=(3, 3))
for bet in Flat, Martingale, OneThreeTwoSix:
    player = Player( SomeStrategy, bet(), 100, 25 )
    for sample in range(5):
        setup_q.put((table, player))


# 为了能够胡序地终止这4个模拟器，需要为每个模拟器定义哨兵对象并插入队列中
for sim in simulators:
    setup_q.put( (None, None) )

for sim in simulators:
    sim.join()

"""
在队列中为每个模拟器各添加一个可以消费的哨兵对象。一旦所有的模拟器都消费了哨兵对象，
等待进程结束后回到父进程中。
一旦Process.join（）操作结束，不会再创建任何模拟数据。我们可以在模拟操作的结果队列
中也添加一个哨兵对象

result_q.put((None, None, None))
result.join()

一旦结果队列中的哨兵对象处理完成，Summarize进程会停止接受输入，这时我们可以使用join()
将其返回父进程。


用multiprocessing将对象从一个进程传输到另外一个进程，这让我们可以用一种相对简单的方法
创建高性能、多重处理的数据管道。由于multiprocessing模块使用了pickle,因此对于对象行为
的限制几乎无法通过管道实现。
"""