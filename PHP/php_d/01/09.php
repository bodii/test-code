<?php

# 接口常量

interface a
{
    const b = 'Interface constant';
}

// 输出常量
echo a::b;


# example 02
interface IConnectInfo
{
    const HOST = 'localhost';
    const UNAME = 'phpWorker';
    const DBNAME = 'dpPatt';
    const PW = 'easyWay';
    function testConnection();
}

class ConSQL implements IConnectInfo
{
    //使用作用域解析操作符传递值
    private $server = IConnectInof::HOST;
    private $currentDB = IConnectInfo::DBNAME;
    private $user = IConnectInfo::UNAME;
    private $pass = IConnectInfo::PW;

    public function testConnection()
    {
        $hookup = new mysqli($this->server, $this->user, $this->pass
        , $this->currentDB);

        if (mysqli_connect_error())
        {
            die('bad mojo');
        }

        echo "You're hooked up Ace!<br>" . $hookup->host_info;
        $hookup->close();
    }
}

$useConstant = new ConSQL();
$useConstant->testConnection();