<?php

class a
{
    public function __toString()
    {
        list($a) = func_get_args();
        return $a;
    }
}

$a = new a;
echo $a->__tostring('PHP');