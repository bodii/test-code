### 构建单线程Web服务器
Web服务器中涉及到的两个主要协议是**超文本传输协议(Hypertext Transfer Protocol, HTTP)**和**传输控制协议(Transmission Control Protocol, TCP)**。
这两者者是**请求-响应(request-response)**协议，也就是说，有**客户端(client)**来初始化请求，并有**服务端(server)**监听请求并向客户端提供响应。请求与响应的内容由协议本身定义。

HTTP通过TCP传输。

#### 监听TCP连接
`流(stream)`代表一个客户端和服务端之间打开的连接。
`连接(connection)`代表客户端连接服务端、服务端生成响应以及服务端关闭连接的全部请求/响应过程。

#### 读取请求
实际读取流，分两步：首先，在栈上声明一个buffer来存放读取到的数据。这里创建一个1024字节的缓冲区，它足以存放基本请求的数据并满足本章的目的。接着将缓冲区传递给stream.read，它会从TcpSteam中读取字节并放入缓冲区中。

#### 观察HTTP请求
`
Method Request-URI HTTP-Version CRLF
headers CRLF
message-body
`
第一行叫做`请求行`(request line)，它存放了客户端请求了什么的信息。

请求行接下来的是/，它代表客户端请求的`统一资源标识符(Uniform Resource Identifier, URI)`

最后一部分是客户端使用的HTTP版本，然后请求行以CRLF序列（回车和换行，carriage return line feed)结束。

#### 编写响应
响应格式如下:
`
HTTP-Version Status-Code Reason-Phrase CRLF
headers CRLF
message-body
`
第一行叫做`状态行(status line)`，它包含响应的HTTP版本、一个数字状态码用以总结请求的结果和一个描述之前状态码的文本原因短语。
之后是任意Header，和响应的body

这里是一个使用HTTP 1.1版本的响应例子，其状态码为200，原因短语为Ok，没有header，也没有body:
`
HTTP/1.1 200 Ok\r\n\r\n
`

#### 验证请求并有选择的进行响应
