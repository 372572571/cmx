# create model

## 项目基本配置

``` yaml
# 配置 

# 初始化相关配置
db_config:
  host: mysql.rds.aliyuncs.com
  port: 3306
  dbname: db_dev
  user: username
  pwd: userpasswd
# 项目路径
project_path: /home/user/project_patn_name
# 数据库字段可空时修改为指针
enable_go_null_point: true
# 启用gorm 软删
enable_gorm_soft_delete: true
# 启用gorm tag
enable_gorm_tag: true
# 数据库中的 bigint转string
enable_big_int_to_string: false
apis:
  # 关联的服务
  admin:
    is_join: false
    api_yaml_path: "/home/lyld/src/git/createmodel/project/test/admin"
    proto_config: # api yaml 后生成的文件配置
      # 输出路径
      output_path: "/home/lyld/src/git/createmodel/project/test/proto/apis-admin/admin/${group}/v1"
      pkg_name: "admin.${group}.v1" # api模块包名称
      # proto go 包名称
      go_pkg_name: "lbck/gen/admin/${group}/v1;${group}"
      import_pkgs:
        - path: "google/api/field_behavior.proto"
        - path: "protoc-gen-openapiv2/options/annotations.proto"
      options_import_pkgs: # 如果额外引用了其他的message
          - path: "incorporate.proto"
  # 关联的服务可以是多个
  # lbck: 

```

## init 初始化项目
```sh

createmodel init --f build.yml
? Confirm initialization  [Use arrows to move, type to filter]
> yes - yes
  no - no

├── admin
│   └── admin.conf
├── build.yml
├── create-openapi.sh
├── default
├── lbck.conf
├── link.conf
├── message
└── model
    ├── announce.yaml
    ├── apply_user_product.yaml
    ├── app_version.yaml

```

## api 编辑一个yaml api 并生成proto 

``` yaml
# build.yal 追加配置 stores_config
stores_config:
  is_enable: true
  stores_name: "incorporate" # incorporate.Xxx subject = 1;
  proto_config:
    output_path: "/home/lyld/src/git/createmodel/project/test/proto/incorporate"
    pkg_name: "incorporate"
    go_pkg_name: "lbck/gen/incorporate"
    import_pkgs:
      - path: "google/api/field_behavior.proto"
      - path: "protoc-gen-openapiv2/options/annotations.proto"

├── admin
│   ├── admin.conf
│   └── user
│       └── refresh_token.yaml
# file refresh_token.yaml 
api_definition:
  user:
  - name: refresh_token
    http:
      method: post
      path: user/refresh-token
      body: '*'
      summary: "刷新用户token"
    request: admin.user.refresh_token.refresh_token_request
    response: admin.user.refresh_token.refresh_token_response
    sign_type: ""
    description: ""
message_definition:
  refresh_token_request:
  refresh_token_response:
  - column_name: token
    type: string
    comment: "token"

```

```sh
➜ createmodel api --f build.yml --s admin
? Confirm api init Yes
/home/lyld/src/git/createmodel/project/test/adminp/proto/admin/user/v1/refresh_token_alone.proto

├── adminp
│   └── proto
│       └── admin
│           └── user
│               └── v1
│                   └── refresh_token_alone.proto
└── proto
    └── incorporate
        └── incorporate.proto
```
# createmodel types --f build.yml
```yaml
# 追加配置
type_config:
  output_path: "/home/lyld/src/git/createmodel/project/test/go/types"
  go_pkg_name: "types"
```
``` sh
createmodel types --f build.yml
# 收集文件中的定义
enums_definition:
  auth:
    - key: authorization_code
      value: 1
      desc: 授权码
      zh: 授权码
    - key: refresh_token
      value: 2
      desc: 刷新token
      zh: 刷新token
├── go
│   └── types
│       └── types.go
```

# createmodel repo --f build.yml 操作模型生成
``` yaml
repo_config:
  output_path: "/home/lyld/src/git/createmodel/project/test/go/repo"
  pkg_name: "newrepo"
  enable_model_to_proto: false
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

```sh
createmodel repo --f build.yml
├── go
│   ├── repo
│   │   ├── announce.go
│   │   ├── apply_user_product.go
│   │   ├── app_version.go
```


## 测试项目路径
```
│   ├── example
│   │   └── test
```

`自行修改路径配置`