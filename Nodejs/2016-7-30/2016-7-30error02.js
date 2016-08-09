function c () {
	b();
};

function b () {
	a();
};

function a () {
	setTimeout (function () {
		throw new Error('here');
	}, 10);
};

c(); 