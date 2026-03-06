#!/bin/bash


a=$(awk 'BEGIN {
  for (i=7; i<10; i++) {
    print "8.8.8." i
  }
}')
for i in $a; do
  ping -c 2 -i 0.1 $i > /dev/null
  if [ $? == 0 ]; then
    echo "node $i is yes"
  else
    echo "node $i is no"
  fi
done


awk 'BEGIN {
  for (i=7; i<10; i++) {
    print "8.8.8." i
  }
}' | \
while read i; do
  ping -c 2 -i 0.1 $i > /dev/null
  if [ $? -eq 0 ]; then
    echo "node $i is yes"
  else
    echo "node $i is no"
  fi
done



#awk 'BEGIN {
#  for (i=1; i<10; i++) {
#    print "8.8.8." i
#  }
#}' | \
#xargs -I {} -P 10 sh -c "
#    ping -c 2 -i 0.3 {} > /dev/null
#    if [ \$? -eq 0 ]; then
#        echo '{} is yes'
#    else
#        echo '{} is no'
#    fi
#"
