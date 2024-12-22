## 常用快捷键

### shell
- 参考：https://www.man7.org/linux/man-pages/man3/readline.3.html#EDITING_COMMANDS
- 参考help命令：stty -a

光标移动
- Ctrl + A 将光标移动到行首
- Ctrl + E 将光标移动到行尾
- Ctrl + F 将光标移动到后一个字符
- Ctrl + B 将光标移动到前一个字符
- Option + F 将光标移动到后一个单词（需要在终端中配置 Left/Right Option Key 为 Meta/Esc+）
- Option + B 将光标移动到前一个单词

文本编辑
- Ctrl + D 删除光标后的字符
- Ctrl + H 删除光标前的字符
- Ctrl + W 删除光标前的一个单词
- Option + Backspace 删除光标前的一个单词
- Option + D 删除光标后的一个单词
- Ctrl + K 删除光标后到行尾的所有内容
- Ctrl + U 删除整行

搜索与历史记录
- Ctrl + R 在历史命令中搜索
- Ctrl + G 退出历史搜索模式
- Ctrl + P 向上浏览历史命令（相当于向上箭头）
- Ctrl + N 向下浏览历史命令（相当于向下箭头）

屏幕与输出
- Ctrl + L 清除屏幕只保留当前行
- Ctrl + S 暂停屏幕输出
- Ctrl + Q 恢复屏幕输出

进程与控制
- Ctrl + C 终止当前正在运行的进程
- Ctrl + Z 将当前进程挂起（放入后台）
- Ctrl + M 就是 Enter


### 其他
> 超级简单文本编辑器

nano text.txt

ctrl + X

Y/N

Enter


> 输出重定向

ifconfig > test.log

cat test.log

ls >> test.log 2>&1

cat test.log
