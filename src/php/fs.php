<?php

function scan_dir($dir)
{
    $a = scandir($dir);
    foreach ($a as $v) {
        if ($v=='.' || $v=='..') continue;
        $f = $dir.'/'.$v;
        if (is_dir($f)) scan_dir($f);
        // do file
    }
}

function composer_dep()
{
    $links = []; // ecCharts json data
    $node = [];
    $hc = [];

    $d = '/Users/dayday/hub/hyperf/hyperf/src';
    $r = scandir($d);
    foreach ($r as $v) {
        if ($v=='.'||$v=='..') continue;
        $json = file_get_contents($d. '/'. $v.'/composer.json');
        $deps = json_decode($json, true)['require'] ?? [];
        foreach ($deps as $k=>$ver) {
            $node[$k] = 1;
            $links[] = [
                'source' => 'hyperf/'.$v, // hyperf component
                'target' => $k, // dep of component
            ];
            $hc[] = ['hyperf/'.$v, $k];
        }
    }

    $n = [];
    foreach ($node as $k=>$v) {
        $n[] = ['id'=>$k];
    }
    echo json_encode($hc), "\n";
// echo json_encode($links);
}