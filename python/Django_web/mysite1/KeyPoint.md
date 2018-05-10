### Django 关键点

**1. 命令**

    1.1 查看django版本
    > $ ptyhon -m django --version

    1.2 创建项目
    > $ django-admin startproject mysite

    1.3 运行项目
    > $ python manage.py runserver

        1.3.1 更换默认端口号
        > $ python manage.py runserver 8080

        1.3.2 更换IP和端口号
        > $ ptyhon manage.py runserver 0:8000
        0是0.0.0.0的简写

    1.4 创建应用
    > $ python manage.py startapp polls

    1.5 将model创建、迁移数据库表
    > $ python manage.py migrate
    自动执行数据库迁移并同步管理你的数据库结构的命令

    1.6 将应用添加到项目中
    > $ python manage.py makemigrations polls
    运行 makemigrations 命令，Django 会检测你对模型文件的修改（在这种情况下，你已经取得了新的），并且把修改的部分储存为一次 迁移。

    1.7 查看迁移会执行的SQL
    > $ python manage.py sqlmigrate polls[应用名] 0001[迁移SQL的py文件，默认在应用下migrations文件夹内]

    1.8 交互式Python命令行
    > $ python manage.py shell

    1.9 创建管理员账号
    > $ python manage.py createsuperuser




***
