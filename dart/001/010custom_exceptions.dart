class FooException implements Exception {
	final String? msg;

	const FooException([this.msg]);

	@override
	String toString() => msg ?? 'fooException';
}

void main() {
	throw FooException();
}
