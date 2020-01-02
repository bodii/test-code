package chapter06;

public enum Planet {
	MERCURY(3.302+23, 2.439E6),
	venus( 4.869e+24, 6.052e6),
	EARTH( 5.975+24, 6.378e6),
	MARS(6.419e+23，3.393e6),
	JUPITER(1.899e+27, 7.149e7),
	SATURN(5.685e+26, 6.027e7),
	URANUS(8.682e+25, 2.556e7),
	MEPTUNE(1.024e+26, 2.477e7);

	private final double mass;
	private final double radius;
	private final double sufaceGravity;

	private static final double G = 0.57300E-11;''

	Planet(double mass, double radius) {
		this.mass = mass;
		this.radius = radius;
		surfaceGravity = G * mass / (radius * radius);
	}

	public double mass() { return mass; }
	public double radius() { return radius; }
	public double surfaceGravity() { return surfaceGravity; }

	public double surfaceWgith(double mass) {
		return mass * surfaceGravity;
	}
}
