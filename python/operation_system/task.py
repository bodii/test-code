#!/usr/bin/env python3
# -*- encoding=utf-8 -*-


import uuid
from threading import Condition

__all__ = ['Task', 'AsyncTask']

'''
任务类
'''
class Task:
    def __init__(self, func, *args, **kwargs):
        self.callable = func
        self.id = uuid.uuid4()
        self.args = args
        self.kwargs = kwargs

    def __str__(self):
        return f"Task id: {self.id}"


'''
异步任务类
'''
class AsyncTask(Task):
    def __init__(self, func, *args, **kwargs):
        self.result = None
        self.condition = Condition()
        super().__init__(func, *args, **kwargs)

    def set_result(self, result):
        """
        设置运行结果
        """
        self.condition.acquire()
        self.result = result
        self.condition.notify()
        self.condition.release()

    def get_result(self):
        """
        获取运行结果
        """
        self.condition.acquire()
        # 如果没有，则等待
        if not self.result:
            self.condition.wait()
        result = self.result
        self.condition.release()

        return result


'''
测试任务的事件方法
'''
def my_function():
    print('this is a task test.')


if __name__ == '__main__':
    task = Task(func=my_function)
    print(task)