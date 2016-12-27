#!/bin/bash

checkip (){
	case $1 in
		"" | *[!0-9.] | *.[!0-9])
			return 1
			;;
	esac

	local IFS=.

	set -- $1

	if [ $# -eq 4 ] && [ ${1:-999} -lt 255 ] && [ ${2:-999} -lt 255 ] && \
		[ ${3:-999} -lt 255 ] && [ ${4:-999} -lt 255 ]
	then
		return 0
	else
		return 1
	fi
}

for IP in 127.0.0.1 127.0.0.0.1 .0.0.1 ...1 172.299.38.1 172.38.1 \
	192.168.10.1 127.0.0.
do
	checkip ${IP}

	if [ $? -eq 0 ]
	then
		printf "%15s: %s\n" "${IP}" "valid ip address"
	else
		printf "%15s: %s\n" "${IP}" "invalid ip address"
	fi
done


while [ true ]
do
	echo -n " Please enter IP address you want to \
check(type \"quit\" to finish):"
	read IPADDRESS
	if [ "$IPADDRESS" = "quit" ]
	then
		break;
	fi

	checkip ${IPADDRESS}

	if [ $? -eq 0 ]
	then
		printf "%15s: %s\n" "${IPADDRESS}" "valid ip address"
	else
		printf "%15s: %s\n" "${IPADDRESS}" "invalid ip address"
	fi
done

exit 0
