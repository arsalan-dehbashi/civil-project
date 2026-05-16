CREATE TABLE IF NOT EXISTS users (
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	phone_number VARCHAR(32) NOT NULL,
	password_hash VARCHAR(255) NULL,
	is_phone_verified BOOLEAN NOT NULL DEFAULT FALSE,
	is_active BOOLEAN NOT NULL DEFAULT TRUE,
	last_login_at DATETIME(3) NULL,
	created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
	deleted_at DATETIME(3) NULL,
	PRIMARY KEY (id),
	UNIQUE KEY idx_users_phone_number (phone_number),
	KEY idx_users_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS roles (
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	name VARCHAR(64) NOT NULL,
	display_name VARCHAR(128) NOT NULL,
	level INT NOT NULL,
	description VARCHAR(255) NULL,
	is_system BOOLEAN NOT NULL DEFAULT FALSE,
	created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
	PRIMARY KEY (id),
	UNIQUE KEY idx_roles_name (name),
	KEY idx_roles_level (level)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS permissions (
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	name VARCHAR(128) NOT NULL,
	resource VARCHAR(64) NOT NULL,
	action VARCHAR(64) NOT NULL,
	description VARCHAR(255) NULL,
	is_system BOOLEAN NOT NULL DEFAULT FALSE,
	created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
	PRIMARY KEY (id),
	UNIQUE KEY idx_permissions_name (name),
	KEY idx_permissions_resource (resource),
	KEY idx_permissions_action (action)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS user_roles (
	user_id BIGINT UNSIGNED NOT NULL,
	role_id BIGINT UNSIGNED NOT NULL,
	created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	PRIMARY KEY (user_id, role_id),
	CONSTRAINT fk_user_roles_user FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT fk_user_roles_role FOREIGN KEY (role_id) REFERENCES roles (id) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS role_permissions (
	role_id BIGINT UNSIGNED NOT NULL,
	permission_id BIGINT UNSIGNED NOT NULL,
	created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	PRIMARY KEY (role_id, permission_id),
	CONSTRAINT fk_role_permissions_role FOREIGN KEY (role_id) REFERENCES roles (id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT fk_role_permissions_permission FOREIGN KEY (permission_id) REFERENCES permissions (id) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS auth_sessions (
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	user_id BIGINT UNSIGNED NOT NULL,
	jti VARCHAR(64) NOT NULL,
	refresh_token_hash VARCHAR(255) NOT NULL,
	user_agent VARCHAR(255) NULL,
	ip_address VARCHAR(64) NULL,
	expires_at DATETIME(3) NOT NULL,
	revoked_at DATETIME(3) NULL,
	created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
	PRIMARY KEY (id),
	UNIQUE KEY idx_auth_sessions_jti (jti),
	UNIQUE KEY idx_auth_sessions_refresh_token_hash (refresh_token_hash),
	KEY idx_auth_sessions_user_id (user_id),
	KEY idx_auth_sessions_expires_at (expires_at),
	KEY idx_auth_sessions_revoked_at (revoked_at),
	CONSTRAINT fk_auth_sessions_user FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS login_attempts (
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	phone_number VARCHAR(32) NOT NULL,
	user_id BIGINT UNSIGNED NULL,
	ip_address VARCHAR(64) NULL,
	user_agent VARCHAR(255) NULL,
	success BOOLEAN NOT NULL DEFAULT FALSE,
	reason VARCHAR(128) NULL,
	created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	PRIMARY KEY (id),
	KEY idx_login_attempts_phone_number (phone_number),
	KEY idx_login_attempts_user_id (user_id),
	KEY idx_login_attempts_ip_address (ip_address),
	KEY idx_login_attempts_success (success),
	KEY idx_login_attempts_created_at (created_at),
	CONSTRAINT fk_login_attempts_user FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
