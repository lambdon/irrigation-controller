#!/bin/bash


#enable the 4 gpio pins that we'll use

GPIODIR="/sys/class/gpio"

cd $GPIODIR

for item in $(seq 2 5)
do
	echo $item > export
	pushd gpio$item
	echo out > direction
	echo 0 > value
	popd
done
