CREATE TABLE "basic_information" (
    "birthday" DATE NOT NULL,
    "gender" VARCHAR NOT NULL DEFAULT 'MALE',
    "academic_background" VARCHAR NOT NULL DEFAULT ''
);
COMMENT ON COLUMN "basic_information"."gender" IS 'MALE or FEMALE';
