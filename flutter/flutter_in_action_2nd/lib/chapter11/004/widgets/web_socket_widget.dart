import 'package:flutter/material.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

class WebSocketWidget extends StatefulWidget {
  const WebSocketWidget({Key? key}) : super(key: key);

  @override
  _WebSocketWidgetState createState() => _WebSocketWidgetState();
}

class _WebSocketWidgetState extends State<WebSocketWidget> {
  final TextEditingController _controller = TextEditingController();
  late WebSocketChannel channel;
  String _text = '';

  @override
  void initState() {
    super.initState();
    channel =
        WebSocketChannel.connect(Uri.parse('wss://echo.websocket.events'));
  }

  @override
  void dispose() {
    channel.sink.close();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('websocket')),
      body: Padding(
        padding: const EdgeInsets.all(20.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Form(
              child: TextFormField(
                controller: _controller,
                decoration: const InputDecoration(labelText: 'Send a message'),
              ),
            ),
            StreamBuilder(
              stream: channel.stream,
              builder: (context, snapshot) {
                // 如果访问不了
                if (snapshot.hasError) {
                  _text = '访问失败...';
                } else if (snapshot.hasData) {
                  _text = 'echo: ${snapshot.data.toString()}';
                }

                return Padding(
                  padding: const EdgeInsets.symmetric(vertical: 24.0),
                  child: Text(_text),
                );
              },
            ),
          ],
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _sendMessage,
        tooltip: 'Send message',
        child: const Icon(Icons.send),
      ),
    );
  }

  void _sendMessage() {
    if (_controller.text.isNotEmpty) {
      channel.sink.add(_controller.text);
    }
  }
}
