use async_session::MemoryStore;
use axum::extract::FromRef;
use oauth2::basic::BasicClient;

#[derive(Clone)]
pub struct AppState {
    pub store: MemoryStore,
    pub oauth_client: BasicClient,
}

impl FromRef<AppState> for MemoryStore {
    fn from_ref(input: &AppState) -> Self {
        input.store.clone()
    }
}

impl FromRef<AppState> for BasicClient {
    fn from_ref(input: &AppState) -> Self {
        input.oauth_client.clone()
    }
}