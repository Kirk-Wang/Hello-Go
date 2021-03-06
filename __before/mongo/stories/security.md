# 数据库安全

* 认证
* 验证身份

* 授权
* 操作权限

创建第一个用户
```sh
use admin
db.createUser({
  user: "myUserAdmin",
  pwd: "passwd",
  roles: [ "userAdminAnyDatabase" ]
})
```

启用身份认证
```sh
docker-compose down
# command: mongod --auth 加入 docker-compose.yml
docker-compose up -d
```

使用用户名和密码进行身份验证
```sh
# authenticationDatabase
# 对应的验证数据库
mongo -u "myUserAdmin" -p "passwd" --authenticationDatabase "admin"

docker exec -it a811efa08b1d mongo -u myUserAdmin -p passwd --authenticationDatabase admin

> use test
> db.accounts.find()
Error: error: {
        "ok" : 0,
        "errmsg" : "not authorized on test to execute command { find: \"accounts\", filter: {}, lsid: { id: UUID(\"4065e318-7334-426d-a7e0-2a31e210bf23\") }, $db: \"test\" }",
        "code" : 13,
        "codeName" : "Unauthorized"
# 有管理的权限，但并没有读取的权限
```

使用 db.auth() 进行身份验证
```sh
docker exec -it a811efa08b1d mongo
> db
test
> use admin;
switched to db admin
> db.auth("myUserAdmin","passwd")
1
```

授权

权限
* = 在哪里 + 做什么
```sh
{
  resource: {
    db: "test",
    collection: ""
  },
  actions: [
    "find",
    "update"
  ]
}
# 在test数库，你可以进行 find & update

{
  resource: {
    cluster: true
  },
  actions: [
    "shutdown"
  ]
}
# 可以停止整个集群
```

角色
* 角色 = 一组权限的集合
* read - 读取当前数据库中所有非系统集合
* readWrite - 读写当前数据库中所有非系统集合
* dbAdmin - 管理当前数据库
* userAdmin - 管理当前数据库中的用户和角色
* read/readWrite/dbAdmin/userAdminAnyDatabase - 对所有数据库执行操作（只在admin数据库中提供）

授权
* 将角色赋予用户

"创建一个只能读取test数据库的用户"
```sh
docker exec -it a811efa08b1d mongo -u myUserAdmin -p passwd --authenticationDatabase admin
# 它才有createUser的权限
use test
db.createUser(
  {
    user: "testReader",
    pwd: "passwd",
    roles: [ { role: "read", db: "test" } ]
  }
)
```

验证
```sh
docker exec -it a811efa08b1d mongo -u testReader -p passwd --authenticationDatabase test

> db.accounts.insert({name: "newUser"})
WriteCommandError({
        "ok" : 0,
        "errmsg" : "not authorized on test to execute command { insert: \"accounts\", ordered: true, lsid: { id: UUID(\"1d006147-5d62-479d-a504-e188151447b8\") }, $db: \"test\" }",
        "code" : 13,
        "codeName" : "Unauthorized"
})
```

自定义角色(更加精细的控制)

"创建一个只能读取 accounts 集合的用户"

注意，在执行创建用户和创建角色的时候，要使用一个具备用户管理权限的用户
```sh
docker exec -it a811efa08b1d mongo -u myUserAdmin -p passwd --authenticationDatabase admin

use test
db.createRole({
  role: "readAccounts",
  privileges: [
    {
      resource: {
        db: "test",
        collection: "accounts"
      },
      actions: ["find"]
    }
  ],
  roles: []
})

use test
db.createUser(
  {
    user: "accountsReader",
    pwd: "passwd",
    roles: [ "readAccounts" ]
  }
)
```

验证
```sh
docker exec -it a811efa08b1d mongo -u accountsReader -p passwd --authenticationDatabase test
> show collections
accounts
# 它只看到 accounts 集合
```