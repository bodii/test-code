import 'package:flutter/material.dart' hide Page;

import 'text_water_mark_painter.dart';
import 'stagger_text_water_mark_painter.dart';
import 'water_mark.dart';
import 'page.dart';

class UseTextWaterMarkPainterWidget extends StatelessWidget {
  const UseTextWaterMarkPainterWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ListPage(
      children: [
        Page('文本水印', wTextWaterMark(), padding: false),
        Page('交错文本水印', wStaggerTextWaterMark(), padding: false),
      ],
    );
  }

  Widget wTextWaterMark() {
    return Stack(
      children: [
        wPage(),
        IgnorePointer(
            child: WaterMark(
          painter: TextWaterMarkPainter(
            text: 'Flutter @wendux',
            textStyle: const TextStyle(
              fontSize: 15,
              fontWeight: FontWeight.w200,
              color: Colors.black38,
            ),
            rotate: -20,
          ),
        )),
      ],
    );
  }

  Widget wStaggerTextWaterMark() {
    return Stack(
      children: [
        wPage(),
        IgnorePointer(
          child: WaterMark(
            painter: StaggerTextWaterMarkPainter(
              text: '《Flutter in action》',
              text2: 'wendux',
              textStyle: const TextStyle(
                color: Colors.black38,
              ),
              padding2: const EdgeInsets.only(left: 40),
              rotate: -10,
            ),
          ),
        ),
      ],
    );
  }

  Widget wPage() {
    return Center(
      child: ElevatedButton(
        child: const Text('按钮'),
        onPressed: () => debugPrint('tab'),
      ),
    );
  }
}
