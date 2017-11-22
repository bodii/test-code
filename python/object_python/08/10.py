#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 针对可靠的pickle处理进行类设计 '''

import builtins
import pickle

class RestrictedUnpickler(pickle.Unpickler):
    def find_class( self, module, name ):
        if module == 'builtins':
            if name not in ('exec', 'eval'):
                return getattr(builtins, name)

        elif module == '__main__':
            return globals()[name]
        # elif module in any of our application modules...
        raise pickle.UnpicklingError(
            "global '{module}.{name}' is forbidden".format( 
                module = module,
                name = name 
                )
        )