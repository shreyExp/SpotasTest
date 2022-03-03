select avg(rating) from "MY_TABLE";
select max(rating) from "MY_TABLE";
select sum(rating) from "MY_TABLE";
select upper(website) from "MY_TABLE" where name = 'Simpsons Restaurant';
select regexp_replace('John Doe', '(.*) (.*)','\2, \1');
--select regexp_replace(website, '', '/^(?:https?:\/\/)?(?:[^@\n]+@)?(?:www\.)?([^:\/\n?]+)/img') from "MY_TABLE" where name = 'CELINE London New Bond Street Store';
-- select regexp_replace(website, '/^(?:https?:\/\/)?(?:[^@\n]+@)?(?:www\.)?([^:\/\n?]+)/img', '\1') as result from "MY_TABLE" where name = 'CELINE London New Bond Street Store';
--select regexp_replace(website, '^(?:.*://)?(?:www\.)?([^:/]*).*$', '\1') as result from "MY_TABLE" where name = 'CELINE London New Bond Street Store';
 --select regexp_replace(website, '^(?:.*://)?(?:www\.)?([^:/]*).*$', '\1') as result from "MY_TABLE";
--select regexp_substr(website, '/^(?:https?:\/\/)?(?:[^@\n]+@)?(?:www\.)?([^:\/\n?]+)/img') from "MY_TABLE" where name = 'Simpsons Restaurant';
-- update "MY_TABLE"
--    set website = regexp_replace(website, '^(?:.*://)?(?:www\.)?([^:/]*).*$', '\1') where name = 'CELINE London New Bond Street Store';
--select * from "MY_TABLE" where name = 'CELINE London New Bond Street Store';
--update "MY_TABLE"
--   set website = regexp_replace(website, '^(?:.*://)?(?:www\.)?([^:/]*).*$', '\1');

