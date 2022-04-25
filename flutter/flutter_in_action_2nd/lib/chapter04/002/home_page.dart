import 'package:flutter/material.dart';

// import 'widgets/test_rows_widget.dart';
import 'widgets/use_column_widget.dart';
import 'widgets/test_column_container_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('row and column demo')),
      body: const TestColumnContainerWidget(),
      // Column(
      //   children: const [
      //     // TestRowsWidget(),
      //     // UseColumnWidget(),
      //     TestColumnContainerWidget(),
      //   ],
      // ),
    );
  }
}
