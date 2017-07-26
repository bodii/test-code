#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic:根据用户输入的登录名和口令模拟用户注册
Desc:
"""

import hashlib
import pickle
import time


# 模拟用户数据表文件 文件路径
db_user_path = 'db_user.txt'
# 模拟被锁用户数据表文件 文件路径
db_lock_user_path = 'db_lock_user.txt'

# 读取用户数据表信息
def read_db_userinfo():
    try:
        with open(db_user_path,'rb') as db_user:
            db_user = db_user
    except FileNotFoundError:
        with open(db_user_path,'wb') as db_user:
            userinfo={}
            pickle.dump(userinfo,db_user)
    finally:
        with open(db_user_path,'rb') as db_user:
            db_userinfo = pickle.load(db_user)
            return db_userinfo


# 读取被锁用户数据表信息
def read_db_lock_userinfo():
    try:
        with open(db_lock_user_path,'rb') as db_lock_user:
            db_lock_user = db_lock_user
    except FileNotFoundError:
        with open(db_lock_user_path,'wb') as db_lock_user:
            lock_userinfo={}
            pickle.dump(lock_userinfo,db_lock_user)
    finally:
        with open(db_lock_user_path,'rb') as db_lock_user:
            db_lock_userinfo = pickle.load(db_lock_user)
            return db_lock_userinfo

# 将用户数据写入数据表
def register_in_db(username, password):
    md5_password = get_md5(password + username + 'the-Salt')
    with open(db_user_path,'rb') as db_user:
        userinfo = pickle.load(db_user)
    with open(db_user_path,'wb') as db_user:
        userinfo[username] = md5_password
        pickle.dump(userinfo,db_user)


# 将被锁用户的数据从用户数据表移入被锁用户数据表
def lock_user_in_db(username):
    try:
        with open(db_lock_user_path,'rb') as db_lock_user:
            db_lock_user = db_lock_user
    except FileNotFoundError:
        with open(db_lock_user_path,'wb') as db_lock_user:
            lock_userinfo={}
            pickle.dump(lock_userinfo,db_lock_user)
    finally:
        with open(db_user_path,'rb') as db_user:
            userinfo = pickle.load(db_user)
            lock_username = username
            lock_password = userinfo[username]
            userinfo.pop(lock_username)
        with open(db_user_path,'wb') as db_user:
            pickle.dump(userinfo,db_user)
        with open(db_lock_user_path,'rb') as db_lock_user:
            lock_userinfo_data = pickle.load(db_lock_user)
            lock_userinfo_data[lock_username] = lock_password
        with open(db_lock_user_path,'wb') as db_lock_user:
            pickle.dump(lock_userinfo_data,db_lock_user)
            return lock_userinfo_data

# 注册用户--并将用户信息写入数据表
def register_user():
    db = read_db_userinfo()  # 读取用户数据表信息
    lock_db = read_db_lock_userinfo()  # 读取被锁用户数据表信息
    is_register = input('你要注册吗：yes or no:\n')
    if is_register == 'yes':
        register_username = input('请输入你要注册的用户名：')
        if register_username in db.keys():
            print('用户名已存在')
            time.sleep(1)
            return
        elif register_username in lock_db.keys():
            print('此用户名已被锁，')
            time.sleep(1)
            return
        elif register_username is None:
            print('用户名不能为空')
            time.sleep(1)
            return
        else:
            register_password = input('请输入密码：')
            if register_password is None:
                print('密码不能为空')
                time.sleep(1)
                return
            else:
                register_in_db(register_username.strip(), register_password.strip())
                print('注册成功！')
                time.sleep(1)
                is_login = input('你现在要登录吗: yes or no ?\n')

                if is_login == 'yes':
                    login()
                else:
                    return
    else:
        return


# 输入并判断登录时输入的用户名是否存在
def exist_user(afresh_username=None):
    db = read_db_userinfo()  # 读取用户数据表信息
    if afresh_username is None:
        username = input('用户名：')
    else:
        username = input('请再次输入登录时的用户名(或输入"exit"退出)：\n')
    if username in db.keys():
        exist_password(username)  # 调用“判断登录时输入的密码是否正确”方法
    elif username == 'exit':
        time.sleep(1)
        register_user() # 调用注册用户方法
    else:
        print('用户名不正确!')
        exist_user(afresh_username=True)



# 输入并判断登录时输入的密码是否正确
def exist_password(username, afresh_password = None, password_num = 1):
    db = read_db_userinfo()  # 读取用户数据表信息
    if afresh_password is None:
        password = input('密码:')
    else:
        password = input('请再次输入登录时的密码(或输入"exit"退出):\n')
    login_password = get_md5(username + password + 'the-Salt')
    if login_password == db[username]:
        print('登录成功!')
    elif password == 'exit':
        return
    else:
        if password_num < 3:  # 记录用户输入的,当大于3次时，退出程序
            print('密码不正确！')
            exist_password(username, afresh_password = True, password_num = password_num+1)
        else:
            lock_user_in_db(username)  # 锁用户方法
            print('密码输入错误超过3次，用户名已被锁，请联系网站管理员！')
            return


# md5加密
def get_md5(data):
    md5 = hashlib.md5()
    md5.update(data.encode('utf-8'))
    return md5.hexdigest()


# 登录 -- 用户信息判断
def login():
    db = read_db_userinfo()  # 读取用户数据表信息
    print('[ 登 录 ]')
    exist_user()



# 测试登录程序
if __name__ == '__main__':
    login()
