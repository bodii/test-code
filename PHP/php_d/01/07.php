<?php

# 接口练习 02
# 可扩充的接口

interface a
{
    public function foo();
}

interface b extends a
{
    public function baz(Baz $baz);
}

class c implements b
{
    public function foo()
    {

    }

    public function baz(Baz $baz)
    {
        
    }
}
