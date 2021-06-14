// 获取阿里云产品列表信息: ZH + EN
function get_aliyun_product_list() {
    dom = $("[data-spm='products-grouped']")
    cnt = dom.childElementCount
    // 简单粗暴定义数据结构: 参照 dom 结构
    ans = []
    for (let i = 0; i < cnt; i++) {
        sub_cnt = dom.children[i].childElementCount
        head = dom.children[i].children[0].outerText
        ans[i] = {}
        ans[i]['head'] = head
        ans[i]['list'] = []
        // console.log('head', head)
        for (let j = 1; j < sub_cnt; j++) {
            url = dom.children[i].children[j].children[0].href
            text = dom.children[i].children[j].children[0].innerText
            console.log(head, text, url)
            ans[i]['list'][j-1] = {
                'text': text,
                'url': url
            }
        }
    }
    console.log(JSON.stringify(ans))
}

// todo: 获取到 zh en
zh_obj = JSON.parse(zh)
en_obj = JSON.parse(en)
for (let i = 0; i < zh_obj.length; i++) {
    json_h2 = '## ' + en_obj[i]['head'] + ' ' + zh_obj[i]['head']
    console.log(json_h2)
    for (const j in zh_obj[i]['list']) {
        json_ul = '- [' + en_obj[i]['list'][j]['text'] + ' ' + zh_obj[i]['list'][j]['text'] + '](' + en_obj[i]['list'][j]['url'] + ')'
        console.log(json_ul)
    }
}