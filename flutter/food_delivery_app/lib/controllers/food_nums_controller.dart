import 'package:get/get.dart';

class FoodNumsController extends GetxController {
  RxInt _nums = 0.obs;
  static const int _maxNums = 99;

  void increase() {
    if (_nums < _maxNums) {
      ++_nums;
    }
  }

  void decrease() {
    if (_nums > 0) {
      --_nums;
    }
  }

  RxInt get nums => _nums;
}
