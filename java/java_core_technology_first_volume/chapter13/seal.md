#### 密封
可以将Java包密封(seal) 以保证不会有其他的类加入到其中。
如果在代码中使用了包可见的类、方法和域，就可能希望密封包。如果不密封，其他类就有可能
放在这个包中，进而访问包可见的特性。

例如，如果密封了com.mycompany.util包，就不能用下面的语句顶替密封包之个的类:
package com.myCompany.util;
要想密封一个包，需要将包中的所有类放到一个JAR文件中。在默认情况下，JAR文件中的包是没有
密封的。可以在清单文件的主节中加入下面一行:
```
Sealed: true
```
来改变全局的默认设定。对于每个单独的包，可以通过在JAR文件的清单中增加一节，来指定是否
想要密封这个包:
```
Name: com/mycompany/util/
Sealed: true

Name: com/mycompany/misc/
Sealed: false
```
要想密封一个包，需要创建一个包含清单指令的文本文件。然后用常规的方式运行jar命令:
```shell
jar cvfm MyArchive.jar manifest.mf files to add 
```
