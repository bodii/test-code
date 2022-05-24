import 'package:flutter/material.dart';

class UseShowModalBottomSheetWidget extends StatelessWidget {
  const UseShowModalBottomSheetWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    Future<int?> _showModalBottomSheet() {
      return showModalBottomSheet<int>(
          context: context,
          builder: (BuildContext context) {
            return ListView.builder(
              itemCount: 30,
              itemBuilder: (BuildContext context, int index) {
                return ListTile(
                  title: Text('$index'),
                  onTap: () => Navigator.of(context).pop(index),
                );
              },
            );
          });
    }

    return ElevatedButton(
      child: const Text('显示底部菜单列表'),
      onPressed: () async {
        int? i = await _showModalBottomSheet();
        debugPrint(i.toString());
      },
    );
  }
}
