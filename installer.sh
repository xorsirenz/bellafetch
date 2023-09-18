#!/bin/sh
if [ $(whoami) = "root" ]; then
  :
else
  echo -e "\033[31mError: \033[0mNot running as root!"
  exit 1
fi

if cp ./bellafetch.sh /usr/bin/bellafetch; then
  echo "Successfully installed!"
  exit 0 
else 
  echo "Something went wrong!"
fi
