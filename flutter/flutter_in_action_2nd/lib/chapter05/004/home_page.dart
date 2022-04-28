import 'package:flutter/material.dart';

import 'widgets/use_matrix4_01_widget.dart';
import 'widgets/use_transform_translate_01_widget.dart';
import 'widgets/use_transform_rotate_01_widget.dart';
import 'widgets/use_transform_scale_01_widget.dart';
import 'widgets/use_transform_scale_02_widget.dart';
import 'widgets/use_rotated_box_01_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('变换transform')),
      body: Container(
        margin: const EdgeInsets.all(50),
        child: Column(
          children: const [
            UseMatrix401Widget(),
            SizedBox(height: 10),
            UseTransformTranslate01Widget(),
            SizedBox(height: 50),
            UseTransformRotate01Widget(),
            SizedBox(height: 50),
            UseTransformScale01Widget(),
            SizedBox(height: 20),
            UseTransformScale02Widget(),
            SizedBox(height: 30),
            UseRotatedBox01Widget(),
          ],
        ),
      ),
    );
  }
}
