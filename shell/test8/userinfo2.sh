#!/bin/bash

echo
read -p "Enter user name (or quit): "

echo
until [ "$REPLY" = "quit" ] || [ "$REPLY" = "Quit" ] || [ "$REPLY" = "q" ]
do
	echo
	awk '
		BEGIN{
			FS=":";
		}
	
		$1 == USERNAME {
			printf "User:%-8s UID:%-5s GID:%-5s Home Dir:%-5s shell:%-10s\n",$1,$3,$4,$6,$7;
		}
	
		' USERNAME="$REPLY" /etc/passwd 
	
	echo 
	read -p "Enter your name (or quit): "
done


echo "Bye."

exit 0

