# alipay ILLEGAL_SIGN 错误解决

> alipay ILLEGAL_SIGN 错误解决: https://www.jianshu.com/p/28585a6454b2

大概是 2 周前周五遇到的问题吧，现在回想起来，当时追踪并解决问题的过程挺有意思的，记录下来给大家看看。

## 事件回顾

我们的产品提供消费分期的能力给航旅商户，当天下午 2 点，有一个商户反馈他们调起我们的 sdk 后选择 alipay 时，跳出支付宝 **ILLEGAL_SIGN** 提示页面。

![](http://g.hiphotos.baidu.com/exp/w=500/sign=4c28f919372ac65c67056673cbf3b21d/4ec2d5628535e5dd91be80a674c6a7efce1b623c.jpg)

接到反馈后，我们进行了 alipay 的单元测试、自家机酒产品进入 sdk 后使用 alipay 以及使用商户传递过来的订单信息进行 alipay 单元测试，都没有复现这个问题。

> PS：支付系统我开发完，已经稳定运行 3 个月了，这是第一次收到这样问题的反馈。

我们反馈给商户，希望商户确认一下自己的环境，我们也根据搜索引擎的结果，增加了：

```php
header("Content-type:text/html;charset=utf-8");
```

随后配合商家测试，问题依旧存在。

## 定位问题

由于周五有各种例会（项目会、技术团队会、周报等），等到下午 6 点才有时间继续跟进。

根据上面代码 `charset=utf-8`，我尝试用 chrome 的 `Charset` 插件，通过修改字符集为 **GBK**，终于稳定复现了这个问题。

![chrome-ext-charset](https://upload-images.jianshu.io/upload_images/567399-d62641fc22873103.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/700)

![charset-gbk](https://upload-images.jianshu.io/upload_images/567399-809ff266b8157989.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/425)

其实就是确定了，问题出在 `charset` 上面。

## round one：fail

于是，围绕 `charset` 做了多个尝试：

- html 通过 meta 限定 charset
- form 表单提交数据的 4 种格式，以及数据对应的编码

```
<meta charset="UTF-8"> // html meta
<form action="" method="post" enctype="multipart/form-data;charset=utf-8"></form>

// http header 'Content-type' 标准形式
Content-type: application/x-www-form-urlencoded; charset=UTF-8
```

> PS: http 协议中的 key 是不区分大小写的，所以写 'content-type' 也是可以的

但是无论怎么改，还是会一直跳出支付宝 **ILLEGAL_SIGN** 提示页面。

## 还是插件

仔细一回顾，本来我这里是不会出现这个问题的，是通过chrome charset 插件才稳定复现的，于是又尝试用插件将 charset 改回 **utf-8**，果然好了！！！

所以，这插件干了什么？！！！

```bash
# show me the code
git clone https://github.com/jinliming2/Chrome-Charset.git
```

```js
// 原来插件接管了所有请求
chrome.tabs.onUpdated.addListener((tabId, changeInfo, tab) => {
    if(tab.url.startsWith('file://') && changeInfo.status === 'complete' && localStorage.getItem('tab' + tabId)) {
        let xmlHttp = new XMLHttpRequest();
        xmlHttp.overrideMimeType('text/plain; charset=' + localStorage.getItem('tab' + tabId)); // 看这里
        xmlHttp.onload = () => {
            const is_html = /\.html?$/.test(tab.url);
            const data = is_html ? encodeURI(xmlHttp.responseText) : encodeURI(html_special_chars(xmlHttp.responseText));
            chrome.tabs.executeScript(tabId, {
                code: `const _t = document.open('text/${is_html ? 'html' : 'plain'}', 'replace');
                _t.write(${is_html ? `decodeURI('${data}')` : `'<pre>' + decodeURI('${data}') + '</pre>'`});
                _t.close();`,
                runAt: 'document_start'
            });
        };
        xmlHttp.onerror = () => {
            alert(chrome.i18n.getMessage('cannotLoadLocalFile'));
        };
        xmlHttp.open('GET', tab.url, true);
        xmlHttp.send();
    }
}
```

所以，果断禁用掉。

> PS: chrome 的「隐身模式」简直是调试神器，欢迎尝试。

## round two： fail

禁用插件之后，继续按照 **round one** 的思路进行修改，然后，这问题就没办法复现了。

所以，现在变成了 **世纪难题** 了：

- 如果不启用 charset 插件，就无法复现问题了
- 如果启动插件，所有 chrome 的请求都会被接管，修改的代码不会生效，也无法验证和修复问题

## final：月光就在眼前

进入死胡同的时候，千万不要怀疑人生，你要相信：

- 我书读得少，你不要骗我。
- 那些看似无解的问题，往往只是你孤陋寡闻。
- 知识这东西就是这样：知道就是知道，不知道就是不知道，所以还是多知道一点比较好

所以，翻开了[《http 权威指南》](https://book.douban.com/subject/10746113/)，仔细查阅之后，你就会发现，在 http协议里面，只有 2 个地方会影响到 charset：

- 客户端：`accept-charset='utf-8'`
- 服务端：`content-type: text/plain;charset:utf-8`

所以，在代码里面修改为：

```html
<form id='alipaysubmit' name='alipaysubmit' accept-charset='utf-8'></form>
```

晚上 10 点的月光，正好。