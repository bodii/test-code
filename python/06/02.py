#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import sys
import getopt

cmdline_params = sys.argv[1:]

opts, args = getopt.getopt(cmdline_params, 'hc:', ['help', 'config='])

print(args)
for option, parameter in opts:
    if option in ('-h', '--help'):
        print('This parogram can be run with either -h or --help for this' +
                'message,')
        print('or with -c or --config=<file> to specify a different' +
                'configuration file')

    if option in ('-c', '--config'):
        print('Using configuration file %s' % parameter)
