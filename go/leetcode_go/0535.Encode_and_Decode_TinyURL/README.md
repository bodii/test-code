### 535. Encode and Decode TinyURL
> Note: This is a companion problem to the System Design problem: Design TinyURL.

TinyURL is a URL shortening service where you enter a URL such as https://leetcode.com/problems/design-tinyurl and it returns a short URL such as http://tinyurl.com/4e9iAk. Design a class to encode a URL and decode a tiny URL.

There is no restriction on how your encode/decode algorithm should work. You just need to ensure that a URL can be encoded to a tiny URL and the tiny URL can be decoded to the original URL.

Implement the Solution class:

* Solution() Initializes the object of the system.
* String encode(String longUrl) Returns a tiny URL for the given longUrl.
* String decode(String shortUrl) Returns the original long URL for the given shortUrl. It is guaranteed that the given shortUrl was encoded by the same object.



Example 1:

	Input: url = "https://leetcode.com/problems/design-tinyurl"
	Output: "https://leetcode.com/problems/design-tinyurl"

	Explanation:
	Solution obj = new Solution();
	string tiny = obj.encode(url); // returns the encoded tiny url.
	string ans = obj.decode(tiny); // returns the original url after deconding it.



Constraints:

* 1 <= url.length <= 10^4
* url is guranteed to be a valid URL.

----
### 535. TinyURL 的加密与解密
TinyURL是一种URL简化服务， 比如：当你输入一个URL https://leetcode.com/problems/design-tinyurl 时，它将返回一个简化的URL http://tinyurl.com/4e9iAk.

要求：设计一个 TinyURL 的加密 encode 和解密 decode 的方法。你的加密和解密算法如何设计和运作是没有限制的，你只需要保证一个URL可以被加密成一个TinyURL，并且这个TinyURL可以用解密方法恢复成原本的URL。

