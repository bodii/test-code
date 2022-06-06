import 'package:flutter/material.dart';
import 'custom_checkbox.dart';

class UseCustomCheckboxWidget extends StatefulWidget {
  const UseCustomCheckboxWidget({Key? key}) : super(key: key);

  @override
  _UseCustomCheckboxWidgetState createState() =>
      _UseCustomCheckboxWidgetState();
}

class _UseCustomCheckboxWidgetState extends State<UseCustomCheckboxWidget> {
  bool _checked = false;

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          CustomCheckbox(
            value: _checked,
            onChanged: _onChange,
          ),
          Padding(
            padding: const EdgeInsets.all(18.0),
            child: SizedBox(
              width: 16,
              height: 16,
              child: CustomCheckbox(
                strokeWidth: 1,
                radius: 1,
                value: _checked,
                onChanged: _onChange,
              ),
            ),
          ),
          SizedBox(
            width: 30,
            height: 30,
            child: CustomCheckbox(
              strokeWidth: 3,
              radius: 3,
              value: _checked,
              onChanged: _onChange,
            ),
          ),
        ],
      ),
    );
  }

  void _onChange(value) {
    setState(() => _checked = value);
  }
}
