#!/usr/bin/env python3
# -*- encoding=utf-8 -*-


import uuid

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
测试任务的事件方法
'''
def my_function():
    print('this is a task test.')


if __name__ == '__main__':
    task = Task(func=my_function)
    print(task)