## Homebrew
- https://github.com/Homebrew
- https://brew.sh
- https://github.com/Homebrew/brew
- https://github.com/Homebrew/homebrew-cask
- https://formulae.brew.sh
- https://formulae.brew.sh/formula/


## 命令

```
brew 安装的程序可以通过

brew services -h 帮助命令

brew services list 查看所有安装

brew services start APP_NAME

brew services stop APP_NAME

brew services restart APP_NAME

brew安装的程序配置文件都放在类似目录下

/usr/local/etc/APP_NAME
```

### brew 使用代理安装软件

all_proxy=socks5://127.0.0.1:1090 brew install APP_NAME

或在 bash 中设置 brew 别名

alias brew="all_proxy=socks5://127.0.0.1:1090 brew”

### Version
Homebrew 4.0.0
- https://sspai.com/post/78587
