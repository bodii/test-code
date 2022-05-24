import 'package:flutter/material.dart';

class UseShowLoadingDialogWidget extends StatelessWidget {
  const UseShowLoadingDialogWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    Future<void> showLoadingDialog() {
      return showDialog(
          context: context,
          barrierDismissible: false,
          builder: (context) {
            return AlertDialog(
              content: Column(
                mainAxisSize: MainAxisSize.min,
                children: const [
                  CircularProgressIndicator(),
                  Padding(
                    padding: EdgeInsets.only(top: 26.0),
                    child: Text('正在加载，请稍后...'),
                  ),
                ],
              ),
            );
          });
    }

    return ElevatedButton(
      child: const Text('显示loading框'),
      onPressed: () {
        showLoadingDialog();
      },
    );
  }
}
