早期的Unicode设计将所有的字符都使用16位来表示，不过16位最多只能表示65532个字符，因此人们
对其进行了扩展，以便能表示超过16位的字符。
十六进制数0x0000~0xFFFF的范围内的字符被称为基本多文种平面，即BMP(Basic Multi-lingual Plane),
1个字符可以使用16位来表示。当在Java中使用该范围内的字符时，1个字符就是16位的char型。
这样，BMP这外的增补字符(supplementary character)就要使用超过16位来表示。当在Java中使用该
范围内的字符时，1个字符要使用2个char来表示（为此，Java中提供了用来进行转换的API）。
Unicode的编码（符号化）方式分为UTF-8、UTF-16、UTF-32等诸多形式，Java中使用UTF-16.
在UTF-16中，BMP使用16位，而增补字符则采用代理对(surrogate pairs)方式，使用32位。
