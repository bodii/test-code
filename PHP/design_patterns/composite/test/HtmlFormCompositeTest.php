<?php

// 命名空间
namespace design_patterns\composite\test;

use design_patterns\composite\{
    Form,
    InputElement,
    TextElement,
};

// 定义目录分隔符
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入composer vender autoload 文件
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

// 使用vender 下的 phpunit测试类
use PHPUnit\Framework\TestCase;

/**
 * HtmlFormCompositeTest class
 * html form表单组合测试类
 */
class HtmlFormCompositeTest extends TestCase
{
    /**
     * TestRendorEmailInputContent function
     * 测试Email input 组合方法
     * 
     * @return void
     */
    public function testRendorEmailInputContent()
    {
        $form = new Form();
        $form->addElement(new TextElement('Email:'));
        $form->addElement(new InputElement());

        $this->assertEquals(
            '<form>Email:<input type="text" /></form>',
            $form->render()
        );
    }

    /**
     * TestRendorPasswordInputContent function
     * 测试password input 组合方法
     *
     * @return void
     */
    public function testRendorPasswordInputContent()
    {
        $embed = new Form();
        $embed->addElement(new TextElement('Password:'));
        $embed->addElement(new InputElement());
        $form = new Form();
        $form->addElement($embed);

        $this->assertEquals(
            '<form><form>Password:<input type="text" /></form></form>',
            $form->render()
        );
    }
}