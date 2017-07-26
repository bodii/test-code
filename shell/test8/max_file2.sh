#!/bin/bash

echo
read -p "Find the biggest file in which directory(or quit):"

until [ "$REPLY" = "quit" ] || [ "$REPLY" = "Quit" ] || [ "$REPLY" = "q" ]
do
	if [ -d "$REPLY" ] && [ -r "$REPLY" ] && [ -x "$REPLY" ]
	then
		ls -lR "$REPLY" | awk '
		$1 !~/总用量/{
				if($5 > max_size){
					max_size = $5;
					max_file = $9;
				}
			}
		END{
			printf("In directory %s,\nFile %s is the biggest,has %d bytes.\n",dir,max_file,max_size);
		}

		' dir=$REPLY

	else
		echo "$REPLY can not be accessed!"
	fi
	
	echo 
	read -p "Find the biggest file in which directory(or quit):"
done

	echo "Bye."

	exit 0
