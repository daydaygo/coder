# tech| 技术分享: 加解密那些事儿

关于加解密, 需要补充的基础知识:

- [https权威指南读书笔记](https://www.jianshu.com/p/b6719f6606d0)
- 密码分类与密码加密模式
- 对称加密与非对称加密

内容速览:
- 八一八有趣的人和事
- 详解常见的几种用法(PHP语言实现)

# 八一八有趣的人和事

- 使用 openssl 替代 mcrypt

mcrypt过时且停止维护

```
Warning
This feature was DEPRECATED in PHP 7.1.0, and REMOVED in PHP 7.2.0.
```

- heartblood 心脏出血

> 抱歉这次我们搞砸了, 要不我退你们钱? 哦, 对了, 你们也没付钱.
> Tim Hudson: 愿意奉献, 持之以恒的品质, 难能可贵!
> [OpenSSL“大管家”Steve Marquess：“心脏出血”重创后的OpenSSL（图灵访谈）](http://www.ituring.com.cn/article/497271)

![HeartBleed](http://file.ituring.com.cn/Original/17090a309b8f677561d9)

# 秘钥格式

在讲解秘钥格式之前, 请先确认自己是否知道 **对称加密/非对称加密** **公钥/私钥** 等基础知识
常用的三种格式转换: [PEM P12 JKS](https://myssl.com/cert_convert.html)

- 生成公钥/私钥

安装好 git 后, 就可以使用 `ssh-keygen`, 一顿回车就能帮我们搞定 **公钥/私钥**

```
ssh-keygen
# 选择秘钥保存的路径
# 是否设置密码

cat ~/.ssh/id_rsa # 私钥
-----BEGIN RSA PRIVATE KEY-----

cat ~/.ssh/id_rsa.pub # 公钥
```

- openssl 支持的秘钥的格式

查看 [php manual 中 openssl 的函数](http://php.net/manual/en/function.openssl-sign.php), 会查看到下面这样一致的说明:

```
priv_key_id
	resource - a key, returned by openssl_get_privatekey()
	string - a PEM formatted key
```

推荐使用 `a PEM formatted key`, 使用 php 中的 nowdoc 字符串表达式即可:

```php
$xxxPublicKey = <<<'KEY'
-----BEGIN PUBLIC KEY-----
 (内容被删除)
-----END PUBLIC KEY-----
KEY;

$xxxPrivateKey = <<<'KEY'
-----BEGIN RSA PRIVATE KEY-----
 (内容被删除)
-----END RSA PRIVATE KEY-----
KEY;
```

- 第一种格式转换: 字符串到 PEM format

有时候我们获取的秘钥是不带 `-----BEGIN RSA PRIVATE KEY-----` 等标识的, 需要转换成 PEM format, 很简单, 参照下面的函数:

```php
function format_key($str, $type = 1)
{
    $str = chunk_split($str, 64, "\n");
    if ($type) {
        $key = "-----BEGIN RSA PRIVATE KEY-----\n$str-----END RSA PRIVATE KEY-----\n";
    } else {
        $key = "-----BEGIN PUBLIC KEY-----\n$str-----END PUBLIC KEY-----\n";
    }
    echo $key;
}
```

- 第二种格式转换: pkcs12 到 PEM 格式

有时候我们获取的秘钥是 `.pfx` 或者 `.cer` 的文件, 里面的内容是二进制, 私钥很可能被加密了:

```php
// 私钥
openssl_pkcs12_read(file_get_contents('xxx-pri.pfx'), $pkfArr, 'pri_key_passwd');
var_dump($pkfArr['pkey']);

// 公钥
echo openssl_pkey_get_details(openssl_get_publickey(file_get_contents('xxx-pub.cer')))['key'];
```

## 使用 openssl 注意点

- 推荐 PEM 格式秘钥
- 了解使用的 **加密模式**: 比如 'DES-ECB`
- 了解使用的其他的一些编码处理: 比如 `hex2bin()` `base64_encode()`
- `openssl_encrypt()` 处理的数据, 已经经过 `base64_decode()` 处理

## DES

我们经常会使用 DES 来给敏感信息加密(比如用户四要素), 避免敏感信息在信息泄露后直接被暴露
加密方式: **对称加密**

- 使用 openssl

```php
// 加密
return bin2hex(base64_decode(openssl_encrypt($str, 'DES-ECB', $enKey)));
// 解密
return hex2bin(base64_encode(openssl_decrypt($str, 'DES-ECB', $enKey)));
// 获取 openssl 支持的加密模式
openssl_get_cipher_methods()
```

mcrypt vs openssl: 看完下面的对比代码, 对以前使用 mcrypt 的程序员致以敬意(或者说同情)

- 对应的 mcrypt

```php
<?php
class SecurityUtil{
    public static function DesEncrypt($dataStr, $key) {
        if (!function_exists( 'bin2hex')) {
            throw new Exception("bin2hex PHP5.4及以上版本支持此函数，也可自行实现！");
        }
        $Block_Size = mcrypt_get_block_size(MCRYPT_DES, MCRYPT_MODE_ECB);
        $InStr = self::pkcs5Pad($dataStr, $Block_Size);
        $Td = mcrypt_module_open(MCRYPT_DES, '', MCRYPT_MODE_ECB, '');
        $Iv = mcrypt_create_iv (mcrypt_enc_get_iv_size($Td), MCRYPT_RAND);
        try {
            ob_start();//关闭警告提示
            mcrypt_generic_init($Td, $key, $Iv);
            ob_end_clean();
        } catch (Exception $exc) {
            $exc->getTraceAsString();
            throw new Exception("解密异常：".$exc->getMessage());
        }
        $dataStr = mcrypt_generic($Td, $InStr);
        mcrypt_generic_deinit($Td);
        mcrypt_module_close($Td);
       return bin2hex($dataStr);
    }

    public static  function DesDecrypt($dataStr, $key) {
        if (!function_exists( 'hex2bin')) {
            throw new Exception("hex2bin PHP5.4及以上版本支持此函数，也可自行实现！");
        }
        $StrBin = \hex2bin($dataStr);
        try {
            ob_start();//关闭警告提示
            $DeStr = mcrypt_decrypt(MCRYPT_DES, $key, $StrBin, MCRYPT_MODE_ECB);
            ob_end_clean();
        } catch (Exception $exc) {
            $exc->getTraceAsString();
            throw new Exception("解密异常：".$exc->getMessage());
        }
        $ReturnStr = self::pkcs5Unpad($DeStr);
        return $ReturnStr;
    }

    private static function pkcs5Pad($text, $blocksize) {
        $pad = $blocksize - (strlen($text) % $blocksize);
        return $text . str_repeat(chr($pad), $pad);
    }

    private static function pkcs5Unpad($text) {
        $pad = ord($text {strlen($text) - 1});
        if ($pad > strlen($text)) {
            return false;
        }
        if (\strspn($text, chr($pad), strlen($text) - $pad) != $pad) {
            return false;
        }
        return substr($text, 0, - 1 * $pad);
    }
}
```

- mcrypt: 我可以更复杂点

```php
<?php
class DesHelper
{

	/**
	 * des 加密
	 * @param string $key 加密key
	 * @param string $message 加密前参数
	 * @return string
	 */
	public static function desEncrypt($key, $message){
		$ciphertext = self::des($key, $message, 1,0, null, null);
		return self::stringToHex($ciphertext);
	}

	/**
	 * 解密
	 * @param string $key 加密key
	 * @param string $hexString 加密后参数
	 * @return string
	 */
	public static function desDecrypt($key, $hexString){
		$ciphertext = self::hexToString($hexString);
		return self::des($key, $ciphertext, 0, 0, null, null);
	}

	/**
	 * @param string|mixed $s
	 * @return string
	 */
	private static function stringToHex ($s) {
		$r = "0x";
		$hexes = array ("0","1","2","3","4","5","6","7","8","9","a","b","c","d","e","f");
		for ($i=0; $i<strlen($s); $i++) {$r .= ($hexes [(ord($s{$i}) >> 4)] . $hexes [(ord($s{$i}) & 0xf)]);}
		return $r;
	}

	/**
	 * @param string|mixed $h
	 * @return string
	 */
	private static function hexToString ($h) {
		$r = "";
		for ($i= (substr($h, 0, 2)=="0x")?2:0; $i<strlen($h); $i+=2) {$r .= chr (base_convert (substr ($h, $i, 2), 16, 10));}
		return $r;
	}

	private static function des ($key, $message, $encrypt, $mode, $iv, $padding) {
		$spfunction1 = array (0x1010400,0,0x10000,0x1010404,0x1010004,0x10404,0x4,0x10000,0x400,0x1010400,0x1010404,0x400,0x1000404,0x1010004,0x1000000,0x4,0x404,0x1000400,0x1000400,0x10400,0x10400,0x1010000,0x1010000,0x1000404,0x10004,0x1000004,0x1000004,0x10004,0,0x404,0x10404,0x1000000,0x10000,0x1010404,0x4,0x1010000,0x1010400,0x1000000,0x1000000,0x400,0x1010004,0x10000,0x10400,0x1000004,0x400,0x4,0x1000404,0x10404,0x1010404,0x10004,0x1010000,0x1000404,0x1000004,0x404,0x10404,0x1010400,0x404,0x1000400,0x1000400,0,0x10004,0x10400,0,0x1010004);
		$spfunction2 = array (-0x7fef7fe0,-0x7fff8000,0x8000,0x108020,0x100000,0x20,-0x7fefffe0,-0x7fff7fe0,-0x7fffffe0,-0x7fef7fe0,-0x7fef8000,-0x80000000,-0x7fff8000,0x100000,0x20,-0x7fefffe0,0x108000,0x100020,-0x7fff7fe0,0,-0x80000000,0x8000,0x108020,-0x7ff00000,0x100020,-0x7fffffe0,0,0x108000,0x8020,-0x7fef8000,-0x7ff00000,0x8020,0,0x108020,-0x7fefffe0,0x100000,-0x7fff7fe0,-0x7ff00000,-0x7fef8000,0x8000,-0x7ff00000,-0x7fff8000,0x20,-0x7fef7fe0,0x108020,0x20,0x8000,-0x80000000,0x8020,-0x7fef8000,0x100000,-0x7fffffe0,0x100020,-0x7fff7fe0,-0x7fffffe0,0x100020,0x108000,0,-0x7fff8000,0x8020,-0x80000000,-0x7fefffe0,-0x7fef7fe0,0x108000);
		$spfunction3 = array (0x208,0x8020200,0,0x8020008,0x8000200,0,0x20208,0x8000200,0x20008,0x8000008,0x8000008,0x20000,0x8020208,0x20008,0x8020000,0x208,0x8000000,0x8,0x8020200,0x200,0x20200,0x8020000,0x8020008,0x20208,0x8000208,0x20200,0x20000,0x8000208,0x8,0x8020208,0x200,0x8000000,0x8020200,0x8000000,0x20008,0x208,0x20000,0x8020200,0x8000200,0,0x200,0x20008,0x8020208,0x8000200,0x8000008,0x200,0,0x8020008,0x8000208,0x20000,0x8000000,0x8020208,0x8,0x20208,0x20200,0x8000008,0x8020000,0x8000208,0x208,0x8020000,0x20208,0x8,0x8020008,0x20200);
		$spfunction4 = array (0x802001,0x2081,0x2081,0x80,0x802080,0x800081,0x800001,0x2001,0,0x802000,0x802000,0x802081,0x81,0,0x800080,0x800001,0x1,0x2000,0x800000,0x802001,0x80,0x800000,0x2001,0x2080,0x800081,0x1,0x2080,0x800080,0x2000,0x802080,0x802081,0x81,0x800080,0x800001,0x802000,0x802081,0x81,0,0,0x802000,0x2080,0x800080,0x800081,0x1,0x802001,0x2081,0x2081,0x80,0x802081,0x81,0x1,0x2000,0x800001,0x2001,0x802080,0x800081,0x2001,0x2080,0x800000,0x802001,0x80,0x800000,0x2000,0x802080);
		$spfunction5 = array (0x100,0x2080100,0x2080000,0x42000100,0x80000,0x100,0x40000000,0x2080000,0x40080100,0x80000,0x2000100,0x40080100,0x42000100,0x42080000,0x80100,0x40000000,0x2000000,0x40080000,0x40080000,0,0x40000100,0x42080100,0x42080100,0x2000100,0x42080000,0x40000100,0,0x42000000,0x2080100,0x2000000,0x42000000,0x80100,0x80000,0x42000100,0x100,0x2000000,0x40000000,0x2080000,0x42000100,0x40080100,0x2000100,0x40000000,0x42080000,0x2080100,0x40080100,0x100,0x2000000,0x42080000,0x42080100,0x80100,0x42000000,0x42080100,0x2080000,0,0x40080000,0x42000000,0x80100,0x2000100,0x40000100,0x80000,0,0x40080000,0x2080100,0x40000100);
		$spfunction6 = array (0x20000010,0x20400000,0x4000,0x20404010,0x20400000,0x10,0x20404010,0x400000,0x20004000,0x404010,0x400000,0x20000010,0x400010,0x20004000,0x20000000,0x4010,0,0x400010,0x20004010,0x4000,0x404000,0x20004010,0x10,0x20400010,0x20400010,0,0x404010,0x20404000,0x4010,0x404000,0x20404000,0x20000000,0x20004000,0x10,0x20400010,0x404000,0x20404010,0x400000,0x4010,0x20000010,0x400000,0x20004000,0x20000000,0x4010,0x20000010,0x20404010,0x404000,0x20400000,0x404010,0x20404000,0,0x20400010,0x10,0x4000,0x20400000,0x404010,0x4000,0x400010,0x20004010,0,0x20404000,0x20000000,0x400010,0x20004010);
		$spfunction7 = array (0x200000,0x4200002,0x4000802,0,0x800,0x4000802,0x200802,0x4200800,0x4200802,0x200000,0,0x4000002,0x2,0x4000000,0x4200002,0x802,0x4000800,0x200802,0x200002,0x4000800,0x4000002,0x4200000,0x4200800,0x200002,0x4200000,0x800,0x802,0x4200802,0x200800,0x2,0x4000000,0x200800,0x4000000,0x200800,0x200000,0x4000802,0x4000802,0x4200002,0x4200002,0x2,0x200002,0x4000000,0x4000800,0x200000,0x4200800,0x802,0x200802,0x4200800,0x802,0x4000002,0x4200802,0x4200000,0x200800,0,0x2,0x4200802,0,0x200802,0x4200000,0x800,0x4000002,0x4000800,0x800,0x200002);
		$spfunction8 = array (0x10001040,0x1000,0x40000,0x10041040,0x10000000,0x10001040,0x40,0x10000000,0x40040,0x10040000,0x10041040,0x41000,0x10041000,0x41040,0x1000,0x40,0x10040000,0x10000040,0x10001000,0x1040,0x41000,0x40040,0x10040040,0x10041000,0x1040,0,0,0x10040040,0x10000040,0x10001000,0x41040,0x40000,0x41040,0x40000,0x10041000,0x1000,0x40,0x10040040,0x1000,0x41040,0x10001000,0x40,0x10000040,0x10040000,0x10040040,0x10000000,0x40000,0x10001040,0,0x10041040,0x40040,0x10000040,0x10040000,0x10001000,0x10001040,0,0x10041040,0x41000,0x41000,0x1040,0x1040,0x40040,0x10000000,0x10041000);
		$masks = array (4294967295,2147483647,1073741823,536870911,268435455,134217727,67108863,33554431,16777215,8388607,4194303,2097151,1048575,524287,262143,131071,65535,32767,16383,8191,4095,2047,1023,511,255,127,63,31,15,7,3,1,0);

		$keys = self::des_createKeys ($key);
		$m=0;
		$len = strlen($message);
		$chunk = 0;
		$iterations = ((count($keys) == 32) ? 3 : 9);
		if ($iterations == 3) {$looping = (($encrypt) ? array (0, 32, 2) : array (30, -2, -2));}
		else {$looping = (($encrypt) ? array (0, 32, 2, 62, 30, -2, 64, 96, 2) : array (94, 62, -2, 32, 64, 2, 30, -2, -2));}

		if ($padding == 2) $message .= "        ";
		else if ($padding == 1) {$temp = 8-($len%8); $message .= chr($temp) . chr($temp) . chr($temp) . chr($temp) . chr($temp) . chr($temp) . chr($temp) . chr($temp); if ($temp==8) $len+=8;} //PKCS7 padding
		else if (!$padding) $message .= (chr(0) . chr(0) . chr(0) . chr(0) . chr(0) . chr(0) . chr(0) . chr(0)); //pad the message out with null bytes

		$result = "";
		$tempresult = "";

		if ($mode == 1) { //CBC mode
			$cbcleft = (ord($iv{$m++}) << 24) | (ord($iv{$m++}) << 16) | (ord($iv{$m++}) << 8) | ord($iv{$m++});
			$cbcright = (ord($iv{$m++}) << 24) | (ord($iv{$m++}) << 16) | (ord($iv{$m++}) << 8) | ord($iv{$m++});
			$m=0;
		}

		while ($m < $len) {
			$left = (ord($message{$m++}) << 24) | (ord($message{$m++}) << 16) | (ord($message{$m++}) << 8) | ord($message{$m++});
			$right = (ord($message{$m++}) << 24) | (ord($message{$m++}) << 16) | (ord($message{$m++}) << 8) | ord($message{$m++});

			if ($mode == 1) {if ($encrypt) {$left ^= $cbcleft; $right ^= $cbcright;} else {$cbcleft2 = $cbcleft; $cbcright2 = $cbcright; $cbcleft = $left; $cbcright = $right;}}

			$temp = (($left >> 4 & $masks[4]) ^ $right) & 0x0f0f0f0f; $right ^= $temp; $left ^= ($temp << 4);
			$temp = (($left >> 16 & $masks[16]) ^ $right) & 0x0000ffff; $right ^= $temp; $left ^= ($temp << 16);
			$temp = (($right >> 2 & $masks[2]) ^ $left) & 0x33333333; $left ^= $temp; $right ^= ($temp << 2);
			$temp = (($right >> 8 & $masks[8]) ^ $left) & 0x00ff00ff; $left ^= $temp; $right ^= ($temp << 8);
			$temp = (($left >> 1 & $masks[1]) ^ $right) & 0x55555555; $right ^= $temp; $left ^= ($temp << 1);

			$left = (($left << 1) | ($left >> 31 & $masks[31]));
			$right = (($right << 1) | ($right >> 31 & $masks[31]));

			for ($j=0; $j<$iterations; $j+=3) {
				$endloop = $looping[$j+1];
				$loopinc = $looping[$j+2];
				for ($i=$looping[$j]; $i!=$endloop; $i+=$loopinc) {
					$right1 = $right ^ $keys[$i];
					$right2 = (($right >> 4 & $masks[4]) | ($right << 28 & 0xffffffff)) ^ $keys[$i+1];
					$temp = $left;
					$left = $right;
					$right = $temp ^ ($spfunction2[($right1 >> 24 & $masks[24]) & 0x3f] | $spfunction4[($right1 >> 16 & $masks[16]) & 0x3f]
							| $spfunction6[($right1 >>  8 & $masks[8]) & 0x3f] | $spfunction8[$right1 & 0x3f]
							| $spfunction1[($right2 >> 24 & $masks[24]) & 0x3f] | $spfunction3[($right2 >> 16 & $masks[16]) & 0x3f]
							| $spfunction5[($right2 >>  8 & $masks[8]) & 0x3f] | $spfunction7[$right2 & 0x3f]);
				}
				$temp = $left; $left = $right; $right = $temp;
			}


			$left = (($left >> 1 & $masks[1]) | ($left << 31));
			$right = (($right >> 1 & $masks[1]) | ($right << 31));


			$temp = (($left >> 1 & $masks[1]) ^ $right) & 0x55555555; $right ^= $temp; $left ^= ($temp << 1);
			$temp = (($right >> 8 & $masks[8]) ^ $left) & 0x00ff00ff; $left ^= $temp; $right ^= ($temp << 8);
			$temp = (($right >> 2 & $masks[2]) ^ $left) & 0x33333333; $left ^= $temp; $right ^= ($temp << 2);
			$temp = (($left >> 16 & $masks[16]) ^ $right) & 0x0000ffff; $right ^= $temp; $left ^= ($temp << 16);
			$temp = (($left >> 4 & $masks[4]) ^ $right) & 0x0f0f0f0f; $right ^= $temp; $left ^= ($temp << 4);


			if ($mode == 1) {if ($encrypt) {$cbcleft = $left; $cbcright = $right;} else {$left ^= $cbcleft2; $right ^= $cbcright2;}}
			$tempresult .= (chr($left>>24 & $masks[24]) . chr(($left>>16 & $masks[16]) & 0xff) . chr(($left>>8 & $masks[8]) & 0xff) . chr($left & 0xff) . chr($right>>24 & $masks[24]) . chr(($right>>16 & $masks[16]) & 0xff) . chr(($right>>8 & $masks[8]) & 0xff) . chr($right & 0xff));

			$chunk += 8;
			if ($chunk == 512) {$result .= $tempresult; $tempresult = ""; $chunk = 0;}
		}

		return ($result . $tempresult);
	}

	private static function des_createKeys ($key) {
		$pc2bytes0  = array (0,0x4,0x20000000,0x20000004,0x10000,0x10004,0x20010000,0x20010004,0x200,0x204,0x20000200,0x20000204,0x10200,0x10204,0x20010200,0x20010204);
		$pc2bytes1  = array (0,0x1,0x100000,0x100001,0x4000000,0x4000001,0x4100000,0x4100001,0x100,0x101,0x100100,0x100101,0x4000100,0x4000101,0x4100100,0x4100101);
		$pc2bytes2  = array (0,0x8,0x800,0x808,0x1000000,0x1000008,0x1000800,0x1000808,0,0x8,0x800,0x808,0x1000000,0x1000008,0x1000800,0x1000808);
		$pc2bytes3  = array (0,0x200000,0x8000000,0x8200000,0x2000,0x202000,0x8002000,0x8202000,0x20000,0x220000,0x8020000,0x8220000,0x22000,0x222000,0x8022000,0x8222000);
		$pc2bytes4  = array (0,0x40000,0x10,0x40010,0,0x40000,0x10,0x40010,0x1000,0x41000,0x1010,0x41010,0x1000,0x41000,0x1010,0x41010);
		$pc2bytes5  = array (0,0x400,0x20,0x420,0,0x400,0x20,0x420,0x2000000,0x2000400,0x2000020,0x2000420,0x2000000,0x2000400,0x2000020,0x2000420);
		$pc2bytes6  = array (0,0x10000000,0x80000,0x10080000,0x2,0x10000002,0x80002,0x10080002,0,0x10000000,0x80000,0x10080000,0x2,0x10000002,0x80002,0x10080002);
		$pc2bytes7  = array (0,0x10000,0x800,0x10800,0x20000000,0x20010000,0x20000800,0x20010800,0x20000,0x30000,0x20800,0x30800,0x20020000,0x20030000,0x20020800,0x20030800);
		$pc2bytes8  = array (0,0x40000,0,0x40000,0x2,0x40002,0x2,0x40002,0x2000000,0x2040000,0x2000000,0x2040000,0x2000002,0x2040002,0x2000002,0x2040002);
		$pc2bytes9  = array (0,0x10000000,0x8,0x10000008,0,0x10000000,0x8,0x10000008,0x400,0x10000400,0x408,0x10000408,0x400,0x10000400,0x408,0x10000408);
		$pc2bytes10 = array (0,0x20,0,0x20,0x100000,0x100020,0x100000,0x100020,0x2000,0x2020,0x2000,0x2020,0x102000,0x102020,0x102000,0x102020);
		$pc2bytes11 = array (0,0x1000000,0x200,0x1000200,0x200000,0x1200000,0x200200,0x1200200,0x4000000,0x5000000,0x4000200,0x5000200,0x4200000,0x5200000,0x4200200,0x5200200);
		$pc2bytes12 = array (0,0x1000,0x8000000,0x8001000,0x80000,0x81000,0x8080000,0x8081000,0x10,0x1010,0x8000010,0x8001010,0x80010,0x81010,0x8080010,0x8081010);
		$pc2bytes13 = array (0,0x4,0x100,0x104,0,0x4,0x100,0x104,0x1,0x5,0x101,0x105,0x1,0x5,0x101,0x105);
		$masks = array (4294967295,2147483647,1073741823,536870911,268435455,134217727,67108863,33554431,16777215,8388607,4194303,2097151,1048575,524287,262143,131071,65535,32767,16383,8191,4095,2047,1023,511,255,127,63,31,15,7,3,1,0);

		$iterations = ((strlen($key) > 8) ? 3 : 1);
		$keys = array ();
		$shifts = array (0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0);
		$m=0;
		$n=0;

		for ($j=0; $j<$iterations; $j++) {
			$left = (@ord($key{$m++}) << 24) | (@ord($key{$m++}) << 16) | (@ord($key{$m++}) << 8) | ord(@$key{$m++});
			$right = (@ord($key{$m++}) << 24) | (@ord($key{$m++}) << 16) | (@ord($key{$m++}) << 8) | ord(@$key{$m++});

			$temp = (($left >> 4 & $masks[4]) ^ $right) & 0x0f0f0f0f; $right ^= $temp; $left ^= ($temp << 4);
			$temp = (($right >> 16 & $masks[16]) ^ $left) & 0x0000ffff; $left ^= $temp; $right ^= ($temp << 16);
			$temp = (($left >> 2 & $masks[2]) ^ $right) & 0x33333333; $right ^= $temp; $left ^= ($temp << 2);
			$temp = (($right >> 16 & $masks[16]) ^ $left) & 0x0000ffff; $left ^= $temp; $right ^= ($temp << 16);
			$temp = (($left >> 1 & $masks[1]) ^ $right) & 0x55555555; $right ^= $temp; $left ^= ($temp << 1);
			$temp = (($right >> 8 & $masks[8]) ^ $left) & 0x00ff00ff; $left ^= $temp; $right ^= ($temp << 8);
			$temp = (($left >> 1 & $masks[1]) ^ $right) & 0x55555555; $right ^= $temp; $left ^= ($temp << 1);

			$temp = ($left << 8) | (($right >> 20 & $masks[20]) & 0x000000f0);
			$left = ($right << 24) | (($right << 8) & 0xff0000) | (($right >> 8 & $masks[8]) & 0xff00) | (($right >> 24 & $masks[24]) & 0xf0);
			$right = $temp;

			for ($i=0; $i < count($shifts); $i++) {
				if ($shifts[$i] > 0) {
					$left = (($left << 2) | ($left >> 26 & $masks[26]));
					$right = (($right << 2) | ($right >> 26 & $masks[26]));
				} else {
					$left = (($left << 1) | ($left >> 27 & $masks[27]));
					$right = (($right << 1) | ($right >> 27 & $masks[27]));
				}
				$left = $left & -0xf;
				$right = $right & -0xf;

				$lefttemp = $pc2bytes0[$left >> 28 & $masks[28]] | $pc2bytes1[($left >> 24 & $masks[24]) & 0xf]
					| $pc2bytes2[($left >> 20 & $masks[20]) & 0xf] | $pc2bytes3[($left >> 16 & $masks[16]) & 0xf]
					| $pc2bytes4[($left >> 12 & $masks[12]) & 0xf] | $pc2bytes5[($left >> 8 & $masks[8]) & 0xf]
					| $pc2bytes6[($left >> 4 & $masks[4]) & 0xf];
				$righttemp = $pc2bytes7[$right >> 28 & $masks[28]] | $pc2bytes8[($right >> 24 & $masks[24]) & 0xf]
					| $pc2bytes9[($right >> 20 & $masks[20]) & 0xf] | $pc2bytes10[($right >> 16 & $masks[16]) & 0xf]
					| $pc2bytes11[($right >> 12 & $masks[12]) & 0xf] | $pc2bytes12[($right >> 8 & $masks[8]) & 0xf]
					| $pc2bytes13[($right >> 4 & $masks[4]) & 0xf];
				$temp = (($righttemp >> 16 & $masks[16]) ^ $lefttemp) & 0x0000ffff;
				$keys[$n++] = $lefttemp ^ $temp; $keys[$n++] = $righttemp ^ ($temp << 16);
			}
		}
		return $keys;
	}
}
```

## sign

添加签名, 验证数据是否有效, 以及通信双方是否正确
加密方式: **非对称加密**
使用场景: webapi的安全性, 使用 `token(access_token + refresh_token) + sign`

- 使用 openssl

```php
// 一种常用的签名方式
public function sign($params)
{
	ksort($params);
	$arr = [];
	foreach ($params as $k => $v) {
		$arr[] = "$k=$v";
	}
	$str = join('&', $arr);
	// openssl 其实就一行
	openssl_sign($str, $sign, $priKey); // 私钥加密
	// return base64_encode($sign);
	return bin2hex($sign);
}

return openssl_verify($str, hex2bin($sign), $pubkey); // 公钥解密
```

- 当然, 也有人把 openssl 玩得足够复杂

```php
<?php
class SignatureUtils{
    public static function Sign($Data,$PfxPath,$Pwd)
    {
        if (!function_exists( 'bin2hex')) {
            throw new Exception("bin2hex PHP5.4及以上版本支持此函数，也可自行实现！");
        }
        if(!file_exists($PfxPath)) {
           throw new Exception("私钥文件不存在！");
        }

        $pkcs12 = file_get_contents($PfxPath);
        $PfxPathStr=array();
        if (openssl_pkcs12_read($pkcs12, $PfxPathStr, $Pwd)) {
            $PrivateKey = $PfxPathStr['pkey'];
            $BinarySignature=NULL;
            if (openssl_sign($Data, $BinarySignature, $PrivateKey, OPENSSL_ALGO_SHA1)) {
                return bin2hex($BinarySignature);
            } else {
                throw new Exception("加签异常！");
            }
        } else {
            throw new Exception("私钥读取异常【密码和证书不匹配】！");
        }
    }

    public static function VerifySign($Data,$CerPath,$SignaTure)
    {
        if (!function_exists( 'hex2bin')) {
            throw new Exception("hex2bin PHP5.4及以上版本支持此函数，也可自行实现！");
        }
        if(!file_exists($CerPath)) {
            throw new Exception("私钥文件不存在！");
        }
        $PubKey = file_get_contents($CerPath);
        $Certs = openssl_get_publickey($PubKey);
        $ok = openssl_verify($Data,hex2bin($SignaTure), $Certs);
        if ($ok == 1) {
            return true;
        }
        return false;
    }
 }
?>
```

## other

这一部分, 我推荐称之为 **其他编码方式**, 而不应该放在 **加解密** 的范畴

- 使用 md5

```php
public function sign($params)
{
	ksort($params);
	$stringToBeSigned = $this->secret;
	foreach ($params as $k => $v)
	{
		if(!is_array($v) && "@" != substr($v, 0, 1))
		{
			$stringToBeSigned .= "$k$v";
		}
	}
	unset($k, $v);
	$stringToBeSigned .= $this->secret;
	$str = strtoupper(md5($stringToBeSigned)); // 统一使用大写
	return $str;
}

public function apiSign($arr)
{
	ksort($arr);
	$str = '';
	foreach ($arr as $k => $v) {
		$str .="$k$v";
	}
	$str = $this->app_secret . $str . $this->app_secret;
	return strtoupper(md5($str));
}
```

- base64: [魔鬼在细节中：Base64 你可能不知道的几个细节](https://liudanking.com/sitelog/%E9%AD%94%E9%AC%BC%E5%9C%A8%E7%BB%86%E8%8A%82%E4%B8%AD%EF%BC%9Abase64-%E4%BD%A0%E5%8F%AF%E8%83%BD%E4%B8%8D%E7%9F%A5%E9%81%93%E7%9A%84%E5%87%A0%E4%B8%AA%E7%BB%86%E8%8A%82/)

- `urlencode()/http_build_query()`
- `json_encode()`
