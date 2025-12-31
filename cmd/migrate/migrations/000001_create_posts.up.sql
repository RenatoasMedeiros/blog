CREATE TABLE IF NOT EXISTS posts(
	id bigserial PRIMARY KEY,
	title varchar(255) NOT NULL,
	content text NOT NULL,
	created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
	updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
	is_deleted boolean NOT NULL DEFAULT FALSE
);
