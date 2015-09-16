#!/bin/bash

for package in `ls -1 src`
do
      go install $package
  done
