use std::{
    collections::HashMap,
    sync::{
        atomic::{AtomicU64, Ordering},
        Arc, Mutex,
    },
};

use axum::{
    extract::{rejection::JsonRejection, FromRequest, MatchedPath, Request, State},
    http::StatusCode,
    response::{IntoResponse, Response},
    routing::post,
    Router,
};
use serde::{Deserialize, Serialize};
use time_library::TimeStamp;
use tower_http::trace::TraceLayer;
use tracing_subscriber::{layer::SubscriberExt, util::SubscriberInitExt};

#[tokio::main]
async fn main() {
    subscriber();

    let listener = tokio::net::TcpListener::bind("127.0.0.1:3000")
        .await
        .unwrap();

    tracing::debug!("listening on {}", listener.local_addr().unwrap());
    axum::serve(listener, app()).await.unwrap();
}

fn subscriber() {
    let layer = tracing_subscriber::EnvFilter::try_from_default_env()
        .unwrap_or_else(|_| "error_handling=debug,tower_http=debug".into());

    tracing_subscriber::registry()
        .with(layer)
        .with(tracing_subscriber::fmt::layer())
        .init();
}

fn app() -> Router {
    let state = AppState::default();
    let layer = TraceLayer::new_for_http()
        .make_span_with(|req: &Request| {
            let method = req.method();
            let uri = req.uri();

            let matched_path = req
                .extensions()
                .get::<MatchedPath>()
                .map(|matched_path| matched_path.as_str());

            tracing::debug_span!("request", %method, %uri, matched_path)
        })
        .on_failure(());

    Router::new()
        .route("/users", post(users_create))
        .layer(layer)
        .with_state(state)
}

#[derive(FromRequest)]
#[from_request(via(axum::Json), rejection(AppError))]
struct AppJson<T>(T);

impl<T> IntoResponse for AppJson<T>
where
    axum::Json<T>: IntoResponse,
{
    fn into_response(self) -> Response {
        axum::Json(self.0).into_response()
    }
}

#[derive(Default, Clone)]
struct AppState {
    next_id: Arc<AtomicU64>,
    users: Arc<Mutex<HashMap<u64, User>>>,
}

#[derive(Deserialize)]
struct UserParams {
    name: String,
}

#[derive(Serialize, Clone)]
struct User {
    id: u64,
    name: String,
    created_at: TimeStamp,
}

async fn users_create(
    State(state): State<AppState>,
    AppJson(params): AppJson<UserParams>,
) -> Result<AppJson<User>, AppError> {
    let id = state.next_id.fetch_add(1, Ordering::SeqCst);

    let created_at = TimeStamp::now()?;

    let user = User {
        id,
        name: params.name,
        created_at,
    };

    state.users.lock().unwrap().insert(id, user.clone());

    Ok(AppJson(user))
}

enum AppError {
    JsonRejection(JsonRejection),
    TimeError(time_library::Error),
}

impl IntoResponse for AppError {
    fn into_response(self) -> Response {
        #[derive(Serialize)]
        struct ErrorResponse {
            message: String,
        }

        let (status, message) = match self {
            AppError::JsonRejection(rejection) => (rejection.status(), rejection.body_text()),
            AppError::TimeError(err) => {
                tracing::error!(%err, "error from time_library");

                (
                    StatusCode::INTERNAL_SERVER_ERROR,
                    "Something went wrong".to_owned(),
                )
            }
        };

        (status, AppJson(ErrorResponse { message })).into_response()
    }
}

impl From<JsonRejection> for AppError {
    fn from(value: JsonRejection) -> Self {
        Self::JsonRejection(value)
    }
}

impl From<time_library::Error> for AppError {
    fn from(value: time_library::Error) -> Self {
        Self::TimeError(value)
    }
}

mod time_library {
    use std::{
        sync::atomic::{AtomicU64, Ordering},
        time::{SystemTime, UNIX_EPOCH},
    };

    use serde::Serialize;

    #[derive(Serialize, Clone)]
    pub struct TimeStamp(u64);

    impl TimeStamp {
        pub fn now() -> Result<Self, Error> {
            static COUNTER: AtomicU64 = AtomicU64::new(0);

            if COUNTER.fetch_add(1, Ordering::SeqCst) % 3 == 0 {
                Err(Error::FailedToGetTime)
            } else {
                let time = SystemTime::now()
                    .duration_since(UNIX_EPOCH)
                    .unwrap()
                    // .as_millis()
                    .as_secs();
                Ok(Self(time))
            }
        }
    }

    #[derive(Debug)]
    pub enum Error {
        FailedToGetTime,
    }

    impl std::fmt::Display for Error {
        fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
            write!(f, "failed to get time")
        }
    }
}
