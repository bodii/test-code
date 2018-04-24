/**
 * 命名空间和模块
 * 
 */
/**
 * 模块解析
 */

 /*
 相对 vs. 非相对模块导入

 相对导入是以/, ./或../开头的：
 example:
 import Entry from './components/Entry';
 import { DefaultHeaders } from '../constants/http';
 import '/mod';
 */

 /*
 模块解析策略
 共有两种可用的模块解析策略：Node和Classic。可以使用 --moduleResolution标记来指定使用哪种模块解析策略。
 若未指定， 那么在使用了 --module AMD | System | ES2015时的默认值为Classic，其它情况时则为None.


 Classic
 这种策略在以前是TypeScript默认的解析策略。现在，它存在的理由主要是为了向后兼容。

 相对导入的模块是相对于导入它的文件进行解析的。

 对于非相对模块的导入，编译器则会从包含导入文件的目录开始依次向上级目录遍历，尝试定位匹配的声明文件

 examplea:
有一个对moduleB的非相对导入import { b } from 'moduleB', 它是在/root/src/folder/A.ts文件里，会以如下
的方式定位'moduleB':
    1. /root/src/folder/moduleB.ts
    2. /root/src/folder/moduleB.d.ts
    3. /root/src/moduleB.ts
    ...
    8. /moduleB.d.ts
    9. /moduleB.ts
*/


/*
路径映射

TypeScript编译器通过使用tsconfig.json文件里的"paths"来支持这样的声明映射。

// example　jquery的paths
{
    'compilerOptions': {
        'baseUrl': '.',
        'paths': {
            'jquery': ['node_modules/jquery/dist/jquery'] // 此处映射是相对于'baseUrl'
        }
    }
}

通过“paths”可以指定复杂的映射，包括指定多个回退位置。
工程结构可能如下：
projectRoot
|-- folder1
|   |-- file1.ts
|   └── file2.ts
├── generated
│   ├── folder1
│   └── folder2
│       └── file3.ts
└── tsconfig.json


相应的tsconfig.json文件：
{
    'complierOptions': {
        'baseUrl': '.',
        'paths': {
            '*': [
                '*',
                'generated/*'
            ]
        }
    }
}

它告诉编译器所有匹配'*'(所有的值)模式的模块导入会在以下两个位置查找：
    1. '*': 表示名字不发生改变，所以映射为<moduleName> => <baseUrl>/<moduleName>
    2. 'generated/*'表示模块名添加了'generated'前缀，所以映射为<moduleName> => <baseUrl>/generated/<moduleName>

*/


/*
利用rootDirs指定虚拟目录

有时多个目录下的工程源文件在编译时会进行合并放在某个输出目录下。这可以看做一些源目录创建了一个“虚拟”目录

如下：
src
 └── views
     └── view1.ts (imports './template1')
     └── view2.ts

 generated
 └── templates
         └── views
             └── template1.ts (imports './view2')


：tsconfig.json:
{
  "compilerOptions": {
    "rootDirs": [
      "src/views",
      "generated/templates/views"
    ]
  }
}

设想这样一个国际化的场景，构建工具自动插入特定的路径记号来生成针对不同区域的捆绑，
比如将#{locale}做为相对模块路径./#{locale}/messages的一部分。在这个假定的设置下，
工具会枚举支持的区域，将抽像的路径映射成./zh/messages，./de/messages等。


假设每个模块都会导出一个字符串的数组：如 ./zh/message可能包含：
export default [
    '您好吗'，
    '很高兴认识你'
];

利用rootDirs我们可以让编译器了解这个映射关系，从而也允许编译器能够安全地解析./#{locale}/messages，就算这个目录永远都不存在。比如，使用下面的tsconfig.json：
{
  "compilerOptions": {
    "rootDirs": [
      "src/zh",
      "src/de",
      "src/#{locale}"
    ]
  }
}

编译器现在可以将import messages from './#{locale}/messages'解析为import messages from './zh/messages'用做工具支持的目的，并允许在开发时不必了解区域信息。




跟踪模块解析

如之前讨论，编译器在解析模块时可能访问当前文件夹外的文件。 这会导致很难诊断模块为什么没有被解析，或解析到了错误的位置。 通过 --traceResolution启用编译器的模块解析跟踪，它会告诉我们在模块解析过程中发生了什么。

假设我们有一个使用了typescript模块的简单应用。 app.ts里有一个这样的导入import * as ts from "typescript"。

│   tsconfig.json
├───node_modules
│   └───typescript
│       └───lib
│               typescript.d.ts
└───src
        app.ts



使用--noResolve

正常来讲编译器会在开始编译之前解析模块导入。 每当它成功地解析了对一个文件 import，这个文件被会加到一个文件列表里，以供编译器稍后处理。

--noResolve编译选项告诉编译器不要添加任何不是在命令行上传入的文件到编译列表。 编译器仍然会尝试解析模块，但是只要没有指定这个文件，那么它就不会被包含在内。

比如
app.ts

import * as A from "moduleA" // OK, moduleA passed on the command-line
import * as B from "moduleB" // Error TS2307: Cannot find module 'moduleB'.

tsc app.ts moduleA.ts --noResolve

使用--noResolve编译app.ts：

    可能正确找到moduleA，因为它在命令行上指定了。
    找不到moduleB，因为没有在命令行上传递。


*/

