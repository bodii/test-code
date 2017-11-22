#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 参数和选项 '''

"""
选项和参数的基本规则：
1. 选项先出现。它们以-或--为前缀。有两种格式：-letter和--word。有两种选项：
带参数选项和无参选项。无参选项的例子是用-v或者--version显示版本。带参选项的一个
个参数是-m module, -m选项后面必须带有一个模块名。

2. 不带参数的短格式（单一字母）选项可以组合在单个的-之后。方便起见，我们可以用-bqv作
为-b -q -v 选项的组合。

3. 参数最后出现。它们没有前置的-或者--。有两种常用的参数：
    3.1 对于位置参数，位置是有意义的。如一个是输入文件名和一个输出文件名
    3.2 参数列表，这些参数在语义上是平等的。当我们说process.py *.html时，shell会将
    *.html命令扩展为代表文件名的位置参数
"""

''' ArgumentParser对象 '''

"""

class argparse.ArgumentParser(prog=None, usage=None, description=None, 
epilog=None, parents=[], formatter_class=argparse.HelpFormatter, 
prefix_chars='-', fromfile_prefix_chars=None, argument_default=None, 
conflict_handler='error', add_help=True, allow_abbrev=True)

    创建一个新的 ArgumentParser 对象。所有参数都应作为关键字参数传递。每个参数都有
    自己更详细的描述如下，但总之他们是：

        prog - 程序的名称（默认值：sys.argv[0]）

        usage - 描述程序用法的字符串（默认：从添加到解析器的参数生成）

        description - 在参数help之前显示的文本（默认值：none）

        epilog - 在参数help后显示的文本（默认值：none）

        parents - 还应包括其参数的 ArgumentParser 对象的列表

        formatter_class - 用于自定义帮助输出的类

        prefix_chars - 前缀可选参数的字符集（默认值：’ - ‘）

        fromfile_prefix_chars - 应该读取其他参数的文件前缀的字符集（默认值：None）

        argument_default - 参数的全局默认值（默认值：None）

        conflict_handler - 解决冲突可选项的策略（通常不必要）

        add_help - 为解析器添加一个-h/–help选项（默认值：True）

        allow_abbrev - 如果缩写是无歧义的，允许缩短长选项。 （默认：True）

    在 3.5 版更改: 添加了 allow_abbrev 参数。

"""


''' add_argument()方法的使用 '''
"""
add_argument()方法

ArgumentParser.add_argument(name or flags...[, action]
[, nargs][, const][, default][, type][, choices][, required]
[, help][, metavar][, dest])

    定义应如何解析单个命令行参数。每个参数都有自己更详细的描述如下，但总之他们是：

        name or flags - 名称或选项字符串列表，例如 foo 或 -f, --foo。

        action - 在命令行遇到此参数时要执行的操作的基本类型。
            -- 'store' - 这只是存储参数的值。这是默认操作。
            -- 'store_const' - 存储由 const 关键字参数指定的值。
            -- 'store_true' 和 'store_false' - 这些是用于分别存储值
              True 和 False 的 'store_const' 的特殊情况。此外，它们分别创建
              False 和 True 的默认值。
            -- 'append' - 这存储一个列表，并将每个参数值附加到列表。这对于允许多
              次指定选项很有用。
            -- 'append_const' - 存储一个列表，并将 const 关键字参数指定的值追
              加到列表中。 （注意，const 关键字参数默认为 None。）
            -- 'count' - 这计算关键字参数发生的次数。
            -- 'help' - 这将为当前解析器中的所有选项输出完整的帮助消息，然后退出。
            您还可以通过传递Action子类或实现相同接口的其他对象来指定任意操作。推荐的方法是扩展 Action，覆盖 __call__ 方法和可选的 __init__ 方法。

            自定义操作的示例:
            >>> class FooAction(argparse.Action):
            ...     def __init__(self, option_strings, dest, nargs=None, **kwargs):
            ...         if nargs is not None:
            ...             raise ValueError("nargs not allowed")
            ...         super(FooAction, self).__init__(option_strings, dest, **kwargs)
            ...     def __call__(self, parser, namespace, values, option_string=None):
            ...         print('%r %r %r' % (namespace, values, option_string))
            ...         setattr(namespace, self.dest, values)

        nargs - 应该使用的命令行参数的数量。

        const - 某些 action 和 nargs 选择所需的常量值。

        default - 如果参数在命令行中不存在，则生成的值。

        type - 应转换命令行参数的类型。

        choices - 参数的允许值的容器。

        required - 是否可以省略命令行选项（仅可选）。

        help - 参数的作用的简要描述。

        metavar - 使用消息中参数的名称。

        dest - 要添加到 parse_args() 返回的对象的属性的名称。


"""