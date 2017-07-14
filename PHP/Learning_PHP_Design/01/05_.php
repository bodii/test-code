<?php 

/*
	数据对象映射模式

    将对象和数据存储映射起来，对一个对象的操作会映射为数据存储的操作
*/

interface IFactory
{
    const HOST = 'localhost';
    const UNAME = 'uname';
    const PW = 'password';
    const DBNAME = 'db';
    public static function getDateabase();
}

class Factory implements IFactory
{
    private static $server = IFactory::HOST;
    private static $userName = IFactory::UNAME;
    private static $password = IFactory::PW;
    private static $dbname = IFactory::DBNAME;
    private static $hookup;

    public static function getDateabase()
    {
        self::$hookup = mysqli_connect(
            self::$server, 
            self::$userName, 
            self::$password, 
            self::$dbname
            );

        if (self::$hookup) 
        {
            // 连接成功提示 或其他
        } 
        elseif (mysqli_connect_error(self::$hookup))
        {
            echo 'Here\'s why it faild: ' . mysqli_connect_error();
        }

        return self::$hookup;
    }
}


class User
{
    private $db;
    private $data = [];
    private $id;
    private $sql;
    private $change = false;
    private $tableMaster;

    public function __construct(int $id)
    {
        $this->db = Factory::getDateabase();
        $this->id = $this->db->real_escape_string($id);
        $this->sql = "select * from {$this->tableMaster} where id={$this->id} limit 1";
        $res = $this->db->query($sql);
        $this->data = $res->fetch_assoc();
    }

    function __get($key)
    {
        if (isset($this->data[$key]))
        {
            return $this->data[$key];
        }
    }

    function __set($key, $value)
    {
        $this->data[$key] = $value;
        $this->change = true;
    }
    
    function __destruct()
    {
        $fileds = [];
        if ($this->change)
        {
            foreach($this->data as $k=>$v)
            {
                $fileds[] = "$k=$v";
            }
            $changeSql = "update  {$this->tableMaster} set " . implode(', ', $fileds) 
            . "where id={$this->id} limit 1";
            $this->db->query($changeSql);
        }
    }
}

$user = new User(1);
$user->mobile = '188 0000 0000';
$user->name = 'test';