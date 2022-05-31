import 'package:flutter/material.dart';
import '../home_page.dart';

class PageA extends StatelessWidget {
  const PageA({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('a page demo')),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Text('This is page a'),
            ElevatedButton(
              child: const Text('back'),
              onPressed: () => Navigator.push(
                context,
                PageRouteBuilder(
                  transitionDuration: const Duration(milliseconds: 500),
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
