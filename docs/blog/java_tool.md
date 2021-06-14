# java| tool

## gradle

- 语言: kotlin groovy

## maven

- 项目管理利器——maven: <https://www.imooc.com/learn/443>
- maven repository/mirror 详解: <https://www.sojson.com/blog/168.html>
- maven repository/mirror 配置: <https://blog.csdn.net/joewolf/article/details/4876604>
- <https://maven.aliyun.com/>
- lifecycle: clean(pre - post) default(compile test package install) site(pre post deploy)

```sh
# 镜像加速
~/.m2/settings.xml
<settings>
    <mirror>
        <id>nexus-aliyun</id>
        <mirrorOf>central</mirrorOf>
        <name>Nexus aliyun</name>
        <url>http://maven.aliyun.com/nexus/content/groups/public</url>
    </mirror>
</settings>

mvn compile
mvn test
mvn package
mvn clean
mvn install
mvn archetype:generate
```
