use super::{into_next, Department, Patient};

pub struct Reception {
    next: Option<Box<dyn Department>>,
}

impl Reception {
    pub fn new(next: impl Department + 'static) -> Self {
        Self {
            next: into_next(next),
        }
    }
}

impl Department for Reception {
    fn handle(&mut self, patient: &mut Patient) {
        if patient.registration_done {
            println!("Patient registration is already done");
        } else {
            println!("Reception registration a patient {}", patient.name);
            patient.registration_done = true;
        }
    }

    fn next(&mut self) -> &mut Option<Box<dyn Department>> {
        &mut self.next
    }
}
