<?php
trait TInnerClosuresInvoke 
{
    public function __call($method, $args)
    {
        if (isset($this->$method) && is_callable($method)) {
            $closure = $this->$method;
            call_user_func_array($closure, $args);
        } 
        else {
            trigger_error(
                'Call to undefined method '. __CLASS__.'::'.$method.'()', 
                E_USER_ERROR
            );
        }
    }
}

class A
{
    use TInnerClosuresInvoke;
}

$a = new A;
$a->vv('aa');