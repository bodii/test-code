import 'package:flutter/material.dart';

import 'widgets/use_constrained_box_widget.dart';
import 'widgets/use_unconstrained_box_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('ConstraineBox、SizedBox'),
        actions: const [
          UnconstrainedBox(
            child: SizedBox(
              width: 20,
              height: 20,
              child: CircularProgressIndicator(
                strokeWidth: 3,
                valueColor: AlwaysStoppedAnimation(
                  Colors.white70,
                ),
              ),
            ),
          ),
        ],
      ),
      body: Column(
        children: const [
          UseConstrainedBoxWidget(),
          UseUnconstrainedBoxWidget(),
        ],
      ),
    );
  }
}
