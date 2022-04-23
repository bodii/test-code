import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';

class TextSpanWidget extends StatelessWidget {
  const TextSpanWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    TapGestureRecognizer _tapRecognizer = TapGestureRecognizer();

    // return RichText();
    return Text.rich(
      TextSpan(
        children: [
          const TextSpan(
            text: 'Home: ',
          ),
          TextSpan(
            text: 'https://flutter-io.cn/',
            style: const TextStyle(color: Colors.blue),
            recognizer: _tapRecognizer..onTap = null,
          ),
        ],
      ),
    );
  }
}
