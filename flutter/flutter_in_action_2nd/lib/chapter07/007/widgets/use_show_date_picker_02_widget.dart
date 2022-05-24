import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class UseShowDatePicker02Widget extends StatelessWidget {
  const UseShowDatePicker02Widget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    Future<DateTime?> _showDatePicker() {
      DateTime date = DateTime.now();
      return showCupertinoModalPopup(
          context: context,
          builder: (context) {
            return SizedBox(
              height: 200,
              child: CupertinoDatePicker(
                mode: CupertinoDatePickerMode.dateAndTime,
                minimumDate: date,
                maximumDate: date.add(const Duration(days: 30)),
                maximumYear: date.year + 1,
                onDateTimeChanged: (DateTime value) {
                  debugPrint(value.toLocal().toString());
                },
              ),
            );
          });
    }

    return ElevatedButton(
      child: const Text('显示ios风格日历'),
      onPressed: () {
        _showDatePicker();
      },
    );
  }
}
