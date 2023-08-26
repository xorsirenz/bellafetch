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
pkg_total=$(pacman -Q |wc -l)

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

## output
clear

	printf '%s\n' "
          bellafetch
     [Github: madison-isa]

   host   :: ${HOST}
   os     :: ${PRETTY_NAME}
   ver    :: ${kernel}
   uptime :: ${uptime}
   pkgs   :: ${pkg_total}
   wm     :: ${wm}
   cpu    :: ${cpu}
   gpu    :: ${gpu}
   mem    :: ${mem}
"

exit 0
