<?php


# 尽管PHP没有强类型的签名，不过完全可以在你的计划和实现中加以强制。可以在接口中
# 使用注释语句来做到。
# 如果接口（包括抽象类）中有一个扩展的虚签名（使用注释），就能提供和利用强类型语言
# 的优点。

// PHP接口
interface interface
{
    //@return 字符串
    function stringMethod();

    //@return 整数并使用整数属性
    function numMethod($intProp);

    //不使用return
    function noReturnMethod();
}
