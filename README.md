# cli

## 基本配置 build.yml 


```yaml
# 数据库配置
db_config:
  host: "your mysql host"
  port: "3306"
  dbname: "db name"
  user: "user name"
  pwd: "mysql password"
project_path: ".model" 配置路径
# 可为字段空标记为指针
enable_go_null_point: true
# 开启软删
enable_gorm_soft_delete: true
# 开启gorm序列化
enable_gorm_serializer: true
# 开启tag
enable_gorm_tag: true
# 模型中的proto bigint转string
enable_big_int_to_string: false
# 枚举.go 文件
type_config:
  output_path: "logic/types"
  go_pkg_name: "types"
# model repo .go 文件
repo_config:
  output_path: "logic/repo"
  pkg_name: "repo"
  # 模型需要导入的库
  import_pkgs:
    - path: "context"
    - path: "gorm.io/gorm"
    - path: "gorm.io/gorm/clause"
    - path: "gorm.io/gorm/schema"
    - path: "gorm.io/gen"
    - path: "gorm.io/gen/field"
    - path: "gorm.io/plugin/dbresolver"
    - path: "gorm.io/plugin/soft_delete"
      default_ref: "var _ soft_delete.DeletedAt"
    - path: "time"
      default_ref: "var _ = time.ANSIC"
    - path: "github.com/spf13/cast"
      default_ref: "var _ = cast.StringToDate"
    - path: "encoding/json"
      default_ref: "var _ = json.Marshal"
```

### 初始化项目
```sh
# 假设现在有一个项目名称叫 `test_service`

# 进入项目
cd /your/path/test_service

# 当前路径
pwd
`/your/path/cmx/example/test_service`

mkdir cmx_model

cat << EOF > cmx_model/build.yml 
db_config:
  host: "your mysql host"
  port: "3306"
  dbname: "db name"
  user: "user name"
  pwd: "mysql password"
project_path: ".model"
enable_go_null_point: true
enable_gorm_soft_delete: true
enable_gorm_serializer: true
enable_gorm_tag: true
enable_big_int_to_string: false
type_config:
  output_path: "logic/types"
  go_pkg_name: "types"
repo_config:
  output_path: "logic/repo"
  pkg_name: "repo"
  import_pkgs:
    - path: "context"
    - path: "gorm.io/gorm"
    - path: "gorm.io/gorm/clause"
    - path: "gorm.io/gorm/schema"
    - path: "gorm.io/gen"
    - path: "gorm.io/gen/field"
    - path: "gorm.io/plugin/dbresolver"
    - path: "gorm.io/plugin/soft_delete"
      default_ref: "var _ soft_delete.DeletedAt"
    - path: "time"
      default_ref: "var _ = time.ANSIC"
    - path: "github.com/spf13/cast"
      default_ref: "var _ = cast.StringToDate"
    - path: "encoding/json"
      default_ref: "var _ = json.Marshal"
EOF

# 修改配置
vim cmx_model/build.yml
# project_path: ".model"
# 修改为应为相对于项目的模型配置存放路径是这个目录 
# project_path: "cmx_model"
# 以下配置修改成自己的数据库配置
# db_config:
#   host: "your mysql host"
#   port: "3306"
#   dbname: "db name"
#   user: "user name"
#   pwd: "mysql password"
❯ cmx init --f ./cmx_model/build.yml
? Confirm initialization yes
# 此时会读取数据库并初始化
# ....
# 完成后目录结构
❯ tree
.
└── cmx_model
    ├── build.yml
    ├── default # 存放自定义消息
    ├── link.conf # 关联的模型
    ├── local # 本地模型修改
    ├── message
    └── model
        ├── dict_item.yaml # 数据库中同步下来的模型
        ├── dict.yaml
```

### 生成模型

- 模型生成

```sh
mkdir -p logic/repo
cmx repo --f ./cmx_model/build.yml
# echo:
# write path: logic/repo model: geo_map.go
# write path: logic/repo model: oauth2_access_tokens.go
# write path: logic/repo model: oauth2_authorization_codes.go
# write path: logic/repo model: oauth2_clients.go
# write path: logic/repo model: oauth2_refresh_tokens.go
# write path: logic/repo model: user.go
❯ tree logic
logic
└── repo
    ├── geo_map.go
    ├── oauth2_access_tokens.go
    ├── oauth2_authorization_codes.go
    ├── oauth2_clients.go
    ├── oauth2_refresh_tokens.go
    └── user.go
```

- 生成枚举

```sh
mkdir -p logic/types
cmx types --f ./cmx_model/build.yml
❯ tree logic
logic
├── repo
│   ├── geo_map.go
│   ├── oauth2_access_tokens.go
│   ├── oauth2_authorization_codes.go
│   ├── oauth2_clients.go
│   ├── oauth2_refresh_tokens.go
│   └── user.go
└── types
    └── types.go
# 枚举定义与使用参考项目中的
`example/test_service/cmx_model/default/group.yaml`
```

### 本地修改模型并更新

```sh
# 比如更新user模型

# 创建对应模型的sql create 语句
touch cmx_model/local/user.sql

CREATE TABLE `user` (
    `id` bigint unsigned NOT NULL COMMENT '用户ID',
    `type` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '1' COMMENT '用户类型',
    `passwd` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码',
    `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '姓名',
    `avatar` varchar(104) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '头像',
    `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '昵称',
    `mobile` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '手机号',
    `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '邮箱',
    `switch` char(1) NOT NULL DEFAULT '1' COMMENT '开关',
    `created_at` datetime NOT NULL COMMENT '创建时间',
    `updated_at` datetime NOT NULL COMMENT '更新时间',
    `deleted_at` bigint NOT NULL DEFAULT '0' COMMENT '软删标识符',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `user_mobile_email` (`email`, `mobile`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC COMMENT = '用户'
# 语句中添加性别字段
`sex` char(1) NOT NULL DEFAULT '0' COMMENT '性别,[0:未知,1:男,2:女]',    
❯ cmx init_local --f ./cmx_model/build.yml
? Confirm initialization yes
echo
# init
# 字段新增 user.sex string 性别,[0:未知,1:男,2:女] 
# 此时`user.yaml`model文件已经被本地的sql更新
# 而后调用模型更新 cmx repo --f ./cmx_model/build.yml 即可
```

<!-- ### 关联proto

#### 添加分组配置

- build.yml 添加配置
```yaml

apis:
  admin:
    is_join: false
    api_yaml_path: "cmx_model/admin" # 分组存放路径
    proto_config:
      output_path: "proto/apis-admin/admin/${group}/v1" # proto位置
      pkg_name: "admin.${group}.v1" # proto 包名
      go_pkg_name: "test_service/gen/admin/${group}/v1;${group}" # go 包名
      import_pkgs:
        - path: "google/api/field_behavior.proto"
        - path: "protoc-gen-openapiv2/options/annotations.proto"
      options_import_pkgs: # 引用的合并proto
        - path: "incorporate.proto"
  app:
    is_join: false
    api_yaml_path: "cmx_model/app"
    proto_config:
      output_path: "proto/apis-app/app/${group}/v1"
      pkg_name: "app.${group}.v1"
      go_pkg_name: "test_service/gen/app/${group}/v1;${group}"
      import_pkgs:
        - path: "google/api/field_behavior.proto"
        - path: "protoc-gen-openapiv2/options/annotations.proto"
      options_import_pkgs:
        - path: "incorporate.proto"

```

- 创建必要目录

```sh
mkdir admin
mkdir app
mkdir gen
mkdir -p proto/apis-admin/admin
mkdir -p proto/apis-app/app
``` -->

