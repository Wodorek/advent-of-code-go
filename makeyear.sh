#!/bin/bash

year=$1

if [ -z "$1" ]; then
  echo "Usage: $0 <year>"
  exit 1
fi

mkdir $year

cd "${PWD}/$1"

counter=1

while [ $counter -le 25 ]
do
  if [[ $counter -le 9 ]]
  then
    mkdir "0$counter"
    cd "0$counter"

  else
    mkdir "$counter"
    cd "$counter"
  fi
    touch main.go
    touch input.txt
    cd ..
  ((counter++))
done