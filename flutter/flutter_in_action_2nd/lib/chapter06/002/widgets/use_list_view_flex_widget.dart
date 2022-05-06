import 'package:flutter/material.dart';

class UseListViewFlexWidget extends StatelessWidget {
  const UseListViewFlexWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        const ListTile(title: Text('商品列表')),
        Expanded(
          child:
              ListView.builder(itemBuilder: (BuildContext context, int index) {
            return ListTile(title: Text('$index'));
          }),
        ),
      ],
    );
  }
}
