#!/bin/bash

APP_NAME=httpclient

GO=/mnt/d/soft/go/bin/go.exe
UPX=/mnt/d/soft/upx-3.96-win64/upx.exe

$GO generate

rm -rf bin/*.exe
$GO build -ldflags "-H windowsgui -s -w" -o bin/tmp.exe
$UPX -9 -o bin/${APP_NAME}.exe bin/tmp.exe
rm -f bin/tmp.exe