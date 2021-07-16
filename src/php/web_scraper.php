<?php

require_once __DIR__ . "/vendor/autoload.php";

use Goutte\Client;
use Symfony\Component\HttpClient\HttpClient;

$client = new Client(HttpClient::create(['timeout' => 60]));
$crawler = $client->request('GET', 'https://www.symfony.com/blog/');

$link = $crawler->selectLink('Security Advisories')->link();
$crawler = $client->click($link);

$crawler->filter('h2 > a')->each(function ($node) {
    print $node->text()."\n";
});

$crawler = $client->request('GET', 'https://github.com/');
$crawler = $client->click($crawler->selectLink('Sign in')->link());
$form = $crawler->selectButton('Sign in')->form();
$crawler = $client->submit($form, ['login' => 'fabpot', 'password' => 'xxxxxx']);
$crawler->filter('.flash-error')->each(function ($node) {
    print $node->text()."\n";
});

// wechat
$crawler->filter('title')->text();
$crawler->filter('.rich_media_content')->text();
$crawler->filter('#post-date')->nextAll()->text();
$crawler->filter('#post-date')->text();