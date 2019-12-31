package chapter04;

/**
 * Constant interface antipattern - do not use!
 * 常量接口模式是对接口的不良使用。
 */

public interface PhysicalConstants {
    static final double AVOGADROS_NUMBER = 6.022_140_857e23;

    static final double BOLTZMANN_CONSTANT = 1.380_648_52e-23;

    static final double ELECTRON_MASS = 9.109_383_56e-31;
}