trait TemlateMethod {
    fn temlate_method(&self) {
        self.base_operation1();
        self.required_operation1();
        self.base_operation2();
        self.hook1();
        self.required_operation2();
        self.base_operation3();
        self.hook2();
    }

    fn base_operation1(&self) {
        println!("TemplateMehod says: I am doing the bulk of the work");
    }

    fn base_operation2(&self) {
        println!("TemplatedMethod says: But I let subclasses override some operations");
    }

    fn base_operation3(&self) {
        println!("TemplateMethod says: But I am doing the bulk of the work anyway");
    }

    fn hook1(&self) {}
    fn hook2(&self) {}

    fn required_operation1(&self);
    fn required_operation2(&self);
}

struct ConcreteStruct1;

impl TemlateMethod for ConcreteStruct1 {
    fn required_operation1(&self) {
        println!("ConcreteStruct1 says: Implemented Operation1")
    }

    fn required_operation2(&self) {
        println!("ConcreteStruct1 says: Implemented Operation2")
    }
}

struct ConcreteStruct2;

impl TemlateMethod for ConcreteStruct2 {
    fn required_operation1(&self) {
        println!("ConcreteStruct2 says: Implemented Operation1")
    }

    fn required_operation2(&self) {
        println!("ConcreteStruct2 says: Implemented Operation2")
    }
}

fn client_code(concrete: impl TemlateMethod) {
    concrete.temlate_method()
}

fn main() {
    println!("Same client code can work with different concrete implementations:");
    client_code(ConcreteStruct1);
    println!();

    println!("Same client code can work with different concrete implementations:");
    client_code(ConcreteStruct2);
}
