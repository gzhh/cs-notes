## 常用快捷键

超级简单文本编辑器:

nano text.txt

ctrl + X

Y/N

Enter

### shell
- 参考：https://www.man7.org/linux/man-pages/man3/readline.3.html#EDITING_COMMANDS
- 参考help命令：stty -a

+ a 将光标移动到行首
+ e 将光标移动到行尾
+ f 将光标移动到后一个字符
+ b 将光标移动到前一个字符

+ l 清除屏幕只保留当前行
+ r 在历史命令中搜索

Ctrl+W 删除前一个单词

Ctrl+U 删除整行

Ctrl+C 终止目前的命令

Ctrl+Z 暂停目前命令

Ctrl+D 输入结束

Ctrl+M 就是Enter

Ctrl+S 暂停屏幕输出

Ctrl+Q 恢复屏幕输出


输出重定向

ifconfig > test.log

cat test.log

ls >> test.log 2>&1

cat test.log
