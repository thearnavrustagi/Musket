#!/usr/bin/bash
if [[ $1 == "" ]]
then
	./main
else
	./main $1 $2
fi
