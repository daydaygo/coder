# -*- coding: UTF-8 -*-
import datetime
import os

import requests
from bs4 import BeautifulSoup

def get_html():
    headers = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36'
    }
    url = 'https://baike.baidu.com/item/青春有你第二季'
    resp = requests.get(url, headers)
    fd = open('tmp.html', 'w', encoding='utf-8').write(resp.text)

def bs_table():
    with open('tmp.html') as fd:
        soup = BeautifulSoup(fd.read(), 'lxml')
        tables = soup.find_all('table', {'log-set-param':'table_view'})
        table_title = '花语'
        for table in tables:
            table_th = table.find_all('th')
            for title in table_th:
                if table_title in title.get_text():
                    return table

today = datetime.date.today().strftime('%Y%m%d')
print('hello')