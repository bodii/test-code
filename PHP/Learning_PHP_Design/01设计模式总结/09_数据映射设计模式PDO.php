<?php

/*
    [ 数据映射设计模式 ]

    数据映射模式（Data Mapper Pattern） 
    描述如何创建提供透明访问任何数据源的对象。数据映射模式，也叫数据访问对象模式，或数据对象映射模式。

    为什么需要数据映射模式
    数据映射模式的目的是让持久化数据存储层、驻于内存的数据表现层、以及数据映射本身三者相互独立、互不
    依赖。这个数据访问层由一个或多个映射器（或者数据访问对象）组成，用于实现数据传输。通用的数据访问
    层可以处理不同的实例类型，而专用的则处理一个或多几个。
*/

interface IFactory
{
    const HOST = 'localhost';
    const UNAME = 'root';
    const PW = '123456';
    const DBNAME = 'test';

    public function goConnect();
}

class Factory implements IFactory
{
    private static $server = IFactory::HOST;
    private static $username = IFactory::UNAME;
    private static $password = IFactory::PW;
    private static $dbname = IFactory::DBNAME;
    private static $dbtype = 'mysql';
    private static $hookup;

    public function goConnect()
    {
        $dsn = self::$dbtype . ":dbname=" . self::$dbname . ";host=" . self::$server;
        try 
        {
            self::$hookup = new PDO($dsn, self::$username, self::$password);
        } 
        catch (PDOException $e)
        {
            echo '数据库连接失败： ' . $e->getMessage();
        }
        return self::$hookup;
    }
}


class User 
{
    const TABLENAME = 'user';
    protected $id;
    protected $data;
    protected $db;
    
    protected $change = false;

    public function __construct($id)
    {
        $this->id = $id;
        $this->db = Factory::goConnect();

        $this->select();
    }

    public function __get($key)
    {
        if ($this->data[$key])
        {
            return $this->data[$key];
        }
    }

    public function __set($key, $value)
    {
        $this->data[$key] = $value;
        $this->change = true;
    }

    public function __destruct()
    {
        // 如果对象属性改变过，则更新数据库
        $this->change && $this->update();
    }

    public function select()
    {
        $selSql = "select * from " . self::TABLENAME . " where id='{$his->id}' limit 1";
        try
        {
            $this->data = $this->db->query($selSql);
        }
        catch (PDOException $e)
        {
            echo $e->getMessage();
        }
    }

    public function update()
    {
        $upFields = '';
        foreach($this->data as $key => $val)
        {
            $upFields .= "{$key}='{$val}',";
        }
        $upFields = rtrim(',', $upFields);
        $upSql = "update " . self::TABLENAME . " set {$upFields} where id={$this->id} limit 1"; 
        try
        {
            $upStatus = $this->db->exec($upSql);
        }
        catch (PDOException $e)
        {
            echo $e->getMessage();
        }

        return $upStatus ? true : false;
    }
}

$user = new User(1);
// 改变名字
$user->name = 'admin';

/*
    如果我们要实现实时更新，也可以不要chnage属性，直接在__set方法中调用update方法，
    不用等到对象销毁前再统一更新。当然实现更新时更新方法可以精简不需要的foreach,只写
    更新一字段指令就好了，但是这样也带来频繁操作数据库的问题。
*/