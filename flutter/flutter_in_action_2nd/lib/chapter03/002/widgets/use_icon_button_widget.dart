import 'package:flutter/material.dart';

class UseIconButtonWidget extends StatelessWidget {
  const UseIconButtonWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.symmetric(vertical: 15),
      child: Center(
        child: IconButton(
          icon: const Icon(Icons.thumb_up),
          onPressed: () {},
        ),
      ),
    );
  }
}
