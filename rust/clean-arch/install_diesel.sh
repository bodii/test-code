#!/bin/bash

# 在项目下执行启动docker服务
# 根据docker-compose.yml，启动执行配置的服务
sudo docker-compose up -d > /dev/null 2>&1

# 查看docker服务是否加载成功
sudo docker ps

# 安装 mysql、sqlite3、postgresql 客户端驱动
sudo apt install -y libmysqlclient libsqlite3-dev libpq-dev

# 安装diesel
cargo install diesel

# 项目下初始化diesel
diesel setup > /dev/null 2>&1

# 删除diesel初始化示例目录
rm -rf migrations/00000000000000_diesel_initial_setup

# 创建user
diesel migration generate create_user > /dev/null 2>&1

# 添加创建users表的sql语句
d=$(ls migrations);
echo '
CREATE table users (
    id SERIAL primary key,
    name varchar not null,
    email varchar not null unique,
    phone varchar not null,
    address varchar not null
);

create index index_users_on_email on users(email);
' \
>> "migrations/${d}/up.sql";

# 通过diesl将sql语句添加到表中
diesel migration run

# 创建目录
mkdir -p src/application src/domain \
src/infrastructure src/presentation

# end
cargo run --quiet
