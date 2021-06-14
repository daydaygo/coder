<?php

$a = json_decode(file_get_contents('dedao.json'), true)['c']['catalog_list'];
foreach ($a as $v) {
    echo str_repeat('  ', $v['level']).'- '.$v['text'], "\n";
}