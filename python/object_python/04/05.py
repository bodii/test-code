#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 使用小数上下文 '''

"""
小数上下文 ，它定义了很多decimal.Decimal计算的特性，
包含了大量的规则用来求近值或对值进行截取。
"""

import decimal

PENNY = decimal.Decimal('0.00')

price = decimal.Decimal('15.99')
rate = decimal.Decimal('0.0075')
print( 'Tax=', (price * rate).quantize(PENNY), 'Fully=', price * rate )
# Tax= 0.12 Fully= 0.119925


with decimal.localcontext() as ctx:
    ctx.rounding = decimal.ROUND_DOWN
    tax = (price * rate).quantize(PENNY)
    print( 'Tax=', tax )

# Tax= 0.11