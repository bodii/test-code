import 'package:flutter/material.dart';

import 'fade_route.dart';
import 'page_a_02.dart';

class UseCustomFadeRouteAnimationWidget extends StatelessWidget {
  const UseCustomFadeRouteAnimationWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          const Text('Fade Route page'),
          const SizedBox(height: 10.0),
          ElevatedButton(
            child: const Text('go to page A'),
            onPressed: () => Navigator.push(
              context,
              FadeRoute(
                builder: (context) => const PageA02(),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
