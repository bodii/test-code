import 'package:flutter/material.dart';
import 'water_mark.dart';

class WaterMarkDemo extends StatelessWidget {
  const WaterMarkDemo({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return wTextWaterMark();
  }

  Widget wTextWaterMark() {
    return Stack(
      children: [
        wPage(),
        IgnorePointer(
          child: WaterMark(
            painter: TextWaterMarkPainter(
              text: 'Hello 中国 Flutter',
              textStyle: const TextStyle(
                fontSize: 15,
                fontWeight: FontWeight.w200,
                color: Colors.black38,
              ),
              rotate: -20,
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
