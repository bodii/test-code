#!/bin/bash


#access a non exist variable,echo nothing
echo $THERE_IS_NO_THIS_VARIABLE

#assign something into this variable
THERE_IS_NO_THIS_VARIABLE="there is not this variable"

# echo this variable again, output is not empty
echo $THERE_IS_NO_THIS_VARIABLE
