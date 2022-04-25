import 'package:flutter/material.dart';

class TestRowsWidget extends StatelessWidget {
  const TestRowsWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: const [
            Text(' hello world '),
            Text(' I am Jack '),
          ],
        ),
        Row(
          mainAxisSize: MainAxisSize.min,
          mainAxisAlignment: MainAxisAlignment.center,
          children: const [
            Text(' hello world '),
            Text(' I am Jack '),
          ],
        ),
        Row(
          mainAxisAlignment: MainAxisAlignment.end,
          textDirection: TextDirection.rtl,
          children: const [
            Text(' hello world '),
            Text(' I am Jack '),
          ],
        ),
        Row(
          crossAxisAlignment: CrossAxisAlignment.start,
          verticalDirection: VerticalDirection.up,
          children: const [
            Text(
              ' hello world ',
              style: TextStyle(fontSize: 30.0),
            ),
            Text(' I am Jack '),
          ],
        ),
      ],
    );
  }
}
