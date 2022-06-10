import 'package:flutter/material.dart';

import 'widgets/socket_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('socket api')),
      body: const SocketWidget(),
    );
  }
}
