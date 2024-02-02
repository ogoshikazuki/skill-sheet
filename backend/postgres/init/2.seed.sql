INSERT INTO "basic_information"("birthday", "gender", "academic_background") VALUES('1991-07-01', 'MALE', '上智大学卒業');

INSERT INTO "projects"("id", "name", "start_month", "end_month") VALUES
    (1, '人材紹介会社向けクラウド型業務管理システムのリニューアル', '2017-04-01', '2018-08-01'),
    (2, 'オンライン商談システムの管理画面保守開発', '2020-07-01', '2021-03-01'),
    (3, '健診PHR開発プロジェクト', '2021-10-01', NULL);

INSERT INTO "technologies"("id", "name") VALUES
    (1, 'Laravel'),
    (2, 'Vue.js'),
    (3, 'Nuxt'),
    (4, 'Go'),
    (5, 'GraphQL');

INSERT INTO "project_technology"("project_id", "technology_id") VALUES
    (1, 1),
    (1, 2),
    (2, 1),
    (2, 2),
    (2, 3),
    (3, 4),
    (3, 5);
