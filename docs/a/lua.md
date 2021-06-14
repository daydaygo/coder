# lua

- why: game embed compile VM glue胶水/redis/nginx
- syntax
  - comment: `--` `--[[ --]]`
  - type: string table=map nil boolean number function userdata thread `type()`
  - op: `~=`
  - var: `local`
  - ctl: `for fori`
- eco: <https://www.lua.org> <https://lua-users.org>

```sh
brew install lua luarocks
luarocks install mobdebug # debug
lua -i # 交互式
lua xxx.lua
```

```lua
-- string
a=b..c -- no nil
string.format("%s%s",a,b)
[[str]] -- heredoc

-- table
for k, v in pairs(tab1) do
    print(k .. " - " .. v)
end
table.insert(t, {k1=v1})

for product in string.gmatch(content, '<div class="zg_itemRow">(.-)<div class="zg_clear">') do --注意是 gmatch
```

## 文件操作

```lua
function writeFile(file_name,string)
  local f = assert(io.open(file_name, 'w'))
  f:write(string)
  f:close()
end
```

## 模式匹配

```lua
-- 使用 .- 代替 .*?
test = string.gsub(test,"\\\\","\\") -- 使用的 \ 进行转义
local company = string.match(c,"<th>xxx</th><td>(.-)</td>") --匹配单次
--匹配多个数据 , \ 用来转义 , % 用来转义 regexp 中的特殊字符
for shop_name, item_code in string.gmatch(content, "<p class=\"price\"><a href=\"http://item%.rakuten.co%.jp/(.-)/(.-)/\">") do --匹配多次
table.insert(result, {shop_name = shop_name , item_code = item_code}) --插入table
end
```

## cjson

```lua
local cjson = require "cjson" -- 解析 json 数据
local data = cjson.decode(params) -- url = data.url
return cjson.encode({success = 0})
```

## iconv

```lua
local iconv = require "iconv" -- 转码
local cd = iconv.new("utf8","euc-jp") -- cd = iconv.open(to , from) 也可以
cont = cd:iconv(cont)
```

## http

```lua
--get
local http = require "socket.http" -- 相当于Perl 中的 $ua
local headers = {}
headers["User-Agent"] = "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:30.0) Gecko/20100101 Firefox/30.0"
local t = {}
local response, code, head = http.request{url=url, headers=headers, sink = ltn12.sink.table(t)}
if code == 200 then
local c = table.concat(t) -- 这就是网页的源代码了
end

--post 并加上计时
local http = require "socket.http" -- 相当于Perl 中的 $ua
local cjson = require "cjson" -- 解析 json 数据
local socket = require "socket" -- 获取网页访问需要的时间

function run(params)
    local data = cjson.decode(params)
    local reqbody = data.reqbody
    local headers = {}
    headers["User-Agent"] = "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:30.0) Gecko/20100101 Firefox/30.0"
    headers["content-length"] = string.len(reqbody)
    local t = {}
    local t0 = socket.gettime() -- 网页访问开始
    local response, code, head, status = http.request{
        method = "POST",
        url=data.url,
        headers=headers,
        source = ltn12.source.string(reqbody),
        sink = ltn12.sink.table(t)
    }
    local c = table.concat(t) -- 这就是网页的源代码了
    local t1 = socket.gettime() -- 网页访问结束
    local time = math.floor((t1 - t0) * 1000000)
    return cjson.encode({code = code, time = time, c = c})
end
```

## curl

```lua
ocal curl = require 'lcurl'
local iconv = require 'iconv'
function extract_date(headers)
    for _, v in pairs(headers) do
        local date = string.match(v, '^Date: ([^\r\n]+)')
        if date then
            return date
        end
    end
    return nil
end
function curl_get(url)
    local response = {}
    local headers = {}
    local c = curl.easy()
        :setopt_url(url)
        :setopt(curl.OPT_USERAGENT, 'Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0)')
        :setopt(curl.OPT_FOLLOWLOCATION, true)
        :setopt_writefunction(function(buf)
            table.insert(response, buf)
            return #buf
        end)
        :setopt_headerfunction(function(buf)
            table.insert(headers, buf)
            return #buf
        end)
        :perform()
    local code = c:getinfo(curl.INFO_RESPONSE_CODE)
    c:close()
    if code == 200 then
        local response = table.concat(response)
        local err -- 正确使用 iconv
        response, err = iconv.new('utf-8//IGNORE', 'shift-jis'):iconv(response)
        if err then
            return nil
        end
        return response, extract_date(headers)
    else
        return nil
    end
end
```
