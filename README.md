# generate-models

#### 介绍
用于生产go mysql的model文件，不限框架都可以生成

#### 使用说明

  -cfg string
        指定要使用的配置文件路径 (default "configs/")
  -cover string
        是否替换重新生成 (default "n")
  -c string
        是否替换生成 (default "n")
  -db string
        数据库名,不填则按配置文件来
  -d string
        数据库名,不填则按配置文件来
  -file string
        指定要使用的配置文件名 (default "dev") 使用dev.yaml
  -f string
        指定要使用的配置文件名 (default "dev") 使用dev.yaml
  -path string
        指定要生成的model路径 (default "./models/")
  -p string
        指定要生成的model路径 (default "./models/")
  -tables string
        表名，多个使用,分割，不填则生成数据库下所有表的model
  -t string
        表名，多个使用,分割，不填则生成数据库下所有表的model

