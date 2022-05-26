import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';

class UseGestureRecognizerWidget extends StatefulWidget {
  const UseGestureRecognizerWidget({Key? key}) : super(key: key);

  @override
  _UseGestureRecognizerWidgetState createState() =>
      _UseGestureRecognizerWidgetState();
}

class _UseGestureRecognizerWidgetState
    extends State<UseGestureRecognizerWidget> {
  final TapGestureRecognizer _tapGestureRecognizer = TapGestureRecognizer();
  bool _toggle = false;

  @override
  void dispose() {
    // 用到GestureRecognizer的话一定要调用dispose方法释放资源
    _tapGestureRecognizer.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Text.rich(
        TextSpan(children: [
          const TextSpan(text: '你好世界'),
          TextSpan(
            text: '点我变色',
            style: TextStyle(
              color: _toggle ? Colors.blue : Colors.red,
              fontSize: 30.0,
            ),
            recognizer: _tapGestureRecognizer
              ..onTap = () {
                setState(() {
                  _toggle = !_toggle;
                });
              },
          ),
          const TextSpan(text: '你好世界'),
        ]),
      ),
    );
  }
}
