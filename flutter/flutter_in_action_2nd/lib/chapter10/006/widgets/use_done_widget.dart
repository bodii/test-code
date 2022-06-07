import 'package:flutter/material.dart';

import 'done_widget.dart';

class UseDoneWidget extends StatelessWidget {
  const UseDoneWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    double width = MediaQuery.of(context).size.width;
    return Center(
      child: SizedBox(
        width: width * .3,
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceEvenly,
          children: const [
            DoneWidget(
              outline: true,
            ),
            Text(
              '操作成功',
              style: TextStyle(
                fontWeight: FontWeight.bold,
              ),
            ),
            DoneWidget(),
          ],
        ),
      ),
    );
  }
}
