#### String型的方法
char charAt(int index)  返回索引index位置处的字符。与数组一样，索引也是从头开始依次为0,1

int codePointAt(int index)  返回索引index位置处的字符(Unicode代码点)

int codePointBefore(int index) 返回索引index之前的字符(Unicode代码点)

int codePointCount(int beginIndex, int endIndex) 返回heginIndex和endIndex指定的范围内Unicode代码点数

int compareTo(String anotherString) 判断与字符串auotherString的字典顺序的大小关系

int compareToIgnoreCase(String str) 判断与字符串str的字典顺序的大小关系，不区分大小写

String concat(String str) 将字符串str拼接到字符串的末尾

boolean contains(CharSequence s) 判断字符串中是否包含字符序列s。如果包含，返回true

boolean contentEqueals(CharSequence cs) 判断与字符序列cs是否相等

boolean contentEqueals(StringBuffer sb) 判断与StringBuffer sb是否相等

static String copayValueOf(char[] data) 返回将字符数组data转换为字符串后的结果

static String copyValueOf(char[] data, int offset, int count) 返回将字符数组data中的指定部分转换为字符串后的结果

boolean endsWith(String suffix) 判断字符串的末尾是否是字符串suffix

boolean equeals(Object anObject) 判断与对象anObject是否相等

boolean equealsIgnoreCase(String anotherString) 判断与字符串anotherString是否相等

static String format(Locale l, String format, Object... args) 返回按照指定的环境l、格式字符串format及其后面的参数格式化后的字符串

static String format(String format, Object... args) 返回按照格式字符串format及其后面的参数格式化后的字符串

byte[] getBytes() 使用默认字符集，将字符串编码为字节序列，并返回一个保存结果的字节数组

byte[] getBytes(Charset charset) 使用字符集charset，将字符串编码为字节序列，并返回一个保存结果的字节数组

byte[] getBytes(String charsetName) 使用字符集charsetName, 将字符串编码为字节序列，并返回一个保存结果的字节数组

void getChars(int srcBegin, int srcEnd, char[] dst, int dstBegin) 将字符串的指定部分复制到目标字符数组dst的指定位置

int hashCode() 返回字符串的哈希码

int indexOf(int ch) 返回字符ch在字符串中第一次出现处的索引

int indexOf(int ch, int fromIndex) 返回字符ch在以fromIndex为开头索引的字符串第一次出现处的索引

int indexOf(String str) 返回字符串str在字符串第一次出现处的索引

int indexOf(Stringstr, int fromIndex) 返回字符串str在以fromIndex为开头索引的字符串第一次出现处的索引

String intern() 返回将字符串实例驻留后的字符串实例的引用

boolean isEmpty() 判断字符串的长度(字符个数)是否为0,如果为0，则返回true

int lastIndexOf(int ch) 返回字符ch在字符串中最后一次出现的索引

int lastIndexOf(int ch, int fromIndex) 返回字符ch在以fromIndex为开头索引的字符串中最后一次出现处的索引

int lastIndexOf(String str) 返回字符串str在字符串中最后一次出现的索引

int lastIndexOf(String str, int fromIndex) 返回字符串str在以fromIndex为开头索引的字符串中最后一次出现处的索引

int length() 返回字符串的长度(字符个数)

boolean matches(String regex) 判断字符串是否匹配正则表达式regex

int offsetByCodePoints(int index, int codePointOffset) 返回此字符串中从指定的index处偏移codePointOffset个代码点的索引

boolean regionMatches(boolean ignoreCase, int toffset, String other, int coffset, int len) 判断此字符串与字符串other是否相等。当ignoreCase赋为false时，则不区分大小写

boolean regionMatches(int toffset, String other, int coffset, int len) 判断此字符串与字符串other是否相等

String replace(char oldChar, char newChar) 返回将字符串中的所有oldChar替换为newChar后所创建的新的字符串

String replaceAll(String regex, String replacement) 使用指定字符串替换此字符串中匹配正则表达式regex的第一个子字符串

String[] split(String regex) 根据正则表达式regex的匹配位置来拆分此字符串

String[] split(String regex, int limit) 根据正则表达式regex的匹配位置来拆分此字符串limit次以内

boolean startsWith(String prefix) 判断此字符串是否以字符串prefix开始

boolean startsWidth(String prefix, int toffset) 判断此字符串从索引toffset开始的子字符串是否以字符串prefix开始

CharSequence subSequnce(int beginIndex, int endIndex) 返回一个字符序列，它是此序列的一个子序列

String subString(int beginIndex) 返回一个新的字符串，它是此字符串从索引beginIndex处开始的一个子字符串

String subString(int beginIndex, int endIndex) 返回一个新的字符串，它是此字符串从索引beginIndex处开始到endIndex处为止的一个子字符串

char[] toCharArray() 将字符串转换为一个新的字符数组，并返回

String toLowerCase() 使用默认语言环境的规则，将字符串中的所有字符都转换为小写

String toLowerCase(Locale locale) 使用语言环境locale的规则，将字符串中的所有字符都转换为小写

String toString() 直接返回此字符串

String toUpperCase() 使用默认语言环境的规则，将字符串中的所有字符都转换为大写

String trim() 返回字符串的副本

static String valueOf(boolean b) 返回boolean参数b的字符串表示形式

static String valueOf(char c) 返回char参数c的字符串表示形式

static String valueOf(char[] data) 返回char数组参数data的字符串表示形式

static String valueOf(char[] data, int offset, int count) 返回char数组参数data的特定子数组的字符串表示形式

static String valueOf(double d) 返回double参数d的字符串表示形式

static String valueOf(float f) 返回float参数f的字符串表示形式

static String valueOf(int i) 返回int参数i的字符串表示形式

static String valueOf(long l) 返回long参数l的字符串表示形式

static String valueOf(Object obj) 返回Object参数obj的字符串表示形式


