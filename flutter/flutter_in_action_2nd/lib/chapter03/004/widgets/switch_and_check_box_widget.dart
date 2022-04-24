import 'package:flutter/material.dart';

class SwitchAndCheckBoxWidget extends StatefulWidget {
  const SwitchAndCheckBoxWidget({Key? key}) : super(key: key);

  @override
  _SwitchAndCheckBoxWidgetState createState() =>
      _SwitchAndCheckBoxWidgetState();
}

class _SwitchAndCheckBoxWidgetState extends State<SwitchAndCheckBoxWidget> {
  bool _switchSelected = true;
  bool _checkBoxSelected = true;

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Switch(
          value: _switchSelected,
          onChanged: (value) {
            setState(() {
              _switchSelected = value;
            });
          },
        ),
        Checkbox(
          value: _checkBoxSelected,
          activeColor: Colors.red,
          onChanged: (value) {
            setState(() {
              _checkBoxSelected = value!;
            });
          },
        )
      ],
    );
  }
}
