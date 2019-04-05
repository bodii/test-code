function Stopwatch() {
    let startTime, endTime, running, duration = 0;

    Object.defineProperty(this, 'startTime', {
        get: function() {
            return startTime;
        },
        set: function(value) {
            startTime = value
        }
    });

    Object.defineProperty(this, 'endTime', {
        get: function() {
            return endTime;
        },
        set: function(value) {
            endTime = value;
        }
    });

    Object.defineProperty(this, 'duration', {
        get: function() {
            return duration;
        },
        set: function(value) {
            duration = value;
        }
    });

    Object.defineProperty(this, 'running', {
        get: function() {
            return running;
        },
        set: function(value) {
            running = value;
        }
    });
}

Stopwatch.prototype.start = function() {
    if (this.running)
        throw new Error('Stopwatch has already started.');

    this.startTime = new Date();
    this.running = true;
};

Stopwatch.prototype.stop = function() {
    if (!this.running)
        throw new Error('Stopwatch is not stoped.');

    this.endTime = new Date();
    this.running = false;

    const seconds = (this.endTime.getTime() - this.startTime.getTime()) / 1000;
    this.duration += seconds;
};

Stopwatch.prototype.reset = function() {
    this.running = false;
    this.startTime  = this.endTime = null;
    this.duration = 0;
}

const sw = new Stopwatch();
sw.duration = 10;