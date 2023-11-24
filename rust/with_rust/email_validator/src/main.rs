use email_validator::*;

fn main() {
    let my_email = EmailAddress::new(
        "test".to_string(),
        "gmail.com".to_string());
    
    let local_part_validator = UsernameValidator { min_length: 3 };
    let domain_validator = DomainValidator {
        min_length: 3,
        allowed_domains: vec!["gmail.com".to_string()],
    };
    let email_validator = EmailValidatorImpl::new(
        vec![
            Box::new(local_part_validator),
            Box::new(domain_validator),
        ]);

    println!("{}", email_validator.is_valid(&my_email));
}
