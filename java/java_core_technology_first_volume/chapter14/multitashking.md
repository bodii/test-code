#### 多任务(multitasking)
多线程程序在较低的层次上扩展了多任务的概念: 一个程序同时执行多个任务。
通常，每个任务称为一个`线程(thread)`,它是线程控制的简称。
可以同时运行一个以上线程的程序称为`多线程程序(multithreaded)`。

多进程和多线程的区别:
本质的区别在于每个进程拥有自己的一整套变量，而线程则共享数据。
共享变量使线程之间的通信比进程之间的通信更有效、更容易。此外，在有些操作系统中，与进程相
比，线程更“轻量级”，创建、撤销一个线程比启动新进程的开销要小很多。
