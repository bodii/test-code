#!/bin/bash

# ==================================
#                               扫雷游戏
# ==================================
# 游戏必须的几个部分：
# 1. 显示雷区
# 2. 创建游戏逻辑
# 3. 创建判断单元格是否可选的逻辑
# 4. 记录可用和已查明（已排雷）单元格的个数
# 5. 创建游戏结束逻辑
# --------------------------------------------------------


# --------------------------------------------------------
# [  玩法 ]
# 在扫雷中，游戏界面是一个由2D数组(列和行)组成的不透明小方格。每一格下都有可能藏有地雷。玩家的任务
# 是找到这些不含雷的方格，并且在这一过程中，不能点到雷。
# --------------------------------------------------------

score=0
a="1 10 -10 -1"
b="-1 0 1"
c="0 1"
d="-1 0 1 -2 -3"
e="1 2 20 21 10 0 -10 -20 -23 -2 -1"
f="1 2 3 35 30 30 22 10 0 -10 -20 -25 -30 -35 -3 -2 -1"
g="1 4 6 9 10 15 20 30 -30 -24 -11 -10 -9 -8 -7"
RED="\033[0;31m"
GREEN="\033[0;32m"
NC="\033[0m" # no color

# 声明
declare -a room # 声明一个room数组，他用来表示雷区的每一格

function usage() {
    cat <<_EOF_
    This game is made for fun, it has simple rules.
    * Do not cheat.
    * You can re-select an element if only it's - empty.
    * if even, double digits are given, only the
        first will be considered.
    * When the game is over, it is over.

    Shown down here is a metrics 10x10 in size and
    to play you have to enter the coordinates.

    NOTE: To choose col- g, row- 5, give input - g5

_EOF_
}

function time_to_quit() {
    printf '\n\n%s\n\n' "info: Sadly! You opted to quit!!"
    exit 1
} 


function plough() {
    r=0
    printf '\n\n'
    printf '%s'  "     a   b   c   d   e   f   g   h   i   j"
    printf '\n   %s\n' "-----------------------------------------"
    for row in $(seq 0 9);
    do
        printf '%d  ' "$row"
        for col in $(seq 0 9); do
        ((r+=1))
        is_null_field $r
        printf '%s \e[33m%s\e[0m ' "|" "${room[$r]}"
        done
        printf '%s\n' "|"
        printf '   %s\n' "-----------------------------------------"
    done
    printf '\n\n'
}

function is_null_field () {
    local e=$1 #在数组room中，我们已经用过循环变量'r'了，这次我们用 'e'
    if [[ -z "${room[$e]}" ]]
    then
        room[$r]='.'
    fi
}

function get_mines() {
	m=$(shuf -e a b c d e f g X -n 1)
	if [[ "$m" != "X" ]]
	then
		for Limit in ${!m}
		do
			field=$(shuf -i 0-5 -n 1)
			index=$((i+Limit))
			is_free_field $index $field
		done
	elif [[ "$m" = "X" ]]
	then
		g=0
		room[$i]=X
		for j in {42..49}
		do
			out="gameover"
			k=${out:%g:1}
			room[$j]=${k^^}
			((g+=1))
		done
	fi
}

function get_free_fields () {
    free_fields=0
    for n in $(seq 1 ${#roo[@]})
    do
        if [[ "${room[$n]}" = "." ]]
        then
            ((free_fields+=1))
        fi
    done
}

# plough


# --------------------------------------------------------
# 创建玩家逻辑
# --------------------------------------------------------
function get_coordinates() {
    colm=${opt:0:1}     # 得到第一个字符，一个字母
    ro=${opt:1:1}           # 得到第二个字符，一个整数
    case $colm in 
        a ) o=1;;
        b ) o=2;;
        c ) o=3;;
        d ) o=4;;
        e ) o=5;;
        f  ) o=6;;
        g ) o=7;;
        h ) o=8;;
        i  ) o = 9;;
        j  ) o=10;;
    esac

    i=$(((ro*10)+o))    # 遵循运算规则，算出最终值
    is_null_field $i $(shuf -i 0-5 -n 1)    # 调用自定义函数，判断其指向空/可选择单元格
    if [[ $not_allowed -eq 1 ]] || [[ ! "$colm" =~ [a-j] ]]
    then
        printf "$RED \n%s: %s \n$NC" "warning" "not allowed!!!!!"
    else
        get_mines
        plough
        get_free_fields
        if [[ "$m" = "X" ]]
        then
			printf "\n\n\t $RED%s: $NC %s %d\n" "GAME OVER" "your scored" "$score"
			printf "\n\n\t%s\n\n" "You were just $free_fields mines away."
			exit 0
        elif [[ $free_fields -eq 0 ]] 
        then
			printf "\n\n\t $GREEN%s: %s $NC %d\n\n" "You Win" "your scored" "$score"
			exit 0
        fi
    fi
}

# --------------------------------------------------------
# 3. 创建判断单元格是否可选的逻辑
# --------------------------------------------------------
function is_free_field() {
    local f=$1
    local val=$2
    not_allowed=0
    if [[ "${room[$f]}" = "." ]]
    then
        room[$f]=$val
        score=$((score+val))
    else
        not_allowed=1
    fi
}

# --------------------------------------------------------
# --------------------------------------------------------

# --------------------------------------------------------
# --------------------------------------------------------

# --------------------------------------------------------
# --------------------------------------------------------

# main
trap time_to_quit INT
printf "\e[2J\e[H"
usage
read -p "Type Enter to continue. And good luck!"
plough
while true
do
	printf "Remember: to choose col- g, row - 5, give input - g5 \n\n"
	read -p "info: enter the coordinates: " opt
	get_coordinates
done

