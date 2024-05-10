# Interview

## 成绩统计
```
DROP TABLE IF EXISTS `student_score`;

CREATE TABLE `student_score` (

`id` int(10) NOT NULL AUTO_INCREMENT,

`name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',

`subject` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',

`score` int(3) NULL DEFAULT 0,

PRIMARY KEY (`id`) USING BTREE

) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;


INSERT INTO `student_score` VALUES (1, 'zhao', 'chinese', 90);

INSERT INTO `student_score` VALUES (2, 'zhao', 'math', 85);

INSERT INTO `student_score` VALUES (3, 'zhao', 'english', 95);

INSERT INTO `student_score` VALUES (4, 'zhao', 'science', 86);

INSERT INTO `student_score` VALUES (5, 'qian', 'chinese', 89);

INSERT INTO `student_score` VALUES (6, 'qian', 'math', 98);

INSERT INTO `student_score` VALUES (7, 'qian', 'english', 75);

INSERT INTO `student_score` VALUES (8, 'qian', 'science', 96);

INSERT INTO `student_score` VALUES (13, 'sun', 'chinese', 98);

INSERT INTO `student_score` VALUES (14, 'sun', 'math', 89);

INSERT INTO `student_score` VALUES (15, 'sun', 'english', 56);

INSERT INTO `student_score` VALUES (16, 'sun', 'science', 97);

INSERT INTO `student_score` VALUES (17, 'li', 'chinese', 94);

INSERT INTO `student_score` VALUES (18, 'li', 'math', 93);

INSERT INTO `student_score` VALUES (19, 'li', 'english', 95);

INSERT INTO `student_score` VALUES (20, 'li', 'science', 96);

```

1. 请写出每门科目成绩前三的数据。（表：student_score，姓名：name，科目：subject，分数 score）

```
use test;

SELECT

a.*

FROM `student_score` AS a

LEFT JOIN `student_score` AS b ON a.`subject` = b.`subject` AND a.`score` < b.`score`

GROUP BY

a.`id` -- , a.`subject`, a.`score`

HAVING

COUNT(b.`id`) < 3

ORDER BY

a.`subject`, a.`score` DESC;
```

2. 写出删除表中重复数据，并保留一条。

```
DELETE FROM student WHERE

(`name`,`subject`,score) IN (

SELECT t.name,t.subject,t.score FROM (

SELECT `name`,`subject`,score FROM student

GROUP BY `name`,`subject`,score

HAVING COUNT(1)>1

)t

)

AND id not in(

SELECT a.minId FROM (

SELECT id as minId FROM student

GROUP BY `name`,`subject`,score

HAVING COUNT(1)>1

)a

)
```

参考链接：https://blog.csdn.net/n950814abc/article/details/82284838

3. 写出所有科目成绩都大于 80 分的学生数据

```
select name from student group by name having min(score)>80;
```

