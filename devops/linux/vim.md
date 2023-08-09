## vim

### **普通模式(默认模式)：**

Ctrl+f  向下移动一页
Ctrl+b  向上移动一页
0或Home 移到这一行的最前面字符处
$或End  移到这一行的最后面字符处
G       移到这个文件的最后一行
nG      移到这个文件的第n行
gg      移到这个文件的第一行
n+Enter 向下移动n行

按单词移动
w       移动到下一个单词的词首
b       移动到上一个单词的词首
e       移动到下一个单词的词尾
ge      移动到上一个单词的词尾


查找
/word   向下寻找一个名称为word的字符串
?word   向上寻找一个名称为word的字符串
n       向下重复前一个查找
N       向上重复前一个查找

替换
:n1,n2s/word1/word2/g       在第n1行与n2行直接寻找word1这个字符串，并将该字符串替换为word2
:n1,n2s/word1/word2/gc      逐个询问是否替换
:1,第一行到最后一行s/word1/word2/gc        逐个询问

第一行到最后一行
:1,$s/search/replace/g
:%s/search/replace/g

指定行m,n
:m,ns/search/replace/g

当前行到x行
:.,.+xs/search/replace/g

当前行到最后一行
:.,$s/search/replace/g

删除
x,X     向后向前删除字符，(Del, Backspace)
nx
dd      删除光标所在的那一行
ndd     往下删除n行

复制
yy      复制光标所在的那一行
nyy     向下复制n行

粘贴
p,p     向上向下粘贴

复制y、粘贴p、删除d

回退u 、重复ctrl+r

退出vim：ZZ、q、wq

### **编辑模式(普通模式按i,I,o,O,a,A,r,R其中一个可进入编辑模式，编辑模式按Esc可退回一般模式)：**

i,I     光标所在行插入
a,A     光标下一个字符插入，光标行之后一个字符插入
o,O     下一行插入，上一行插入
r,R     替换光标所在字符

返回普通模式esc

### **可视模式**

普通模式下按 v 可进入可视模式

常用功能：

- 复制、粘贴
- 调整缩进【shift + <】，【shift + >】

### **命令行模式(一般模式按 : / ?其中一个可进入命令行模式命令行模式按Esc可退回一般模式)：**

!       强制
:w      将数据写入硬盘
:q      离开vi
:wq     保存后离开
:w [filename]       类似于另存文件
:n1,n2 w [filename] n1到n2行数据存入filename
:ZZ  直接保存退出

:set nu     显示行号
:set nonu   取消行号

:set nu     显示行号
:set nonu   取消行号

### **buffer 相关**

跳转 buffer

- :bp 前一个
- :bn 后一个
- :b[number] 指定序号

关闭 buffer

- (bd 保留当前buffer序号，bw 不保留)
- :bd 或 :bw 当前
- :[number]bw 指定序号
- :n1,n2bw 指定范围序号
- :%bw
- 

### **自定义 vimrc 配置**

~/.vimrc

### **扩展插件**

- [vim-plug 插件管理-新](https://github.com/junegunn/vim-plug)
- [Vundle 插件管理-老](https://github.com/VundleVim/Vundle.vim)
- [nerdtree](https://github.com/preservim/nerdtree)
- [YouCompleteMe](https://github.com/ycm-core/YouCompleteMe)

### **深入 vim**

- [Vim实用技巧](https://book.douban.com/subject/26967597/)



### vim基础操作

2.vi的使用(常用)

1).一般模式(默认模式)：

Ctrl+f 向下移动一页

Ctrl+b 向上移动一页

0或Home 移到这一行的最前面字符处

$或End 移到这一行的最后面字符处

G 移到这个文件的最后一行

nG 移到这个文件的第n行

gg 移到这个文件的第一行

n+Enter 向下移动n行

查找

/word 向下寻找一个名称为word的字符串

?word 向上寻找一个名称为word的字符串

n 向下重复前一个查找

N 向上重复前一个查找

替换

:n1,n2s/word1/word2/g 在第n1行与n2行直接寻找word1这个字符串，并将该字符串替换为word2

:n1,n2s/word1/word2/gc 逐个询问是否替换

:1,$s/word1/word2/g 第一行到最后一行

:1,$s/word1/word2/gc 逐个询问

删除

x,X 向后向前删除字符，(Del, Backspace)

nx

dd 删除光标所在的那一行

ndd 往下删除n行

复制

yy 复制光标所在的那一行

nyy 向下复制n行

粘贴

p,p 向上向下粘贴

复原前一个操作

u

重复上一个操作

Ctrl+r或.

2).编辑模式(一般模式按i,I,o,O,a,A,r,R其中一个可进入编辑模式，编辑模式按Esc可退回一般模式)：

i,I 光标所在行插入

a,A 光标下一个字符插入，光标行之后一个字符插入

o,O 下一行插入，上一行插入

r,R 替换光标所在字符

3).命令行模式(一般模式按:,/,?其中一个可进入命令行模式,命令行模式按Esc可退回一般模式)：

! 强制

:w 将数据写入硬盘

:q 离开vi

:wq 保存后离开

:w [filename] 类似于另存文件

:n1,n2 w [filename] n1到n2行数据存入filename

:set nu 显示行号

:set nonu 取消行号

3.vim的使用

vim filename 进入vim

Ctrl+z 将vim丢到后台执行，返回命令行

1).选择块

v 字符选择，选到反白

V 行选择

Ctrl+v 块选择

y 将反白复制

p 粘贴

d 将反白删除

2).多文件编辑

:n 编辑下一个

:N 上

:files 列出目前这个vim的打开的所有文件

3).多窗口

:sp [filename] 打开一个新窗口

Ctrl+w+j或Ctrl+w+↓ 窗口向下切换

Ctrl+w+k或Ctrl+w+↑ 上

Ctrl+w+q 离开

4).vim环境设置与记录： ~/.vimrc, ~/.viminfo

P288

### vim奇技淫巧

**help 相关**

查看相关命令:help command

**buffer 相关**

airline 插件

跳转 buffer:bp 前一个:bn 后一个:b[number] 指定序号

关闭 buffer(bd 保留当前buffer序号，bw 不保留):bd 或 :bw 当前:[number]bw 指定序号:n1,n2bw 指定范围序号:%bw

**编码效率**

ctags 和 taglist 插件vim taglist 插件，所以安装 taglist 之前首先要在系统上安装 ctags，mac 默认 ctags 版本太旧，使用 brew install ctags 安装新版本

安装:Plug 'vim-scripts/taglist.vim':source %:PlugInstall

配置.vimrc使vim的taglist插件关联ctagslet Tlist_Ctags_Cmd=your_ctags_bin

然后在 vim 中进行:helptaglist~/.vim/doc 操作，完成taglist安装然后在 vim 中进行:help taglist 查看 taglist 文档

常用命令，Tlist开头例如：:TlistToggle
