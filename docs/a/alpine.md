# alpine

```sh
sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
apk add fish mdocml-apropos vim
apk add/search/info
apk add cherokee --update-cache --repository http://dl-3.alpinelinux.org/alpine/edge/testing/ --allow-untrusted
apk search -v --description 'NTP' # show description and search from description
crond # 需要手动启动
apk add tzdata
cp /usr/share/zoneinfo/Asia/Shanghai > /etc/localtime

# better tool
apk add --no-cache fish mdocml-apropos ack htop mtr aria2 iproute2 drill apache2-utils curl gosu
apk add iproute2 # ss vs netstat
ss -ptl
apk add drill # drill vs nslookup&dig
crond # 开启 cron 服务
ibu # alpine local backup

# client
apk add mysql-client

apk add man man-pages mdocml-apropos less less-doc
export PAGER=less
apk add bash bash-doc bash-completion # bash
apk add util-linux pciutils usbutils coreutils binutils findutils grep # grep / awk
apk add build-base gcc abuild binutils binutils-doc gcc-doc # compile
apk add cmake cmake-doc extra-cmake-modules extra-cmake-modules-doc
apk add ccache ccache-doc
```
