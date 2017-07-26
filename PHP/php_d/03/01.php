<?php

# 继承
# 一个类如果扩展了另一个类，就会拥有这个类的所有属性和方法。

class FurryPets
{
    protected $sound;
    protected function fourlegs()
    {
        return 'walk on all fours';
    }

    protected function makesSound($petNois)
    {
        $this->sound = $petNois;
        return $this->sound;
    }
}

class Dogs extends FurryPets
{
    function __construct()
    {
        echo 'Dogs ' . $this->fourlegs() . '<br>';
        echo $this->makesSound('Woof, woof') . '<br>';
        echo $this->guardsHouse() . '<br>';
    }

    private function guardsHouse()
    {
        return 'Grrrrr' . '<br>';
    }
}

class Cats extends FurryPets
{
    function __construct()
    {
        echo 'Cats ' . $this->fourlegs() . '<br>';
        echo $this->makesSound('Meow, purrr') . '<br>';
        echo $this->ownsHouse() . '<br>';
    }

    private function ownsHouse()
    {
        return "I'll just walk on this keyboard." . '<br>';
    }
}

class Client
{
    function __construct()
    {
        $dogs = new Dogs();
        $cats = new Cats();
    }
}

$worker = new Client();
