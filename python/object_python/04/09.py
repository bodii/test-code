#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 上下文管理器的清理 '''

with Updating( 'some_file' ):
    with open( 'some_file', 'w') as target:
        process( target )

"""
目的是将原文件重命名为 some_file copy并备份。如果上下文中一切正常(没有异常)
然后删除备份文件或重命名为some_file old。
"""

# 上下文管理器
import os 
class Updating:
    def __init__( self, filename ):
        self.filename = filename
    
    def __enter__( self ):
        try:
            self.previous = self.filename + ' copy'
            os.rename( self.filename, self.previous )
        except FileNotFoundError:
            # Never existed, no previous copy
            self.previous = None

    def __exit__( self, exc_type, exc_value, traceback ):
        if exc_type is not None:
            try:
                os.rename( self.filename, self.filename + ' error' )
            except FileNotFoundError:
                pass # Never even got created?
            if self.previous:
                os.rename( self.previous, self.filename )


"""
不管是否执行过程中是否发生了异常，执行上下文管理器的 __exit__() 方法，__exit__() 
方法负责执行“清理”工作，如释放资源等。如果执行过程中没有出现异常，或者语句体中执行了
语句 break/continue/return，则以 None 作为参数调用 __exit__(None, None, None) ；
如果执行过程中出现异常，则使用 sys.exc_info 得到的异常信息为参数调用 __exit__(exc_type,
 exc_value, exc_traceback)
"""