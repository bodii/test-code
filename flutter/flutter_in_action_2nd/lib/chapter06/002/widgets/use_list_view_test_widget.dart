import 'package:english_words/english_words.dart';
import 'package:flutter/material.dart';

class UseListViewTestWidget extends StatefulWidget {
  const UseListViewTestWidget({Key? key}) : super(key: key);

  @override
  _UseListViewTestWidgetState createState() => _UseListViewTestWidgetState();
}

class _UseListViewTestWidgetState extends State<UseListViewTestWidget> {
  static const loadingTag = "##loading##";
  final List<String> _words = [loadingTag];

  @override
  void initState() {
    super.initState();
    _retrieveData();
  }

  void _retrieveData() {
    Future.delayed(const Duration(seconds: 2)).then((e) {
      setState(() {
        // 重新构建列表
        _words.insertAll(
          _words.length - 1,
          generateWordPairs().take(20).map((e) => e.asPascalCase).toList(),
        );
      });
    });
  }

  @override
  Widget build(BuildContext context) {
    return ListView.separated(
      itemCount: _words.length,
      itemBuilder: (context, index) {
        if (_words[index] == loadingTag) {
          // 不足100条，继续获取数据
          if (_words.length - 1 < 100) {
            _retrieveData();

            return Container(
              padding: const EdgeInsets.all(16.0),
              alignment: Alignment.center,
              child: const SizedBox(
                width: 24.0,
                height: 24.0,
                child: CircularProgressIndicator(
                  strokeWidth: 2.0,
                ),
              ),
            );
          } else {
            // 已加载100条数据，不再获取数据
            return Container(
              alignment: Alignment.center,
              padding: const EdgeInsets.all(16),
              child: const Text(
                "没有更多了",
                style: TextStyle(color: Colors.grey),
              ),
            );
          }
        }

        return ListTile(title: Text(_words[index]));
      },
      separatorBuilder: (context, index) => const Divider(height: .0),
    );
  }
}
