# robotclick
Use robotgo to implement this robot click app, to do some boring and repetitive tasks

# Configurations
Global setting is `configs/configs.json`, cron is the task setting, tolerance is set for bitmap finding  

**NOTICE**: First, you should config `configs.json` -> `snippets` -> `files` to make sure which snippet should be invoked.

`configs/snippets/test1.json` is just a example snippet, you can generate yours refer: https://github.com/go-vgo/robotgo  

- `name` is the locate map's file name
- `action` is what do you want at this step, click, type or input
- `offset` used in action click, for mouse move, **it is a relative position**
- `double` as you know, if true will make mouse left double click
- `delay` make step nap a bit to wait your process.
- `keys` is a key type step list, have two forms, one for type one key and another for combination types, <kbd>Ctrl</kbd>+<kbd>v</kbd> can be set to: `{"key":"v", "attr":["ctrl"]},` in windows or `{"key":"v", "attr":["command"]},` in MacOS
- `msg` is the message you wanna input

# Cross-compile

## Logs

### Ubuntu 18.04.5 LTS
```
sudo apt install gcc libc6-dev

sudo apt install libx11-dev xorg-dev libxtst-dev libpng++-dev

sudo apt install xcb libxcb-xkb-dev x11-xkb-utils libx11-xcb-dev libxkbcommon-x11-dev
sudo apt install libxkbcommon-dev

sudo apt install xsel xclip

sudo apt install gcc-multilib
sudo apt install gcc-mingw-w64
# Ubuntu solution: fatal error: zlib.h: No such file or directory
sudo apt install libz-mingw-w64-dev
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -x ./
GOOS=windows GOARCH=386 CGO_ENABLED=1 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -x ./
```
