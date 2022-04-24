import 'package:flutter/material.dart';

import 'widgets/use_icon_widget.dart';
import 'widgets/use_image_widget.dart';
import 'widgets/use_network_image_widget.dart';
import 'widgets/use_image_box_fit_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('image&icon demo')),
      body: Container(
        alignment: Alignment.topLeft,
        padding: const EdgeInsets.only(top: 20),
        child: Column(
          children: const [
            UseImageWidget(),
            // UseNetworkImageWidget(),
            // UseImageBoxFitWidget(),
            UseIconWidget()
          ],
        ),
      ),
    );
  }
}
