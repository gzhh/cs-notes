## Iterm2

### 文档

[https://iterm2.com/documentation.html](https://iterm2.com/documentation.html)

[https://iterm2.com/documentation-one-page.html](https://iterm2.com/documentation-one-page.html)

### 标签

新建标签：command + t

关闭标签：command + w

切换标签：command + 数字, command + arrow

切换全屏：command + enter

查找：command + f

### 分屏 pane

垂直分屏：command + d

水平分屏：command + shift + d

切换屏幕：command + option + arrow

查看历史命令：command + ;

查看剪贴板历史：command + shift + h

### 巧用

可以拖拽选中的字符串；

点击 url：调用默认浏览器访问该网址；

点击文件：调用默认程序打开文件；

如果文件名是filename:42，且默认文本编辑器是 Macvim、Textmate或BBEdit，将会直接打开到这一行；

点击文件夹：在 finder 中打开该文件夹；

同时按住option键，可以以矩形选中，类似于vim中的ctrl v操作。

### 其他

清除当前行：ctrl + u

到行首：ctrl + a

到行尾：ctrl + e

前进后退：ctrl + f/b (相当于左右方向键)

上一条命令：ctrl + p

搜索命令历史：ctrl + r

删除当前光标的字符：ctrl + d

删除光标之前的字符：ctrl + h

删除光标之前的单词：ctrl + w

删除到文本末尾：ctrl + k

交换光标处文本：ctrl + t

清屏1：command + r

清屏2：ctrl + l

### 自带有哪些很实用的功能/快捷键

cmd+ 数字在各 tab 标签直接来回切换

选择即复制 + 鼠标中键粘贴，这个很实用

cmd; + f 所查找的内容会被自动复制

cmd + d 横着分屏 / ⌘ + shift + d 竖着分屏

cmd + r = clear，而且只是换到新一屏，不会想 clear 一样创建一个空屏

ctrl + u 清空当前行，无论光标在什么位置

输入开头命令后 按 &#8984; + ; 会自动列出输入过的命令

cmd + shift + h 会列出剪切板历史

可以在 Preferences > keys 设置全局快捷键调出 iterm，这个也可以用过 Alfred 实现

我常用的一些快捷键

cmd + 1 / 2 左右 tab 之间来回切换，这个在 前面 已经介绍过了

cmd← / ⌘→ 到一行命令最左边/最右边 ，这个功能同 C+a / C+e

← / ⌥→ 按单词前移/后移，相当与 C+f / C+b，其实这个功能在Iterm中已经预定义好了，⌥f / ⌥b，看个人习惯了

设置方法如下

当然除了这些可以自定义的也不能忘了 linux 下那些好用的组合

C+a / C+e 这个几乎在哪都可以使用

C+p / !! 上一条命令

C+k 从光标处删至命令行尾 (本来 C+u 是删至命令行首，但iterm中是删掉整行)

C+w A+d 从光标处删至字首/尾

C+h C+d 删掉光标前后的自负

C+y 粘贴至光标后

C+r 搜索命令历史，这个较常用
