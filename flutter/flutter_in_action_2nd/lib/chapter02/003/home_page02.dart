import 'package:flutter/material.dart';
import 'router_test_route.dart';
import 'new_route.dart';

class HomePage02 extends StatelessWidget {
  const HomePage02({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Flutter Demo Home Page')),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            TextButton(
              child: const Text("open new route"),
              onPressed: () {
                Navigator.of(context).pushNamed('echo_page', arguments: 'hi');
              },
            ),
          ],
        ),
      ),
    );
  }
}
