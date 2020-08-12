#### DSL
DSL（Domain-Specific Language, 领域特定语言）指的是专注于特定问题领域的计算机语言。不同于通用的计算机
语言(GPL)，领域特定语言只用在某些特定的领域。

DSL语言能让我们以一种更优雅、更简洁的方式来表达和解决领域问题。之所以能够这样，是因为DSL语言刚好能够
用于这个特定的解决领域中存在的模式。

DSL简单讲就是对一个`特定问题（受限的表达能力)的方案模型更高层次的抽象表达（领域语言）`,使其更加简单易懂
（容易理解的语义及清晰的语义模型）。

DSL只是问题解决方案模型的外部封装，`这个模型可能是一个API库,也可能是一个完整的框架等。DSL提供了思考特定
领域问题的模型`语言，这使得我们可以更加简单、高效地解决问题。DSL聚焦一个特定的领域，简单易懂，功能极简
但完备，更加方便我们理解和使用模型。

比如用来显示网页的HTML语言，在Kotlin生态中有个kotlinx.html可在Web应用程序中用于构建HTML的DSL。它可以作
为传统模板系统（如JSP、FreeMarker等）的替代品。

kotlinx.html分别提供了kotlinx-html-jvm和kotlinx-html-js库的DSL，用于在JVM和浏览器（或其他JavaScript引擎）
中直接使用Kotlin代码来构建HTML，直接解放了原有的HTML标签式的前端代码。这样，我们也可以使用Kotlin来写
传统意义上的HTML页面了。Kotlin web编程将会更加简单纯净。