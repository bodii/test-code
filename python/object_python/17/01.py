#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 质量和文档 '''

''' 为help()函数编写docstrings '''
"""
help()函数提取并显示文档注释，它会对文本执行一些基本的格式化操作。
help()函数通过site包安装在交互式Python环境中，这个功能实际上是定义在pydoc包中。
可以通过导入并且扩展这个包来自定义help()的输出。
"""

"""
编写文档所有必要的元素：总结、API和描述。
"""

''' 用pydoc编写文档 '''
"""
当用pydoc生成文档时，我们会以下面3种方式中的一种来使用它
1. 准备文本模式的文档文件，然后用命令行工具例如more或者less查看它们
2. 准备HTML文档并且保存在文件中以供之后的浏览使用
3. 运行一个HTTP服务器并且创建需要立刻浏览的HTML文件。
"""



''' 编写类和方法函数的文档字符串 '''
"""
我们会用字段列表技术在全局的类文档字符串中记录类变量。这通常会专注于使用
:ivar variable:、 ：cvar variable 和 :var variable:字段列表元素。
"""

class Card:
    """Definition of a numeric rank playing card.
    Subclasses will define "FaceCard" and "AceCard".
    :ivar rank: Rank
    :ivar suit: Suit
    :ivar hard: Hard point total for a card
    :ivar soft: Soft total; same as hard for all cards except Aces.
    """

    def __init__( self, rank, suit, hard, soft=None ):
        """Define the values for this card.

        :param rank: Numeric rank in the range 1-13. 
        :param suit: Suit object (often a character from '♠ ♣ ♥ ♦')
        :param hard: Hard point total (or 10 for FaceCard or 1 for AceCard)
        :param soft: The soft total for AceCard, otherwise defaults to hard.
        """
        self.rank = rank
        self.suit = suit
        self.hard = hard
        self.soft = soft if soft is not None else hard

c = Card(1, 2, 3, 4)


''' 编写函数文档字符串 '''

def card( rank, suit ):
    """Create a "Card" instance from rank and suit.
    :param rank: Numeic rank in the range 1-13. 
    :param suit: Suit object (often a character from '♠ ♣ ♥ ♦')
    :returns: Card instance
    :raises TypeError: rank out of range.

    >>> import p3_c18
    >>> p3_c18.card( 3, '♥')
    3♥
    """
    if rank == 1: return AceCard( rank, suit, 1, 11 )
    elif 2 <= rank < 11: return Card( rank, suit, rank )
    elif 11 <= rank < 14: return FaceCard( rank, suit, 10 )
    else:
        raise TypeError( 'rank out of range' )
