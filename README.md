# go-api

引用第三方库

    "github.com/gorilla/mux"  http 路由库
    "github.com/Unknwon/goconfig" 配置文件解析库
    如需要json解析。推荐"github.com/tidwall/gjson"

功能

    http服务
        服务端口需要在配置文件中配置
        路由配置在router中配置
    日志
        屏幕输出或文件（配置文件配置log选项）
    支持中间件
        自定义中间件，需要在middleware中配置
        默认支持日志
        支持跨域（需要在initrouter中use）
    支持配置文件
        服务端口配置，监听所有地址
        配置日志输出
        自定义配置需要在conflib中加入解析配置代码
        支持配置段解析
    数据库支持
        只有部分实例代码，使用需要修改
        写好方法，在api的handler中引用

     