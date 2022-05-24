import 'package:flutter/material.dart';

class UseShowDatePicker01Widget extends StatelessWidget {
  const UseShowDatePicker01Widget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    Future<DateTime?> _showDatePicker() {
      DateTime date = DateTime.now();
      return showDatePicker(
        context: context,
        initialDate: date,
        firstDate: date,
        lastDate: date.add(const Duration(days: 30)),
      );
    }

    return ElevatedButton(
      child: const Text('显示android风格日历'),
      onPressed: () {
        _showDatePicker();
      },
    );
  }
}
