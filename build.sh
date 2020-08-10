#!/bin/bash

GO=/mnt/d/soft/go/bin/go.exe
UPX=/mnt/d/soft/upx-3.96-win64/upx.exe

rm -rf bin/*.exe
$GO build -ldflags "-H windowsgui -s -w" -o bin/tmp.exe
$UPX -9 -o bin/httpclient.exe bin/tmp.exe
rm -f bin/tmp.exe