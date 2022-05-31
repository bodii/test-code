import 'package:flutter/material.dart';

import '../home_page.dart';

class PageB extends StatelessWidget {
  const PageB({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('b page demo')),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Text('This is page b'),
            ElevatedButton(
              child: const Text('back'),
              onPressed: () => Navigator.push(
                context,
                PageRouteBuilder(
                  transitionDuration:
                      const Duration(microseconds: 500), // 动画500毫秒
                  pageBuilder: (BuildContext context,
                      Animation<double> animation,
                      Animation secondaryAnimation) {
                    return FadeTransition(
                      opacity: animation,
                      child: const HomePage(),
                    );
                  },
                ),
              ),
            ),
          ]
              .map(
                (e) => Padding(
                  padding: const EdgeInsets.symmetric(vertical: 5.0),
                  child: e,
                ),
              )
              .toList(),
        ),
      ),
    );
  }
}
