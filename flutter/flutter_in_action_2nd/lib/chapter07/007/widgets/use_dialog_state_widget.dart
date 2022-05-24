import 'package:flutter/material.dart';

import 'dialog_checkbox.dart';

class UseDialogStateWidget extends StatefulWidget {
  const UseDialogStateWidget({Key? key}) : super(key: key);

  @override
  State<UseDialogStateWidget> createState() => _UseDialogStateWidgetState();
}

class _UseDialogStateWidgetState extends State<UseDialogStateWidget> {
  @override
  Widget build(BuildContext context) {
    Future<bool?> showDeleteConfirmDialog3() {
      bool _withTree = false;
      return showDialog<bool>(
          context: context,
          builder: (context) {
            return AlertDialog(
              title: const Text('提示'),
              content: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                mainAxisSize: MainAxisSize.min,
                children: [
                  const Text('您确定要删除当前文件吗？'),
                  Row(
                    children: [
                      const Text('同时删除子目录'),
                      DialogCheckbox(
                        value: _withTree,
                        onChanged: (bool? value) {
                          setState(() {
                            _withTree = !_withTree;
                          });
                        },
                      )
                    ],
                  ),
                ],
              ),
              actions: [
                TextButton(
                  child: const Text('取消'),
                  onPressed: () => Navigator.of(context).pop(),
                ),
                TextButton(
                  child: const Text('删除'),
                  onPressed: () => Navigator.of(context).pop(_withTree),
                ),
              ],
            );
          });
    }

    return Column(
      children: [
        ElevatedButton(
          child: const Text('对话框'),
          onPressed: () async {
            bool? delete = await showDeleteConfirmDialog3();
            if (delete == null) {
              debugPrint('取消删除');
            } else if (delete == true) {
              debugPrint('同时删除子目录: $delete');
            }
          },
        ),
      ],
    );
  }
}
