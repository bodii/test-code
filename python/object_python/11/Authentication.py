#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 创建安全的REST服务 '''

from hashlib import sha256
import os

class Authentication:
    iterations = 1000

    def __init__( self, username, password):
        """Works with bytes. Not unicode string."""
        self.username = username
        self.salt = os.urandom(24)
        self.hash = self._iter_hash( 
            self.iterations,
            self.salt,
            username,
            password
            )

    @staticmethod
    def _iter_hash( iterations, salt, username, password ):
        seed = salt+b':'+username+b':'+password
        for i in range(iterations):
            seed = sha256(seed).digest()
        return seed

    def __eq__( self, other ):
        return self.username == other.username and self.hash == other.hash

    def __hash__( self, other ):
        return hash(self.hash)

    def __repr__( self ):
        salt_x = ''.join('{0:x}'.format(b) for b in self.salt)
        hash_x = ''.join('{0:x}'.format(b) for b in self.hash)
        return '{username} {iterations:d}:{salt}:{hash}'.format(
            username = self.username,
            iterations = self.iterations,
            salt = salt_x, hash = hash_x
        ) 

    def match( self, password ):
        test = self._iter_hash( self.iterations, self.salt, self.username, password)
        return self.hash == test

