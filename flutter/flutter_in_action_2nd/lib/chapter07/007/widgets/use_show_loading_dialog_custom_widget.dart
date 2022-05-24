import 'package:flutter/material.dart';

class UseShowLoadingDialogCustomWidget extends StatelessWidget {
  const UseShowLoadingDialogCustomWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    Future<void> showLoadingDialog() {
      return showDialog(
          context: context,
          barrierDismissible: false,
          builder: (context) {
            return UnconstrainedBox(
              constrainedAxis: Axis.vertical,
              child: SizedBox(
                width: 280,
                child: AlertDialog(
                  content: Column(
                    mainAxisSize: MainAxisSize.min,
                    children: const [
                      CircularProgressIndicator(
                        value: .8,
                      ),
                      Padding(
                        padding: EdgeInsets.only(top: 26.0),
                        child: Text('正在加载，请稍后...'),
                      ),
                    ],
                  ),
                ),
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
