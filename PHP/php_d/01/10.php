<?php


include_once '../display_errors.php';



# 接口 建立宽松但明确的绑定

interface IProduct
{
    function apples();
    function oranges();
}

# FruitStore 实现 IProduct接口
class FruitStore implements IProduct
{
    public function apples()
    {
        return 'FruitStore sez--We have apples.<br>';
    }

    public function oranges()
    {
        return 'FruitStore sez-We have no citrus fruit.<br>';
    }
}

# CitrusStore 实现 IProduct 接口
class CitrusStore implements IProduct
{
    public function apples()
    {
        return 'CitrusStore sez--We do not sell apples.<br>';
    }

    public function oranges()
    {
        return 'CitrusStore sez--We have citrus fruit.<br>';
    }
}


# 有类型提示的对象
/*
    强制数据类型可以确保倘若给定方法中使用了代码提示，那么其中使用的对象（类
    ）必然有给定的接口。另外，如果把一个接口（可以是一个抽象或接口）作为代码
    提示，绑定会更宽松；它会绑定到接口而不是绑定到一个特定的实现。
    不能使用标量类型（如string或int）作为代码提示，不过可以使用数组、接口
    和类作为代码提示。所以尽量没有另外一些语言那么灵活，但PHP中可以通过类型
    提示实现类型，这在OOP和设计模式编程中起着重要作用。 
*/

class UseProducts
{
    public function __construct()
    {
        $appleSauce = new FruitStore();
        $orangeJuice = new CitrusStore();
        $this->doInterface($appleSauce);
        $this->doInterface($orangeJuice);
    }

    # 类型提示（type hint）IProduct能够识别实现IProduct接口的两个类。
    # 换句话说，并不是把它们分别识别为一个FruitStore和CitrusrStore实例
    # 而会识别它们共同的接口IProduct
    function doInterface(IProduct $product) #IProduct 属于类型提示
    {
        echo $product->apples();
        echo $product->oranges();
    }
}

$worker = new UseProducts();
