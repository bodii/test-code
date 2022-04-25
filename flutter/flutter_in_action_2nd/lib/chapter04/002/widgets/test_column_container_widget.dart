import 'package:flutter/material.dart';

class TestColumnContainerWidget extends StatelessWidget {
  const TestColumnContainerWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.green,
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          mainAxisSize: MainAxisSize.max,
          children: [
            Expanded(
              child: Container(
                color: Colors.red,
                child: Column(
                  mainAxisSize: MainAxisSize.max,
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: const [
                    Text('hello world '),
                    Text('I am Jack '),
                  ],
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
