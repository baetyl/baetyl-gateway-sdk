#!/bin/bash
set -e

ARCH=`uname -m`

cd "$(cd "$(dirname "$0")" && pwd)"
rm -rf build-$ARCH
rm -rf install-$ARCH

mkdir build-$ARCH
cd build-$ARCH
cmake .. -DCMAKE_INSTALL_PREFIX=../install-$ARCH
make -j4
make install
