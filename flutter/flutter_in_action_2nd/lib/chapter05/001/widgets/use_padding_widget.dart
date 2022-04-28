import 'package:flutter/material.dart';

class UsePaddingWidget extends StatelessWidget {
  const UsePaddingWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: const [
          Padding(
            padding: EdgeInsets.only(left: 8.0),
            child: Text('Hello world'),
          ),
          Padding(
            padding: EdgeInsets.symmetric(vertical: 8.0),
            child: Text('I am LiLi'),
          ),
          Padding(
            padding: EdgeInsets.fromLTRB(20.0, .0, 20.0, 20.0),
            child: Text('Your friend'),
          ),
        ],
      ),
    );
  }
}
