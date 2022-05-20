import 'package:flutter/material.dart';

import 'change_notifier_provider.dart';

class Consumer<T> extends StatelessWidget {
  const Consumer({Key? key, required this.builder}) : super(key: key);

  final Widget Function(BuildContext context, T? value) builder;

  @override
  Widget build(BuildContext context) {
    return builder(
      context,
      ChangeNotifierProvider.of<T>(context),
    );
  }
}
