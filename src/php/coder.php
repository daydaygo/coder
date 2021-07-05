<?php

// coder 文档工具

function check_readme()
{
    $d = '/Users/dayday/hub/coder/docs/blog';
    $file = scandir($d);
    $dr = [];
    foreach ($file as $a) {
        if (strpos($a, '.md') !== false) {
            $dr[] = $a;
        }
    }

    $f = '/Users/dayday/hub/coder/docs/blog/readme.md';
    $fr = [];
    $fd = fopen($f, 'r');
    while ($line = fgets($fd)) {
        preg_match_all('/\((.*?\.md)\)/', $line, $m);
        if (!empty($m[1])) $fr = array_merge($fr, $m[1]);
    }

    var_dump(array_diff($dr, $fr));
}

function dedao_json()
{
    $a = json_decode(file_get_contents('dedao.json'), true)['c']['catalog_list'];
    foreach ($a as $v) {
        echo str_repeat('  ', $v['level']) . '- ' . $v['text'], "\n";
    }
}