<?php


for($i=0; $i<=7000000; $i++){
	$num = rand(0,10000000);
	file_put_contents("./demo.txt", $num."\n", FILE_APPEND);
}
echo "end\n";

