FROM node:10.12.0-alpine
LABEL maintainer="1252409767@qq.com"

RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/" /etc/apk/repositories && \
    apk update && rm -rf /var/cache/apk/* && \
    apk add --no-cache tzdata

# npm/cnpm/yarn https://blog.csdn.net/sinat_34682450/article/details/79473658
RUN npm config -g set registry https://registry.npm.taobao.org && \
    apk add --no-cache yarn && \
    yarn config set registry https://registry.npm.taobao.org
# pm2 https://mp.weixin.qq.com/s/0YWNLUoLt3wdIOQ1MKjASQ
# RUN yarn global add pm2
# hexo https://hexo.io/docs
# RUN yarn global hexo

CMD ["top"]
WORKDIR /var/www/
# ENTRYPOINT [ "/bin/sh" ]