import 'package:flutter/material.dart';

class UseSimpleDialogWidget extends StatelessWidget {
  const UseSimpleDialogWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    Future<void> changeLanguage() async {
      int? i = await showDialog<int>(
        context: context,
        builder: (BuildContext context) {
          return SimpleDialog(
            title: const Text('请选择语言'),
            children: [
              SimpleDialogOption(
                  child: const Padding(
                    padding: EdgeInsets.symmetric(vertical: 6),
                    child: Text('中文简体'),
                  ),
                  onPressed: () {
                    Navigator.pop(context, 1);
                  }),
              SimpleDialogOption(
                child: const Padding(
                    padding: EdgeInsets.symmetric(vertical: 6),
                    child: Text('美国英语')),
                onPressed: () {
                  Navigator.pop(context, 2);
                },
              ),
            ],
          );
        },
      );

      if (i != null) {
        debugPrint('选择了:${i == 1 ? "中文简体" : "美国英文"}');
      }
    }

    return ElevatedButton(
      child: const Text('对话框'),
      onPressed: () async {
        await changeLanguage();
      },
    );
  }
}
