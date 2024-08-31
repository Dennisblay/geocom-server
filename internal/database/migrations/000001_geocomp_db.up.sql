CREATE TABLE "user" (
                        "id" bigserial PRIMARY KEY,
                        "first_name" varchar(100) NOT NULL,
                        "last_name" varchar(100) NOT NULL,
                        "email" varchar(255) UNIQUE NOT NULL,
                        "phone" varchar(15) UNIQUE NOT NULL,
                        "address" text,
                        "password_hash" varchar(255) NOT NULL,
                        "password_updated_at" timestamptz  DEFAULT '0001-01-01 00:00:00Z',
                        "created_at"             timestamptz     NOT NULL DEFAULT NOW(),
                        "updated_at"             timestamptz     not null DEFAULT (now())
);

CREATE TABLE "token" (
                         "id" bigserial PRIMARY KEY,
                         "user_id" bigint NOT NULL,
                         "token" text NOT NULL,
                         "created_at" timestamp DEFAULT (now()),
                         "expires_at" timestamp NOT NULL
);

CREATE TABLE "ussd_transaction" (
                                    "id" bigserial PRIMARY KEY,
                                    "user_id" bigint NOT NULL,
                                    "transaction_id" varchar(255) UNIQUE NOT NULL,
                                    "amount" decimal(10,2) NOT NULL,
                                    "status" varchar(50) NOT NULL,
                                    "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "computation" (
                               "id" bigserial PRIMARY KEY,
                               "user_id" bigint NOT NULL,
                               "computation_type" varchar(50) NOT NULL,
                               "input_data" text NOT NULL,
                               "result_data" text NOT NULL,
                               "created_at" timestamp DEFAULT (now()),
                               "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "credit" (
                          "id" bigserial PRIMARY KEY,
                          "user_id" bigint NOT NULL,
                          "total_credits" int DEFAULT 0,
                          "created_at" timestamp DEFAULT (now()),
                          "updated_at" timestamp DEFAULT (now())
);

CREATE INDEX ON "user" ("email");

CREATE INDEX ON "user" ("phone");

CREATE INDEX ON "token" ("user_id");

CREATE INDEX ON "token" ("token");

CREATE INDEX ON "ussd_transaction" ("transaction_id");

CREATE INDEX ON "computation" ("computation_type");

CREATE INDEX ON "credit" ("user_id");

ALTER TABLE "token" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "ussd_transaction" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "computation" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "credit" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");