<?php

// 大delete -> 小delete
while (1) {
    mysqli_query("delete from logs where log_date<='2020-01-01' limit 1000");
    if (!mysqli_affected_rows()) {
        break;
    }
    usleep(50000);
}