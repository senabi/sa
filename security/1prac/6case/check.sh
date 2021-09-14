#!/bin/bash

while IFS="" read -r p || [ -n "$p" ]
do
  arr=($p)
  echo "${arr[0]} $(rg -co ${arr[0]} ./pi.txt)"
done < res.txt
