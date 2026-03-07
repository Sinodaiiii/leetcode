#!/bin/bash

i=$(seq 50)
sum=0
for num in $i; do
  ((sum+=num))
done
echo $sum


seq 50 | awk '{sum+=$1} END {print sum}'