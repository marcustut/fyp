use argon2::{self, Config};
use color_eyre::Result;
use eyre::eyre;
use std::sync::Arc;
use tracing::instrument;

#[derive(Debug, Clone)]
pub struct CryptoService {
    pub key: Arc<String>,
}

impl CryptoService {
    #[instrument]
    pub async fn hash_password(&self, password: String) -> Result<String> {
        let config = Config::default();
        argon2::hash_encoded(password.as_bytes(), self.key.as_bytes(), &config)
            .map_err(|err| eyre!("Hashing error: {:?}", err))
    }
}
