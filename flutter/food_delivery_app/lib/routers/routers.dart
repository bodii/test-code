import 'package:flutter/material.dart';
import 'package:food_delivery_app/models/food.dart';
import 'package:food_delivery_app/screens/detail/widgets/detail.dart';
import 'package:food_delivery_app/screens/home/widgets/home.dart';

const homePage = '/';
const detailPage = '/detail';
const initPage = homePage;

/// getRouters get router list
Map<String, Widget Function(BuildContext)> getRouters(BuildContext context) {
  return {
    homePage: (context) => const HomePage(),
  };
}
