import 'package:flutter/material.dart';

class UseAlertDialogWidget extends StatelessWidget {
  const UseAlertDialogWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
        title: const Text('提示'),
        content: const Text('您确定要删除当前文件吗？'),
        actions: [
          TextButton(
            child: const Text('取消'),
            onPressed: () => Navigator.of(context).pop(),
          ),
          TextButton(
            child: const Text('删除'),
            onPressed: () => Navigator.of(context).pop(true),
          ),
        ]);
  }
}
