# generate-models

#### 介绍
用于生成go mysql的model文件，不限框架都可以生成，如果使用gin开发接口，可以参考另外一个仓库https://github.com/Eric-chy/generate-models   包含redis，es，rabbitmq，mongo，jwt，上传，发邮件，jaeger，sentry，组件齐全

#### 使用说明
go run main.go 加上以下参数或者不加

  -c string
        指定要使用的配置文件路径 (default "configs/")  
        
  -r string
        是否替换旧文件生成 (default "n" "n|y")  
        
  -d string
        数据库名,不填则按配置文件来  
        
  -f string
        指定要使用的配置文件名 (default "dev") 使用dev.yaml  
        
  -m string
        指定要生成的model路径 (default "./models/")  
        
  -t string
        表名，多个使用英文半角,分割，不填则生成数据库下所有表的model  

