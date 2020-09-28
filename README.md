<!--
 * @Author: your name
 * @Date: 2020-09-28 15:54:41
 * @LastEditTime: 2020-09-28 15:55:32
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \gqlserver\README.md
-->
# gqlserver
golang server with echo+gqlgen+gorm 

## 项目说明
>此项目根据go run github.com/99designs/gqlgen init自动生成基础框架<br/> 
>集成echo，gorm，gqlgen库作为建议golang-server架构，方便之后项目开发

- 项目目录

    ├── go.mod<br/> 
    ├── go.sum<br/> 
    ├── gqlgen.yml               - The gqlgen config file, knobs for controlling the generated code.<br/> 
    ├── graph<br/> 
    │   ├── generated            - A package that only contains the generated runtime<br/> 
    │   │   └── generated.go<br/> 
    │   ├── model                - A package for all your graph models, generated or otherwise<br/> 
    │   │   └── models_gen.go<br/> 
    │   ├── resolver.go          - The root graph resolver type. This file wont get regenerated<br/> 
    │   ├── schema.graphqls      - Some schema. You can split the schema into as many graphql files as you like<br/> 
    │   └── schema.resolvers.go  - the resolver implementation for schema.graphql<br/> 
    └── server.go                - The entry point to your app. Customize it however you see fit<br/> 


- 项目提供功能
  
  访问http://localhost:8787进行GraphQL playground的访问，可以网页发送GraphQL命令，查看schema<br/>
  直接访问http://localhost:8787/query进行GraphQL命令的收发



## 开发说明
>- 配置数据结构
  
修改gqlgen.yml，schema.graphqls文件，编辑需要增加的数据结构和数据内容

>- 自动生成代码

使用go run github.com/99designs/gqlgen generate命令生成代码
查看graph/schema.resolvers.go，编辑对应的resolver处理请求内容

>- 运行测试
  
使用go run server.go运行代码，执行相应的命令