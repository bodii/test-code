<?php

/*
    接口
    接口的核心部分由类中操作（函数）定义的所有签名组成。
    签名 包括一个操作的操作名和参数。
*/

# 接口 一般g约定总以字母I或i开头

interface  IMethodHolder
{
    public function getInfo($info);
    public function sendInfo($info);
    public function calculate($frist, $second);
}


# 实现接口  需要使用implements语句。
# 接口中的所有方法必须实现！

class ImplementAlpha implements IMethodHolder
{
    public function getInfo($info){
        echo "This is NEWS! " . $info . "<br>";
    }

    public function sendInfo($info){
        return $info;
    }

    public function calculate($first, $second){
        $calulated = $first * $second;
        return $calulated;
    }

    public function useMethods(){
        $this->getInfo('The sky is falling...');
        echo $this->sendInfo('Vote for Senator Snort!') . '<br>';
        echo 'You make $' . $this->calculate(20, 15) . " in your part-time job<br>";
    }
}

$worker = new ImplementAlpha();
$worker->useMethods();
