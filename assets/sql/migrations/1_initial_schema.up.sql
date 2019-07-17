CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "merchant" (
  "merchantuid" uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
  "storename" varchar NOT NULL,
  "storeemail" varchar UNIQUE NOT NULL,
  "created_at" timestamp,
  "last_modified" timestamp
);

CREATE TABLE "product" (
  "productuid" uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
  "productname" varchar UNIQUE NOT NULL,
  "unit" varchar UNIQUE NOT NULL
);

CREATE TABLE "transaction" (
  "transactionuid" uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
  "merchant_id" uuid NOT NULL,
  "recipient_id" uuid NOT NULL,
  "created_at" timestamp
);

CREATE TABLE "product_transaction" (
  "ptuid" uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
  "product_id" uuid NOT NULL,
  "transaction_id" uuid NOT NULL
);

ALTER TABLE "transaction" ADD FOREIGN KEY ("merchant_id") REFERENCES "merchant" ("merchantuid");

ALTER TABLE "product_transaction" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("productuid");

ALTER TABLE "product_transaction" ADD FOREIGN KEY ("transaction_id") REFERENCES "transaction" ("transactionuid");