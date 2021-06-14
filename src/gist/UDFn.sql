-- userDefineFunction

CREATE FUNCTION vipLevel (amount INT)
RETURNS int(10)

BEGIN
    DECLARE vip int(10);
    IF amount >=1000 THEN
        SET vip=2;
    ELSEIF amount>=600 THEN
        SET vip=1;
    ELSE
        SET vip=0;
    END IF;
    RETURN vip;
END;

-- unicode 转 utf-8：https://gist.github.com/joni/2956080
DROP FUNCTION STRINGDECODE;
CREATE FUNCTION STRINGDECODE(str TEXT CHARSET utf8)
RETURNS text CHARSET utf8 DETERMINISTIC
BEGIN
declare pos         int;
declare escape      char(6) charset utf8;
declare unescape    char(3) charset utf8;
set pos = locate('\u', str);
while pos > 0 do
    set escape = substring(str, pos, 5);
    set unescape = char(conv(substring(escape,2),16,10) using ucs2);
    set str = replace(str, escape, unescape);
    set pos = locate('\u', str, pos+1);
end while;
return str;
END;

SELECT STRINGDECODE("\u9648\u5fd7\u6797")
