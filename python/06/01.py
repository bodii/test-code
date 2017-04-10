#!/usr/bin/env python3
# -*- coding:utf-8 -*-

# getopt -- 从命令行中得到选项

import sys
import getopt

# Remember, the first thing in the sys.argv list is the name of the command
# You don't need that.

cmdline_params = sys.argv[1:]
print(cmdline_params)

opts, args = getopt.getopt(cmdline_params, 'hc:', ['help', 'config='])

for option, parameter in opts:
    
    if option == '-h' or option == '--help':
        print("This program can be run with either -h or --help for this"
                +"message,")
        print("or with -c or --config=<file> to specify a different" +
        "configuration file")

    if option in ('-c', '--config'): # this means the same as the above
        print("Using configuration file %s" % parameter)



