CREATE TABLE "basic_information" (
    "birthday" DATE NOT NULL,
    "gender" VARCHAR NOT NULL DEFAULT 'MALE'
);
COMMENT ON COLUMN "basic_information"."gender" IS 'MALE or FEMALE';
