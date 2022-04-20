import 'package:flutter/material.dart';

class NewRoute extends StatelessWidget {
  const NewRoute({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('New route'),
      ),
      body: const Center(
        child: Text('This is new route'),
      ),
    );
  }
}
