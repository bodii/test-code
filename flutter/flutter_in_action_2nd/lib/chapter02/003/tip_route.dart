import 'package:flutter/material.dart';

class TipRoute extends StatelessWidget {
  const TipRoute({Key? key, required this.text}) : super(key: key);
  final String text;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('提示'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(18),
        child: Center(
          child: Column(
            children: [
              Text(text),
              ElevatedButton(
                onPressed: () => Navigator.pop(context, '我是返回值'),
                child: const Text('返回'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
