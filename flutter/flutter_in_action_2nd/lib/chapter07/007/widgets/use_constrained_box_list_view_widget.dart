import 'package:flutter/material.dart';

class UseConstrainedBoxListViewWidget extends StatelessWidget {
  const UseConstrainedBoxListViewWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    Future<void> showListDialog() async {
      int? index = await showDialog<int>(
        context: context,
        builder: (BuildContext context) {
          Widget child = Column(
            children: [
              const ListTile(title: Text('请选择')),
              Expanded(
                child: ListView.builder(
                  itemCount: 30,
                  itemBuilder: (BuildContext context, int index) {
                    return ListTile(
                      title: Text('$index'),
                      onTap: () => Navigator.of(context).pop(index),
                    );
                  },
                ),
              ),
            ],
          );
          return UnconstrainedBox(
            constrainedAxis: Axis.vertical,
            child: ConstrainedBox(
              constraints: BoxConstraints(
                minWidth: 280,
                maxWidth: MediaQuery.of(context).size.width - 100,
              ),
              child: Material(child: child, type: MaterialType.card),
            ),
          );
        },
      );

      if (index != null) {
        debugPrint('点击了： $index');
      }
    }

    return ElevatedButton(
      child: const Text('ListView对话框'),
      onPressed: () async => await showListDialog(),
    );
  }
}
