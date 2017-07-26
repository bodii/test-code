<?php

# 多态
# 多态就是指多种形态，不过只有在OOP上下文中这么讲才有意义。
# 多态的真正价值在于，可以调用有相同接口的对象来完成不同的工作。

include_once '../display_errors.php';
interface ISpeed
{
    function fast();
    function cruise();
    public function slow();
}

class Jet implements ISpeed
{
    function slow()
    {
        return 120;
    }

    public function cruise()
    {
        return 1200;
    }

    function fast()
    {
        return 1500;
    }
}


class Car implements ISpeed
{
    function slow()
    {
        $carSlow = 15;
        return $carSlow;
    }

    function cruise()
    {
        $carCruise = 65;
        return $carCruise;
    }

    function fast()
    {
        $carZoom = 110;
        return $carZoom;
    }
}

$f22 = new Jet();
$jetSlow = $f22->slow();
$jetCruise = $f22->cruise();
$jetFast = $f22->fast();
echo "<br> My jet can take off at $jetSlow mph and cruises at $jetCruise
mph.However, I can crank it up to $jetFast mph if I'm in a hurry.<br>";

$ford = new Car();
$fordSlow = $ford->slow();
$fordCruise = $ford->cruise();
$fordFast = $ford->fast();
echo "<br> My car pokes along at $fordSlow mph in a school zone and cruise
 at $fordCruise mph on the highway. However, I can crank it up to $fordFast
  mph if I'm in a hurry.<br>";
