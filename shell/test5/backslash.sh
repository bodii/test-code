#!/bin/bash


echo 
echo Enter how many books do you want to buy:
echo


echo The First Book \<\< Linux Shell Programming \>\>

echo Price:\$20
read -p "Order Number:" FIRST_NUM

echo
echo The Second Book \<\< Linux C language Programmin \>\>
echo Price:\$50
read -p "Order Number:" SECOND_NUM

if [ $FIRST_NUM -gt $SECOND_NUM ]
then
	echo Oredr Numbers: $FIRST_NUM \> $SECOND_NUM
	echo The first book is more.
elif [ $FIRST_NUM -lt $SECOND_NUM ]
then
	echo Order Numbers: $FIRST_NUM \< $SECOND_NUM
	echo The second book is more.

elif [ $FIRST_NUM -eq $SECOND_NUM ]
then
	echo Order Numbers: $FIRST_NUM = $SECOND_NUM
	echo the price of two book is equal.
fi

exit 0

