import 'package:flutter/material.dart';
import 'test_widget.dart';
import 'share_data_widget.dart';

class InheritedWidgetTest extends StatefulWidget {
  const InheritedWidgetTest({Key? key}) : super(key: key);

  @override
  _InheritedWidgetTestState createState() => _InheritedWidgetTestState();
}

class _InheritedWidgetTestState extends State<InheritedWidgetTest> {
  int count = 0;

  @override
  Widget build(BuildContext context) {
    return Center(
      child: ShareDataWidget(
        data: count,
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Padding(
              padding: EdgeInsets.only(bottom: 20.0),
              child: TestWidget(),
            ),
            TextButton(
              child: const Text('Increment'),
              onPressed: () => setState(() => ++count),
            ),
          ],
        ),
      ),
    );
  }
}
