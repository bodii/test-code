import 'package:flutter/material.dart';

class UseFittedBox01Widget extends StatelessWidget {
  const UseFittedBox01Widget({Key? key}) : super(key: key);

  Widget wContainer(BoxFit boxFit) {
    return Container(
      width: 50,
      height: 50,
      color: Colors.red,
      child: FittedBox(
        fit: boxFit,
        child: Container(
          width: 60,
          height: 70,
          color: Colors.blue,
        ),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Column(
        children: [
          wContainer(BoxFit.none),
          const Text('Wendux'),
          wContainer(BoxFit.contain),
          const Text('Flutter'),
        ],
      ),
    );
  }
}
