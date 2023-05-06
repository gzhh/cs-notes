## grep

过滤有乱码的行

例子：

grep -Pv "[\x80-\xFF]" input.txt > out.txt

grep -v grep 过滤掉包含 grep 的结果行

例子：

ps aux | grep -v grep | grep php-fpm
