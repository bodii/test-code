### if let 
if let 获取通过等号分隔的一个模式和一个表达式。它的工作方式与match相同，这里的表达式对应match而模式则对应第一个分支。
可以认为if let是match的一个语法糖，它当值匹配某一模式时执行代码而忽略所有其他值。

可以在if let中包含一个else。
else块中的代码与match表达式中的_分支块中的代码相同，这样的match表达式就等同于if let 和else。