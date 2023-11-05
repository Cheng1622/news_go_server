## 特性
`Zap`

DB *gorm.DB // DB 全局mysql数据库变量

	Redis *redis.Client // Redis全局redis数据库变量

	CasbinEnforcer *casbin.Enforcer // CasbinEnforcer 全局CasbinEnforcer

	Log *zap.SugaredLogger // Log 全局日志变量

	Trans ut.Translator // Trans 全局validate翻译器

	AuthMiddleware *jwt.GinJWTMiddleware // AuthMiddleware jwt auth认证

	Gin 一个类似于martini但拥有更好性能的API框架, 由于使用了httprouter, 速度提高了近40倍
MySQL 采用的是MySql数据库
Jwt 使用JWT轻量级认证, 并提供活跃用户Token刷新功能
Casbin Casbin是一个强大的、高效的开源访问控制框架，其权限管理机制支持多种访问控制模型
Gorm 采用Gorm 2.0版本开发, 包含一对多、多对多、事务等操作
Validator 使用validator v10做参数校验, 严密校验前端传入参数
Lumberjack 设置日志文件大小、保存数量、保存时间和压缩等
Viper Go应用程序的完整配置解决方案, 支持配置热更新