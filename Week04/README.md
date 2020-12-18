# awesome-project 

## 项目简介
1. 通过 config 下的 toml 配置文件保存数据库的连接信息
2. 在 internal/dao 中,通过 wire 实现控制反转,方便测试 dao.
3. api 文件夹中放置 api 的定义 proto 文件.方便别的服务调用本服务.
4. 在 service 层撸业务代码即可.
