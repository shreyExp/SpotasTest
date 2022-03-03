-- To update the table for modifying the website values, uncomment the below update query
--update "MY_TABLE"
--   set website = regexp_replace(website, '^(?:.*://)?(?:www\.)?([^:/]*).*$', '\1');
--\echo Hello


-- Below is the query to count how many spots contain the same domain
select website as domain, count(*) as "Number spots for the domain" from "MY_TABLE" where website is not null group by website;

select name as "spots with domain with count greater than 1" from "MY_TABLE"
    where website in
        (select website from
            (select website, count(*) as num from "MY_TABLE" where website is not null group by website) as foo
        where num > 1);
