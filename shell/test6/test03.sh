#!/bin/bash

find '.' -name '*.txt' -print -exec mv '{}' t ';' > files.f
