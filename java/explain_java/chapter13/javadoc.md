### javadoc

#### 文档注释
文档注释只有放在类、接口、构造函数、方法、字段的声明前面才会被识别。

你下面这样，如果在import声明的前面书写注释，将会被看作是类的注释(被略过):
```java
/** 类Day是表示日期的类。*/
import java.util.*;

class Day {
}
```

注释中可以直接使用HTML标签。例如，<b>和</b>括起来的部分是粗体。

#### @auther "作者"
> 一个@auther中可以书写多个"作者名", 也可以为每个"作者名"都加上@author.

#### {@code 代码}
表示此处为程序代码

#### @return "返回值"

#### @param "参数名" "说明"

#### @see "引用目标"

#### * @see "字符串"
添加"字符串",但不生成链接。

#### * @see <a href = "URL#value">lable</a>
添加URL#value定义的链接。URL＃value是相对URL或者绝对URL。

#### * @see package.class#member lable
添加指向与特定名称的成员相关的文档的链接，以及显示的文本lable。label可以省略。当省略
label时，链接目标的成员名称在显示时会进行适当缩减。


### javadoc工具
javadoc工具会基于源程序中记述的文档注释来创建文档。该工具的启动按如下方式进行:
```
javadoc 选项 包名 源文件 @参数文件
```
