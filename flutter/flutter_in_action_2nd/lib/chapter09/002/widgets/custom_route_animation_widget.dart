import 'package:flutter/material.dart';

import 'page_a.dart';
import 'page_b.dart';

class CustomRouteAnimationWidget extends StatelessWidget {
  const CustomRouteAnimationWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          ElevatedButton(
            child: const Text('go to page A'),
            onPressed: () => Navigator.push(
              context,
              PageRouteBuilder(
                transitionDuration: const Duration(microseconds: 500),
                pageBuilder: (BuildContext context, Animation<double> animation,
                    Animation secondaryAnimation) {
                  return FadeTransition(
                    opacity: animation,
                    child: const PageA(),
                  );
                },
              ),
            ),
          ),
          ElevatedButton(
            child: const Text('go to page B'),
            onPressed: () => Navigator.push(
              context,
              PageRouteBuilder(
                transitionDuration: const Duration(microseconds: 500),
                pageBuilder: (BuildContext context, Animation<double> animation,
                    Animation secondaryAnimation) {
                  return FadeTransition(
                    opacity: animation,
                    child: const PageB(),
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
    );
  }
}
