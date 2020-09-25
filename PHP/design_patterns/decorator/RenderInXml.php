<?php

// 命名空间
namespace design_patterns\decorator;

/**
 * RenderInXml class
 * Xml装饰器类
 * 用于将数据生成XML格式内容的字符串
 */
class RenderInXml extends RendererDecorator
{
    /**
     * RenderData function
     * 渲染数据的装饰方法
     * 生成XMl内容的方法
     * 
     * @return string 返回生成的XML内容
     */
    public function renderData(): string
    {
        $doc = new \DOMDocument();
        $data = $this->wrapped->renderData();
        $doc->appendChild(
            $doc->createElement(
                'content', 
                $data
            )
        );

        return $doc->saveXml();
    }
}