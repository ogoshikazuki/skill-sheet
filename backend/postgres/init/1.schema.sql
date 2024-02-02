CREATE TABLE "basic_information" (
    "birthday" DATE NOT NULL,
    "gender" VARCHAR NOT NULL DEFAULT 'MALE',
    "academic_background" VARCHAR NOT NULL DEFAULT ''
);
COMMENT ON COLUMN "basic_information"."gender" IS 'MALE or FEMALE';

CREATE TABLE "projects" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "start_month" DATE NOT NULL,
    "end_month" DATE
);
COMMENT ON COLUMN "projects"."start_month" IS 'day is always 1';
COMMENT ON COLUMN "projects"."end_month" IS 'day is always 1';

CREATE TABLE "technologies" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR NOT NULL
);

CREATE TABLE "project_technology" (
    "project_id" INT NOT NULL,
    "technology_id" INT NOT NULL,
    FOREIGN KEY ("project_id") REFERENCES "projects"("id"),
    FOREIGN KEY ("technology_id") REFERENCES "technologies"("id")
);
