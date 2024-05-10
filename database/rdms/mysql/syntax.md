# Syntax

1.MySQL中and的优先级高于or

2.从文件中导入数据到MySQL, sql语句

LOAD DATA LOCAL INFILE 'D:/temp.txt'

INTO TABLE temp_table (pid,pro_used,@notused,selected);

3.MySQL查询某字段中以逗号分隔的字符串的方法

`field_name` 值例子为 3,6,9

SELECT * FROM `table_name` WHERE find_in_set('3', `field_name`) OR find_in_set('9', `field_name`);
