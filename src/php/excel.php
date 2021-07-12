<?php

// https://opensource.box.com/spout/docs
require_once __DIR__ . '/vendor/autoload.php';

use Box\Spout\Writer\Common\Creator\Style\StyleBuilder;
use Box\Spout\Writer\Common\Creator\WriterEntityFactory;
use Box\Spout\Common\Entity\Row;

$filePath = 'tmp/t.xlsx';
$writer = WriterEntityFactory::createXLSXWriter();
// $writer = WriterEntityFactory::createODSWriter();
// $writer = WriterEntityFactory::createCSVWriter();

$writer->openToFile($filePath); // write data to a file or to a PHP stream
//$writer->openToBrowser($fileName); // stream data directly to the browser

// sheet
$sheet1 = $writer->getCurrentSheet();
$sheet1->setName('sheet1');
$sheet2 = $writer->addNewSheetAndMakeItCurrent('sheet1');
$sheet2->setName('sheet2');
$writer->setCurrentSheet($sheet1);

/** Shortcut: add a row from an array of values */
$values = ['Carl', 'is', 'great!', 99.9];
$rowFromValues = WriterEntityFactory::createRowFromArray($values);
//$style = (new StyleBuilder())->setFormat('0.00')->build();
//$rowFromValues = WriterEntityFactory::createRowFromArray($values, $style);
$writer->addRow($rowFromValues);
$writer->addRows([$rowFromValues, $rowFromValues, $rowFromValues]);

$writer->close();

function get_sheet($writer, $sheet_name)
{

}
