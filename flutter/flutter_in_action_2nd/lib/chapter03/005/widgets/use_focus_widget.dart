import 'package:flutter/material.dart';

class UseFocusWidget extends StatefulWidget {
  const UseFocusWidget({Key? key}) : super(key: key);

  @override
  _UseFocusWidgetState createState() => _UseFocusWidgetState();
}

class _UseFocusWidgetState extends State<UseFocusWidget> {
  FocusNode focusNode1 = FocusNode();
  FocusNode focusNode2 = FocusNode();
  // FocusScopeNode focusScopeNode;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        children: [
          TextField(
            autofocus: true,
            focusNode: focusNode1,
            decoration: const InputDecoration(labelText: 'input1'),
          ),
          TextField(
            focusNode: focusNode2,
            decoration: const InputDecoration(labelText: 'input2'),
          ),
          Builder(builder: (ctx) {
            return Column(
              children: [
                ElevatedButton(
                  child: const Text('移动焦点'),
                  onPressed: () {
                    FocusScope.of(context).requestFocus(focusNode2);
                  },
                ),
                ElevatedButton(
                  child: const Text('隐藏键盘'),
                  onPressed: () {
                    focusNode1.unfocus();
                    focusNode2.unfocus();
                  },
                ),
              ].map((e) {
                return Padding(
                  padding: const EdgeInsets.symmetric(vertical: 8.0),
                  child: e,
                );
              }).toList(),
            );
          }),
        ].map((e) {
          return Padding(
            padding: const EdgeInsets.symmetric(vertical: 5.0),
            child: e,
          );
        }).toList(),
      ),
    );
  }
}
