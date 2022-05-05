import 'package:flutter/material.dart';

class UseSingleChildScrollView extends StatelessWidget {
  const UseSingleChildScrollView({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    String str = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
    double width = MediaQuery.of(context).size.width;

    return Scrollbar(
      child: SingleChildScrollView(
        padding: const EdgeInsets.all(16.0),
        child: Container(
          margin: EdgeInsets.only(left: width * 0.8),
          child: Column(
            children: str
                .split('')
                .map((e) => Text(e, textScaleFactor: 2.0))
                .toList(),
          ),
        ),
      ),
    );
  }
}
