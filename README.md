## consul-watch

简介

    本程序接收consul watch 监控到的变化（服务--json）
    解析json，使用go的模板库。修改nginx upstream配置文件。
    调用dyups的http接口，修改nginx 内存中的upstream的配置，不用重启nginx

使用环境

    consul  使用了服务注册，使用了watch功能。
        watch监控服务变化，向指定api推送可用服务（json格式）。 
    tengine（编译了dyups模块）
        开启dyups模块的http接口

tengine 编译配置

    ./configure --add-module=/root/tengine-2.3.2/modules/ngx_http_upstream_dyups_module/
    
    [root@localhost conf.d]# cat dyups_management.conf 
         server {
            listen  18882; 
            location / {
                dyups_interface;
            }
        }
     
配置文件

    [main]
    #模板路径
    tpl_path=./ups.tpl
    #upstream路径，必须以“/”结尾
    upstream_path=./
    #dyups模块的接口地址，必须以“/”结尾
    dyups_url=http://192.168.3.112:18882/upstream/
    #api 监听端口格式  :9527 |  127.0.0.1:9527
    listen_port=:9527

命令

    main [-c /etc/consu-watch/conf]
    不指定配置文件，找本地路径（main的 ./）的下的conf文件。

本程序编译

    使用了gin框架。可能会出现go get很慢。建议用国外主机编译。
    有时间去掉框架，直接用go的原生http包实现。
    