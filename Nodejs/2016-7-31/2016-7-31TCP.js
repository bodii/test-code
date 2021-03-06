/*
	
	传输控制协议（TCP）
	是一个面向连接的协议， 它保证了两台计算机之间数据传输的可靠性和顺序。
	换言之，TCP是一种传输层协议， 它可以让你将数据从一台计算机完整有序地传输到另一台
	计算机。
	IP是基于数据报的传输。这些数据报是独立进行传输的，送达的顺序也是无序的。

	//面向字节
	TCP对字符以及字符编码是完全无知的。不同的编码会导致传输的字节数不同。
	所以，TCP允许数据以ASCII字符（每个字符一个字节）或者 Unicode(即每个字符四个字节)进行传输。
	正是因为对消息格式没有严格的约束， 使得TCP有很好的灵活性。

	//可靠性
	由于TCP是基于底层不可靠的服务，因此，它必须要基于确认和超时实现一系列的机制来达到可靠性
	的要求。
	当数据发送出去后，发送方就会等待一个确认消息（表示数据包已经收到的简短的确认消息）。
	如果过了指定的窗口时间，还未收到确认消息，发送方就会对数据进行重发。
	这种机制有效地解决了如网络错误或者网络阻塞这样的不可预测的情况。

	//流控制
	TCP会通过一种叫流控制的方式来确保两点之间传输数据的平衡。

	//拥堵控制
	TCP有一种内置的机制能够控制数据包的延迟率及丢包率不会太高，以此来确认服务的质量（QoS）;


	//Telnet 工具
	Telnet是一个早期的网络协议，旨在提供双向的虚拟终端。在SSH出现前，它作为一种控制
	远程计算机的方式被广泛使用，如远程服务器管理。它是TCP协议上层的协议。

	//套接字（socket）
	要往TCP连接中写数据，必须首先创建一个HTTP请求，这就是套接字(socket)

*/


