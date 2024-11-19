#!/bin/bash

year=$1
day=$2

if [ -z "$1" -o -z "$2" ]; then
  echo "Usage: $0 <year> <day>"
  exit 1

fi
cd "${PWD}/$1/$2"

go run . 