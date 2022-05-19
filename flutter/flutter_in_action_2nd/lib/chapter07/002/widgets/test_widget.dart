import 'package:flutter/material.dart';
import 'share_data_widget.dart';

class TestWidget extends StatefulWidget {
  const TestWidget({Key? key}) : super(key: key);

  @override
  _TestWidgetState createState() => _TestWidgetState();
}

class _TestWidgetState extends State<TestWidget> {
  @override
  Widget build(BuildContext context) {
    // 使用InheritedWidget中的共享数据
    return Text(ShareDataWidget.of(context)!.data.toString());
  }

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();
    // 父或祖先widget中的InheritedWidget改变(updateShouldNotify返回true)时会被调用。
    // 如果build中没有信赖InheritedWidget,则此回调不会被调用。
    debugPrint('Dependencies change');
  }
}
