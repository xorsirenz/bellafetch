#!/bin/sh

## system infomration tool
## created by madison

## hostname
HOST=$USER@$HOSTNAME

## os
. /etc/os-release

## kernel
read -r _ _ version _ < /proc/version
kernel=${version%%-*}

## uptime
uptime=$(uptime -p |cut -d' ' -f2-)

## packages 
pkg_total=$(dpkg-query -l 2>/dev/null | grep "^ii")
pkg_total=$(yum list installed 2>/dev/null)
pkg_total=$(pacman -Q 2>/dev/null |wc -l)

## wm detection
id=$(xprop -root -notype _NET_SUPPORTING_WM_CHECK)
id=${id##* }
wm=$(xprop -id "$id" -notype -len 100 -f _NET_WM_NAME 8t |grep '^_NET_WM_NAME' |cut -d\" -f 2)

## memory
mem=$(free -m |awk 'FNR==2 {print $6"MiB / "$2"MiB"}')

## cpu
cpu=$(grep -Po 'model name.*: \K.*' /proc/cpuinfo | uniq | sed -E 's/\([^)]+\)//g')

## gpu
gpu=$(lspci |grep -i vga |cut -d' ' -f5-9)

## storage
storage=$(df -Ph . | tail -1 | awk '{print $4"iB / "$2"iB"}')

## output
clear

	printf '%s\n' "
           bellafetch
      [Github: bootlegwifi]

   host    :: ${HOST}
   os      :: ${PRETTY_NAME}
   ver     :: ${kernel}
   uptime  :: ${uptime}
   pkgs    :: ${pkg_total}
   wm      :: ${wm}
   cpu     :: ${cpu}
   gpu     :: ${gpu}
   storage :: ${storage}
   mem     :: ${mem}
"

exit 0
