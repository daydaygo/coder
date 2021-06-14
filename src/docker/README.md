- docker 安装加速: https://cr.console.aliyun.com/#/accelerator

```sh
docker run -d -p 18080:8080 -e SWAGGER_JSON=/foo/push.swagger.json -v /opt/php-swagger:/foo swaggerapi/swagger-ui
docker run -it -v $PWD:/docs:ro -p 3000:3000 quintoandar/docsify
docker run -p 8080:8080 jupyter/nbviewer
```