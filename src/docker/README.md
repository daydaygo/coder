# docker: 快速配置开发环境依赖

- config
  - nginx.conf -> 启动 nginx 服务使用
  - nginx.conf.demo -> nginx 配置文档

```sh
docker run -d -p 18080:8080 -e SWAGGER_JSON=/foo/push.swagger.json -v /opt/php-swagger:/foo swaggerapi/swagger-ui
docker run -it -v $PWD:/docs:ro -p 3000:3000 quintoandar/docsify
docker run -p 8080:8080 jupyter/nbviewer
```
