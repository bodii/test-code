#!/usr/bin/env python3
# -*- coding:utf-8 -*-


from itsdangerous import TimedJSONWebSignatureSerializer as Serializer
import os
from global_config.app import default_config


class TokenCase:
    def get_token(self, data=dict(), expires=3600):
        s = Serializer(
            default_config.SECRET_KEY, 
            expires_in = expires 
            )
        return s.dumps(data)

    def checkToken(self, token='', id=0):
        s = Serializer(
            default_config.SECRET_KEY,
        )
        try:
            data = s.loads(token)
            if id > 0 and int(data.get('confirm')) == id:
                return True
            return False
        except:
            return False
        

            
    