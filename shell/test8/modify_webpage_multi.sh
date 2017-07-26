#!/bin/bash

for file in "$@"
do
	if [ -f $file ]
	then
		echo "**processing $file**"

		sed 's/emma/jerry/g;s/2012/2013/g' "$file" > "$file.$$"

		if [ -f "$file.$$" ]
		then
			mv -f "$file.$$" "$file"
		fi
	fi
done

echo "All Done."

exit 0
