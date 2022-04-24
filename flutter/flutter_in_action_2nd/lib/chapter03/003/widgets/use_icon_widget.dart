import 'package:flutter/material.dart';

class UseIconWidget extends StatelessWidget {
  const UseIconWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    String icons = "";
    icons += "\uE03e";
    icons += " \uE237";
    icons += " \uE287";
    return Column(
      children: [
        Text(
          icons,
          style: const TextStyle(
            fontFamily: "MaterialIcons",
            fontSize: 24.0,
            color: Colors.green,
          ),
        ),
        Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: const [
            Icon(Icons.accessible, color: Colors.green),
            Icon(Icons.error, color: Colors.green),
            Icon(Icons.fingerprint, color: Colors.green),
          ],
        ),
      ],
    );
  }
}
