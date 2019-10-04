#### 沙箱
你可以选择沙箱权限和完全权限。
包含Permissions: snadbox 或Permaissions: all-permission
清单:
```
Manifest-Version: 1.0
Permissions: all-permissions
```

接下来远行jar工具:
```shell
jar cvfm MyApplet.jar manifest.mf mypackage/*.class
```
HTML文件的applet元素应当有属性archive="MyApplet.jar".
最后，对JAR文件签名。命令如下:
```shell
jarsigner -keystore keytorefile -tsa timestampURL MyApplet.jar keyalias
```

需要向证书发行商询问时间戳URL。密钥别名由证书发行商签名。运行以下命令:
keyool -keystore keystorefile -list
来得到密钥别名。还可以用keytool命令的-changealias选项改别名。
现在把签名的JAR文件和包含applet元素的HTML文件放到你的Web服务器上。
