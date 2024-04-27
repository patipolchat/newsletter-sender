CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE status AS ENUM (
    'pending',
    'completed',
    'failed'
    );

CREATE TABLE "subscribers"
(
    "id"         uuid PRIMARY KEY    NOT NULL DEFAULT gen_random_uuid(),
    "email"      varchar(255) UNIQUE NOT NULL,
    "created_at" timestamp                    DEFAULT (now()),
    "updated_at" timestamp                    DEFAULT (now()),
    "active"     boolean             NOT NULL DEFAULT true
);

CREATE TABLE "newsletters"
(
    "id"         uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    "header"     varchar(255) NOT NULL,
    "body"       varchar      NOT NULL,
    "created_at" timestamp DEFAULT (now()),
    "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "sending_emails"
(
    "id"            uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    "subscriber_id" uuid         NOT NULL,
    "newsletter_id" uuid         NOT NULL,
    "status"        status NOT NULL DEFAULT ('pending'),
    "created_at"    timestamp             DEFAULT (now()),
    "updated_at"    timestamp             DEFAULT (now())
);

ALTER TABLE "sending_emails"
    ADD FOREIGN KEY ("subscriber_id") REFERENCES "subscribers" ("id");

ALTER TABLE "sending_emails"
    ADD FOREIGN KEY ("newsletter_id") REFERENCES "newsletters" ("id");
