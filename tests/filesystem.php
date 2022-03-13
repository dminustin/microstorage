<?php
//Simple test written on PHP
function guidv4($data = null) {
    // Generate 16 bytes (128 bits) of random data or use the data passed into the function.
    $data = $data ?? random_bytes(16);
    assert(strlen($data) == 16);

    // Set version to 0100
    $data[6] = chr(ord($data[6]) & 0x0f | 0x40);
    // Set bits 6-7 to 10
    $data[8] = chr(ord($data[8]) & 0x3f | 0x80);

    // Output the 36 character UUID.
    return vsprintf('%s%s-%s-%s-%s-%s%s%s', str_split(bin2hex($data), 4));
}

$args = $_SERVER['argv'];
$dir = array_pop($args);
if (!file_exists($dir)) {
die(sprintf("Directory %s does not exists", $dir));
}

$files = glob($dir . "*.jpg");
foreach ($files as $file) {
$json = [
'mime'=>'image/jpeg',
'ext'=>'.jpg',
'transform'=>[
'resize'=>["800x800","160x160", "640x480"]
]
];
$uuid = guidv4();
copy($file, './filesystem/files/' . $uuid);
file_put_contents('./filesystem/data/' . $uuid . '.json', json_encode($json));
}