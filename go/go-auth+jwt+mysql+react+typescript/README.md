# `This is an authentication website`

It is based on golang to provide a back-end interface for user registration, login and other interfaces, and based on react javascript to provide a front-end request and display of the website function program.

### `Used tech stack`

> `golang`、`jwt`、`cookie`、`mysql`、`react`、`typescript` ......

### `Used golang formework`
> `fiber`

### `Instructions for use`

1. Create mysql database：

```mysql
create database yt_go_auth default character set utf8mb4 default collate utf8mb4_unicode_ci;
```

2. Start the back-end interface program：

```shell
cd go-auth
vim database/connetction.go (view and modify the database configuration on line 13)
go mod tidy
go run main.go
```

3. Start the front-end interface program：

```shell
cd react-auth
npm i
npm start
```
