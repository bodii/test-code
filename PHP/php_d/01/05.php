<?php

class A
{
    public function create(){
        return new self();
    }
}

class B extends A
{

}

B->create(); // 返回A的实例


class C
{
    public function create(){
        return new static();
    }
}

class D extends C
{

}

D->create() // 返回D类的实例
