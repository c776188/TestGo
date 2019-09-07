# TestMysql

Reference [golang学习之旅：使用go语言操作mysql数据库](https://studygolang.com/articles/3022)

## QS

TestMysql/.env
    `MYSQL_CONFIG=root:password@@/DATABASE`

## SQL
```sql
CREATE TABLE `test` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `text` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1501 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```