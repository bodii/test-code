import 'package:flutter/material.dart';

class EchoRoute extends StatelessWidget {
  const EchoRoute({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    var args = ModalRoute.of(context)?.settings.arguments;
    return Scaffold(
      appBar: AppBar(title: const Text('接收路由参数Demo')),
      body: Center(
        child: Text('路由参数:${args.toString()}'),
      ),
    );
  }
}
