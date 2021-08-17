# robotclick
Use robotgo to implement this robot click app, to do some boring and repetitive tasks

# Cross-compile

## Logs

### Ubuntu 18.04.5 LTS
```
sudo apt install gcc libc6-dev

sudo apt install libx11-dev xorg-dev libxtst-dev libpng++-dev

sudo apt install xcb libxcb-xkb-dev x11-xkb-utils libx11-xcb-dev libxkbcommon-x11-dev
sudo apt install libxkbcommon-dev

sudo apt install xsel xclip

sudo apt-get install gcc-multilib
sudo apt-get install gcc-mingw-w64
# Ubuntu solution: fatal error: zlib.h: No such file or directory
sudo apt install libz-mingw-w64-dev
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -x ./
GOOS=windows GOARCH=386 CGO_ENABLED=1 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -x ./
```
