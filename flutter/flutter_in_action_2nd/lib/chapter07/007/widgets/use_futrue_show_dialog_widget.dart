import 'package:flutter/material.dart';

class UseFutrueShowDialogWidget extends StatelessWidget {
  const UseFutrueShowDialogWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    Future<bool?> showDeleteConfirmDialog01() {
      return showDialog<bool>(
        context: context,
        builder: (context) {
          return AlertDialog(
            title: const Text('提示'),
            content: const Text('你确定要删除当前文件吗?'),
            actions: [
              TextButton(
                child: const Text('取消'),
                onPressed: () => Navigator.of(context).pop(),
              ),
              TextButton(
                child: const Text('确定'),
                onPressed: () => Navigator.of(context).pop(true),
              ),
            ],
          );
        },
      );
    }

    return ElevatedButton(
      child: const Text('对话框'),
      onPressed: () async {
        bool? delete = await showDeleteConfirmDialog01();
        if (delete == null) {
          debugPrint('取消删除');
        } else {
          debugPrint('已确认删除');
        }
      },
    );
  }
}
