abstract class ObjectCache {
  Object getByKey(String key);
  void setByKey(String key, Object value);
}

abstract class Cache<T> {
  T getByKey(String key);
  void setByKey(String key, T value);
}
