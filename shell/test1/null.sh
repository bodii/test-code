#!/bin/bash

echo 
echo "Find some file in the whole filesystem, take very long time...."
echo "You can press Ctrl+C to abort this..."

# 在整个文件系统中查找一个文件
find / -name "somefile" 2> /dev/null

exit 0
