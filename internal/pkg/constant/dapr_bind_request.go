package constant

const MySQLBindName = "dapr-user-service-mysql"

// 对于数据库操作，使用 gorm-gen来进行数据读写。
// dapr building 构建快，关于数据读写提供的API功能有限: 1. 不支持防SQL注入 2. 事务无法保证
