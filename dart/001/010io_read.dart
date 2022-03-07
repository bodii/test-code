import "dart:io";

void main() async {
	var config = File("config.txt");
	var StringContents = await config.readAsString();
	print('The file is ${StringContents.length} characters long.');

	var lines = await config.readAsLines();
	print('The file is ${lines.length} lines long.');
}
