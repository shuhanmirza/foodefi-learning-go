CREATE TABLE "users"
(
    "username"   varchar PRIMARY KEY,
    "password"   varchar     NOT NULL,
    "role"       varchar     NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "blockchains"
(
    "id"         bigserial PRIMARY KEY,
    "name"       varchar     NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "events"
(
    "id"            bigserial PRIMARY KEY,
    "blockchain_id" bigint  NOT NULL,
    "block_number"  bigint  NOT NULL,
    "event_name"    varchar NOT NULL
);

CREATE TABLE "event_fields"
(
    "id"         bigserial PRIMARY KEY,
    "event_id"   bigint      NOT NULL,
    "name"       varchar     NOT NULL,
    "type"       varchar     NOT NULL,
    "value"      varchar     NOT NULL,
    "recorder"   varchar     NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "events"
    ADD FOREIGN KEY ("blockchain_id") REFERENCES "blockchains" ("id");

ALTER TABLE "event_fields"
    ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "event_fields"
    ADD FOREIGN KEY ("recorder") REFERENCES "users" ("username");

CREATE INDEX ON "users" ("role");

CREATE INDEX ON "events" ("blockchain_id", "block_number");

CREATE INDEX ON "event_fields" ("recorder");

CREATE INDEX ON "event_fields" ("event_id");
