#!/bin/bash

for file in "$@"
do
	if [ -f "$file" ]
	then
		echo "**processing $file**"
		sed -e 's/jerry/emma/g' -e 's/2013/2012/g' "$file" > "$file.$$"

		if [ -f "$file.$$" ]
		then
			mv -f "$file.$$" "$file"
		fi
	fi
done

echo "All Done."

exit 0
