<?php

// 命名空间
namespace design_patterns\facade;

/**
 * BiosInferface interface
 * 输入输出系统接口类
 */
interface BiosInferface
{
    /**
     * Execute function
     * 开始执行方法
     *
     * @return void
     */
    public function execute();

    /**
     * WaitForKeyPress function
     * 等待按键方法
     *
     * @return void
     */
    public function waitForKeyPress();

    /**
     * Launch function
     * 启动OS对象
     * 
     * @param OsInterface $os OS实例对象
     * 
     * @return void
     */
    public function launch(OsInterface $os);

    /**
     * PowerDown function
     * 断电方法
     *
     * @return void
     */
    public function powerDown();
}