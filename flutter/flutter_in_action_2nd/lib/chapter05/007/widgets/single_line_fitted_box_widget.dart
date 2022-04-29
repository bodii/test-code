import 'package:flutter/material.dart';

class SingleLineFittedBoxWidget extends StatelessWidget {
  final Widget? child;
  const SingleLineFittedBoxWidget({Key? key, this.child}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(builder: (_, constraints) {
      return FittedBox(
        child: ConstrainedBox(
          constraints: constraints.copyWith(
            minWidth: constraints.maxWidth,
            maxWidth: double.infinity,
          ),
          child: child,
        ),
      );
    });
  }
}
