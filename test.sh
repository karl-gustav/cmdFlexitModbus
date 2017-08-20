#! /bin/bash
set -e

echo "Bulding..."
./build.sh

echo "Copying to server..."
scp cmdFlexitModbus root@smart:~/smart

echo "SSH-ing into server..."
ssh root@smart 'echo "Running code..." && ./smart'


