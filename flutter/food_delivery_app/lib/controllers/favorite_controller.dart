import 'package:get/state_manager.dart';

class FavoriteController extends GetxController {
  RxBool _favorite = false.obs;

  void change() {
    if (_favorite.isTrue) {
      _favorite = false.obs;
    } else {
      _favorite = true.obs;
    }
  }

  RxBool get val => _favorite;

  bool getVal() {
    return _favorite.isTrue;
  }
}
