// 抽象意味着应该隐藏的细节和复杂的部分，只显示或暴露必要的部分
function Circle(radius) {
    this.radius = radius;
    // private proteced
    let defaultLocation = { x: 0, y: 0 };

    // private function
    let computeOptimumLocation = function (factor) {

    }

    this.draw = function() {
        this.computeOptimumLocation(0.1);
        
        console.log('draw');
    }
}

const circle = new Circle(10);
circle.computeOptimumLocation();
circle.draw();