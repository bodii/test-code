pub trait EmailValidator {
    fn is_valid(&self, email: &EmailAddress) -> bool;
}

pub struct EmailAddress {
    pub user_name: String,
    pub domain: String,
}

impl EmailAddress {
    pub fn new(user_name: String, domain: String) -> Self {
        Self {
            user_name,
            domain,
        }
    }
}

pub struct UsernameValidator {
    pub min_length: usize,
}

impl EmailValidator for UsernameValidator {
    fn is_valid(&self, email: &EmailAddress) -> bool {
        email.user_name.len() >= self.min_length
    }
}

pub struct DomainValidator {
    pub min_length: usize,
    pub allowed_domains: Vec<String>,
}

impl EmailValidator for DomainValidator {
    fn is_valid(&self, email: &EmailAddress) -> bool {
        email.domain.len() >= self.min_length
            && self.allowed_domains.contains(&email.domain)
    }
}

pub struct EmailValidatorImpl {
    pub validators: Vec<Box<dyn EmailValidator>>,
}

impl EmailValidatorImpl {
    pub fn new(validators: Vec<Box<dyn EmailValidator>>) -> Self {
        Self { validators }
    }

    pub fn is_valid(&self, email: &EmailAddress) -> bool {
        self.validators
            .iter()
            .all(|validator| validator.is_valid(email))
    }
}
