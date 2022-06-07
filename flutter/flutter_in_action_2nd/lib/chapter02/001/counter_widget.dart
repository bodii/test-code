import 'package:flutter/material.dart';

class CounterWidget extends StatefulWidget {
  final int initValue;
  const CounterWidget({Key? key, this.initValue = 0}) : super(key: key);

  @override
  State<CounterWidget> createState() => _CounterWidgetState();
}

class _CounterWidgetState extends State<CounterWidget> {
  int _counter = 0;

  @override
  void initState() {
    super.initState();
    _counter = widget.initValue;
    debugPrint('initState');
  }

  @override
  Widget build(BuildContext context) {
    debugPrint('build');
    return Scaffold(
      body: Center(
        child: TextButton(
          child: Text('$_counter'),
          onPressed: () => setState(
            () => ++_counter,
          ),
        ),
      ),
    );
  }

  @override
  didUpdateWidget(CounterWidget oldWidget) {
    super.didUpdateWidget(oldWidget);
    debugPrint('didUpdateWidget');
  }

  @override
  void deactivate() {
    super.deactivate();
    debugPrint('deactivate');
  }

  @override
  void dispose() {
    super.dispose();
    debugPrint('dispose');
  }

  @override
  void reassemble() {
    super.reassemble();
    debugPrint('reassemble');
  }

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();
    debugPrint('didChangeDependencies');
  }
}
