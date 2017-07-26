#!/bin/bash

echo 
echo "Sending server data to file desciptor 3..."

# 发送数据到文件描述符3
echo "Server Data from file desciptor 3." >&3
echo 
echo "Sending Login data to file desciptor 4..."

# 发送数据到文件描述符4
echo "Login Data from file descriptor 4." >&4
echo 
echo "Sending Error Message to file descriptor 5..."

# 发送数据到文件描述符5
echo "Error Message from file descriptor 5." >&5
echo 
echo "Sending Statistics data to file descriptor 6..."

# 发送数据到文件描述符6
echo "Statics data from file desciptor 6." >&6
echo 

exit 0
