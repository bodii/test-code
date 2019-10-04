#### Java Web Start
Java Web Start 是一项在Interent上发布应用程序的技术。Java Web Start应用程序包含下列主要
特性：
* Java Web Start 应用程序一般通过浏览器发布。只要Java Web Start应用程序下载到本地就可以
启动它，而不需要浏览器。
* Java Web Start 应用程序并不在浏览器窗口内。它将显示在浏览器外的一个属于自己的框架中。
* Java Web Start 应用程序不使用浏览器的Java实现。浏览器只是在加载Java Web Start应用程序
描述符时启动一个外部应用程序。这与启动诸如Adobe Acrobat 或RealAudio这样的辅助应用所使用
的机制一样。
* 数字签名应用程序可以被赋予访问本地机器的任意权限。未签名的应用程序只能运行在"沙箱"中，
它可以阻止具有潜在危险的操作。


#### 发布Java Web Start应用
要想准备一个通过Java Web Start发布的应用程序，应该将其打包到一个或多个JAR文件中。然后创建
一个Java Network Launch Protocol(JNLP)格式的描述文件。将这些文件放在Web服务器上。
