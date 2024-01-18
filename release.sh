#!/bin/sh
ver=1.2.0
mkdir fstabfmt-$ver
cp -v *.go Makefile LICENSE README.md fstabfmt-$ver/
tar zcvf fstabfmt-$ver.tar.gz fstabfmt-$ver/
rm -r fstabfmt-$ver/
