#!/bin/bash

function hmm() {

cat << EOF
Invoke ".build/envsetup.sh" from your shell to add the following functions
to your environment:
- croot:  Changes directory to the top of the tree.
- m:      Makes from the top of the tree.
- mm:     Builds all of the modules in the crrent directory.
- mmm:    Builds all of the modules in the supplied directory.
- cgrep:  Greps on all local C/C++ files.
- jgrep:  Greps on all local Java files.
- resgrep:Greps on all local res/*.xml files.
- godir:  Go to the directory containing a file.
Look at the source to view more functions. The complete list is:
#delimiter
EOF
T=$(gettop)
local A
A=""
for i in `cat $T/build/envsetup.sh | sed -n "/^function /s/function \([a-z_]*\).*/\1/p" | sort`;
do
	A="$A $i"
done
echo $A
}
