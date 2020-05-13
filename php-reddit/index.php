<?php
require_once("phapper.php");
$phapper = new Phapper("protrandat", "badboyd13", "g1uQYjsE3h1_GA", "59pYu2YpxOAS3pxO1_-Eqa8QNxo", "https://www.reddit.com", "https://oauth.reddit.com");
var_dump($phapper->getSubreddits('popular', 25, null, null));
?>
