## awk

按逗号分割，并打印第三列内容到指定文件

awk -F"," '{print $3}' input.txt > out.txt

按逗号分割，取出第十列小于1的所有行

awk -F"," '$10<1{print $0}' > out.txt
