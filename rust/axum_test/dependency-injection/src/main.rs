use std::{
    collections::HashMap,
    sync::{Arc, Mutex},
};

use axum::{
    extract::{Path, State},
    http::StatusCode,
    routing::{get, post},
    Json, Router,
};
use serde::{Deserialize, Serialize};
use tracing_subscriber::{layer::SubscriberExt, util::SubscriberInitExt};
use uuid::Uuid;

#[tokio::main]
async fn main() {
    subscriber();

    let make_service = app(InMemoryUserRepo::default());
    let tcp_listener = tokio::net::TcpListener::bind("127.0.0.1:3000")
        .await
        .unwrap();

    tracing::debug!("listening on {}", tcp_listener.local_addr().unwrap());

    axum::serve(tcp_listener, make_service).await.unwrap();
}

fn subscriber() {
    let layer = tracing_subscriber::EnvFilter::try_from_default_env()
        .unwrap_or_else(|_| "dependency_injection=debug".into());
    tracing_subscriber::registry()
        .with(layer)
        .with(tracing_subscriber::fmt::layer())
        .init();
}

fn app(user_repo: InMemoryUserRepo) -> Router {
    let using_dyn = Router::new()
        .route("/users/:id", get(get_user_dyn))
        .route("/users", post(create_user_dyn))
        .with_state(AppStateDyn {
            user_repo: Arc::new(user_repo.clone()),
        });

    let using_generic = Router::new()
        .route("/users/:id", get(get_user_generic::<InMemoryUserRepo>))
        .route("/users", post(create_user_generic::<InMemoryUserRepo>))
        .with_state(AppStateGeneric { user_repo });

    Router::new()
        .nest("/dyn", using_dyn)
        .nest("/generic", using_generic)
}

#[derive(Debug, Serialize, Clone)]
struct User {
    id: Uuid,
    name: String,
}

#[derive(Deserialize)]
struct UserParams {
    name: String,
}

trait UserRepo: Send + Sync {
    fn get_user_by_id(&self, id: Uuid) -> Option<User>;

    fn get_user_by_name(&self, name: String) -> Option<User>;

    fn save_user(&self, user: &User);
}

#[derive(Debug, Clone, Default)]
struct InMemoryUserRepo {
    map: Arc<Mutex<HashMap<Uuid, User>>>,
}

impl UserRepo for InMemoryUserRepo {
    fn get_user_by_id(&self, id: Uuid) -> Option<User> {
        self.map.lock().unwrap().get(&id).cloned()
    }

    fn get_user_by_name(&self, name: String) -> Option<User> {
        self.map
            .lock()
            .unwrap()
            .clone()
            .into_values()
            .find(|u| u.name == name)
    }

    fn save_user(&self, user: &User) {
        self.map.lock().unwrap().insert(user.id, user.clone());
    }
}

#[derive(Clone)]
struct AppStateDyn {
    user_repo: Arc<dyn UserRepo>,
}

#[derive(Clone)]
struct AppStateGeneric<T> {
    user_repo: T,
}

async fn create_user_dyn(
    State(state): State<AppStateDyn>,
    Json(params): Json<UserParams>,
) -> Json<User> {
    let old_user = state.user_repo.get_user_by_name(params.name.clone());
    if old_user.is_some() {
        return Json(old_user.unwrap());
    }

    let user = User {
        id: Uuid::new_v4(),
        name: params.name,
    };

    state.user_repo.save_user(&user);

    Json(user)
}

async fn get_user_dyn(
    State(state): State<AppStateDyn>,
    Path(id): Path<Uuid>,
) -> Result<Json<User>, StatusCode> {
    match state.user_repo.get_user_by_id(id) {
        Some(user) => Ok(Json(user)),
        None => Err(StatusCode::NOT_FOUND),
    }
}

async fn create_user_generic<T>(
    State(state): State<AppStateGeneric<T>>,
    Json(params): Json<UserParams>,
) -> Json<User>
where
    T: UserRepo,
{
    let old_user = state.user_repo.get_user_by_name(params.name.clone());
    if old_user.is_some() {
        return Json(old_user.unwrap());
    }

    let user = User {
        id: Uuid::new_v4(),
        name: params.name,
    };

    state.user_repo.save_user(&user);

    Json(user)
}

async fn get_user_generic<T>(
    State(state): State<AppStateGeneric<T>>,
    Path(id): Path<Uuid>,
) -> Result<Json<User>, StatusCode>
where
    T: UserRepo,
{
    match state.user_repo.get_user_by_id(id) {
        Some(user) => Ok(Json(user)),
        None => Err(StatusCode::NOT_FOUND),
    }
}
