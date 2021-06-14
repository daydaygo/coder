<?php
// league/csv https://csv.thephpleague.com/9.0/ https://packagist.org/packages/league/csv

use League\Csv\Reader;
use League\Csv\Writer;

require_once __DIR__ . '/vendor/autoload.php';


$a = '/Users/dayday/Downloads/billing_fix_frozen.csv';
$b = '/Users/dayday/Downloads/billingFixFrozen.csv';

// fgetcsv
$f = fopen($a, 'r');
$rows = [];
$row = fgetcsv($f); // header
while ($row = fgetcsv($f)) {
    $rows[] = $row;
}

// fputcsv
$file = 'tmp.csv';
$fd = fopen($file, 'w');
fputcsv($fd, ['a', 'b']); // header
foreach ($rows as $row) {
    fputcsv($fd, $rows);
}

// r
$csv = Reader::createFromPath($a);
$csv->setHeaderOffset(0);
var_dump($csv->getRecords());
return;

// w
$file = 'tmp.csv';
$csv = Writer::createFromPath($file, 'w+');
$header = ['id', 'name'];
$csv->insertOne($header);
$rows = [['1', 'czl']];
$csv->insertAll($rows);
// $csv->output($file); // http download
