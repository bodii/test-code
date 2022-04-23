import 'dart:developer';

import 'package:flutter/material.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    // debugger(when: true, message: 'start testing debugger use');
    log('warning',
        time: DateTime.now(), name: 'test log', level: 3, sequenceNumber: 1);
    return Scaffold(
      appBar: AppBar(title: const Text('Log and debug demo')),
      body: Center(
        child: TextButton(
          child: const Text('dump app'),
          onPressed: () => debugDumpApp(),
        ),
      ),
    );
  }
}
