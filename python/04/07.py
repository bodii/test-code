#!/usr/bin/env python3
# -*- coding:utf-8 -*-


import os
import shutil

# 复制文件
shutil.copy('06.py', '06.backup.py')

# 重命名或移动文件
shutil.move('06.backup.py', '06.backup1.py')

# 删除文件
os.remove('06.backup1.py')

shutil.copy('06.py', '06.backup1.py')
# 删除文件
os.unlink('06.backup1.py')

shutil.copy('06.py', '06.backup1.py')
# 文件权限设定
os.chmod('06.backup1.py', 777)

