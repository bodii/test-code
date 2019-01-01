<?php

/**
 * 冒泡排序法
 */



/**
 * BubbleSortV1 function
 * 冒泡排序法 第一版本
 *
 * @param array $arr 一个数组
 * 
 * @return void
 */
function bubbleSortV1(array &$arr)
{
    $count = count($arr);

    for ($i = 0; $i < $count; $i++) {
        for ($j = 0; $j < $count-1; $j++) {
            if ($arr[$j+1] < $arr[$j]) {
                list($arr[$j], $arr[$j+1]) = array($arr[$j+1], $arr[$j]);
            }
        }
    }
}


