import 'package:get/state_manager.dart';

class HomeSearchController extends GetxController {
  RxBool _toSearch = false.obs;

  void change() {
    if (_toSearch.isTrue) {
      _toSearch = false.obs;
    } else {
      _toSearch = true.obs;
    }
  }

  RxBool get val => _toSearch;

  bool getVal() {
    return _toSearch.isTrue;
  }
}
