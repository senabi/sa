#!/bin/bash

awk '{print toupper($0)}' < ../in.txt | \
    sed 's/[JH]/I/g;s/Ñ/N/g;s/K/L/g;s/[UÚW]/V/g;s/Y/Z/g' | \
    iconv -f utf8 -t ascii//TRANSLIT//IGNORE | \
    sed -z 's/[.,;:!? ]//g;s/\n//g' | \
    ./main -T
