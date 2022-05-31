import 'package:flutter/material.dart';

import 'fade_route.dart';
import '../home_page.dart';

class PageA02 extends StatelessWidget {
  const PageA02({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('a 02 page')),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Text('Fade Route page'),
            const SizedBox(height: 10.0),
            ElevatedButton(
              child: const Text('back'),
              onPressed: () => Navigator.push(
                context,
                FadeRoute(
                  builder: (context) => const HomePage(),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
