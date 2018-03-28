使用方法：
===========

- 查看当前 db迁移状态
> shell : ...\cmd\cli\migrations>go run main.go sqlite3 ./foo.db status

- 前进
> shell : ...\cmd\cli\migrations>go run main.go sqlite3 ./foo.db up

- 回滚
> shell : ...\cmd\cli\migrations>go run main.go sqlite3 ./foo.db down


其余命令 可以看
> shell : ...\cmd\cli\migrations>go run main.go sqlite3 ./foo.db

查看可用的命令列表