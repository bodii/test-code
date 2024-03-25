use serde::Deserialize;

#[derive(Debug, Deserialize)]
#[allow(dead_code)]
pub struct AuthRequest {
    pub code: String,
    pub state: String,
}