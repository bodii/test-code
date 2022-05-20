import 'package:flutter/material.dart';

class ValueListenableWidget extends StatefulWidget {
  const ValueListenableWidget({Key? key}) : super(key: key);

  @override
  _ValueListenableWidgetState createState() => _ValueListenableWidgetState();
}

class _ValueListenableWidgetState extends State<ValueListenableWidget> {
  final ValueNotifier<int> _counter = ValueNotifier<int>(0);
  static const double textScaleFactor = 1.5;

  @override
  Widget build(BuildContext context) {
    debugPrint('build');

    return Scaffold(
      appBar: AppBar(title: const Text('ValueListenableBuilder')),
      body: Center(
        child: ValueListenableBuilder<int>(
          builder: (context, value, Widget? child) {
            return Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                child!,
                Text(
                  '$value 次',
                  textScaleFactor: textScaleFactor,
                ),
              ],
            );
          },
          valueListenable: _counter,
          child: const Text(
            '点击了 ',
            textScaleFactor: textScaleFactor,
          ),
        ),
      ),
      floatingActionButton: FloatingActionButton(
        child: const Icon(Icons.add),
        onPressed: () => _counter.value += 1,
      ),
    );
  }
}
