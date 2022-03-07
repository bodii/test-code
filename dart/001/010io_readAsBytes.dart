import "dart:io";

void main() async {
	var config = File('config.txt');

	var contents = await config.readAsBytes();
	print('The file is ${contents.length} bytes long.');

	var config2 = File('config2.txt');
	try {
		var contents2 = await config2.readAsString();
		print(contents2);
	} catch(e) {
		print(e);
	}
}
