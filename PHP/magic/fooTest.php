<?php

$sequence = 0;

class Foo
{
    public $stuff;

    public function __construct($param)
    {
        global $sequence;
        echo "Seq: ", $sequence++, " - constructor\n";
    }

    public function __destruct()
    {
        global $sequence;
        echo "Seq: ", $sequence++, " - destructor\n";
    }

    public function __sleep()
    {
        global $sequence;
        echo "Seq: ", $sequence++, " - __sleep\n";
        return array('stuff');
    }

    public function wakeup()
    {
        global $sequence;
        echo "Seq: ", $sequence++, " - __wakeup\n";
    }
}


session_start();
$_SESSION['obj'] = new foo('A foo');